//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/order/orm WrapperOrderInterface
package application

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	Order2 "2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/order/orm"
	profileOrmPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm"
)

type OrderApplicationInterface interface {
	CreateOrder(id int, createOrder Order2.CreateOrder) (int, error)
	GetOrders(id int) (*Order2.HistoryOrderArray, error)
	GetActiveOrder(idClient int, idOrder int) (*Order2.ActiveOrder, error)
	UpdateStatusOrder(id int) error
	CancelOrder(id int, textCancel string) error
}

type Order struct {
	DB        ormPkg.WrapperOrderInterface
	DBProfile profileOrmPkg.WrapperProfileInterface
	IntCh     chan authPkg.WebSocketOrder
}

func (o *Order) CreateOrder(id int, createOrder Order2.CreateOrder) (int, error) {
	cart, err := o.DB.GetCart(id)
	if err != nil {
		return 0, err
	}

	err = o.DB.DeleteCart(id)
	if err != nil {
		return 0, err
	}

	courierId := 1

	addressId, err := o.DBProfile.AddAddress(id, createOrder.Address)
	if err != nil {
		return 0, err
	}

	err = o.DBProfile.DeleteAddress(id, addressId)
	if err != nil {
		return 0, err
	}

	order, err := o.DB.CreateOrder(id, createOrder, addressId, *cart, courierId)
	if err != nil {
		return 0, err
	}

	return order, nil
}

func (o *Order) GetOrders(id int) (*Order2.HistoryOrderArray, error) {
	return o.DB.GetOrders(id)
}

func (o *Order) GetActiveOrder(idClient int, idOrder int) (*Order2.ActiveOrder, error) {
	return o.DB.GetOrder(idClient, idOrder)
}

func (o *Order) UpdateStatusOrder(id int) error {
	status, err := o.DB.UpdateStatusOrder(id)
	o.IntCh <- authPkg.WebSocketOrder{Id: id, Status: status}
	return err
}

func (o *Order) CancelOrder(id int, textCancel string) error {
	return o.DB.CancelStatusOrder(id, textCancel)
}
