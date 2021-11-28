package _interface

import "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"

type WrapperRestaurant interface {
	GetRestaurants() ([]restaurant.Restaurants, error)
	GetStructDishes(dishesId int) ([]restaurant.Ingredients, error)
	GetRadios(dishesId int) ([]restaurant.Radios, error)
	GetDishes(restId int, dishesId int) (*restaurant.Dishes, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
	GetMenu(id int) ([]restaurant.Menu, error)
	GetTagsRestaurant(id int) ([]restaurant.Tag, error)
	GetReview(id int) ([]restaurant.Review, error)
	CreateReview(id int, review restaurant.NewReview) error
	SearchCategory(name string) ([]int, error)
	SearchRestaurant(name string) ([]int, error)
	GetGeneralInfoRestaurant(id int) (*restaurant.Restaurants, error)
}
