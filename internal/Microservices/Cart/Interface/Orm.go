package Interface

import "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"

type WrapperCart interface {
	GetCart(id int) (*Cart.ResponseCartErrors, []Cart.CastDishesErrs, error)
	UpdateCart(dishes Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, []Cart.CastDishesErrs, error)
	DeleteCart(id int) error
	GetPriceDelivery(id int) (int, error)
}
