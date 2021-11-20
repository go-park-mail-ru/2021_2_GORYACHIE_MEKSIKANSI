package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyErrors"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interfaces"
	"context"
	"github.com/jackc/pgx/v4"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) getStructFood(id int) ([]Cart.IngredientCartResponse, error) {
	var ingredients []Cart.IngredientCartResponse
	rows, err := db.Conn.Query(context.Background(),
		"SELECT checkbox FROM cart_structure_food WHERE client_id = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.CGetStructFoodRestaurantNotSelect,
		}
	}
	for rows.Next() {
		var ingredient Cart.IngredientCartResponse
		err = rows.Scan(&ingredient.Id)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CGetStructFoodCheckboxNotScan,
			}
		}

		err = db.Conn.QueryRow(context.Background(),
			"SELECT name, cost FROM structure_dishes WHERE id = $1", ingredient.Id).Scan(
			&ingredient.Name, &ingredient.Cost)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func (db *Wrapper) getStructRadios(id int) ([]Cart.RadiosCartResponse, error) {
	var radios []Cart.RadiosCartResponse
	rows, err := db.Conn.Query(context.Background(),
		"SELECT radios_id, radios FROM cart_radios_food WHERE client_id = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.CGetStructRadiosRadiosNotSelect,
		}
	}
	for rows.Next() {
		var radio Cart.RadiosCartResponse
		err = rows.Scan(&radio.RadiosId, &radio.Id)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CGetStructRadiosRadiosNotScan,
			}
		}

		err = db.Conn.QueryRow(context.Background(),
			"SELECT name FROM structure_radios WHERE id = $1", radio.Id).Scan(&radio.Name)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, &errPkg.Errors{
					Alias: errPkg.CGetStructRadiosStructRadiosNotFound,
				}
			}
			return nil, &errPkg.Errors{
				Alias: errPkg.CGetStructRadiosStructRadiosNotScan,
			}
		}
		radios = append(radios, radio)
	}
	return radios, nil
}

func (db *Wrapper) GetCart(id int) (*Cart.ResponseCartErrors, []Cart.CastDishesErrs, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var result Cart.ResponseCartErrors
	row, err := tx.Query(context.Background(),
		"SELECT cart.id, cart.food, cart.number_item, d.avatar, d.name, cart.count_food, d.cost, d.kilocalorie, d.weight,"+
			" d.description, sr.name, sr.id, sr.radios, sd.name, sd.id, sd.cost, d.restaurant, d.count, sr.kilocalorie, sd.kilocalorie "+
			"FROM cart "+
			"LEFT JOIN dishes d ON d.id = cart.food "+
			"LEFT JOIN cart_structure_food csf ON csf.client_id = cart.client_id and d.id=csf.food and cart.id=csf.cart_id "+
			"LEFT JOIN structure_dishes sd ON sd.id = csf.checkbox and sd.food=cart.food "+
			"LEFT JOIN cart_radios_food crf ON crf.client_id = cart.client_id and cart.id=crf.cart_id "+
			"LEFT JOIN structure_radios sr ON sr.id = crf.radios "+
			"WHERE cart.client_id = $1", id)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartNotSelect,
		}
	}

	m := make(map[int]Cart.DishesCartResponse)
	var restaurant Cart.RestaurantIdCastResponse

	for row.Next() {
		var radios Cart.RadiosCartResponse
		var ingredient Cart.IngredientCartResponse
		var dish Cart.DishesCartResponse
		var count, cartId int

		var ingredientKilocalorie, radiosKilocalorie interface{}
		var radiosName, radiosId, radiosRadiosId interface{}
		var ingredientName, ingredientId, ingredientCost interface{}
		err := row.Scan(&cartId, &dish.Id, &dish.ItemNumber, &dish.Img, &dish.Name, &dish.Count, &dish.Cost, &dish.Kilocalorie,
			&dish.Weight, &dish.Description, &radiosName, &radiosId, &radiosRadiosId, &ingredientName,
			&ingredientId, &ingredientCost, &restaurant.Id, &count, &radiosKilocalorie, &ingredientKilocalorie)

		if err != nil {
			return nil, nil, &errPkg.Errors{
				Alias: errPkg.CGetCartNotScan,
			}
		}

		if radiosName != nil {
			radios.Name = radiosName.(string)
			radios.Id = int(radiosId.(int32))
			radios.RadiosId = int(radiosRadiosId.(int32))
			dish.Kilocalorie += int(radiosKilocalorie.(int32))
		}

		if ingredientName != nil {
			ingredient.Name = ingredientName.(string)
			ingredient.Id = int(ingredientId.(int32))
			ingredient.Cost = int(ingredientCost.(int32))
			dish.Kilocalorie += int(ingredientKilocalorie.(int32))
		}

		if dish.Count > count && count != -1 {
			var dishesErrors []Cart.CastDishesErrs
			var dishesError Cart.CastDishesErrs
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.NameDish = dish.Name
			dishesError.CountAvail = count
			dishesErrors = append(dishesErrors, dishesError)
		}
		dish.Weight = dish.Weight * dish.Count
		dish.Kilocalorie = dish.Kilocalorie * dish.Count

		if val, ok := m[cartId]; ok {
			if ingredient.Id != 0 {
				val.IngredientCart = append(m[cartId].IngredientCart, ingredient)
				m[cartId] = val
			}

			if radios.Id != 0 {
				val.RadiosCart = append(m[cartId].RadiosCart, radios)
				m[cartId] = val
			}
		} else {
			if radiosName != nil {
				dish.RadiosCart = append(dish.RadiosCart, radios)
			}

			if ingredientName != nil {
				dish.IngredientCart = append(dish.IngredientCart, ingredient)
			}
			m[cartId] = dish
		}
	}

	if m == nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartDishesNotFound,
		}
	}
	result.Restaurant = restaurant

	for _, dish := range m {
		result.Dishes = append(result.Dishes, dish)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartNotCommit,
		}
	}

	return &result, nil, nil
}

