package Restaurant

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Wrapper struct {
	Conn *pgxpool.Pool
}

func (db *Wrapper) GetRestaurants() ([]Restaurant, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, errors.New(errorsConst.ErrRestaurantsNotFound)
	}

	restaurant := Restaurant{}
	var result []Restaurant
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			return nil, errors.New(errorsConst.ErrRestaurantScan)
		}
		result = append(result, restaurant)
	}

	return result, nil
}
