package Cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

func calculatePriceDelivery(db Utils.WrapperCart, id int) (int, error) {
	return db.GetPriceDelivery(id)
}

func calculateCost(db Utils.WrapperCart, result *Utils.ResponseCartErrors, rest *Utils.RestaurantId) (*Utils.CostCartResponse, error) {
	var cost Utils.CostCartResponse
	sumCost := 0
	for i, dish := range result.Dishes {
		ingredientCost := 0
		for _, ingredient := range dish.IngredientCart {
			ingredientCost = ingredientCost + ingredient.Cost
		}
		dishCost := (dish.Cost + ingredientCost) * dish.Count
		sumCost = sumCost + dishCost
		result.Dishes[i].Cost = dishCost
	}
	cost.SumCost = sumCost
	if sumCost >= rest.CostForFreeDelivery {
		cost.DCost = 0
	} else {
		var err error
		cost.DCost, err = calculatePriceDelivery(db, rest.Id)
		if err != nil {
			return nil, err
		}
	}
	cost.SumCost = cost.DCost + cost.SumCost
	return &cost, nil
}

func GetCart(db Utils.WrapperCart, id int) (*Utils.ResponseCartErrors, error) {
	result, errorDishes, err := db.GetCart(id)
	if err != nil {
		return nil, err
	}
	wrapper := Restaurant.Wrapper{Conn: db.GetConn()}
	rest, err := wrapper.GetGeneralInfoRestaurant(result.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastToRestaurantId(*rest)

	cost, err := calculateCost(db, result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes

	return result, nil
}

func UpdateCart(db Utils.WrapperCart, dishes Utils.RequestCartDefault, clientId int) (*Utils.ResponseCartErrors, error) {
	if dishes.Restaurant.Id == -1 {
		return nil, DeleteCart(db, clientId)
	}

	err := DeleteCart(db, clientId)
	if err != nil {
		return nil, err
	}

	result, errorDishes, err := db.UpdateCart(dishes, clientId)
	if err != nil {
		return nil, err
	}

	wrapper := Restaurant.Wrapper{Conn: db.GetConn()}
	rest, err := wrapper.GetGeneralInfoRestaurant(dishes.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastToRestaurantId(*rest)

	result.CastToRequestCartDefault(dishes)

	sumCost := 0
	for i, dish := range result.Dishes {
		ingredientCost := 0
		for _, ingredient := range dish.IngredientCart {
			ingredientCost = ingredientCost + ingredient.Cost
		}
		dishCost := (dish.Cost + ingredientCost) * dish.Count
		sumCost = sumCost + dishCost
		result.Dishes[i].Cost = dishCost
	}

	cost, err := calculateCost(db, result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes
	return result, nil
}

func DeleteCart(db Utils.WrapperCart, id int) error {
	return db.DeleteCart(id)
}
