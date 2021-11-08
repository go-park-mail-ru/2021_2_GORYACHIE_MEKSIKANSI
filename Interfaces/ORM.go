package Interfaces

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
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
	UpdateAvatar(id int, newAvatar *Utils.UpdateAvatar) error
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
	DeleteCookie(CSRF string) (string, error)
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
	UpdateCartStructFood(ingredients []Utils.IngredientsCartRequest, clientId int, tx pgx.Tx) ([]Utils.IngredientCartResponse, error)
}

type WrapperMiddleware interface {
	CheckAccess(cookie *Utils.Defense) (bool, error)
	NewCSRF(cookie *Utils.Defense) (string, error)
	GetIdByCookie(cookie *Utils.Defense) (int, error)
}
type Uploader interface {
	Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}