package Interface

import "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"

type RestaurantApplication interface {
	AllRestaurants() ([]Restaurant.Restaurants, error)
	GetRestaurant(id int) (*Restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*Restaurant.Dishes, error)
	CreateReview(id int, review Restaurant.NewReview) error
	GetReview(id int) (*Restaurant.ResReview, error)
	SearchRestaurant(search string) ([]Restaurant.Restaurants, error)
}
