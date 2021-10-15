package Utils

import (
	res "2021_2_GORYACHIE_MEKSIKANSI/Utils/Restaurant"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error) // TODO: replace on interface Tx
}

type WrapperRestaurant interface {
	GetRestaurants() ([]res.Restaurant, error)
}

type WrapperProfile interface {
	GetRoleById(id int) (string, error)
	GetProfileClient(id int) (*Profile, error)
	GetProfileHost(id int) (*Profile, error)
	GetProfileCourier(id int) (*Profile, error)
}

type WrapperAuthorization interface {
	SignupClient(signup *RegistrationRequest) (*Defense, error)
	SignupCourier(signup *RegistrationRequest) (*Defense, error)
	SignupHost(signup *RegistrationRequest) (*Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(cookie *Defense) error
	AddCookie(cookie *Defense, id int) error
}

