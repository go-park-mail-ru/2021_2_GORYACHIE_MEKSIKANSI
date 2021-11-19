package Order

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) CreateOrder(id int, createOrder utils.CreateOrder, addressId int, cart utils.ResponseCartErrors, courierId int) error {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OCreateOrderTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var orderId int

	err = tx.QueryRow(context.Background(),
		"INSERT INTO order_user (client_id, courier_id, address_id, restaurant_id, promocode_id, comment,"+
			" method_pay, dCost, sumCost) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		id, courierId, addressId, cart.Restaurant.Id, 1, createOrder.Comment, createOrder.MethodPay,
		cart.Cost.DCost, cart.Cost.SumCost).Scan(&orderId)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OCreateOrderOrderUserNotInsert,
		}
	}

	for i, dish := range cart.Dishes {
		var listId int
		err = tx.QueryRow(context.Background(),
			"INSERT INTO order_list (order_id, food, count, item_number) VALUES ($1, $2, $3, $4) RETURNING id",
			orderId, dish.Id, dish.Count, dish.ItemNumber).Scan(&listId)
		if err != nil {
			return &errPkg.Errors{
				Alias: errPkg.OCreateOrderOrderListNotInsert,
			}
		}

		if dish.IngredientCart != nil {
			for _, ingredient := range dish.IngredientCart {
				_, err = tx.Exec(context.Background(),
					"INSERT INTO order_structure_list (order_id, food, structure_food) VALUES ($1, $2, $3)",
					orderId, dish.Id, ingredient.Id)
				if err != nil {
					return &errPkg.Errors{
						Alias: errPkg.OCreateOrderOrderStructureListNotInsert,
					}
				}
			}
		}

		if dish.RadiosCart != nil {
			for _, radios := range dish.RadiosCart {
				_, err = tx.Exec(context.Background(),
					"INSERT INTO order_radios_list (order_id, radios_id, radios, food, list_id) VALUES ($1, $2, $3, $4, $5)",
					orderId, radios.RadiosId, radios.Id, cart.Dishes[i].Id, listId)
				if err != nil {
					return &errPkg.Errors{
						Alias: errPkg.OCreateOrderOrderRadiosListUserNotInsert,
					}
				}
			}
		}

		var newCount int
		err = tx.QueryRow(context.Background(),
			"UPDATE dishes SET count = count - $1 WHERE id = $2 RETURNING count",
			dish.Count, dish.Id).Scan(&newCount)
		if err != nil {
			return &errPkg.Errors{
				Alias: errPkg.OCreateOrderCountNotUpdate,
			}
		}

		if newCount < 0 && newCount != -1-dish.Count {
			return &errPkg.Errors{
				Alias: errPkg.OCreateOrderCountNotCorrect,
			}
		}

	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.OCreateOrderNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) GetOrders(id int) (*utils.HistoryOrderArray, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var result utils.HistoryOrderArray
	row, err := tx.Query(context.Background(),
		"SELECT order_user.id, ol.item_number, date_order, status, au.alias, au.city, au.street, au.house,"+
			" au.flat, au.porch, au.floor, au.intercom, au.comment, au.latitude,"+
			" au.longitude, d.id, d.avatar, d.name, ol.count, "+
			"d.cost, d.kilocalorie, d.weight, d.description, sr.name, "+
			"sr.radios, sr.id, sd.name, sd.id, sd.cost, restaurant_id, r.name, r.avatar, r.city, r.street,"+
			" r.house, r.floor, r.latitude, r.longitude, dCost, sumCost "+
			"FROM order_user"+
			" LEFT JOIN address_user au ON au.id = order_user.address_id"+
			" LEFT JOIN order_list ol ON ol.order_id = order_user.id"+
			" LEFT JOIN dishes d ON d.id = ol.food"+
			" LEFT JOIN order_structure_list osl ON osl.order_id = order_user.id and d.id=osl.food"+
			" LEFT JOIN order_radios_list orl ON orl.order_id = order_user.id and ol.food=orl.food and ol.id=orl.list_id"+
			" LEFT JOIN structure_radios sr ON sr.id = orl.radios"+
			" LEFT JOIN structure_dishes sd ON sd.id = osl.structure_food"+
			" LEFT JOIN restaurant r ON r.id = order_user.restaurant_id WHERE order_user.client_id = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesNotSelect,
		}
	}

	m := make(map[int]utils.HistoryOrder)
	mAdditionalInfo := make(map[int]interface{})
	for row.Next() {
		var orderId int
		var address utils.AddressCoordinates
		var dish utils.DishesCartResponse
		var radios utils.RadiosCartResponse
		var ingredient utils.IngredientCartResponse
		var order utils.HistoryOrder
		var restaurant utils.HistoryResOrder

		var srName, srRadios, srId interface{}
		var sdName, sdId, sdCost interface{}

		err := row.Scan(&orderId, &dish.ItemNumber, &order.Date, &order.Status, &address.Alias,
			&address.City, &address.Street, &address.House, &address.Flat, &address.Porch,
			&address.Floor, &address.Intercom, &address.Comment, &address.Coordinates.Latitude,
			&address.Coordinates.Longitude, &dish.Id, &dish.Img, &dish.Name, &dish.Count,
			&dish.Cost, &dish.Kilocalorie, &dish.Weight, &dish.Description, &srName, &srRadios,
			&srId, &sdName, &sdId, &sdCost, &restaurant.Id, &restaurant.Name, &restaurant.Img,
			&restaurant.Address.City, &restaurant.Address.Street, &restaurant.Address.House, &restaurant.Address.Floor,
			&restaurant.Address.Coordinates.Latitude, &restaurant.Address.Coordinates.Longitude,
			&order.Cart.Cost.DCost, &order.Cart.Cost.SumCost)

		if err != nil {
			return nil, err
		}

		if sdName != nil {
			ingredient.Name = sdName.(string)
			ingredient.Id = int(sdId.(int32))
			ingredient.Cost = int(sdCost.(int32))
			dish.IngredientCart = append(dish.IngredientCart, ingredient)
		}

		if srName != nil {
			radios.Name = srName.(string)
			radios.RadiosId = int(srRadios.(int32))
			radios.Id = int(srId.(int32))
			dish.RadiosCart = append(dish.RadiosCart, radios)
		}

		if val, ok := m[orderId]; ok {
			val.Cart.Dishes = append(val.Cart.Dishes, dish)
			m[orderId] = val

		} else {
			order.Cart.Dishes = append(order.Cart.Dishes, dish)
			order.Address = address
			order.Restaurant = restaurant
			mAdditionalInfo[orderId] = orderId

			m[orderId] = order
		}
	}

	for _, order := range m {
		result.Orders = append(result.Orders, order)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.OGetOrdersDishesNotCommit,
		}
	}

	return &result, nil
}

func (db *Wrapper) GetPriceDelivery(id int) (int, error) {
	var price int
	err := db.Conn.QueryRow(context.Background(),
		"SELECT price_delivery FROM restaurant WHERE id = $1", id).Scan(&price)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Alias: errPkg.CGetPriceDeliveryPriceNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.CGetPriceDeliveryPriceNotScan,
		}
	}
	return price, nil
}
