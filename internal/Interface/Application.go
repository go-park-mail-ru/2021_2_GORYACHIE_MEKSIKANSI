package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Order"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
)

type AuthorizationApplication interface {
	SignUp(signup *Authorization.RegistrationRequest) (*Util.Defense, error)
	Login(login *Authorization.Authorization) (*Util.Defense, error)
	Logout(CSRF string) (string, error)
}

type CartApplication interface {
	GetCart(id int) (*Cart.ResponseCartErrors, error)
	UpdateCart(dishes Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, error)
}

type ProfileApplication interface {
	GetProfile(id int) (*Profile.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar *Profile.UpdateAvatar) error
	UpdateBirthday(id int, newBirthday string) error
	UpdateAddress(id int, newAddress Profile.AddressCoordinates) error
	AddAddress(id int, newAddress Profile.AddressCoordinates) (int, error)
	DeleteAddress(id int, addressId int) error
}

type MiddlewareApplication interface {
	CheckAccess(cookie *Util.Defense) (bool, error)
	NewCSRF(cookie *Util.Defense) (string, error)
	GetIdByCookie(cookie *Util.Defense) (int, error)
}

type RestaurantApplication interface {
	AllRestaurants() ([]Restaurant.Restaurants, error)
	GetRestaurant(id int) (*Restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*Restaurant.Dishes, error)
	CreateReview(id int, review Restaurant.NewReview) error
	GetReview(id int) (*Restaurant.ResReview, error)
	SearchRestaurant(search string) ([]Restaurant.Restaurants, error)
}

type OrderApplication interface {
	CreateOrder(id int, createOrder Order.CreateOrder) error
	GetOrders(id int) (*Order.HistoryOrderArray, error)
	GetActiveOrder(idClient int, idOrder int) (*Order.ActiveOrder, error)
	UpdateStatusOrder(id int, status int) error
}
