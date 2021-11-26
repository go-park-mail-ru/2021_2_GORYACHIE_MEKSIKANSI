package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
)

type CartApplication interface {
	CalculatePriceDelivery(id int) (int, error)
	CalculateCost(result *Cart.ResponseCartErrors, rest *Restaurant.RestaurantId) (*Cart.CostCartResponse, error)
	GetCart(id int) (*Cart.ResponseCartErrors, error)
	UpdateCart(dishes Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, error)
}
