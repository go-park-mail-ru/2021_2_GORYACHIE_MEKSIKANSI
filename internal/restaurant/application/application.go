//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/orm WrapperRestaurantServerInterface
package application

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/orm"
)

type RestaurantApplicationInterface interface {
	AllRestaurants() (*resPkg.AllRestaurantsPromo, error)
	GetRestaurant(id int, idClient int) (*resPkg.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*resPkg.Dishes, error)
	CreateReview(id int, review resPkg.NewReview) error
	GetReview(idRestaurant int, idClient int) (*resPkg.ResReview, error)
	SearchRestaurant(search string) ([]resPkg.Restaurants, error)
	GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error)
	EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error)
}

type Restaurant struct {
	DB ormPkg.WrapperRestaurantServerInterface
}

func (r *Restaurant) AllRestaurants() (*resPkg.AllRestaurantsPromo, error) {
	return r.DB.AllRestaurants()
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
