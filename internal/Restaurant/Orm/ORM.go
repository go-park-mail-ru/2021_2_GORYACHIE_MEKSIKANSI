package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectRestaurantService
	Ctx context.Context
}

func (r *Wrapper) AllRestaurants() ([]Restaurant.Restaurants, error) {
	// TODO: add convert func
	return nil, nil
}

func (r *Wrapper) GetRestaurant(id int) (*Restaurant.RestaurantId, error) {
	// TODO: add convert func
	return nil, nil
}

func (r *Wrapper) RestaurantDishes(restId int, dishId int) (*Restaurant.Dishes, error) {
	// TODO: add convert func
	return nil, nil
}

func (r *Wrapper) CreateReview(id int, review Restaurant.NewReview) error {
	// TODO: add convert func
	return nil
}

func (r *Wrapper) GetReview(id int) (*Restaurant.ResReview, error) {
	// TODO: add convert func
	return nil, nil
}

func (r *Wrapper) SearchRestaurant(search string) ([]Restaurant.Restaurants, error) {
	// TODO: add convert func
	return nil, nil
}
