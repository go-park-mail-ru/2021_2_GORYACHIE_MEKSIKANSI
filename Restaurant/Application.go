package Restaurant

func AllRestaurants(db Wrapper) ([]Restaurant, error) {
	result, err := db.GetRestaurants()
	if err != nil {
		return nil, err
	}
	return result, nil
}
