package Restaurant

import (
	rest "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	res "2021_2_GORYACHIE_MEKSIKANSI/Utils/Restaurant"
)

func AllRestaurants(db rest.WrapperRestaurant) ([]res.Restaurant, error) {
	result, err := db.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}
