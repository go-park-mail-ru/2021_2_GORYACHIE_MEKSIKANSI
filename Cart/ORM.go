package Cart

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) GetStructFood(id int) ([]Utils.IngredientCartResponse, error) {
	var ingredients []Utils.IngredientCartResponse
	rows, err := db.Conn.Query(context.Background(),
		"SELECT checkbox FROM cart_structure_food WHERE client_id = $1", id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.CGetStructFoodRestaurantNotSelect,
			Time: time.Now(),
		}
	}
	for rows.Next() {
		var ingredient Utils.IngredientCartResponse
		err = rows.Scan(&ingredient.Id)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.CGetStructFoodCheckboxNotScan,
				Time: time.Now(),
			}
		}

		err = db.Conn.QueryRow(context.Background(),
			"SELECT name, cost FROM structure_dishes WHERE id = $1", ingredient.Id).Scan(
			&ingredient.Name, &ingredient.Cost)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func (db *Wrapper) GetStructRadios(id int) ([]Utils.RadiosCartResponse, error) {
	var radios []Utils.RadiosCartResponse
	rows, err := db.Conn.Query(context.Background(),
		"SELECT radios_id, radios FROM cart_radios_food WHERE client_id = $1", id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.CGetStructRadiosRadiosNotSelect,
			Time: time.Now(),
		}
	}
	for rows.Next() {
		var radio Utils.RadiosCartResponse
		err = rows.Scan(&radio.RadiosId, &radio.Id)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.CGetStructRadiosRadiosNotScan,
				Time: time.Now(),
			}
		}

		err = db.Conn.QueryRow(context.Background(),
			"SELECT name FROM structure_radios WHERE id = $1", radio.Id).Scan(&radio.Name)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return nil, &errorsConst.Errors{
					Text: errorsConst.CGetStructRadiosStructRadiosNotFound,
					Time: time.Now(),
				}
			}
			return nil, &errorsConst.Errors{
				Text: errorsConst.CGetStructRadiosStructRadiosNotScan,
				Time: time.Now(),
			}
		}
		radios = append(radios, radio)
	}
	return radios, nil
}

func (db *Wrapper) GetCart(id int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error) {
	var cart *Utils.ResponseCartErrors
	cart = &Utils.ResponseCartErrors{}
	var dishes []Utils.DishesCartResponse
	var radios []Utils.RadiosCartResponse
	var ingredients []Utils.IngredientCartResponse
	var dishesError Utils.CastDishesErrs
	var dishesErrors []Utils.CastDishesErrs

	var restaurant Utils.RestaurantIdCastResponse
	rows, err := db.Conn.Query(context.Background(),
		"SELECT food, count_food, number_item, name, cost, description, avatar, restaurant_id, count, weight, kilocalorie FROM cart"+
			" JOIN dishes ON cart.food = dishes.id WHERE client_id = $1", id)

	var dish *Utils.DishesCartResponse
	dish = &Utils.DishesCartResponse{}
	count := 0

	for rows.Next() {
		err = rows.Scan(&dish.Id, &dish.Count, &dish.ItemNumber, &dish.Name, &dish.Cost, &dish.Description, &dish.Img, &restaurant.Id, &count, &dish.Weight, &dish.Kilocalorie)
		if err != nil {
			return nil, nil, &errorsConst.Errors{
				Text: errorsConst.CGetCartDishesNotScan,
				Time: time.Now(),
			}
		}

		if dish.Count > count && count != -1 {
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.NameDish = dish.Name
			dishesError.CountAvail = count
			dishesErrors = append(dishesErrors, dishesError)
		}
		dish.Weight = dish.Weight * dish.Count
		dish.Kilocalorie = dish.Kilocalorie * dish.Count

		ingredients, err = db.GetStructFood(id)
		if err != nil {
			return nil, nil, err
		}
		dish.IngredientCart = ingredients

		radios, err = db.GetStructRadios(id)
		dish.RadiosCart = radios
		if err != nil {
			return nil, nil, err
		}

		dishes = append(dishes, *dish)
	}
	if cart == nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CGetCartDishesNotFound,
			Time: time.Now(),
		}
	}
	cart.Restaurant = restaurant
	cart.Dishes = dishes
	return cart, dishesErrors, nil
}

func (db *Wrapper) DeleteCart(id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM cart WHERE client_id = $1", id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.CDeleteCartCartNotDelete,
			Time: time.Now(),
		}
	}
	_, err = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_structure_food WHERE client_id = $1", id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.CDeleteCartStructureFoodNotDelete,
			Time: time.Now(),
		}
	}
	_, err = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_radios_food WHERE client_id = $1", id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.CDeleteCartRadiosFoodNotDelete,
			Time: time.Now(),
		}
	}
	return nil
}

