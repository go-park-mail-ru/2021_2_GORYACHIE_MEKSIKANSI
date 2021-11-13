package Order

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"strings"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) CreateOrder(id int, createOrder utils.CreateOrder, addressId int, cart utils.ResponseCartErrors, courierId int) error {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.OCreateOrderTransactionNotCreate,
			Time: time.Now(),
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var orderId int

	err = tx.QueryRow(context.Background(),
		"INSERT INTO order_user (client_id, courier_id, address_id, restaurant_id, promocode_id, comment," +
		" method_pay, dCost, sumCost) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		id, courierId, addressId, cart.Restaurant.Id, 1, createOrder.Comment, createOrder.MethodPay, cart.Cost.DCost, cart.Cost.SumCost).Scan(&orderId)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.OCreateOrderOrderUserNotInsert,
			Time: time.Now(),
		}
	}

	for _, dish := range cart.Dishes {
		_, err = tx.Exec(context.Background(),
			"INSERT INTO order_list (order_id, food, count) VALUES ($1, $2, $3)",
			orderId, dish.Id, dish.Count)
		if err != nil {
			return &errorsConst.Errors{
				Text: errorsConst.OCreateOrderOrderListNotInsert,
				Time: time.Now(),
			}
		}

		if dish.IngredientCart != nil {
			for _, ingredient := range dish.IngredientCart {
				_, err = tx.Exec(context.Background(),
					"INSERT INTO order_structure_list (order_id, food, structure_food) VALUES ($1, $2, $3)",
					orderId, dish.Id, ingredient.Id)
				if err != nil {
					return &errorsConst.Errors{
						Text: errorsConst.OCreateOrderOrderStructureListNotInsert,
						Time: time.Now(),
					}
				}
			}
		}

		if dish.RadiosCart != nil {
			for _, radios := range dish.RadiosCart {
				_, err = tx.Exec(context.Background(),
					"INSERT INTO order_radios_list (order_id, radios_id, radios) VALUES ($1, $2, $3)",
					orderId, dish.Id, radios.Id)
				if err != nil {
					return &errorsConst.Errors{
						Text: errorsConst.OCreateOrderOrderRadiosListUserNotInsert,
						Time: time.Now(),
					}
				}
			}
		}

		var newCount int
		err = tx.QueryRow(context.Background(),
			"UPDATE dishes SET count = count - $1 WHERE id = $2 RETURNING count",
			dish.Count, dish.Id).Scan(&newCount)
		if err != nil {
			return &errorsConst.Errors{
				Text: errorsConst.OCreateOrderCountNotUpdate,
				Time: time.Now(),
			}
		}

		// TODO: make define
		if newCount < 0 && newCount != -1 - dish.Count {
			return &errorsConst.Errors{
				Text: errorsConst.OCreateOrderCountNotCorrect,
				Time: time.Now(),
			}
		}

	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.OCreateOrderNotCommit,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) GetOrders(id int) (*utils.HistoryOrderArray, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.OGetOrdersTransactionNotCreate,
			Time: time.Now(),
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var result utils.HistoryOrderArray
	row, err := tx.Query(context.Background(),
		"SELECT order_user.id, date_order, status, au.alias, au.city, au.street, au.house," +
		" au.flat, au.porch, au.floor, au.intercom, au.comment, au.latitude," +
		" au.longitude, d.id, d.avatar, d.name, d.cost, ol.count, " +
		"d.cost, d.kilocalorie, d.weight, d.description, sr.name, " +
		"sr.radios, sr.id, sd.name, sd.id, sd.cost, restaurant_id, r.name, r.avatar, r.city, r.street," +
		" r.house, r.floor, r.latitude, r.longitude, dCost, sumCost " +
		"FROM order_user" +
		" LEFT JOIN address_user au ON au.id = order_user.address_id" +
		" LEFT JOIN order_list ol ON ol.order_id = order_user.id" +
		" LEFT JOIN order_structure_list osl ON osl.order_id = order_user.id" +
		" LEFT JOIN order_radios_list orl ON orl.order_id = order_user.id" +
		" LEFT JOIN structure_radios sr ON sr.radios = orl.radios_id" +
		" LEFT JOIN structure_dishes sd ON sd.food = osl.food" +
		" LEFT JOIN dishes d ON d.id = ol.food" +
		" LEFT JOIN restaurant r ON r.id = order_user.restaurant_id WHERE order_user.client_id = $1", id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetMenuDishesNotSelect,
			Time: time.Now(),
		}
	}

	m := make(map[int]utils.HistoryOrder)
	for row.Next() {
		var idOrder int
		var address utils.AddressCoordinates
		var dish utils.DishesCartResponse
		var radios utils.RadiosCartResponse
		var ingredient utils.IngredientCartResponse
		var order utils.HistoryOrder
		var restaurant utils.HistoryResOrder

		var srName, srRadios, srId interface{}
		var sdName, sdId, sdCost interface{}

		err := row.Scan(&idOrder, &order.Date, &order.Status, &address.Alias,
			&address.City, &address.Street, &address.House, &address.Flat, &address.Porch,
			&address.Floor, &address.Intercom, &address.Comment, &address.Coordinates.Latitude,
			&address.Coordinates.Longitude, &dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Count,
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

		if _, ok := m[idOrder]; ok {
			maxIndex := len(m[idOrder].Cart.Dishes) - 1
			if sdName != nil {
				m[idOrder].Cart.Dishes[maxIndex].IngredientCart = append(m[idOrder].Cart.Dishes[maxIndex].IngredientCart, ingredient)
			}

			if srName != nil {
				m[idOrder].Cart.Dishes[maxIndex].RadiosCart = append(m[idOrder].Cart.Dishes[maxIndex].RadiosCart, radios)
			}

		} else {
			order.Address = address
			order.Restaurant = restaurant
			order.Cart.Dishes = append(order.Cart.Dishes, dish)
			m[idOrder] = order
		}
	}

	for _, order := range m {
		result.Orders = append(result.Orders, order)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.OGetOrdersDishesNotCommit,
			Time: time.Now(),
		}
	}

	return &result, nil
}

func (db *Wrapper) GetPriceDelivery(id int) (int, error) {
	var price int
	err := db.Conn.QueryRow(context.Background(),
		"SELECT price_delivery FROM restaurant WHERE id = $1", id).Scan(&price)
	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "no rows") {
			return 0, &errorsConst.Errors{
				Text: errorsConst.CGetPriceDeliveryPriceNotFound,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.CGetPriceDeliveryPriceNotScan,
			Time: time.Now(),
		}
	}
	return price, nil
}
