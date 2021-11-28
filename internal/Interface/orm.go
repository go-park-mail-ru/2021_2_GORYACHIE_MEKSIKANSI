package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/proto"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/cart/proto"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/order"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"context"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
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

type ConnectRestaurantService interface {
	AllRestaurants(ctx context.Context, in *resProto.Empty, opts ...grpc.CallOption) (*resProto.Restaurants, error)
	GetRestaurant(ctx context.Context, in *resProto.RestaurantId, opts ...grpc.CallOption) (*resProto.RestaurantInfo, error)
	RestaurantDishes(ctx context.Context, in *resProto.DishInfo, opts ...grpc.CallOption) (*resProto.Dishes, error)
	CreateReview(ctx context.Context, in *resProto.NewReview, opts ...grpc.CallOption) (*resProto.Error, error)
	GetReview(ctx context.Context, in *resProto.RestaurantId, opts ...grpc.CallOption) (*resProto.ResReview, error)
	SearchRestaurant(ctx context.Context, in *resProto.SearchRestaurantText, opts ...grpc.CallOption) (*resProto.Restaurants, error)
}

type WrapperRestaurantServer interface {
	AllRestaurants() ([]restaurant.Restaurants, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error)
	CreateReview(id int, review restaurant.NewReview) error
	GetReview(id int) (*restaurant.ResReview, error)
	SearchRestaurant(search string) ([]restaurant.Restaurants, error)
}

type WrapperRestaurant interface {
	GetRestaurants() ([]restaurant.Restaurants, error)
	GetStructDishes(dishesId int) ([]restaurant.Ingredients, error)
	GetRadios(dishesId int) ([]restaurant.Radios, error)
	GetDishes(restId int, dishesId int) (*restaurant.Dishes, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
	GetMenu(id int) ([]restaurant.Menu, error)
	GetTagsRestaurant(id int) ([]restaurant.Tag, error)
	GetReview(id int) ([]restaurant.Review, error)
	CreateReview(id int, review restaurant.NewReview) error
	SearchCategory(name string) ([]int, error)
	SearchRestaurant(name string) ([]int, error)
	GetGeneralInfoRestaurant(id int) (*restaurant.Restaurants, error)
}

type WrapperProfile interface {
	GetRoleById(id int) (string, error)
	GetProfileClient(id int) (*profile.Profile, error)
	GetProfileHost(id int) (*profile.Profile, error)
	GetProfileCourier(id int) (*profile.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar *profile.UpdateAvatar, newFileName string) error
	UpdateBirthday(id int, newBirthday string) error
	UpdateAddress(id int, newAddress profile.AddressCoordinates) error
	AddAddress(id int, newAddress profile.AddressCoordinates) (int, error)
	DeleteAddress(id int, addressId int) error
}

type ConnectAuthService interface {
	SignUp(ctx context.Context, in *authProto.RegistrationRequest, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, in *authProto.Authorization, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, in *authProto.CSRF, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
}

type WrapperAuthorization interface {
	SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error)
	Login(login *authorization.Authorization) (*Utils2.Defense, error)
	Logout(CSRF string) (string, error)
}

type ConnectCartService interface {
	GetCart(ctx context.Context, in *cartProto.CartId, opts ...grpc.CallOption) (*cartProto.ResponseCartErrors, error)
	UpdateCart(ctx context.Context, in *cartProto.RequestCartDefault, opts ...grpc.CallOption) (*cartProto.ResponseCartErrors, error)
}

type WrapperCartServer interface {
	GetCart(id int) (*cart.ResponseCartErrors, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error)
}

type ConnectionMiddleware interface {
	CheckAccessUser(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.CheckAccess, error)
	NewCSRFUser(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
	GetIdByCookie(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.IdClientResponse, error)
}

type WrapperMiddleware interface {
	CheckAccess(cookie *Utils2.Defense) (bool, error)
	NewCSRF(cookie *Utils2.Defense) (string, error)
	GetIdByCookie(cookie *Utils2.Defense) (int, error)
}

type WrapperCart interface {
	GetCart(id int) (*cart.ResponseCartErrors, []cart.CastDishesErrs, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, []cart.CastDishesErrs, error)
	DeleteCart(id int) error
	GetPriceDelivery(id int) (int, error)
}

type Uploader interface {
	Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

type WrapperOrder interface {
	CreateOrder(id int, createOrder order.CreateOrder, addressId int, cart cart.ResponseCartErrors, courierId int) error
	GetOrders(id int) (*order.HistoryOrderArray, error)
	GetOrder(idClient int, idOrder int) (*order.ActiveOrder, error)
	UpdateStatusOrder(id int, status int) error
	CheckRun(id int) (bool, error)
	DeleteCart(id int) error
	GetCart(id int) (*cart.ResponseCartErrors, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
}
