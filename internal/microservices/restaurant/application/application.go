package application

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant/orm"
)

type RestaurantApplicationInterface interface {
	AllRestaurants() ([]resPkg.Restaurants, error)
	GetRestaurant(id int) (*resPkg.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*resPkg.Dishes, error)
	CreateReview(id int, review resPkg.NewReview) error
	GetReview(id int) (*resPkg.ResReview, error)
	SearchRestaurant(search string) ([]resPkg.Restaurants, error)
}

type Restaurant struct {
	DB ormPkg.WrapperRestaurantInterface
}

func (r *Restaurant) AllRestaurants() ([]resPkg.Restaurants, error) {
	result, err := r.DB.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Restaurant) GetRestaurant(id int) (*resPkg.RestaurantId, error) {
	restInfo, err := r.DB.GetRestaurant(id)
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
	dishes, err := r.DB.GetDishes(restId, dishId)
	if err != nil {
		return nil, err
	}

	dishes.Ingredient, err = r.DB.GetStructDishes(dishId)
	if err != nil {
		return nil, err
	}

	dishes.Radios, err = r.DB.GetRadios(dishId)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *Restaurant) CreateReview(id int, review resPkg.NewReview) error {
	err := r.DB.CreateReview(id, review)
	if err != nil {
		return err
	}
	return nil

}

func (r *Restaurant) GetReview(id int) (*resPkg.ResReview, error) {
	var review resPkg.ResReview
	reviewInfo, err := r.DB.GetReview(id)
	if err != nil {
		return nil, err
	}

	restInfo, err := r.DB.GetRestaurant(id)
	if err != nil {
		return nil, err
	}

	tags, err := r.DB.GetTagsRestaurant(id)
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
