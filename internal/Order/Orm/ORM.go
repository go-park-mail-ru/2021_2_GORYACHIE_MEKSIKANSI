package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Order"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
	"time"
)

type Wrapper struct {
	ConnService Interface.ConnectCartService
	Ctx         context.Context
	Conn        Interface.ConnectionInterface
}

func (db *Wrapper) CreateOrder(id int, createOrder Order.CreateOrder, addressId int, cart Cart.ResponseCartErrors, courierId int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
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
		return &errPkg.Errors{
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
			return &errPkg.Errors{
				Alias: errPkg.OCreateOrderOrderListNotInsert,
			}
		}

		if dish.RadiosCart != nil {
			for _, radios := range dish.RadiosCart {
				_, err = tx.Exec(contextTransaction,
					"INSERT INTO order_radios_list (order_id, radios_id, radios, food, list_id, place) VALUES ($1, $2, $3, $4, $5, $6)",
					orderId, radios.RadiosId, radios.Id, cart.Dishes[i].Id, listId, elementPlace)
				if err != nil {
					return &errPkg.Errors{
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
					return &errPkg.Errors{
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
			return &errPkg.Errors{
				Alias: errPkg.OCreateOrderCountNotUpdate,
			}
		}

		if newCount < 0 && newCount != Util.UnlimitedCount-dish.Count {
			return &errPkg.Errors{
				Alias: errPkg.OCreateOrderCountNotCorrect,
			}
		}
		dishPlace++
		elementPlace = 0
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OCreateOrderNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) GetOrders(id int) (*Order.HistoryOrderArray, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var result Order.HistoryOrderArray
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

	m := make(map[int]Order.HistoryOrder)
	placeOrder := make(map[int]int)
	place := make(map[int]struct {
		StructureDish map[int]map[int]interface{}
		InfoDishes    map[int]Cart.DishesCartResponse
	})
	numberPlaceOrder := 0
	for row.Next() {
		var address Profile.AddressCoordinates
		var dish Cart.DishesCartResponse
		var order Order.HistoryOrder
		var restaurant Order.HistoryResOrder

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

		order.Date, order.Time = Util.FormatDate(date)

		placeDishes := Util.ConvertInt32ToInt(getPlaceDishes)
		placeRadios := Util.ConvertInt32ToInt(getPlaceRadios)
		placeIngredient := Util.ConvertInt32ToInt(getPlaceIngredient)

		var ingredient Cart.IngredientCartResponse
		if sdName != nil {
			ingredient.Name = *sdName
			ingredient.Id = int(*sdId)
			ingredient.Cost = int(*sdCost)
		}

		var radios Cart.RadiosCartResponse
		if srName != nil {
			radios.Name = *srName
			radios.RadiosId = int(*srRadios)
			radios.Id = int(*srId)
		}

		temp := place[order.Id]
		if temp.InfoDishes == nil {
			temp.InfoDishes = make(map[int]Cart.DishesCartResponse)
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
				case Cart.RadiosCartResponse:
					dish.RadiosCart = append(dish.RadiosCart, structDish[k].(Cart.RadiosCartResponse))
				case Cart.IngredientCartResponse:
					dish.IngredientCart = append(dish.IngredientCart, structDish[k].(Cart.IngredientCartResponse))
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

	return &result, nil
}

func (db *Wrapper) GetOrder(idClient int, idOrder int) (*Order.ActiveOrder, error) {
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
	infoDishes := make(map[int]Cart.DishesCartResponse)

	var order Order.ActiveOrder

	for row.Next() {
		var address Profile.AddressCoordinates
		var dish Cart.DishesCartResponse
		var restaurant Order.HistoryResOrder

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

		order.Date, order.Time = Util.FormatDate(date)

		placeDishes := Util.ConvertInt32ToInt(getPlaceDishes)
		placeRadios := Util.ConvertInt32ToInt(getPlaceRadios)
		placeIngredient := Util.ConvertInt32ToInt(getPlaceIngredient)

		var ingredient Cart.IngredientCartResponse
		if sdName != nil {
			ingredient.Name = *sdName
			ingredient.Id = int(*sdId)
			ingredient.Cost = int(*sdCost)
		}

		var radios Cart.RadiosCartResponse
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
			case Cart.RadiosCartResponse:
				dish.RadiosCart = append(dish.RadiosCart, structDish[k].(Cart.RadiosCartResponse))
			case Cart.IngredientCartResponse:
				dish.IngredientCart = append(dish.IngredientCart, structDish[k].(Cart.IngredientCartResponse))
			}
		}
		order.Cart.Dishes = append(order.Cart.Dishes, dish)
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

func (db *Wrapper) GetCart(id int) (*Cart.ResponseCartErrors, error) {
	var cartId *cartProto.CartId
	cartId = &cartProto.CartId{}
	cartId.Id = int64(id)
	receivedCart, err := db.ConnService.GetCart(db.Ctx, cartId)
	if err != nil {
		return nil, err
	}
	cart := cast.CastResponseCartErrorsProtoToResponseCartErrors(receivedCart)

	if cart.DishErr != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetCartCartNoActual,
		}
	}

	return cart, nil
}

func (db *Wrapper) GetRestaurant(id int) (*Restaurant.RestaurantId, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var restaurant Restaurant.RestaurantId
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
