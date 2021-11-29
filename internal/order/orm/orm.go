package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	cartOrmPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/cart/orm"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	orderPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type WrapperOrderInterface interface {
	CreateOrder(id int, createOrder orderPkg.CreateOrder, addressId int, cart cart.ResponseCartErrors, courierId int) (int, error)
	GetOrders(id int) (*orderPkg.HistoryOrderArray, error)
	GetOrder(idClient int, idOrder int) (*orderPkg.ActiveOrder, error)
	UpdateStatusOrder(id int, status int) error
	CheckRun(id int) (bool, error)
	DeleteCart(id int) error
	GetCart(id int) (*cart.ResponseCartErrors, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Wrapper struct {
	ConnService cartOrmPkg.ConnectCartServiceInterface
	Ctx         context.Context
	Conn        ConnectionInterface
}

func (db *Wrapper) CreateOrder(id int, createOrder orderPkg.CreateOrder, addressId int, cart cart.ResponseCartErrors, courierId int) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0,  &errPkg.Errors{
			Alias: errPkg.OCreateOrderTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var orderId int

	err = tx.QueryRow(contextTransaction,
		"INSERT INTO order_user (client_id, courier_id, address_id, restaurant_id, promocode_id, comment,"+
			" method_pay, dCost, sumCost) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		id, courierId, addressId, cart.Restaurant.Id, 1, createOrder.Comment, createOrder.MethodPay,
		cart.Cost.DCost, cart.Cost.SumCost).Scan(&orderId)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.OCreateOrderOrderUserNotInsert,
		}
	}

	dishPlace := 0
	elementPlace := 0
	for i, dish := range cart.Dishes {
		var listId int
		err = tx.QueryRow(contextTransaction,
			"INSERT INTO order_list (order_id, food, count_dishes, item_number, place) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			orderId, dish.Id, dish.Count, dish.ItemNumber, dishPlace).Scan(&listId)
		if err != nil {
			return 0, &errPkg.Errors{
				Alias: errPkg.OCreateOrderOrderListNotInsert,
			}
		}

		if dish.RadiosCart != nil {
			for _, radios := range dish.RadiosCart {
				_, err = tx.Exec(contextTransaction,
					"INSERT INTO order_radios_list (order_id, radios_id, radios, food, list_id, place) VALUES ($1, $2, $3, $4, $5, $6)",
					orderId, radios.RadiosId, radios.Id, cart.Dishes[i].Id, listId, elementPlace)
				if err != nil {
					return 0, &errPkg.Errors{
						Alias: errPkg.OCreateOrderOrderRadiosListUserNotInsert,
					}
				}
			}
		}

		if dish.IngredientCart != nil {
			for _, ingredient := range dish.IngredientCart {
				_, err = tx.Exec(contextTransaction,
					"INSERT INTO order_structure_list (order_id, food, structure_food, list_id, place) VALUES ($1, $2, $3, $4, $5)",
					orderId, dish.Id, ingredient.Id, listId, elementPlace)
				if err != nil {
					return 0, &errPkg.Errors{
						Alias: errPkg.OCreateOrderOrderStructureListNotInsert,
					}
				}
				elementPlace++
			}
		}

		var newCount int
		err = tx.QueryRow(contextTransaction,
			"UPDATE dishes SET count = count - $1 WHERE id = $2 RETURNING count",
			dish.Count, dish.Id).Scan(&newCount)
		if err != nil {
			return 0, &errPkg.Errors{
				Alias: errPkg.OCreateOrderCountNotUpdate,
			}
		}

		if newCount < 0 && newCount != util.UnlimitedCount-dish.Count {
			return 0, &errPkg.Errors{
				Alias: errPkg.OCreateOrderCountNotCorrect,
			}
		}
		dishPlace++
		elementPlace = 0
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.OCreateOrderNotCommit,
		}
	}

	return orderId, nil
}

