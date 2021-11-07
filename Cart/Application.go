package Cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

type Cart struct {
	DB Interfaces.WrapperCart
	DBRestaurant Interfaces.WrapperRestaurant
}

func (c *Cart) CalculatePriceDelivery(id int) (int, error) {
	return c.DB.GetPriceDelivery(id)
}

func (c *Cart) CalculateCost(result *Utils.ResponseCartErrors, rest *Utils.RestaurantId) (*Utils.CostCartResponse, error) {
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
		cost.DCost, err = c.CalculatePriceDelivery(rest.Id)
		if err != nil {
			return nil, err
		}
	}
	cost.SumCost = cost.DCost + cost.SumCost
	return &cost, nil
}

func (c *Cart) GetCart(id int) (*Utils.ResponseCartErrors, error) {
	result, errorDishes, err := c.DB.GetCart(id)
	if err != nil {
		return nil, err
	}
	rest, err := c.DBRestaurant.GetGeneralInfoRestaurant(result.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastToRestaurantId(*rest)

	cost, err := c.CalculateCost(result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes

	return result, nil
}

func (c *Cart) UpdateCart(dishes Utils.RequestCartDefault, clientId int) (*Utils.ResponseCartErrors, error) {
	if dishes.Restaurant.Id == -1 {
		return nil, c.DeleteCart(clientId)
	}

	err := c.DeleteCart(clientId)
	if err != nil {
		return nil, err
	}

	result, errorDishes, err := c.DB.UpdateCart(dishes, clientId)
	if err != nil {
		return nil, err
	}

	rest, err := c.DBRestaurant.GetGeneralInfoRestaurant(dishes.Restaurant.Id)
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

	cost, err := c.CalculateCost(result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes
	return result, nil
}

func (c *Cart) DeleteCart(id int) error {
	return c.DB.DeleteCart(id)
}
