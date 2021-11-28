package _interface

import "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"

type RestaurantApplication interface {
	AllRestaurants() ([]restaurant.Restaurants, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error)
	CreateReview(id int, review restaurant.NewReview) error
	GetReview(id int) (*restaurant.ResReview, error)
	SearchRestaurant(search string) ([]restaurant.Restaurants, error)
}
