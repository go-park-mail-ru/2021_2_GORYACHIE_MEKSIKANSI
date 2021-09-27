package Restaurant

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ERRQUERY = "Error query"
	ERRSCAN = "Error scan"
)

type Wrapper struct {
	Conn *pgxpool.Pool
}

func (db *Wrapper) GetRestaurants() ([]Restaurant, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
	if err != nil {
		return nil, errors.New(ERRQUERY)
	}

	restaurant := Restaurant{}
	var result []Restaurant
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			return nil, errors.New(ERRSCAN)
		}
		result = append(result, restaurant)
	}

	return result, nil
}
