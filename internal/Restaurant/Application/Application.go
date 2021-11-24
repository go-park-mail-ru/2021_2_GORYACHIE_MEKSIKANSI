package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
)

type Restaurant struct {
	DB Interface.WrapperRestaurant
}

func (r *Restaurant) AllRestaurants() ([]restaurant.Restaurants, error) {
	result, err := r.DB.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Restaurant) GetRestaurant(id int) (*restaurant.RestaurantId, error) {
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

func (r *Restaurant) RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error) {
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

func (r *Restaurant) CreateReview(id int, review restaurant.NewReview) error {
	return r.DB.CreateReview(id, review)
}

func (r *Restaurant) GetReview(id int) (*restaurant.ResReview, error) {
	var review restaurant.ResReview
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

func (r *Restaurant) SearchRestaurant(search string) ([]restaurant.Restaurants, error) {
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

	var searchResult []restaurant.Restaurants
	for _, id := range result {
		restaurantInfo, err := r.DB.GetGeneralInfoRestaurant(id)
		if err != nil {
			return nil, err
		}
		searchResult = append(searchResult, *restaurantInfo)
	}
	return searchResult, nil
}
