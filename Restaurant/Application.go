package Restaurant

import (
	rest "2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

func AllRestaurants(db rest.WrapperRestaurant) ([]rest.Restaurants, error) {
	result, err := db.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetRestaurant(db rest.WrapperRestaurant, id int) (*rest.RestaurantId, error) {
	restInfo, err := db.GetGeneralInfoRestaurant(id)
	if err != nil {
		return nil, err
	}

	tags, err := db.GetTagsRestaurant(id)
	if err != nil {
		return nil, err
	}

	dishes, err := db.GetMenu(id)
	if err != nil {
		return nil, err
	}

	restInfo.Menu = dishes
	restInfo.Tags = tags
	return restInfo, nil
}

func RestaurantDishes(db rest.WrapperRestaurant, restId int, dishId int) (*rest.Dishes, error) {
	dishes, err := db.GetDishes(restId, dishId)
	if err != nil {
		return nil, err
	}

	dishes.Ingredient, err = db.GetStructureDishes(dishId)
	if err != nil {
		return nil, err
	}

	dishes.Radios, err = db.GetRadios(dishId)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}
