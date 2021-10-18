package Restaurant

import (
	rest "2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

func AllRestaurants(db rest.WrapperRestaurant) ([]rest.Restaurant, error) {
	result, err := db.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetRestaurant(db rest.WrapperRestaurant, id int) (*rest.RestaurantAndCategory, []rest.Dishes, error) {
	restInfo, dishes, err := db.GetRestaurant(id)
	if err != nil {
		return nil, nil, err
	}
	return restInfo, dishes, nil
}
