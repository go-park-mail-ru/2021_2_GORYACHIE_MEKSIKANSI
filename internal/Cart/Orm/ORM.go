package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
	"github.com/jackc/pgx/v4"
)

type Wrapper struct {
	Conn Interface.ConnectionInterface
}

func (db *Wrapper) GetCart(id int) (*Cart.ResponseCartErrors, []Cart.CastDishesErrs, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var result Cart.ResponseCartErrors
	row, err := tx.Query(contextTransaction,
		"SELECT cart_food.id, cart_food.food, cart_food.number_item, d.avatar, d.name, cart_food.count_food, d.cost, d.kilocalorie, d.weight,"+
			" d.description, sr.name, sr.id, sr.radios, sd.name, sd.id, sd.cost, d.restaurant, d.count, sr.kilocalorie, sd.kilocalorie,"+
			" cart_food.place, crf.place, sd.place "+
			"FROM cart_food "+
			"LEFT JOIN dishes d ON d.id = cart_food.food "+
			"LEFT JOIN cart_structure_food csf ON csf.client_id = cart_food.client_id and d.id=csf.food and cart_food.id=csf.cart_id "+
			"LEFT JOIN structure_dishes sd ON sd.id = csf.checkbox and sd.food=cart_food.food "+
			"LEFT JOIN cart_radios_food crf ON crf.client_id = cart_food.client_id and cart_food.id=crf.cart_id "+
			"LEFT JOIN structure_radios sr ON sr.id = crf.radios "+
			"WHERE cart_food.client_id = $1", id)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartNotSelect,
		}
	}

	place := make(map[int]map[int]interface{})
	infoDishes := make(map[int]Cart.DishesCartResponse)
	var restaurant Cart.RestaurantIdCastResponse

	for row.Next() {
		var dish Cart.DishesCartResponse
		var count, cartId int

		var getPlaceDishes, getPlaceRadios, getPlaceIngredient *int32
		var ingredientKilocalorie, radiosKilocalorie *int32
		var radiosId, radiosRadiosId, ingredientId, ingredientCost *int32
		var radiosName, ingredientName *string
		err := row.Scan(&cartId, &dish.Id, &dish.ItemNumber, &dish.Img, &dish.Name, &dish.Count, &dish.Cost, &dish.Kilocalorie,
			&dish.Weight, &dish.Description, &radiosName, &radiosId, &radiosRadiosId, &ingredientName,
			&ingredientId, &ingredientCost, &restaurant.Id, &count, &radiosKilocalorie, &ingredientKilocalorie,
			&getPlaceDishes, &getPlaceRadios, &getPlaceIngredient)

		if err != nil {
			return nil, nil, &errPkg.Errors{
				Alias: errPkg.CGetCartNotScan,
			}
		}

		placeDishes := Util.ConvertInt32ToInt(getPlaceDishes)
		placeRadios := Util.ConvertInt32ToInt(getPlaceRadios)
		placeIngredient := Util.ConvertInt32ToInt(getPlaceIngredient)

		var radios Cart.RadiosCartResponse
		if radiosName != nil {
			radios.Name = *radiosName
			radios.Id = int(*radiosId)
			radios.RadiosId = int(*radiosRadiosId)
			dish.Kilocalorie += int(*radiosKilocalorie)
		}

		var ingredient Cart.IngredientCartResponse
		if ingredientName != nil {
			ingredient.Name = *ingredientName
			ingredient.Id = int(*ingredientId)
			ingredient.Cost = int(*ingredientCost)
			dish.Kilocalorie += int(*ingredientKilocalorie)
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

		temp := place[placeDishes]
		if temp == nil {
			temp = make(map[int]interface{})
		}

		if placeIngredient != -1 {
			temp[placeIngredient] = ingredient
		}
		if placeRadios != -1 {
			temp[placeRadios] = radios
		}

		place[placeDishes] = temp
		infoDishes[placeDishes] = dish
	}

	for i := 0; i < len(place); i++ {
		dish := infoDishes[i]
		for j := 0; j < len(place[i]); j++ {
			switch place[i][j].(type) {
			case Cart.RadiosCartResponse:
				dish.RadiosCart = append(dish.RadiosCart, place[i][j].(Cart.RadiosCartResponse))
			case Cart.IngredientCartResponse:
				dish.IngredientCart = append(dish.IngredientCart, place[i][j].(Cart.IngredientCartResponse))
			}
		}
		result.Dishes = append(result.Dishes, dish)
	}

	if len(place) == 0 {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartCartNotFound,
		}
	}
	result.Restaurant = restaurant

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CGetCartNotCommit,
		}
	}

	return &result, nil, nil
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

