package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
)

type WrapperCart interface {
	GetCart(id int) (*cart.ResponseCartErrors, []cart.CastDishesErrs, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, []cart.CastDishesErrs, error)
	DeleteCart(id int) error
	GetPriceDelivery(id int) (int, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
}
