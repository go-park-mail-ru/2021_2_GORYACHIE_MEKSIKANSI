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

func (db *Wrapper) GetRestaurants() ([]Utils.Restaurant, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotSelect,
			Time: time.Now(),
		}
	}

	restaurant := Utils.Restaurant{}
	var result []Utils.Restaurant
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

func (db *Wrapper) GetRestaurant(id int) (*Utils.RestaurantAndCategory, []Utils.Dishes, error) {
	var restaurant Utils.RestaurantAndCategory
	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantNotFound,
			Time: time.Now(),
		}
	}

	rowCategory, err := db.Conn.Query(context.Background(),
		"SELECT category FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotSelect,
			Time: time.Now(),
		}
	}
	var tags []string
	var tag string
	for rowCategory.Next() {
		err := rowCategory.Scan(&tag)
		if err != nil {
			return nil, nil, &errorsConst.Errors{
				Text: errorsConst.ErrCategoryRestaurantScan,
				Time: time.Now(),
			}
		}
		tags = append(tags, tag)
	}

	if tags == nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotFound,
			Time: time.Now(),
		}
	}

	// TODO: remake struct
	s := make([]interface{}, len(tags))
	for i, v := range tags {
		s[i] = v
	}
	restaurant.Tags = s

	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, cost, ccal, category_restaurant FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsDishesNotSelect,
			Time: time.Now(),
		}
	}

	dishes := Utils.Dishes{}
	var result []Utils.Dishes
	for rowDishes.Next() {
		err := rowDishes.Scan(&dishes.Id, &dishes.Img, &dishes.Cost, &dishes.Ccal)
		if err != nil {
			return nil, nil, &errorsConst.Errors{
				Text: errorsConst.ErrRestaurantDishesScan,
				Time: time.Now(),
			}
		}
		result = append(result, dishes)
	}

	if result == nil {
		return nil, nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantDishesNotFound,
			Time: time.Now(),
		}
	}

	return &restaurant, result, nil
}
