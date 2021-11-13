package Restaurant

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

func (db *Wrapper) GetRestaurants() ([]Utils.Restaurants, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetRestaurantsRestaurantsNotSelect,
			Time: time.Now(),
		}
	}

	restaurant := Utils.Restaurants{}
	var result []Utils.Restaurants
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetRestaurantsRestaurantsNotScan,
				Time: time.Now(),
			}
		}
		result = append(result, restaurant)
	}

	if result == nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetRestaurantsRestaurantsNotFound,
			Time: time.Now(),
		}
	}

	return result, nil
}

func (db *Wrapper) GetGeneralInfoRestaurant(id int) (*Utils.RestaurantId, error) {
	var restaurant Utils.RestaurantId
	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(
		&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
		&restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetGeneralInfoRestaurantNotFound,
			Time: time.Now(),
		}
	}
	return &restaurant, nil
}

func (db *Wrapper) GetTagsRestaurant(id int) ([]Utils.Tag, error) {
	rowCategory, err := db.Conn.Query(context.Background(),
		"SELECT id, category FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetTagsCategoryNotSelect,
			Time: time.Now(),
		}
	}
	var tags []Utils.Tag
	tag := Utils.Tag{}
	for rowCategory.Next() {
		err := rowCategory.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetTagsCategoryRestaurantNotScan,
				Time: time.Now(),
			}
		}
		tags = append(tags, tag)
	}
	if tags == nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetTagsTagsNotFound,
			Time: time.Now(),
		}
	}
	return tags, nil
}

func getDishesRestaurant(db *Wrapper, name string, id int) ([]Utils.DishesMenu, error) {
	var dishes []Utils.DishesMenu
	dish := Utils.DishesMenu{}
	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, cost, kilocalorie FROM dishes WHERE category_restaurant = $1 AND restaurant = $2",
		name, id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetMenuDishesNotSelect,
			Time: time.Now(),
		}
	}

	for rowDishes.Next() {
		err := rowDishes.Scan(&dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Kilocalorie)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetDishesRestaurantDishesNotScan,
				Time: time.Now(),
			}
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (db *Wrapper) GetMenu(id int) ([]Utils.Menu, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetMenuTransactionNotCreate,
			Time: time.Now(),
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var result []Utils.Menu
	rowDishes, err := tx.Query(context.Background(),
		"SELECT DISTINCT category_restaurant FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetMenuDishesNotSelect,
			Time: time.Now(),
		}
	}

	for rowDishes.Next() {
		var menu Utils.Menu
		err := rowDishes.Scan(&menu.Name)
		if err != nil {
			return nil, err
		}

		dishes, err := getDishesRestaurant(db, menu.Name, id)
		if err != nil {
			return nil, err
		}

		menu.DishesMenu = dishes
		result = append(result, menu)
	}

	if result == nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetMenuDishesNotFound,
			Time: time.Now(),
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetMenuDishesNotCommit,
			Time: time.Now(),
		}
	}

	return result, nil
}

func (db *Wrapper) GetStructDishes(dishesId int) ([]Utils.Ingredients, error) {
	var ingredients []Utils.Ingredients
	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT id, name, cost FROM structure_dishes WHERE food = $1", dishesId)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetStructDishesStructDishesNotSelect,
			Time: time.Now(),
		}
	}

	for rowDishes.Next() {
		var ingredient Utils.Ingredients
		err := rowDishes.Scan(&ingredient.Id, &ingredient.Title, &ingredient.Cost)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetStructDishesStructDishesNotScan,
				Time: time.Now(),
			}
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func getStructRadios(db *Wrapper, radId int) ([]Utils.CheckboxesRows, error) {
	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT id, name FROM structure_radios WHERE radios = $1", radId)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetStructRadiosStructRadiosNotSelect,
			Time: time.Now(),
		}
	}

	var rows []Utils.CheckboxesRows
	for rowDishes.Next() {
		var row Utils.CheckboxesRows
		err := rowDishes.Scan(&row.Id, &row.Name)
		if err != nil {
			errorText := err.Error()
			if strings.Contains(errorText, "no rows") {
				return nil, &errorsConst.Errors{
					Text: errorsConst.RGetStructRadiosStructRadiosNotFound,
					Time: time.Now(),
				}
			}
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetStructRadiosStructRadiosNotScan,
				Time: time.Now(),
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (db *Wrapper) GetDishes(restId int, dishesId int) (*Utils.Dishes, error) {
	var dishes Utils.Dishes
	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, cost, kilocalorie, description FROM dishes WHERE id = $1 AND restaurant = $2",
		dishesId, restId).Scan(
		&dishes.Id, &dishes.Img, &dishes.Title, &dishes.Cost, &dishes.Ccal, &dishes.Description)
	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "no rows") {
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetDishesDishesNotFound,
				Time: time.Now(),
			}
		}
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetDishesDishesNotScan,
			Time: time.Now(),
		}
	}
	return &dishes, nil
}

func (db *Wrapper) GetRadios(dishesId int) ([]Utils.Radios, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetRadiosNotCreate,
			Time: time.Now(),
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var radios []Utils.Radios

	rowDishes, err := tx.Query(context.Background(),
		"SELECT id, name FROM radios WHERE food = $1", dishesId)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetRadiosRadiosNotSelect,
			Time: time.Now(),
		}
	}

	for rowDishes.Next() {
		var rad Utils.Radios
		err := rowDishes.Scan(&rad.Id, &rad.Title)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.RGetRadiosRadiosNotScan,
				Time: time.Now(),
			}
		}

		rows, err := getStructRadios(db, rad.Id)
		if err != nil {
			return nil, err
		}

		if rows != nil {
			rad.Rows = rows
			radios = append(radios, rad)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.RGetRadiosNotCommit,
			Time: time.Now(),
		}
	}

	return radios, nil
}
