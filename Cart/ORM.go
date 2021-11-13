package Cart

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"strings"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) getStructFood(id int) ([]Utils.IngredientCartResponse, error) {
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

func (db *Wrapper) getStructRadios(id int) ([]Utils.RadiosCartResponse, error) {
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
			errorText := err.Error()
			if strings.Contains(errorText, "no rows") {
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

//func (db *Wrapper) GetCart(id int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error) {
//	var cart *Utils.ResponseCartErrors
//	cart = &Utils.ResponseCartErrors{}
//	var dishes []Utils.DishesCartResponse
//	var radios []Utils.RadiosCartResponse
//	var ingredients []Utils.IngredientCartResponse
//	var dishesError Utils.CastDishesErrs
//	var dishesErrors []Utils.CastDishesErrs
//
//	var restaurant Utils.RestaurantIdCastResponse
//	rows, err := db.Conn.Query(context.Background(),
//		"SELECT food, count_food, number_item, name, cost, description, avatar, restaurant_id, count, weight, kilocalorie FROM cart"+
//			" JOIN dishes ON cart.food = dishes.id WHERE client_id = $1", id)
//
//	var dish *Utils.DishesCartResponse
//	dish = &Utils.DishesCartResponse{}
//	count := 0
//
//	for rows.Next() {
//		err = rows.Scan(&dish.Id, &dish.Count, &dish.ItemNumber, &dish.Name, &dish.Cost, &dish.Description, &dish.Img, &restaurant.Id, &count, &dish.Weight, &dish.Kilocalorie)
//		if err != nil {
//			return nil, nil, &errorsConst.Errors{
//				Text: errorsConst.CGetCartDishesNotScan,
//				Time: time.Now(),
//			}
//		}
//
//		if dish.Count > count && count != -1 {
//			dishesError.ItemNumber = dish.ItemNumber
//			dishesError.NameDish = dish.Name
//			dishesError.CountAvail = count
//			dishesErrors = append(dishesErrors, dishesError)
//		}
//		dish.Weight = dish.Weight * dish.Count
//		dish.Kilocalorie = dish.Kilocalorie * dish.Count
//
//		ingredients, err = db.getStructFood(id)
//		if err != nil {
//			return nil, nil, err
//		}
//		dish.IngredientCart = ingredients
//
//		radios, err = db.getStructRadios(id)
//		dish.RadiosCart = radios
//		if err != nil {
//			return nil, nil, err
//		}
//
//		dishes = append(dishes, *dish)
//	}
//	if cart == nil {
//		return nil, nil, &errorsConst.Errors{
//			Text: errorsConst.CGetCartDishesNotFound,
//			Time: time.Now(),
//		}
//	}
//	cart.Restaurant = restaurant
//	cart.Dishes = dishes
//	return cart, dishesErrors, nil
//}


func (db *Wrapper) GetCart(id int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CGetCartTransactionNotCreate,
			Time: time.Now(),
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var result Utils.ResponseCartErrors
	row, err := tx.Query(context.Background(),
		"SELECT cart.food, cart.number_item, d.avatar, d.name, cart.count_food, d.cost, d.kilocalorie, d.weight," +
		" d.description, sr.name, sr.id, sr.radios, sd.name, sd.id, sd.cost, d.restaurant, d.count " +
		"FROM cart " +
		"LEFT JOIN dishes d ON d.id = cart.food " +
		"LEFT JOIN cart_structure_food csf ON csf.client_id = cart.client_id " +
		"LEFT JOIN structure_dishes sd ON sd.id = csf.checkbox " +
		"LEFT JOIN cart_radios_food crf ON crf.client_id = cart.client_id " +
		"LEFT JOIN structure_radios sr ON sr.id = crf.radios " +
		"WHERE cart.client_id = $1", id)
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CGetCartNotSelect,
			Time: time.Now(),
		}
	}

	m := make(map[int]Utils.DishesCartResponse)
	var restaurant Utils.RestaurantIdCastResponse

	for row.Next() {
		var radios Utils.RadiosCartResponse
		var ingredient Utils.IngredientCartResponse
		var dish Utils.DishesCartResponse
		var dishesError Utils.CastDishesErrs
		var dishesErrors []Utils.CastDishesErrs
		var count int

		var radiosName, radiosId, radiosRadiosId interface{}
		var ingredientName, ingredientId, ingredientCost interface{}
		err := row.Scan(&dish.Id, &dish.ItemNumber, &dish.Img, &dish.Name, &dish.Count, &dish.Cost, &dish.Kilocalorie,
			&dish.Weight, &dish.Description, &radiosName, &radiosId, &radiosRadiosId, &ingredientName,
			&ingredientId, &ingredientCost, &restaurant.Id, &count)

		if err != nil {
			return nil, nil, err
		}

		if radiosName != nil {
			radios.Name = radiosName.(string)
			radios.Id = int(radiosId.(int32))
			radios.RadiosId = int(radiosRadiosId.(int32))
		}

		if ingredientName != nil {
			ingredient.Name = ingredientName.(string)
			ingredient.Id = int(ingredientId.(int32))
			ingredient.Cost = int(ingredientCost.(int32))
		}

		if dish.Count > count && count != -1 {
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.NameDish = dish.Name
			dishesError.CountAvail = count
			dishesErrors = append(dishesErrors, dishesError)
		}
		dish.Weight = dish.Weight * dish.Count
		dish.Kilocalorie = dish.Kilocalorie * dish.Count

		if val, ok := m[dish.Id]; ok {
			if ingredient.Id != 0 {
				val.IngredientCart = append(m[dish.Id].IngredientCart, ingredient)
				m[dish.Id] = val
			}

			if radios.Id != 0 {
				val.RadiosCart = append(m[dish.Id].RadiosCart, radios)
				m[dish.Id] = val
			}
		} else {
			if radiosName != nil {
				dish.RadiosCart = append(dish.RadiosCart, radios)
			}

			if ingredientName != nil {
				dish.IngredientCart = append(dish.IngredientCart, ingredient)
			}
			m[dish.Id] = dish
		}
	}

	if m == nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CGetCartDishesNotFound,
			Time: time.Now(),
		}
	}
	result.Restaurant = restaurant

	for _, dish := range m {
		result.Dishes = append(result.Dishes, dish)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CGetCartNotCommit,
			Time: time.Now(),
		}
	}

	return &result, nil, nil
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

func (db *Wrapper) updateCartStructFood(ingredients []Utils.IngredientsCartRequest, clientId int, tx pgx.Tx) ([]Utils.IngredientCartResponse, error) {
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

func (db *Wrapper) updateCartRadios(radios []Utils.RadiosCartRequest, clientId int, tx pgx.Tx) ([]Utils.RadiosCartResponse, error) {
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
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.CUpdateCartTransactionNotCreate,
			Time: time.Now(),
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

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
			errorText := err.Error()
			if strings.Contains(errorText, "no rows") {
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
		cart.Dishes[i].IngredientCart, err = db.updateCartStructFood(dish.Ingredients, clientId, tx)
		if err != nil {
			return nil, nil, err
		}

		cart.Dishes[i].RadiosCart, err = db.updateCartRadios(dish.Radios, clientId, tx)
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
