//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internal/order/orm WrapperOrderInterface
package application

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	Order2 "2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/order/orm"
	profileOrmPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/profile/orm"
	"time"
)

type OrderApplicationInterface interface {
	CreateOrder(id int, createOrder Order2.CreateOrder) (int, error)
	GetOrders(id int) (*Order2.HistoryOrderArray, error)
	GetActiveOrder(idClient int, idOrder int) (*Order2.ActiveOrder, error)
	UpdateStatusOrder(id int, status int) error
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

	rest, err := o.DB.GetRestaurant(cart.Restaurant.Id)
	if err != nil {
		return 0, err
	}

	cart.CastFromRestaurantId(*rest)

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

	//err = o.DB.DoPromoCode(cart.PromoCode, rest.Id, cart)
	//if err != nil {
	//	return nil, err
	//}

	order, err := o.DB.CreateOrder(id, createOrder, addressId, *cart, courierId)
	if err != nil {
		return 0, err
	}

	//TODO: delete
	check, _ := o.DB.CheckRun(order)
	if check {
		go o.UpdateStatusOrder(order, 4)
	}

	return order, nil
}

func (o *Order) GetOrders(id int) (*Order2.HistoryOrderArray, error) {
	return o.DB.GetOrders(id)
}

func (o *Order) GetActiveOrder(idClient int, idOrder int) (*Order2.ActiveOrder, error) {
	return o.DB.GetOrder(idClient, idOrder)
}

func (o *Order) UpdateStatusOrder(id int, status int) error {
	for i := 1; i <= 4; i++ {
		time.Sleep(time.Second * 15)
		o.IntCh <- authPkg.WebSocketOrder{Id: id, Status: i}
		o.DB.UpdateStatusOrder(id, i)
	}
	// o.IntCh <- authPkg.WebSocketOrder{Id: id, Status: status}
	return o.DB.UpdateStatusOrder(id, status)
}
