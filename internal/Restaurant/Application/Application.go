package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
)

type Restaurant struct {
	DB Interface.WrapperRestaurant
}

func (r *Restaurant) AllRestaurants() ([]utils.Restaurants, error) {
	result, err := r.DB.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Restaurant) GetRestaurant(id int) (*utils.RestaurantId, error) {
	restInfo, err := r.DB.GetGeneralInfoRestaurant(id)
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

func (r *Restaurant) RestaurantDishes(restId int, dishId int) (*utils.Dishes, error) {
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
