package application

import (
	Cart2 "2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/cart/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
)

type CartInterface interface {
	CalculatePriceDelivery(id int) (int, error)
	CalculateCost(result *Cart2.ResponseCartErrors, rest *restaurant.RestaurantId) (*Cart2.CostCartResponse, error)
	GetCart(id int) (*Cart2.ResponseCartErrors, error)
	UpdateCart(dishes Cart2.RequestCartDefault, clientId int) (*Cart2.ResponseCartErrors, error)
}

type Cart struct {
	DB ormPkg.WrapperCartInterface
}

func (c *Cart) CalculatePriceDelivery(id int) (int, error) {
	return c.DB.GetPriceDelivery(id)
}

func (c *Cart) CalculateCost(result *Cart2.ResponseCartErrors, rest *restaurant.RestaurantId) (*Cart2.CostCartResponse, error) {
	var cost Cart2.CostCartResponse
	sumCost := 0
	for i, dish := range result.Dishes {
		ingredientCost := 0
		for _, ingredient := range dish.IngredientCart {
			ingredientCost += ingredient.Cost
		}
		dishCost := (dish.Cost + ingredientCost) * dish.Count
		sumCost += dishCost
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
	cost.SumCost += cost.DCost
	return &cost, nil
}

func (c *Cart) GetCart(id int) (*Cart2.ResponseCartErrors, error) {
	result, errorDishes, err := c.DB.GetCart(id)
	if err != nil {
		return nil, err
	}

	rest, err := c.DB.GetRestaurant(result.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastFromRestaurantId(*rest)

	cost, err := c.CalculateCost(result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes

	return result, nil
}

func (c *Cart) UpdateCart(dishes Cart2.RequestCartDefault, clientId int) (*Cart2.ResponseCartErrors, error) {
	if dishes.Restaurant.Id == -1 {
		return nil, c.DB.DeleteCart(clientId)
	}

	err := c.DB.DeleteCart(clientId)
	if err != nil {
		return nil, err
	}

	result, errorDishes, err := c.DB.UpdateCart(dishes, clientId)
	if err != nil {
		return nil, err
	}

	rest, err := c.DB.GetRestaurant(dishes.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastFromRestaurantId(*rest)

	result.CastFromRequestCartDefault(dishes)

	cost, err := c.CalculateCost(result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes

	return result, nil
}
