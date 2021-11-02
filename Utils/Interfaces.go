package Utils

import (
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
	GetRestaurants() ([]Restaurants, error)
	GetStructureDishes(dishesId int) ([]Ingredients, error)
	GetRadios(dishesId int) ([]Radios, error)
	GetDishes(restId int, dishesId int) (*Dishes, error)
	GetGeneralInfoRestaurant(id int) (*RestaurantId, error)
	GetMenu(id int) ([]Menu, error)
	GetTagsRestaurant(id int) ([]Tag, error)
}

type WrapperProfile interface {
	GetRoleById(id int) (string, error)
	GetProfileClient(id int) (*Profile, error)
	GetProfileHost(id int) (*Profile, error)
	GetProfileCourier(id int) (*Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar string) error
	UpdateBirthday(id int, newBirthday time.Time) error
	UpdateAddress(id int, newAddress AddressCoordinates) error
}

type WrapperAuthorization interface {
	SignupClient(signup *RegistrationRequest) (*Defense, error)
	SignupCourier(signup *RegistrationRequest) (*Defense, error)
	SignupHost(signup *RegistrationRequest) (*Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(cookie *Defense) error
	GenerateNew() *Defense
	AddCookie(cookie *Defense, id int) error
}

type WrapperCart interface {
	GetCart(id int) (*ResponseCartErrors, []CastDishesErrs, error)
	UpdateCart(dishes RequestCartDefault, clientId int) (*ResponseCartErrors, []CastDishesErrs, error)
	DeleteCart(id int) error
	GetConn() ConnectionInterface
	GetPriceDelivery(id int) (int, error)
	UpdateCartRadios(radios []RadiosCartRequest, clientId int, tx pgx.Tx) ([]RadiosCartResponse, error)
	GetStructureRadios(id int) ([]RadiosCartResponse, error)
	GetCartStructureFood(id int) ([]IngredientCartResponse, error)
	UpdateCartStructureFood(ingredients []IngredientsCartRequest, clientId int, tx pgx.Tx) ([]IngredientCartResponse, error)
}
