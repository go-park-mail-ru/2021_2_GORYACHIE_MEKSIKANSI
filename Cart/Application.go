package Cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

func calculatePriceDelivery(db Utils.WrapperCart, id int) (int, error) {
	return db.GetPriceDelivery(id)
}

func GetCart(db Utils.WrapperCart, id int) (*Utils.ResponseCart, error) {
	var cost Utils.CostCartResponse
	result, err := db.GetCart(id)
	if err != nil {
		return nil, err
	}
	wrapper := Restaurant.Wrapper{Conn: db.GetConn()}
	rest, err := wrapper.GetGeneralInfoRestaurant(result.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.Restaurant.Id = rest.Id
	result.Restaurant.Img = rest.Img
	result.Restaurant.Rating = rest.Rating
	result.Restaurant.CostForFreeDelivery = rest.CostForFreeDelivery
	result.Restaurant.Name = rest.Name
	result.Restaurant.MaxDelivery = rest.MaxDelivery
	result.Restaurant.MinDelivery = rest.MinDelivery

	sumCost := 0
	for i, dish := range result.Dishes {
		ingredientCost := 0
		for _, ingredient := range dish.IngredientCart {
			ingredientCost = ingredientCost + ingredient.Cost
		}
		dishCost := dish.Cost * dish.Count + ingredientCost
		sumCost = sumCost + dishCost
		result.Dishes[i].Cost = dishCost
	}
	cost.SumCost = sumCost
	if sumCost >= rest.CostForFreeDelivery {
		cost.DCost = 0
	} else {
		cost.DCost, _ = calculatePriceDelivery(db, rest.Id)
	}
	cost.SumCost = cost.DCost + cost.SumCost
	result.Cost = cost
	return result, nil
}

func UpdateCart(db Utils.WrapperCart, dishes Utils.CartRequest, clientId int) (*Utils.ResponseCart, []Utils.CastDishesErrs, error) {
	if dishes.Restaurant.Id == -1 {
		return nil, nil, DeleteCart(db, clientId)
	}

	err := DeleteCart(db, clientId)
	if err != nil {
		return nil, nil, err
	}
	result, errCast, err := db.UpdateCart(dishes, clientId)
	if err != nil || errCast != nil {
		return nil, errCast, err
	}
	wrapper := Restaurant.Wrapper{Conn: db.GetConn()}
	rest, err := wrapper.GetGeneralInfoRestaurant(dishes.Restaurant.Id)
	if err != nil {
		return nil, nil, err
	}

	result.Restaurant.Id = rest.Id
	result.Restaurant.Img = rest.Img
	result.Restaurant.Rating = rest.Rating
	result.Restaurant.CostForFreeDelivery = rest.CostForFreeDelivery
	result.Restaurant.Name = rest.Name
	result.Restaurant.MaxDelivery = rest.MaxDelivery
	result.Restaurant.MinDelivery = rest.MinDelivery

	result.Cast(dishes)

	sumCost := 0
	for i, dish := range result.Dishes {
		ingredientCost := 0
		for _, ingredient := range dish.IngredientCart {
			ingredientCost = ingredientCost + ingredient.Cost
		}
		dishCost := dish.Cost * dish.Count + ingredientCost
		sumCost = sumCost + dishCost
		result.Dishes[i].Cost = dishCost
	}
	var cost Utils.CostCartResponse
	cost.SumCost = sumCost
	if sumCost >= rest.CostForFreeDelivery {
		cost.DCost = 0
	} else {
		cost.DCost, _ = calculatePriceDelivery(db, rest.Id)
	}
	cost.SumCost = cost.DCost + cost.SumCost
	result.Cost = cost
	return result, errCast, nil
}

func DeleteCart(db Utils.WrapperCart, id int) error {
	return db.DeleteCart(id)
}