func (db *Wrapper) UpdateCartStructFood(ingredients []Utils.IngredientsCartRequest, clientId int, tx pgx.Tx) ([]Utils.IngredientCartResponse, error) {
	var result []Utils.IngredientCartResponse
	for _, ingredient := range ingredients {
		var checkedIngredient Utils.IngredientCartResponse
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, name, cost FROM structure_dishes WHERE id = $1", ingredient.Id).Scan(
			&checkedIngredient.Id, &checkedIngredient.Name, &checkedIngredient.Cost)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.CUpdateCartStructureFoodStructureFoodNotSelect,
				Time: time.Now(),
			}
		}
		result = append(result, checkedIngredient)

		_, err = tx.Exec(context.Background(),
			"INSERT INTO cart_structure_food (checkbox, client_id) VALUES ($1, $2)",
			ingredient.Id, clientId)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.CUpdateCartStructFoodStructureFoodNotInsert,
				Time: time.Now(),
			}
		}
	}
	return result, nil
}
func (db *Wrapper) UpdateCartRadios(radios []Utils.RadiosCartRequest, clientId int, tx pgx.Tx) ([]Utils.RadiosCartResponse, error) {
	var result []Utils.RadiosCartResponse
	for _, radio := range radios {
		var checkedRadios Utils.RadiosCartResponse
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, name FROM structure_radios WHERE id = $1", radio.Id).Scan(
			&checkedRadios.Id, &checkedRadios.Name)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.CUpdateCartStructRadiosStructRadiosNotSelect,
				Time: time.Now(),
			}
		}
		result = append(result, checkedRadios)

		_, err = tx.Exec(context.Background(),
			"INSERT INTO cart_radios_food (radios_id, radios, client_id) VALUES ($1, $2, $3)",
			radio.RadiosId, radio.Id, clientId)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.CUpdateCartRadiosRadiosNotInsert,
				Time: time.Now(),
			}
		}
	}
	return result, nil
}

func (db *Wrapper) UpdateCart(newCart Utils.RequestCartDefault, clientId int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error) {
	tx, err := db.Conn.Begin(context.Background())
	defer func(tx pgx.Tx) {
		switch err {
		case nil:

		default:
			err := tx.Rollback(context.Background())
			if err != nil {
				return
			}
		}
	}(tx)
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CUpdateCartTransactionNotCreate,
			Time: time.Now(),
		}
	}

	var dishesErrors []Utils.CastDishesErrs
	var cart Utils.ResponseCartErrors
	for i, dish := range newCart.Dishes {
		var dishes Utils.DishesCartResponse
		var dishesError Utils.CastDishesErrs
		count := 0
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, avatar, cost, name, description, count, weight, kilocalorie FROM dishes WHERE id = $1 AND restaurant = $2",
			dish.Id, newCart.Restaurant.Id).Scan(
			&dishes.Id, &dishes.Img, &dishes.Cost, &dishes.Name, &dishes.Description, &count, &dishes.Weight, &dishes.Kilocalorie)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return nil, nil, &errorsConst.Errors{
					Text: errorsConst.CUpdateCartCartNotFound,
					Time: time.Now(),
				}
			}
			return nil, nil, &errorsConst.Errors{
				Text: errorsConst.CUpdateCartCartNotScan,
				Time: time.Now(),
			}
		}

		dishes.Count = dish.Count

		if dish.Count > count && count != -1 {
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.NameDish = dishes.Name
			dishesError.CountAvail = count
			dishesErrors = append(dishesErrors, dishesError)

			dishes.Count = count
			dish.Count = count
		}
		dishes.Weight = dishes.Weight * dishes.Count
		dishes.Kilocalorie = dishes.Kilocalorie * dishes.Count

		cart.Dishes = append(cart.Dishes, dishes)

		_, err = tx.Exec(context.Background(),
			"INSERT INTO cart (client_id, food, count_food, restaurant_id, number_item) VALUES ($1, $2, $3, $4, $5)",
			clientId, dish.Id, dish.Count, newCart.Restaurant.Id, newCart.Dishes[i].ItemNumber)
		if err != nil {
			return nil, nil, &errorsConst.Errors{
				Text: errorsConst.CUpdateCartCartNotInsert,
				Time: time.Now(),
			}
		}
		cart.Dishes[i].IngredientCart, err = db.UpdateCartStructFood(dish.Ingredients, clientId, tx)
		if err != nil {
			return nil, nil, err
		}

		cart.Dishes[i].RadiosCart, err = db.UpdateCartRadios(dish.Radios, clientId, tx)
		if err != nil {
			return nil, nil, err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CUpdateCartNotCommit,
			Time: time.Now(),
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
