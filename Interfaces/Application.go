package Interfaces

import (
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"time"
)

type AuthorizationApplication interface {
	SignUp(signup *utils.RegistrationRequest) (*utils.Defense, error)
	Login(login *utils.Authorization) (*utils.Defense, error)
	Logout(CSRF string) (string, error)
}

type CartApplication interface {
	CalculatePriceDelivery(id int) (int, error)
	CalculateCost(result *utils.ResponseCartErrors, rest *utils.RestaurantId) (*utils.CostCartResponse, error)
	GetCart(id int) (*utils.ResponseCartErrors, error)
	UpdateCart(dishes utils.RequestCartDefault, clientId int) (*utils.ResponseCartErrors, error)
	DeleteCart(id int) error
}

type ProfileApplication interface {
	GetProfile(id int) (*utils.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar *utils.UpdateAvatar) error
	UpdateBirthday(id int, newBirthday time.Time) error
	UpdateAddress(id int, newAddress utils.AddressCoordinates) error
}

type MiddlewareApplication interface {
	CheckAccess(cookie *utils.Defense) (bool, error)
	NewCSRF(cookie *utils.Defense) (string, error)
	GetIdByCookie(cookie *utils.Defense) (int, error)
}

type RestaurantApplication interface {
	AllRestaurants() ([]utils.Restaurants, error)
	GetRestaurant(id int) (*utils.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*utils.Dishes, error)
}
