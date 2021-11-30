package orm

import (
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
	"google.golang.org/grpc"
)

type WrapperRestaurantServerInterface interface {
	AllRestaurants() ([]restaurant.Restaurants, error)
	GetRestaurant(id int) (*restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error)
	CreateReview(id int, review restaurant.NewReview) error
	GetReview(id int) (*restaurant.ResReview, error)
	SearchRestaurant(search string) ([]restaurant.Restaurants, error)
	GetFavouriteRestaurants(id int) ([]restaurant.Restaurants, error)
	AddRestaurantInFavourite(idRestaurant int, idClient int) (bool, error)
}

type ConnectRestaurantServiceInterface interface {
	AllRestaurants(ctx context.Context, in *resProto.Empty, opts ...grpc.CallOption) (*resProto.Restaurants, error)
	GetRestaurant(ctx context.Context, in *resProto.RestaurantId, opts ...grpc.CallOption) (*resProto.RestaurantInfo, error)
	RestaurantDishes(ctx context.Context, in *resProto.DishInfo, opts ...grpc.CallOption) (*resProto.Dishes, error)
	CreateReview(ctx context.Context, in *resProto.NewReview, opts ...grpc.CallOption) (*resProto.Error, error)
	GetReview(ctx context.Context, in *resProto.RestaurantId, opts ...grpc.CallOption) (*resProto.ResReview, error)
	SearchRestaurant(ctx context.Context, in *resProto.SearchRestaurantText, opts ...grpc.CallOption) (*resProto.Restaurants, error)
	GetFavouriteRestaurants(ctx context.Context, clientId *resProto.UserId, opts ...grpc.CallOption) (*resProto.Restaurants, error)
	AddRestaurantInFavourite(ctx context.Context, restaurant *resProto.AddRestaurantInFavouriteRequest, opts ...grpc.CallOption) (*resProto.ResponseAddRestaurantInFavourite, error)
}

type Wrapper struct {
	Conn ConnectRestaurantServiceInterface
	Ctx  context.Context
}

func (r *Wrapper) AllRestaurants() ([]restaurant.Restaurants, error) {
	restaurants, err := r.Conn.AllRestaurants(r.Ctx, &resProto.Empty{})
	if err != nil {
		return nil, err
	}
	if restaurants.Error != "" {
		return nil, &errPkg.Errors{Alias: restaurants.Error}
	}
	return cast.CastRestaurantsProtoToRestaurants(restaurants.Restaurants), nil
}

func (r *Wrapper) GetRestaurant(id int) (*restaurant.RestaurantId, error) {
	var restaurantId *resProto.RestaurantId
	restaurantId = &resProto.RestaurantId{}
	restaurantId.Id = int64(id)
	rest, err := r.Conn.GetRestaurant(r.Ctx, restaurantId)
	if err != nil {
		return nil, err
	}
	if rest.Error != "" {
		return nil, &errPkg.Errors{Alias: rest.Error}
	}
	return cast.CastRestaurantInfoToRestaurantIdProto(rest), nil
}

func (r *Wrapper) RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error) {
	var info *resProto.DishInfo
	info = &resProto.DishInfo{}
	info.DishId = int64(dishId)
	info.RestaurantId = int64(restId)
	dishes, err := r.Conn.RestaurantDishes(r.Ctx, info)
	if err != nil {
		return nil, err
	}
	if dishes.Error != "" {
		return nil, &errPkg.Errors{Alias: dishes.Error}
	}
	return cast.CastDishesProtoToDishes(dishes), nil
}

func (r *Wrapper) CreateReview(id int, review restaurant.NewReview) error {
	createReview, err := r.Conn.CreateReview(r.Ctx, cast.CastNewReviewToNewReviewProto(review, id))
	if err != nil {
		return err
	}
	if createReview.Error != "" {
		return &errPkg.Errors{Alias: createReview.Error}
	}
	return &errPkg.Errors{
		Alias: createReview.Error,
	}
}

func (r *Wrapper) GetReview(id int) (*restaurant.ResReview, error) {
	var restId *resProto.RestaurantId
	restId = &resProto.RestaurantId{}
	restId.Id = int64(id)
	review, err := r.Conn.GetReview(r.Ctx, restId)
	if err != nil {
		return nil, err
	}
	if review.Error != "" {
		return nil, &errPkg.Errors{Alias: review.Error}
	}
	return cast.CastResReviewProtoToResReview(review), nil
}

func (r *Wrapper) SearchRestaurant(search string) ([]restaurant.Restaurants, error) {
	var searchText *resProto.SearchRestaurantText
	searchText = &resProto.SearchRestaurantText{}
	searchText.Text = search
	rest, err := r.Conn.SearchRestaurant(r.Ctx, searchText)
	if err != nil {
		return nil, err
	}
	if rest.Error != "" {
		return nil, &errPkg.Errors{Alias: rest.Error}
	}
	return cast.CastRestaurantsProtoToRestaurant(rest), nil
}

func (r *Wrapper) GetFavouriteRestaurants(id int) ([]restaurant.Restaurants, error) {
	restaurants, err := r.Conn.GetFavouriteRestaurants(r.Ctx, &resProto.UserId{Id: int64(id)})
	if err != nil {
		return nil, err
	}
	if restaurants.Error != "" {
		return nil, &errPkg.Errors{Alias: restaurants.Error}
	}
	return cast.CastRestaurantsProtoToRestaurants(restaurants.Restaurants), nil
}

func (r *Wrapper) AddRestaurantInFavourite(idRestaurant int, idClient int) (bool, error) {
	restaurants, err := r.Conn.AddRestaurantInFavourite(r.Ctx, &resProto.AddRestaurantInFavouriteRequest{IdRestaurant: int64(idRestaurant), IdClient: int64(idClient)})
	if err != nil {
		return false, err
	}
	if restaurants.Error != "" {
		return false, &errPkg.Errors{Alias: restaurants.Error}
	}
	return restaurants.Status, nil
}
