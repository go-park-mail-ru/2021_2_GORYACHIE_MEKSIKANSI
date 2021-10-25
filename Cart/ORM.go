package Cart

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"time"
)

type Wrapper struct {
	Conn Utils.ConnectionInterface
}

func (db *Wrapper) GetCart(id int) (*Utils.CartResponse, error) {
	var cart *Utils.CartResponse
	cart = &Utils.CartResponse{}
	var dishes []Utils.DishesCartResponse
	var radios []Utils.RadiosCartResponse
	var ingredients []Utils.IngredientCartResponse

	var restaurant Utils.RestaurantIdCastResponse
	restaurant.Id = id
	err := db.Conn.QueryRow(context.Background(),
		"SELECT restaurant_id FROM cart WHERE client_id = $1", id).Scan(&restaurant.Id)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, &errorsConst.Errors{
				Text: errorsConst.GetCartRestaurantNotFound,
				Time: time.Now(),
			}
		}
		return nil, &errorsConst.Errors{
			Text: errorsConst.GetCartRestaurantNotScan,
			Time: time.Now(),
		}
	}
	cart.Restaurant = restaurant

	rows, err := db.Conn.Query(context.Background(),
		"SELECT food, count_food, number_item FROM cart WHERE client_id = $1", id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.GetCartCartNotSelect,
			Time: time.Now(),
		}
	}

	for rows.Next() {
		var dish *Utils.DishesCartResponse
		dish = &Utils.DishesCartResponse{}
		err = rows.Scan(&dish.Id, &dish.Count, &dish.ItemNumber)
		rows, err = db.Conn.Query(context.Background(),
			"SELECT name, cost, description, avatar FROM dishes WHERE id = $1", dish.Id)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.GetCartDishesNotScan,
				Time: time.Now(),
			}
		}
		for rows.Next() {
			err = rows.Scan(&dish.Name, &dish.Cost, &dish.Description, &dish.Img)

			rows, err := db.Conn.Query(context.Background(),
				"SELECT checkbox FROM cart_structure_food WHERE client_id = $1", id)
			if err != nil {
				return nil, &errorsConst.Errors{
					Text: errorsConst.GetCartRestaurantNotSelect,
					Time: time.Now(),
				}
			}
			for rows.Next() {
				var ingredient Utils.IngredientCartResponse
				err = rows.Scan(&ingredient.Id)
				if err != nil {
					return nil, &errorsConst.Errors{
						Text: errorsConst.GetCartCheckboxNotScan,
						Time: time.Now(),
					}
				}

				err = db.Conn.QueryRow(context.Background(),
					"SELECT name, cost FROM structure_dishes WHERE id = $1", ingredient.Id).Scan( // TODO(N): не запускать 2 раза
						&ingredient.Name, &ingredient.Cost)
				ingredients = append(ingredients, ingredient)
			}
			if ingredients != nil {
				dish.IngredientCart = ingredients
			}

			rows, err = db.Conn.Query(context.Background(),
				"SELECT radios_id, radios FROM cart_radios_food WHERE client_id = $1", id)
			if err != nil {
				return nil, &errorsConst.Errors{
					Text: errorsConst.GetCartRadiosNotSelect,
					Time: time.Now(),
				}
			}
			for rows.Next() {
				var radio Utils.RadiosCartResponse
				err = rows.Scan(&radio.RadiosId, &radio.Id)
				if err != nil {
					return nil, &errorsConst.Errors{
						Text: errorsConst.GetCartRadiosNotScan,
						Time: time.Now(),
					}
				}

				err = db.Conn.QueryRow(context.Background(),
					"SELECT name FROM structure_radios WHERE id = $1", radio.Id).Scan(&radio.Name)
				if err != nil {
					if err.Error() == "no rows in result set" {
						return nil, &errorsConst.Errors{
							Text: errorsConst.GetCartStructRadiosNotFound,
							Time: time.Now(),
						}
					}
					return nil, &errorsConst.Errors{
						Text: errorsConst.GetCartStructRadiosNowScan,
						Time: time.Now(),
					}
				}
				radios = append(radios, radio)
			}
			if radios != nil {
				dish.RadiosCart = radios
			}
			dishes = append(dishes, *dish)
		}
	}
	if cart == nil {
		if err.Error() == "no rows in result set" {
			return nil, &errorsConst.Errors{
				Text: errorsConst.GetCartDishesNotFound,
				Time: time.Now(),
			}
		}
	}
	cart.Dishes = dishes
	return cart, nil
}


