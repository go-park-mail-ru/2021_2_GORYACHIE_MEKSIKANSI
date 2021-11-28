package application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	Order2 "2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	"time"
)

type Order struct {
	DB Interface.WrapperOrder
	//DBCart       interface.WrapperCartOld
	DBProfile Interface.WrapperProfile
	//DBRestaurant interface.WrapperRestaurant
}

func (o *Order) CreateOrder(id int, createOrder Order2.CreateOrder) error {
	cart, err := o.DB.GetCart(id)
	if err != nil {
		return err
	}

	rest, err := o.DB.GetRestaurant(cart.Restaurant.Id)
	if err != nil {
		return err
	}

	cart.CastFromRestaurantId(*rest)

	err = o.DB.DeleteCart(id)
	if err != nil {
		return err
	}

	courierId := 1

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