func (db *Wrapper) DeleteCart(id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM cart WHERE client_id = $1", id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.CDeleteCartCartNotDelete,
		}
	}
	_, err = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_structure_food WHERE client_id = $1", id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.CDeleteCartStructureFoodNotDelete,
		}
	}
	_, err = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_radios_food WHERE client_id = $1", id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.CDeleteCartRadiosFoodNotDelete,
		}
	}
	return nil
}

func (db *Wrapper) updateCartStructFood(ingredients []Cart.IngredientsCartRequest, clientId int, tx pgx.Tx, cartId int) ([]Cart.IngredientCartResponse, error) {
	var result []Cart.IngredientCartResponse
	for _, ingredient := range ingredients {
		var checkedIngredient Cart.IngredientCartResponse
		var dishId int
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, name, cost, food  FROM structure_dishes WHERE id = $1", ingredient.Id).Scan(
			&checkedIngredient.Id, &checkedIngredient.Name, &checkedIngredient.Cost, &dishId)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartStructureFoodStructureFoodNotSelect,
			}
		}
		result = append(result, checkedIngredient)

		_, err = tx.Exec(context.Background(),
			"INSERT INTO cart_structure_food (checkbox, client_id, food, cart_id) VALUES ($1, $2, $3, $4)",
			ingredient.Id, clientId, dishId, cartId)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartStructFoodStructureFoodNotInsert,
			}
		}
	}
	return result, nil
}

func (db *Wrapper) updateCartRadios(radios []Cart.RadiosCartRequest, clientId int, tx pgx.Tx, cartId int) ([]Cart.RadiosCartResponse, error) {
	var result []Cart.RadiosCartResponse
	for _, radio := range radios {
		var checkedRadios Cart.RadiosCartResponse
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, name FROM structure_radios WHERE id = $1", radio.Id).Scan(
			&checkedRadios.Id, &checkedRadios.Name)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartStructRadiosStructRadiosNotSelect,
			}
		}
		result = append(result, checkedRadios)

		_, err = tx.Exec(context.Background(),
			"INSERT INTO cart_radios_food (radios_id, radios, client_id, cart_id) VALUES ($1, $2, $3, $4)",
			radio.RadiosId, radio.Id, clientId, cartId)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartRadiosRadiosNotInsert,
			}
		}
	}
	return result, nil
}

func (db *Wrapper) UpdateCart(newCart Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, []Cart.CastDishesErrs, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CUpdateCartTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var dishesErrors []Cart.CastDishesErrs
	var cart Cart.ResponseCartErrors
	for i, dish := range newCart.Dishes {
		var dishes Cart.DishesCartResponse
		count := 0
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id, avatar, cost, name, description, count, weight, kilocalorie FROM dishes WHERE id = $1 AND restaurant = $2",
			dish.Id, newCart.Restaurant.Id).Scan(
			&dishes.Id, &dishes.Img, &dishes.Cost, &dishes.Name, &dishes.Description, &count, &dishes.Weight, &dishes.Kilocalorie)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, &errPkg.Errors{
					Alias: errPkg.CUpdateCartCartNotFound,
				}
			}
			return nil, nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartCartNotScan,
			}
		}

		dishes.Count = dish.Count

		if dish.Count > count && count != -1 {
			var dishesError Cart.CastDishesErrs
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

		var idCart int
		err = tx.QueryRow(context.Background(),
			"INSERT INTO cart (client_id, food, count_food, restaurant_id, number_item) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			clientId, dish.Id, dish.Count, newCart.Restaurant.Id, newCart.Dishes[i].ItemNumber).Scan(&idCart)
		if err != nil {
			return nil, nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartCartNotInsert,
			}
		}
		cart.Dishes[i].IngredientCart, err = db.updateCartStructFood(dish.Ingredients, clientId, tx, idCart)
		if err != nil {
			return nil, nil, err
		}

		cart.Dishes[i].RadiosCart, err = db.updateCartRadios(dish.Radios, clientId, tx, idCart)
		if err != nil {
			return nil, nil, err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CUpdateCartNotCommit,
		}
	}
	return &cart, dishesErrors, nil
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