func (db *Wrapper) GetOrders(id int) (*orderPkg.HistoryOrderArray, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var result orderPkg.HistoryOrderArray
	row, err := tx.Query(contextTransaction,
		"SELECT order_user.id, ol.item_number, date_order, status, au.alias, au.city, au.street, au.house,"+
			" au.flat, au.porch, au.floor, au.intercom, au.comment, au.latitude,"+
			" au.longitude, d.id, d.avatar, d.name, ol.count_dishes, "+
			"d.cost, d.kilocalorie, d.weight, d.description, sr.name, "+
			"sr.radios, sr.id, sd.name, sd.id, sd.cost, restaurant_id, r.name, r.avatar, r.city, r.street,"+
			" r.house, r.floor, r.latitude, r.longitude, dCost, sumCost, ol.place, orl.place, osl.place "+
			"FROM order_user"+
			" LEFT JOIN address_user au ON au.id = order_user.address_id"+
			" LEFT JOIN order_list ol ON ol.order_id = order_user.id"+
			" LEFT JOIN dishes d ON d.id = ol.food"+
			" LEFT JOIN order_structure_list osl ON osl.order_id = order_user.id and d.id=osl.food and ol.id=osl.list_id"+
			" LEFT JOIN order_radios_list orl ON orl.order_id = order_user.id and ol.food=orl.food and ol.id=orl.list_id"+
			" LEFT JOIN structure_radios sr ON sr.id = orl.radios"+
			" LEFT JOIN structure_dishes sd ON sd.id = osl.structure_food"+
			" LEFT JOIN restaurant r ON r.id = order_user.restaurant_id WHERE order_user.client_id = $1 ORDER BY date_order", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersNotSelect,
		}
	}

	m := make(map[int]orderPkg.HistoryOrder)
	placeOrder := make(map[int]int)
	place := make(map[int]struct {
		StructureDish map[int]map[int]interface{}
		InfoDishes    map[int]cart.DishesCartResponse
	})
	numberPlaceOrder := 0
	for row.Next() {
		var address profile.AddressCoordinates
		var dish cart.DishesCartResponse
		var order orderPkg.HistoryOrder
		var restaurant orderPkg.HistoryResOrder

		var getPlaceDishes, getPlaceRadios, getPlaceIngredient *int32
		var srRadios, srId, sdId, sdCost *int32
		var srName, sdName *string

		var date time.Time
		err := row.Scan(&order.Id, &dish.ItemNumber, &date, &order.Status, &address.Alias,
			&address.City, &address.Street, &address.House, &address.Flat, &address.Porch,
			&address.Floor, &address.Intercom, &address.Comment, &address.Coordinates.Latitude,
			&address.Coordinates.Longitude, &dish.Id, &dish.Img, &dish.Name, &dish.Count,
			&dish.Cost, &dish.Kilocalorie, &dish.Weight, &dish.Description, &srName, &srRadios,
			&srId, &sdName, &sdId, &sdCost, &restaurant.Id, &restaurant.Name, &restaurant.Img,
			&restaurant.Address.City, &restaurant.Address.Street, &restaurant.Address.House,
			&restaurant.Address.Floor, &restaurant.Address.Coordinates.Latitude, &restaurant.Address.Coordinates.Longitude,
			&order.Cart.Cost.DCost, &order.Cart.Cost.SumCost, &getPlaceDishes, &getPlaceRadios, &getPlaceIngredient)

		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.OGetOrdersNotScan,
			}
		}
		switch order.Status {
		case 0:
			order.Status = 1
		case 1, 2, 3:
			order.Status = 2
		case 4:
			order.Status = 3
		}

		order.Date, order.Time = util.FormatDate(date)

		placeDishes := util.ConvertInt32ToInt(getPlaceDishes)
		placeRadios := util.ConvertInt32ToInt(getPlaceRadios)
		placeIngredient := util.ConvertInt32ToInt(getPlaceIngredient)

		var ingredient cart.IngredientCartResponse
		if sdName != nil {
			ingredient.Name = *sdName
			ingredient.Id = int(*sdId)
			ingredient.Cost = int(*sdCost)
		}

		var radios cart.RadiosCartResponse
		if srName != nil {
			radios.Name = *srName
			radios.RadiosId = int(*srRadios)
			radios.Id = int(*srId)
		}

		temp := place[order.Id]
		if temp.InfoDishes == nil {
			temp.InfoDishes = make(map[int]cart.DishesCartResponse)
		}

		if placeRadios != -1 {
			if temp.StructureDish == nil {
				temp.StructureDish = make(map[int]map[int]interface{})
			}
			if temp.StructureDish[placeDishes] == nil {
				temp.StructureDish[placeDishes] = make(map[int]interface{})
			}
			temp.StructureDish[placeDishes][placeRadios] = radios
		}

		if placeIngredient != -1 {
			if temp.StructureDish == nil {
				temp.StructureDish = make(map[int]map[int]interface{})
			}
			if temp.StructureDish[placeDishes] == nil {
				temp.StructureDish[placeDishes] = make(map[int]interface{})
			}
			temp.StructureDish[placeDishes][placeIngredient] = ingredient
		}

		temp.InfoDishes[placeDishes] = dish

		place[order.Id] = temp

		if _, ok := m[order.Id]; !ok {
			order.Address = address
			order.Restaurant = restaurant
			m[order.Id] = order
			placeOrder[numberPlaceOrder] = order.Id
			numberPlaceOrder++
		}
	}

	for i := 0; i < len(placeOrder); i++ {
		order := m[placeOrder[i]]
		for j := 0; j < len(place[placeOrder[i]].InfoDishes); j++ {
			structDish := place[placeOrder[i]].StructureDish[j]
			dish := place[placeOrder[i]].InfoDishes[j]
			for k := 0; k < len(structDish); k++ {
				switch structDish[k].(type) {
				case cart.RadiosCartResponse:
					dish.RadiosCart = append(dish.RadiosCart, structDish[k].(cart.RadiosCartResponse))
				case cart.IngredientCartResponse:
					dish.IngredientCart = append(dish.IngredientCart, structDish[k].(cart.IngredientCartResponse))
				}
			}
			order.Cart.Dishes = append(order.Cart.Dishes, dish)
		}
		result.Orders = append(result.Orders, order)
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersNotCommit,
		}
	}

	if result.Orders == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersOrdersIsVoid,
		}
	}
	return &result, nil
}

