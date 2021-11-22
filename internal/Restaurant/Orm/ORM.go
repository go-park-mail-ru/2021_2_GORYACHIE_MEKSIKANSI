package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"context"
	"github.com/jackc/pgx/v4"
)

type Wrapper struct {
	Conn Interface.ConnectionInterface
}

func (db *Wrapper) GetRestaurants() ([]Restaurant.Restaurants, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	row, err := tx.Query(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsRestaurantsNotSelect,
		}
	}

	restaurant := Restaurant.Restaurants{}
	var result []Restaurant.Restaurants
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetRestaurantsRestaurantsNotScan,
			}
		}
		result = append(result, restaurant)
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsRestaurantsNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetGeneralInfoRestaurant(id int) (*Restaurant.RestaurantId, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var restaurant Restaurant.RestaurantId
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(
		&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
		&restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoRestaurantNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoNotCommit,
		}
	}

	return &restaurant, nil
}

func (db *Wrapper) GetTagsRestaurant(id int) ([]Restaurant.Tag, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsRestaurantTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	rowCategory, err := tx.Query(contextTransaction,
		"SELECT id, category, place FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsRestaurantNotSelect,
		}
	}

	var tags []Restaurant.Tag
	var tag Restaurant.Tag

	place := make(map[int]Restaurant.Tag)
	for rowCategory.Next() {
		var placeCategory int
		err := rowCategory.Scan(&tag.Id, &tag.Name, &placeCategory)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetTagsRestaurantRestaurantNotScan,
			}
		}
		place[placeCategory] = tag
	}

	for i := 0; i < len(place); i++ {
		tags = append(tags, place[i])
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsRestaurantNotCommit,
		}
	}

	return tags, nil
}

func (db *Wrapper) GetMenu(id int) ([]Restaurant.Menu, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var result []Restaurant.Menu

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT category_restaurant, id, avatar, name, cost, kilocalorie, place, place_category FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesCategoryNotSelect,
		}
	}

	categoryPlace := make(map[int]Restaurant.Menu)
	place := make(map[int]map[int]Restaurant.DishesMenu)

	for rowDishes.Next() {
		var menu Restaurant.Menu
		var dish Restaurant.DishesMenu
		var placeDish, placeCategory int
		err := rowDishes.Scan(&menu.Name, &dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Kilocalorie, &placeDish, &placeCategory)
		if err != nil {
			return nil, err
		}

		temp := place[placeCategory]
		if temp == nil {
			temp = make(map[int]Restaurant.DishesMenu)
		}
		temp[placeDish] = dish
		place[placeCategory] = temp
		categoryPlace[placeCategory] = menu
	}

	for i := 0; i < len(place); i++ {
		result = append(result, categoryPlace[i])
		for j := 0; j < len(place[i]); j++ {
			result[i].DishesMenu = append(result[i].DishesMenu, place[i][j])
		}
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetStructDishes(dishesId int) ([]Restaurant.Ingredients, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var ingredients []Restaurant.Ingredients
	rowDishes, err := tx.Query(contextTransaction,
		"SELECT id, name, cost, place FROM structure_dishes WHERE food = $1", dishesId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesStructDishesNotSelect,
		}
	}

	place := make(map[int]Restaurant.Ingredients)
	for rowDishes.Next() {
		var placeDish int
		var ingredient Restaurant.Ingredients
		err := rowDishes.Scan(&ingredient.Id, &ingredient.Title, &ingredient.Cost, &placeDish)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetStructDishesStructDishesNotScan,
			}
		}
		place[placeDish] = ingredient
	}

	for i := 0; i < len(place); i++ {
		ingredients = append(ingredients, place[i])
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesNotCommit,
		}
	}

	return ingredients, nil
}

func (db *Wrapper) GetDishes(restId int, dishesId int) (*Restaurant.Dishes, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetDishesTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var dishes Restaurant.Dishes
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, cost, kilocalorie, description FROM dishes WHERE id = $1 AND restaurant = $2",
		dishesId, restId).Scan(
		&dishes.Id, &dishes.Img, &dishes.Title, &dishes.Cost, &dishes.Ccal, &dishes.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetDishesDishesNotFound,
			}
		}
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetDishesDishesNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetDishesNotCommit,
		}
	}

	return &dishes, nil
}

func (db *Wrapper) GetRadios(dishesId int) ([]Restaurant.Radios, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT r.id, r.name, sr.id, sr.name, r.place, sr.place FROM radios r "+
			"LEFT JOIN structure_radios sr ON sr.radios=r.id WHERE r.food = $1", dishesId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosRadiosNotSelect,
		}
	}

	var radios []Restaurant.Radios
	place := make(map[int]map[int]Restaurant.CheckboxesRows)
	radiosInfo := make(map[int]Restaurant.Radios)

	for rowDishes.Next() {
		var rad Restaurant.Radios
		var elementRadios Restaurant.CheckboxesRows
		var placeRadios, placeElementRadios int
		err := rowDishes.Scan(&rad.Id, &rad.Title, &elementRadios.Id, &elementRadios.Name, &placeRadios, &placeElementRadios)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetRadiosRadiosNotScan,
			}
		}

		temp := place[placeRadios]
		if temp == nil {
			temp = make(map[int]Restaurant.CheckboxesRows)
		}
		temp[placeElementRadios] = elementRadios
		place[placeRadios] = temp
		radiosInfo[placeRadios] = rad
	}
	for i := 0; i < len(place); i++ {
		radios = append(radios, radiosInfo[i])
		var rows []Restaurant.CheckboxesRows
		for j := 0; j < len(place[i]); j++ {
			rows = append(rows, place[i][j])
		}
		radios[i].Rows = rows
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosNotCommit,
		}
	}

	return radios, nil
}
