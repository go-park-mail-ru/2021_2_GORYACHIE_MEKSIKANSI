package Interface

import "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"

type WrapperRestaurant interface {
	GetRestaurants() ([]Restaurant.Restaurants, error)
	GetStructDishes(dishesId int) ([]Restaurant.Ingredients, error)
	GetRadios(dishesId int) ([]Restaurant.Radios, error)
	GetDishes(restId int, dishesId int) (*Restaurant.Dishes, error)
	GetRestaurant(id int) (*Restaurant.RestaurantId, error)
	GetMenu(id int) ([]Restaurant.Menu, error)
	GetTagsRestaurant(id int) ([]Restaurant.Tag, error)
	GetReview(id int) ([]Restaurant.Review, error)
	CreateReview(id int, review Restaurant.NewReview) error
	SearchCategory(name string) ([]int, error)
	SearchRestaurant(name string) ([]int, error)
	GetGeneralInfoRestaurant(id int) (*Restaurant.Restaurants, error)
}
