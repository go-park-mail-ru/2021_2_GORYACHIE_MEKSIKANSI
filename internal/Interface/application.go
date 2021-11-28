package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
)

type AuthorizationApplication interface {
	SignUp(signup *authorization.RegistrationRequest) (*util.Defense, error)
	Login(login *authorization.Authorization) (*util.Defense, error)
	Logout(CSRF string) (string, error)
}

type CartApplication interface {
	GetCart(id int) (*cart.ResponseCartErrors, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error)
}

type ProfileApplication interface {
	GetProfile(id int) (*profile.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar *profile.UpdateAvatar) error
	UpdateBirthday(id int, newBirthday string) error
	UpdateAddress(id int, newAddress profile.AddressCoordinates) error
	AddAddress(id int, newAddress profile.AddressCoordinates) (int, error)
	DeleteAddress(id int, addressId int) error
}

type MiddlewareApplication interface {
	CheckAccess(cookie *util.Defense) (bool, error)
	NewCSRF(cookie *util.Defense) (string, error)
	GetIdByCookie(cookie *util.Defense) (int, error)
}

type RestaurantApplication interface {
	AllRestaurants() ([]restaurant.Restaurants, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error)
	CreateReview(id int, review restaurant.NewReview) error
	GetReview(id int) (*restaurant.ResReview, error)
	SearchRestaurant(search string) ([]restaurant.Restaurants, error)
}

type OrderApplication interface {
	CreateOrder(id int, createOrder order.CreateOrder) error
	GetOrders(id int) (*order.HistoryOrderArray, error)
	GetActiveOrder(idClient int, idOrder int) (*order.ActiveOrder, error)
	UpdateStatusOrder(id int, status int) error
}
