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
		"SELECT id, avatar, name, cost, kilocalorie, category_restaurant FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsDishesNotSelect,
			Time: time.Now(),
		}
	}

	dishes := Utils.DishesMenu{}
	var result []Utils.Menu
	for rowDishes.Next() {
		var menu Utils.Menu
		err := rowDishes.Scan(&dishes.Id, &dishes.Img, &dishes.Name, &dishes.Cost, &dishes.Kilocalorie, &menu.Name)
		if err != nil {
			return nil, nil, nil, &errorsConst.Errors{
				Text: errorsConst.ErrRestaurantDishesScan,
				Time: time.Now(),
			}
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
	var checkboxes []Utils.Ingredients
	_ = db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, cost, kilocalorie, description FROM dishes WHERE id = $1 AND restaurant = $2",
		dishesId, restId).Scan(
		&dishes.Id, &dishes.Img, &dishes.Title, &dishes.Cost, &dishes.Ccal, &dishes.Description)

	rowDishes, _ := db.Conn.Query(context.Background(),
		"SELECT structure_dishes.id, structure_dishes.name, structure_dishes.cost, radios.id, radios.name FROM structure_dishes JOIN radios ON structure_dishes.food = radios.food WHERE structure_dishes.food = $1", dishesId)
	for rowDishes.Next() {
		var checkbox Utils.Ingredients
		var rad Utils.Radios

		_ = rowDishes.Scan(&checkbox.Id, &checkbox.Title, &checkbox.Cost, &rad.Id, &rad.Title)

		checkboxes = append(checkboxes, checkbox)
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

	return &dishes, radios, checkboxes, nil
}