func (db *Wrapper) GetOrder(idClient int, idOrder int) (*orderPkg.ActiveOrder, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrderTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	row, err := tx.Query(contextTransaction,
		"SELECT order_user.id, ol.item_number, date_order, status, au.alias, au.city, au.street, au.house,"+
			" au.flat, au.porch, au.floor, au.intercom, au.comment, au.latitude,"+
			" au.longitude, d.id, d.avatar, d.name, ol.count_dishes, "+
			"d.cost, d.kilocalorie, d.weight, d.description, sr.name, "+
			"sr.radios, sr.id, sd.name, sd.id, sd.cost, restaurant_id, r.name, r.avatar, r.city, r.street,"+
			" r.house, r.floor, r.latitude, r.longitude, dCost, sumCost, ol.place, orl.place, osl.place "+
			"FROM order_user"+
			" LEFT JOIN address_user au ON au.id = order_user.address_id"+
			" LEFT JOIN order_list ol ON ol.order_id = order_user.id"+
			" LEFT JOIN dishes d ON d.id = ol.food"+
			" LEFT JOIN order_structure_list osl ON osl.order_id = order_user.id and d.id=osl.food and ol.id=osl.list_id"+
			" LEFT JOIN order_radios_list orl ON orl.order_id = order_user.id and ol.food=orl.food and ol.id=orl.list_id"+
			" LEFT JOIN structure_radios sr ON sr.id = orl.radios"+
			" LEFT JOIN structure_dishes sd ON sd.id = osl.structure_food"+
			" LEFT JOIN restaurant r ON r.id = order_user.restaurant_id WHERE order_user.client_id = $1 AND order_user.id = $2", idClient, idOrder)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrderNotSelect,
		}
	}

	structureDish := make(map[int]map[int]interface{})
	infoDishes := make(map[int]cart.DishesCartResponse)

	var order orderPkg.ActiveOrder

	for row.Next() {
		var address profile.AddressCoordinates
		var dish cart.DishesCartResponse
		var restaurant orderPkg.HistoryResOrder

		var getPlaceDishes, getPlaceRadios, getPlaceIngredient *int32
		var srRadios, srId, sdId, sdCost *int32
		var srName, sdName *string

		var date time.Time
		err := row.Scan(&order.Id, &dish.ItemNumber, &date, &order.Status, &address.Alias,
			&address.City, &address.Street, &address.House, &address.Flat, &address.Porch,
			&address.Floor, &address.Intercom, &address.Comment, &address.Coordinates.Latitude,
			&address.Coordinates.Longitude, &dish.Id, &dish.Img, &dish.Name, &dish.Count,
			&dish.Cost, &dish.Kilocalorie, &dish.Weight, &dish.Description, &srName, &srRadios,
			&srId, &sdName, &sdId, &sdCost, &restaurant.Id, &restaurant.Name, &restaurant.Img,
			&restaurant.Address.City, &restaurant.Address.Street, &restaurant.Address.House,
			&restaurant.Address.Floor, &restaurant.Address.Coordinates.Latitude, &restaurant.Address.Coordinates.Longitude,
			&order.Cart.Cost.DCost, &order.Cart.Cost.SumCost, &getPlaceDishes, &getPlaceRadios, &getPlaceIngredient)

		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.OGetOrderNotScan,
			}
		}

		order.Date, order.Time = util.FormatDate(date)

		placeDishes := util.ConvertInt32ToInt(getPlaceDishes)
		placeRadios := util.ConvertInt32ToInt(getPlaceRadios)
		placeIngredient := util.ConvertInt32ToInt(getPlaceIngredient)

		var ingredient cart.IngredientCartResponse
		if sdName != nil {
			ingredient.Name = *sdName
			ingredient.Id = int(*sdId)
			ingredient.Cost = int(*sdCost)
		}

		var radios cart.RadiosCartResponse
		if srName != nil {
			radios.Name = *srName
			radios.RadiosId = int(*srRadios)
			radios.Id = int(*srId)
		}

		if placeRadios != -1 {
			if structureDish == nil {
				structureDish = make(map[int]map[int]interface{})
			}
			if structureDish[placeDishes] == nil {
				structureDish[placeDishes] = make(map[int]interface{})
			}
			structureDish[placeDishes][placeRadios] = radios
		}

		if placeIngredient != -1 {
			if structureDish == nil {
				structureDish = make(map[int]map[int]interface{})
			}
			if structureDish[placeDishes] == nil {
				structureDish[placeDishes] = make(map[int]interface{})
			}
			structureDish[placeDishes][placeIngredient] = ingredient
		}

		infoDishes[placeDishes] = dish

		order.Address = address
		order.Restaurant = restaurant
	}

	for j := 0; j < len(infoDishes); j++ {
		structDish := structureDish[j]
		dish := infoDishes[j]
		for k := 0; k < len(structDish); k++ {
			switch structDish[k].(type) {
			case cart.RadiosCartResponse:
				dish.RadiosCart = append(dish.RadiosCart, structDish[k].(cart.RadiosCartResponse))
			case cart.IngredientCartResponse:
				dish.IngredientCart = append(dish.IngredientCart, structDish[k].(cart.IngredientCartResponse))
			}
		}
		order.Cart.Dishes = append(order.Cart.Dishes, dish)
	}

	if order.Id == 0 {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrderNotExist,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrderNotCommit,
		}
	}

	return &order, nil
}