func (db *Wrapper) updateCartStructFood(ingredients []Cart.IngredientsCartRequest, clientId int, cartId int, place int, tx Interface.TransactionInterface, contextTransaction context.Context) ([]Cart.IngredientCartResponse, error) {
	var result []Cart.IngredientCartResponse
	for _, ingredient := range ingredients {
		var checkedIngredient Cart.IngredientCartResponse
		var dishId int
		err := tx.QueryRow(contextTransaction,
			"SELECT id, name, cost, food  FROM structure_dishes WHERE id = $1", ingredient.Id).Scan(
			&checkedIngredient.Id, &checkedIngredient.Name, &checkedIngredient.Cost, &dishId)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartStructureFoodStructureFoodNotSelect,
			}
		}
		result = append(result, checkedIngredient)

		_, err = tx.Exec(contextTransaction,
			"INSERT INTO cart_structure_food (checkbox, client_id, food, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
			ingredient.Id, clientId, dishId, cartId, place)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartStructFoodStructureFoodNotInsert,
			}
		}
		place++
	}
	return result, nil
}

func (db *Wrapper) updateCartRadios(radios []Cart.RadiosCartRequest, clientId int, cartId int, tx Interface.TransactionInterface, contextTransaction context.Context) ([]Cart.RadiosCartResponse, int, error) {
	var result []Cart.RadiosCartResponse
	radiosPlace := 0
	for _, radio := range radios {
		var checkedRadios Cart.RadiosCartResponse
		err := tx.QueryRow(contextTransaction,
			"SELECT id, name FROM structure_radios WHERE id = $1", radio.Id).Scan(
			&checkedRadios.Id, &checkedRadios.Name)
		if err != nil {
			return nil, 0, &errPkg.Errors{
				Alias: errPkg.CUpdateCartStructRadiosStructRadiosNotSelect,
			}
		}
		result = append(result, checkedRadios)

		_, err = tx.Exec(contextTransaction,
			"INSERT INTO cart_radios_food (radios_id, radios, client_id, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
			radio.RadiosId, radio.Id, clientId, cartId, radiosPlace)
		if err != nil {
			return nil, 0, &errPkg.Errors{
				Alias: errPkg.CUpdateCartRadiosRadiosNotInsert,
			}
		}
		radiosPlace++
	}
	return result, radiosPlace, nil
}

func (db *Wrapper) UpdateCart(newCart Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, []Cart.CastDishesErrs, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CUpdateCartTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var dishesErrors []Cart.CastDishesErrs
	var cart Cart.ResponseCartErrors

	structureDishesPlace := 0
	for i, dish := range newCart.Dishes {
		var dishes Cart.DishesCartResponse
		count := 0
		err := tx.QueryRow(contextTransaction,
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

		if dish.Count > count && count != Util.UnlimitedCount {
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
		err = tx.QueryRow(contextTransaction,
			"INSERT INTO cart_food (client_id, food, count_food, restaurant_id, number_item, place) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			clientId, dish.Id, dish.Count, newCart.Restaurant.Id, newCart.Dishes[i].ItemNumber, i).Scan(&idCart)
		if err != nil {
			return nil, nil, &errPkg.Errors{
				Alias: errPkg.CUpdateCartCartNotInsert,
			}
		}
		cart.Dishes[i].RadiosCart, structureDishesPlace, err = db.updateCartRadios(dish.Radios, clientId, idCart, tx, contextTransaction)
		if err != nil {
			return nil, nil, err
		}

		cart.Dishes[i].IngredientCart, err = db.updateCartStructFood(dish.Ingredients, clientId, idCart, structureDishesPlace, tx, contextTransaction)
		if err != nil {
			return nil, nil, err
		}

		structureDishesPlace = 0
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Alias: errPkg.CUpdateCartNotCommit,
		}
	}
	return &cart, dishesErrors, nil
}

func (db *Wrapper) GetPriceDelivery(id int) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.CGetPriceDeliveryTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var price int
	err = tx.QueryRow(contextTransaction,
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

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.CGetPriceDeliveryNotCommit,
		}
	}

	return price, nil
}
