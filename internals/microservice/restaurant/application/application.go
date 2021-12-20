//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/orm WrapperRestaurantInterface
package application

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/orm"
)

type RestaurantApplicationInterface interface {
	AllRestaurantsPromo() (*resPkg.AllRestaurantsPromo, error)
	RecommendedRestaurants() (*resPkg.AllRestaurants, error)
	GetRestaurant(id int, idClient int) (*resPkg.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*resPkg.Dishes, error)
	CreateReview(id int, review resPkg.NewReview) error
	GetReview(idRestaurant int, idClient int) (*resPkg.ResReview, error)
	SearchRestaurant(search string) ([]resPkg.Restaurants, error)
	GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error)
	EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error)
}

type Restaurant struct {
	DB ormPkg.WrapperRestaurantInterface
}

func (r *Restaurant) AllRestaurantsPromo() (*resPkg.AllRestaurantsPromo, error) {
	restaurants, err := r.DB.GetRestaurants()
	if err != nil {
		return nil, err
	}
	restaurants.AllPromo, err = r.DB.GetPromoCodes()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *Restaurant) RecommendedRestaurants() (*resPkg.AllRestaurants, error) {
	return r.DB.GetRecommendedRestaurants()
}

func (r *Restaurant) GetRestaurant(id int, idClient int) (*resPkg.RestaurantId, error) {
	restInfo, err := r.DB.GetRestaurant(id)
	if err != nil {
		return nil, err
	}

	restInfo.Favourite, err = r.DB.IsFavoriteRestaurant(idClient, id)
	if err != nil {
		return nil, err
	}

	tags, err := r.DB.GetTagsRestaurant(id)
	if err != nil {
		return nil, err
	}

	dishes, err := r.DB.GetMenu(id)
	if err != nil {
		return nil, err
	}

	restInfo.Menu = dishes
	restInfo.Tags = tags
	return restInfo, nil
}

func (r *Restaurant) RestaurantDishes(restId int, dishId int) (*resPkg.Dishes, error) {
	return r.DB.GetDishes(restId, dishId)
}

func (r *Restaurant) CreateReview(id int, review resPkg.NewReview) error {
	return r.DB.CreateReview(id, review)
}

func (r *Restaurant) GetReview(idRestaurant int, idClient int) (*resPkg.ResReview, error) {
	var review resPkg.ResReview
	reviewInfo, err := r.DB.GetReview(idRestaurant)
	if err != nil {
		return nil, err
	}

	review.Status, err = r.DB.IsFavoriteRestaurant(idClient, idRestaurant)
	if err != nil {
		return nil, err
	}

	restInfo, err := r.DB.GetRestaurant(idRestaurant)
	if err != nil {
		return nil, err
	}

	tags, err := r.DB.GetTagsRestaurant(idRestaurant)
	if err != nil {
		return nil, err
	}
	restInfo.Tags = tags

	review.CastFromRestaurantId(*restInfo)
	review.Tags = tags
	review.Reviews = reviewInfo

	return &review, nil

}

func (r *Restaurant) SearchRestaurant(search string) ([]resPkg.Restaurants, error) {
	result, err := r.DB.SearchCategory(search)
	if err != nil {
		return nil, err
	}

	if result == nil {
		result, err = r.DB.SearchRestaurant(search)
		if err != nil {
			return nil, err
		}
	}

	var searchResult []resPkg.Restaurants
	for _, id := range result {
		restaurantInfo, err := r.DB.GetGeneralInfoRestaurant(id)
		if err != nil {
			return nil, err
		}
		searchResult = append(searchResult, *restaurantInfo)
	}
	return searchResult, nil

}

func (r *Restaurant) GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error) {
	return r.DB.GetFavoriteRestaurants(id)
}

func (r *Restaurant) EditRestaurantInFavorite(idRestaurants int, idClient int) (bool, error) {
	return r.DB.EditRestaurantInFavorite(idRestaurants, idClient)
}
