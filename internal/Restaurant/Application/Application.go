package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
)

type Restaurant struct {
	DB Interface.WrapperRestaurantServer
}

func (r *Restaurant) AllRestaurants() ([]restaurant.Restaurants, error) {
	return r.DB.AllRestaurants()
}

func (r *Restaurant) GetRestaurant(id int) (*restaurant.RestaurantId, error) {
	return r.DB.GetRestaurant(id)
}

func (r *Restaurant) RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error) {
	return r.DB.RestaurantDishes(restId, dishId)
}

func (r *Restaurant) CreateReview(id int, review restaurant.NewReview) error {
	return r.DB.CreateReview(id, review)
}

func (r *Restaurant) GetReview(id int) (*restaurant.ResReview, error) {
	return r.DB.GetReview(id)
}

func (r *Restaurant) SearchRestaurant(search string) ([]restaurant.Restaurants, error) {
	return r.DB.SearchRestaurant(search)
}
