package Restaurant

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"time"
)

type Wrapper struct {
	Conn Utils.ConnectionInterface
}

func (db *Wrapper) GetRestaurants() ([]Utils.Restaurants, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotSelect,
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
				Text: errorsConst.ErrRestaurantsScan,
				Time: time.Now(),
			}
		}
		result = append(result, restaurant)
	}

	if result == nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotFound,
			Time: time.Now(),
		}
	}

	return result, nil
}

func (db *Wrapper) GetRestaurant(id int) (*Utils.RestaurantId, []Utils.Tag, []Utils.Menu, error) {
	var restaurant Utils.RestaurantId
	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
		&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantNotFound,
			Time: time.Now(),
		}
	}

	rowCategory, err := db.Conn.Query(context.Background(),
		"SELECT id, category FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotSelect,
			Time: time.Now(),
		}
	}

	var tags []Utils.Tag
	tag := Utils.Tag{}
	for rowCategory.Next() {
		err := rowCategory.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, nil, nil, &errorsConst.Errors{
				Text: errorsConst.ErrCategoryRestaurantScan,
				Time: time.Now(),
			}
		}
		tags = append(tags, tag)
	}

	if tags == nil {
		return nil, nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotFound,
			Time: time.Now(),
		}
	}

	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT DISTINCT category_restaurant FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsDishesNotSelect,// TODO: new error
			Time: time.Now(),
		}
	}

	var result []Utils.Menu
	for rowDishes.Next() {
		var menu Utils.Menu
		var dishes []Utils.DishesMenu
		dish := Utils.DishesMenu{}
		err := rowDishes.Scan(&menu.Name)

		rowDishes, err := db.Conn.Query(context.Background(),
			"SELECT id, avatar, name, cost, kilocalorie FROM dishes WHERE category_restaurant = $1", menu.Name)
		if err != nil {
			return nil, nil, nil, &errorsConst.Errors{
				Text: errorsConst.ErrRestaurantsDishesNotSelect,
				Time: time.Now(),
			}
		}

		for rowDishes.Next() {
			err := rowDishes.Scan(&dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Kilocalorie)
			if err != nil {
				return nil, nil, nil, &errorsConst.Errors{
					Text: errorsConst.ErrRestaurantDishesScan,
					Time: time.Now(),
				}
			}
			dishes = append(dishes, dish)
		}

		menu.DishesMenu = dishes
		result = append(result, menu)
	}

	if result == nil {
		return nil, nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantDishesNotFound,
			Time: time.Now(),
		}
	}

	return &restaurant, tags, result, nil
}

func (db *Wrapper) RestaurantDishes(restId int, dishesId int) (*Utils.Dishes, []Utils.Radios, []Utils.Ingredients, error) {
	var dishes Utils.Dishes
	var radios []Utils.Radios
	var ingredients []Utils.Ingredients
	_ = db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, cost, kilocalorie, description FROM dishes WHERE id = $1 AND restaurant = $2",
		dishesId, restId).Scan(
		&dishes.Id, &dishes.Img, &dishes.Title, &dishes.Cost, &dishes.Ccal, &dishes.Description)

	rowDishes, _ := db.Conn.Query(context.Background(),
		"SELECT id, name, cost FROM structure_dishes WHERE food = $1", dishesId)
	for rowDishes.Next() {
		var ingredient Utils.Ingredients
		_ = rowDishes.Scan(&ingredient.Id, &ingredient.Title, &ingredient.Cost)
		ingredients = append(ingredients, ingredient)
	}
	dishes.Ingredient = ingredients

	rowDishes, _ = db.Conn.Query(context.Background(),
		"SELECT id, name FROM radios WHERE food = $1", dishesId)
	for rowDishes.Next() {
		var rad Utils.Radios
		_ = rowDishes.Scan(&rad.Id, &rad.Title)
		radios = append(radios, rad)
	}

	for i, rad := range radios {
		rowDishes, _ := db.Conn.Query(context.Background(),
			"SELECT id, name FROM structure_radios WHERE radios = $1", rad.Id)

		var rows []Utils.CheckboxesRows
		for rowDishes.Next() {
			var row Utils.CheckboxesRows
			_ = rowDishes.Scan(&row.Id, &row.Name)
			rows = append(rows, row)
		}
		radios[i].Rows = rows
	}

	return &dishes, radios, ingredients, nil
}
