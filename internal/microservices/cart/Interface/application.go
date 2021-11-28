package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
)

type CartApplication interface {
	CalculatePriceDelivery(id int) (int, error)
	CalculateCost(result *cart.ResponseCartErrors, rest *restaurant.RestaurantId) (*cart.CostCartResponse, error)
	GetCart(id int) (*cart.ResponseCartErrors, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error)
}
