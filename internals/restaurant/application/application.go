//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/orm WrapperRestaurantServerInterface
package application

import (
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/orm"
)

type RestaurantApplicationInterface interface {
	AllRestaurants() (*resPkg.AllRestaurantsPromo, error)
	RecommendedRestaurants() (*resPkg.AllRestaurants, error)
	GetRestaurant(id int, idClient int) (*resPkg.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*resPkg.Dishes, error)
	CreateReview(id int, review resPkg.NewReview) error
	GetReview(idRestaurant int, idClient int) (*resPkg.ResReview, error)
	SearchRestaurant(search string) ([]resPkg.Restaurants, error)
	GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error)
	EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error)
	DeleteDish(idDish int) error
	AddDish(dish restaurant.DishHost) error
	AddRadios(dishId int, dish []restaurant.CreateRadios) error
	AddIngredient(dishId int, dish []restaurant.CreateIngredients) error
	UpdateDish(dish restaurant.DishHost) error
	UpdateIngredient(dishId int, ingredients []restaurant.CreateIngredients) error
	UpdateRadios(dishId int, radios []restaurant.CreateRadios) error
}

type Restaurant struct {
	DB ormPkg.WrapperRestaurantServerInterface
}

func (r *Restaurant) AllRestaurants() (*resPkg.AllRestaurantsPromo, error) {
	return r.DB.AllRestaurants()
}

func (r *Restaurant) RecommendedRestaurants() (*resPkg.AllRestaurants, error) {
	return r.DB.RecommendedRestaurants()
}

func (r *Restaurant) GetRestaurant(id int, idClient int) (*resPkg.RestaurantId, error) {
	return r.DB.GetRestaurant(id, idClient)
}

func (r *Restaurant) RestaurantDishes(restId int, dishId int) (*resPkg.Dishes, error) {
	return r.DB.RestaurantDishes(restId, dishId)
}

func (r *Restaurant) CreateReview(id int, review resPkg.NewReview) error {
	return r.DB.CreateReview(id, review)
}

func (r *Restaurant) GetReview(idRestaurant int, idClient int) (*resPkg.ResReview, error) {
	return r.DB.GetReview(idRestaurant, idClient)
}

func (r *Restaurant) SearchRestaurant(search string) ([]resPkg.Restaurants, error) {
	return r.DB.SearchRestaurant(search)
}

func (r *Restaurant) GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error) {
	return r.DB.GetFavoriteRestaurants(id)
}

func (r *Restaurant) EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error) {
	return r.DB.EditRestaurantInFavorite(idRestaurant, idClient)
}

func (r *Restaurant) DeleteDish(idDish int) error {
	return r.DB.DeleteDish(idDish)
}

func (r *Restaurant) AddDish(dish restaurant.DishHost) error {
	return r.DB.AddDish(dish)
}

func (r *Restaurant) AddRadios(dishId int, dish []restaurant.CreateRadios) error {
	return r.DB.AddRadios(dishId, dish)
}

func (r *Restaurant) AddIngredient(dishId int, dish []restaurant.CreateIngredients) error {
	return r.DB.AddIngredient(dishId, dish)
}

func (r *Restaurant) UpdateDish(dish restaurant.DishHost) error {
	return r.DB.UpdateDish(dish)
}

func (r *Restaurant) UpdateIngredient(dishId int, ingredients []restaurant.CreateIngredients) error {
	return r.DB.UpdateIngredient(dishId, ingredients)
}

func (r *Restaurant) UpdateRadios(dishId int, radios []restaurant.CreateRadios) error {
	return r.DB.UpdateRadios(dishId, radios)
}
