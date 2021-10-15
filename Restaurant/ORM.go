package Restaurant

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	res "2021_2_GORYACHIE_MEKSIKANSI/Utils/Restaurant"
	"context"
	//"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Wrapper struct {
	Conn Utils.ConnectionInterface
}

func (db *Wrapper) GetRestaurants() ([]res.Restaurant, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrRestaurantsNotSelect,
			Time: time.Now(),
		}
	}

	restaurant := res.Restaurant{}
	var result []res.Restaurant
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrRestaurantScan,
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