func (db *Wrapper) DeleteCart(id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM cart WHERE client_id = $1", id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.CartNotDelete,
			Time: time.Now(),
		}
	}
	_, err = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_structure_food WHERE client_id = $1", id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.StructureFoodNotDelete,
			Time: time.Now(),
		}
	}
	_, err = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_radios_food WHERE client_id = $1", id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.CartRadiosFoodNotDelete,
			Time: time.Now(),
		}
	}
	return nil
}


func (db *Wrapper) GetConn() Utils.ConnectionInterface {
	return db.Conn
}


func (db *Wrapper) UpdateCart(newCart Utils.CartRequest, clientId int) (*Utils.CartResponse, []Utils.CastDishesErrs, error) {
	var dishesErrors []Utils.CastDishesErrs
	var cart Utils.CartResponse
	for i, dish := range newCart.Dishes {
		var dishes Utils.DishesCartResponse
		var dishesError Utils.CastDishesErrs
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, avatar, cost, name, description FROM dishes WHERE id = $1 ", dish.Id).Scan(
				&dishes.Id, &dishes.Img, &dishes.Cost, &dishes.Name, &dishes.Description)
		if err != nil {
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.Explain = dishes.Name
			dishesErrors = append(dishesErrors, dishesError)
			continue
		}
		cart.Dishes = append(cart.Dishes, dishes)

		_, err = db.Conn.Exec(context.Background(),
			"INSERT INTO cart (client_id, food, count_food, restaurant_id, number_item) VALUES ($1, $2, $3, $4, $5)",
			clientId, dish.Id, dish.Count, newCart.Restaurant.Id, newCart.Dishes[i].ItemNumber)
		if err != nil {
			return nil, nil, &errorsConst.Errors{
				Text: errorsConst.UpdateCartCartNotInsert,
				Time: time.Now(),
			}
		}

		for j, ingredient := range dish.Ingredients {
			var ingredients Utils.IngredientCartResponse
			err := db.Conn.QueryRow(context.Background(),
				"SELECT id, name, cost FROM structure_dishes WHERE id = $1 ", ingredient.Id).Scan(
					&ingredients.Id, &ingredients.Name, &ingredients.Cost)
			if err != nil {
				dishesError.ItemNumber = dish.ItemNumber
				dishesError.Explain = cart.Dishes[i].IngredientCart[j].Name
				dishesErrors = append(dishesErrors, dishesError)
				continue
			}
			cart.Dishes[i].IngredientCart = append(cart.Dishes[i].IngredientCart, ingredients)

			_, err = db.Conn.Exec(context.Background(),
				"INSERT INTO cart_structure_food (checkbox, client_id) VALUES ($1, $2)",
				ingredient.Id, clientId)
			if err != nil {
				return nil, nil, &errorsConst.Errors{
					Text: errorsConst.UpdateCartStructureFoodNotInsert,
					Time: time.Now(),
				}
			}
		}

		for j, radio := range dish.Radios {
			var radios Utils.RadiosCartResponse
			err := db.Conn.QueryRow(context.Background(),
				"SELECT id, name FROM structure_radios WHERE id = $1", radio.Id).Scan(
					&radios.Id, &radios.Name)
			if err != nil {
				dishesError.ItemNumber = dish.ItemNumber
				dishesError.Explain = cart.Dishes[i].RadiosCart[j].Name
				dishesErrors = append(dishesErrors, dishesError)
				continue
			}
			cart.Dishes[i].RadiosCart = append(cart.Dishes[i].RadiosCart, radios)

			_, err = db.Conn.Exec(context.Background(),
				"INSERT INTO cart_radios_food (radios_id, radios, client_id) VALUES ($1, $2, $3)",
				radio.RadiosId, radio.Id, clientId)
			if err != nil {
				return nil, nil, &errorsConst.Errors{
					Text: errorsConst.UpdateCartRadiosNotInsert,
					Time: time.Now(),
				}
			}
		}
	}

	return &cart, dishesErrors, nil
}

func (db *Wrapper) GetPriceDelivery(id int) (int, error) {
	var price int
	err := db.Conn.QueryRow(context.Background(),
		"SELECT price_delivery FROM restaurant WHERE id = $1", id).Scan(&price)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, &errorsConst.Errors{
				Text: errorsConst.GetPriceDeliveryNotFound,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.GetPriceDeliveryNotScan,
			Time: time.Now(),
		}
	}
	return price, nil
}