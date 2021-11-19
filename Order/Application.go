package Order

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

type Order struct {
	DB           Interfaces.WrapperOrder
	DBCart       Interfaces.WrapperCart
	DBProfile    Interfaces.WrapperProfile
	DBRestaurant Interfaces.WrapperRestaurant
}

func (o *Order) CalculatePriceDelivery(id int) (int, error) {
	return o.DB.GetPriceDelivery(id)
}

func (o *Order) CalculateCost(result *utils.ResponseCartErrors, rest *utils.RestaurantId) (*utils.CostCartResponse, error) {
	var cost utils.CostCartResponse
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
		cost.DCost, err = o.CalculatePriceDelivery(rest.Id)
		if err != nil {
			return nil, err
		}
	}
	cost.SumCost += cost.DCost
	return &cost, nil
}

func (o *Order) CreateOrder(id int, createOrder utils.CreateOrder) error {
	cart, errDish, err := o.DBCart.GetCart(id)
	if err != nil || errDish != nil || cart.Dishes == nil {
		return &errPkg.Errors{
			Alias: errPkg.OCreateOrderCartIsVoid,
		}
	}
	if errDish != nil {
		return err
	}

	rest, err := o.DBRestaurant.GetGeneralInfoRestaurant(cart.Restaurant.Id)
	if err != nil {
		return err
	}

	cart.CastToRestaurantId(*rest)

	cost, err := o.CalculateCost(cart, rest)
	if err != nil {
		return err
	}
	cart.Cost = *cost

	courierId := 1

	//err = o.DBCart.DeleteCart(id)
	//if err != nil {
	//	return err
	//}

	addressId, err := o.DBProfile.AddAddress(id, createOrder.Address)
	if err != nil {
		return err
	}

	err = o.DBProfile.DeleteAddress(id, addressId)
	if err != nil {
		return err
	}

	return o.DB.CreateOrder(id, createOrder, addressId, *cart, courierId)
}

func (o *Order) GetOrders(id int) (*utils.HistoryOrderArray, error) {
	orders, err := o.DB.GetOrders(id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
