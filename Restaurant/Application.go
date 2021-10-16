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
