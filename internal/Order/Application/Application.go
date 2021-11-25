package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	Order2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Order"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"time"
)

type Order struct {
	DB           Interface.WrapperOrder
	DBCart       Interface.WrapperCart
	DBProfile    Interface.WrapperProfile
	DBRestaurant Interface.WrapperRestaurant
}

func (o *Order) CalculatePriceDelivery(id int) (int, error) {
	return o.DBCart.GetPriceDelivery(id)
}

func (o *Order) CalculateCost(result *Cart.ResponseCartErrors, rest *Restaurant.RestaurantId) (*Cart.CostCartResponse, error) {
	var cost Cart.CostCartResponse
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

func (o *Order) CreateOrder(id int, createOrder Order2.CreateOrder) error {
	cart, errDish, err := o.DBCart.GetCart(id)
	if err != nil || errDish != nil || cart.Dishes == nil {
		return &errPkg.Errors{
			Alias: errPkg.OCreateOrderCartIsVoid,
		}
	}
	if errDish != nil {
		return err
	}

	rest, err := o.DBRestaurant.GetRestaurant(cart.Restaurant.Id)
	if err != nil {
		return err
	}

	cart.CastFromRestaurantId(*rest)

	cost, err := o.CalculateCost(cart, rest)
	if err != nil {
		return err
	}
	cart.Cost = *cost

	courierId := 1

	err = o.DBCart.DeleteCart(id)
	if err != nil {
		return err
	}

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

func (o *Order) GetOrders(id int) (*Order2.HistoryOrderArray, error) {
	return o.DB.GetOrders(id)
}

func (o *Order) GetActiveOrder(idClient int, idOrder int) (*Order2.ActiveOrder, error) {
	order, err := o.DB.GetOrder(idClient, idOrder)
	if err != nil {
		return nil, err
	}
	check, _ := o.DB.CheckRun(idOrder)
	if check {
		go o.UpdateStatusOrder(idOrder, 4)
	}

	return order, nil
}

func (o *Order) UpdateStatusOrder(id int, status int) error {
	for i := 1; i <= 4; i++ {
		time.Sleep(time.Second * 15)
		o.DB.UpdateStatusOrder(id, i)
	}
	return o.DB.UpdateStatusOrder(id, status)
}
