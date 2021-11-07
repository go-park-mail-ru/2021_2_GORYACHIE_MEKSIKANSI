package Interfaces

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type WrapperRestaurant interface {
	GetRestaurants() ([]Utils.Restaurants, error)
	GetStructDishes(dishesId int) ([]Utils.Ingredients, error)
	GetRadios(dishesId int) ([]Utils.Radios, error)
	GetDishes(restId int, dishesId int) (*Utils.Dishes, error)
	GetGeneralInfoRestaurant(id int) (*Utils.RestaurantId, error)
	GetMenu(id int) ([]Utils.Menu, error)
	GetTagsRestaurant(id int) ([]Utils.Tag, error)
}

type WrapperProfile interface {
	GetRoleById(id int) (string, error)
	GetProfileClient(id int) (*Utils.Profile, error)
	GetProfileHost(id int) (*Utils.Profile, error)
	GetProfileCourier(id int) (*Utils.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar string) error
	UpdateBirthday(id int, newBirthday time.Time) error
	UpdateAddress(id int, newAddress Utils.AddressCoordinates) error
}

type WrapperAuthorization interface {
	GeneralSignUp(signup *Utils.RegistrationRequest, transaction pgx.Tx) (int, error)
	SignupClient(signup *Utils.RegistrationRequest, cookie *Utils.Defense) (*Utils.Defense, error)
	SignupCourier(signup *Utils.RegistrationRequest, cookie *Utils.Defense) (*Utils.Defense, error)
	SignupHost(signup *Utils.RegistrationRequest, cookie *Utils.Defense) (*Utils.Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(cookie *Utils.Defense) error
	GenerateNew() *Utils.Defense
	AddCookie(cookie *Utils.Defense, id int) error
	AddTransactionCookie(cookie *Utils.Defense, Transaction pgx.Tx, id int) error
}

type WrapperCart interface {
	GetCart(id int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error)
	UpdateCart(dishes Utils.RequestCartDefault, clientId int) (*Utils.ResponseCartErrors, []Utils.CastDishesErrs, error)
	DeleteCart(id int) error
	GetPriceDelivery(id int) (int, error)
	UpdateCartRadios(radios []Utils.RadiosCartRequest, clientId int, tx pgx.Tx) ([]Utils.RadiosCartResponse, error)
	GetStructRadios(id int) ([]Utils.RadiosCartResponse, error)
	GetStructFood(id int) ([]Utils.IngredientCartResponse, error)
	UpdateCartStructureFood(ingredients []Utils.IngredientsCartRequest, clientId int, tx pgx.Tx) ([]Utils.IngredientCartResponse, error)
}

type WrapperMiddleware interface {
	CheckAccess(cookie *Utils.Defense) (bool, error)
	NewCSRF(cookie *Utils.Defense) (string, error)
	GetIdByCookie(cookie *Utils.Defense) (int, error)
}