func (db *Wrapper) UpdateStatusOrder(id int, status int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OUpdateStatusOrderTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE order_user SET status = $1 WHERE id = $2",
		status, id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OUpdateStatusOrderNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OUpdateStatusOrderNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) CheckRun(id int) (bool, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.OUpdateStatusOrderTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var check bool
	err = tx.QueryRow(contextTransaction,
		"SELECT check_run FROM order_user WHERE id = $1",
		id).Scan(&check)
	_, err = tx.Exec(contextTransaction,
		"UPDATE order_user SET check_run = false WHERE id = $1",
		id)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.OUpdateStatusOrderNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.OUpdateStatusOrderNotCommit,
		}
	}

	return check, nil
}

func (db *Wrapper) GetCart(id int) (*cart.ResponseCartErrors, error) {
	var cartId *cartProto.CartId
	cartId = &cartProto.CartId{}
	cartId.Id = int64(id)
	receivedCart, err := db.ConnService.GetCart(db.Ctx, cartId)
	if err != nil {
		return nil, err
	}
	if receivedCart.Error!= "" {
		return nil, &errPkg.Errors{
			Alias: receivedCart.Error,
		}
	}
	cart := cast.CastResponseCartErrorsProtoToResponseCartErrors(receivedCart)

	if cart.DishErr != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetCartCartNoActual,
		}
	}

	return cart, nil
}

func (db *Wrapper) GetRestaurant(id int) (*restaurant.RestaurantId, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var restaurant restaurant.RestaurantId
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(
		&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
		&restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantRestaurantNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantNotCommit,
		}
	}

	return &restaurant, nil
}

func (db *Wrapper) DeleteCart(id int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.CDeleteCartTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"DELETE FROM cart_food CASCADE WHERE client_id = $1", id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.CDeleteCartCartNotDelete,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.CDeleteCartNotCommit,
		}
	}
	return nil
}
