//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/cart/orm WrapperCartServerInterface
package application

import (
	Cart2 "2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/cart/orm"
)

type CartApplicationInterface interface {
	GetCart(id int) (*Cart2.ResponseCartErrors, error)
	UpdateCart(dishes Cart2.RequestCartDefault, clientId int) (*Cart2.ResponseCartErrors, error)
	AddPromoCode(promoCode string, restaurantId int, clientId int) error
}

type Cart struct {
	DB ormPkg.WrapperCartServerInterface
}

func (c *Cart) GetCart(id int) (*Cart2.ResponseCartErrors, error) {
	return c.DB.GetCart(id)
}

func (c *Cart) UpdateCart(dishes Cart2.RequestCartDefault, clientId int) (*Cart2.ResponseCartErrors, error) {
	return c.DB.UpdateCart(dishes, clientId)
}

func (c *Cart) AddPromoCode(promoCode string, restaurantId int, clientId int) error {
	return c.DB.AddPromoCode(promoCode, restaurantId, clientId)
}
