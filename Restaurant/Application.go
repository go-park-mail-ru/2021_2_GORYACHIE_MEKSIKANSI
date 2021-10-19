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
	restInfo, tags, dishes, err := db.GetRestaurant(id)
	if err != nil {
		return nil, err
	}
	restInfo.Menu = dishes
	restInfo.Tags = tags
	return restInfo, nil
}
