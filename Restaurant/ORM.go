package Restaurant

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Wrapper struct {
	Conn *pgxpool.Pool
}

func (db *Wrapper) GetRestaurants() ([]Restaurant, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant LIMIT 50")
	if err != nil {
		return nil, err
	}

	restaurant := Restaurant{}
	var result []Restaurant
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			panic(err)
		}
		result = append(result, restaurant)
	}

	return result, nil
}
