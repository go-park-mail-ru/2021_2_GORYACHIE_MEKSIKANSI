//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/orm WrapperRestaurantServerInterface,ConnectRestaurantServiceInterface
package orm

import (
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internals/util/cast"
	"context"
	"google.golang.org/grpc"
)

type WrapperRestaurantServerInterface interface {
	AllRestaurants() (*restaurant.AllRestaurantsPromo, error)
	RecommendedRestaurants() (*restaurant.AllRestaurants, error)
	GetRestaurant(id int, idClient int) (*restaurant.RestaurantId, error)
	RestaurantDishes(restId int, dishId int) (*restaurant.Dishes, error)
	CreateReview(id int, review restaurant.NewReview) error
	GetReview(idRestaurant int, idClient int) (*restaurant.ResReview, error)
	SearchRestaurant(search string) ([]restaurant.Restaurants, error)
	GetFavoriteRestaurants(id int) ([]restaurant.Restaurants, error)
	EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error)
	DeleteDish(idDish int) error
}

type ConnectRestaurantServiceInterface interface {
	AllRestaurants(ctx context.Context, in *resProto.Empty, opts ...grpc.CallOption) (*resProto.RestaurantsTagsPromo, error)
	GetRecommendedRestaurants(ctx context.Context, in *resProto.Empty, opts ...grpc.CallOption) (*resProto.RecommendedRestaurants, error)
	GetRestaurant(ctx context.Context, in *resProto.RestaurantId, opts ...grpc.CallOption) (*resProto.RestaurantInfo, error)
	RestaurantDishes(ctx context.Context, in *resProto.DishInfo, opts ...grpc.CallOption) (*resProto.Dishes, error)
	CreateReview(ctx context.Context, in *resProto.NewReview, opts ...grpc.CallOption) (*resProto.Error, error)
	GetReview(ctx context.Context, in *resProto.RestaurantClientId, opts ...grpc.CallOption) (*resProto.ResReview, error)
	SearchRestaurant(ctx context.Context, in *resProto.SearchRestaurantText, opts ...grpc.CallOption) (*resProto.Restaurants, error)
	GetFavoriteRestaurants(ctx context.Context, clientId *resProto.UserId, opts ...grpc.CallOption) (*resProto.Restaurants, error)
	EditRestaurantInFavorite(ctx context.Context, restaurant *resProto.EditRestaurantInFavoriteRequest, opts ...grpc.CallOption) (*resProto.ResponseEditRestaurantInFavorite, error)
	DeleteDish(ctx context.Context, restaurant *resProto.DishId, opts ...grpc.CallOption) (*resProto.Error, error)
}

type Wrapper struct {
	Conn ConnectRestaurantServiceInterface
	Ctx  context.Context
}

func (r *Wrapper) AllRestaurants() (*restaurant.AllRestaurantsPromo, error) {
	restaurants, err := r.Conn.AllRestaurants(r.Ctx, &resProto.Empty{})
	if err != nil {
		return nil, err
	}
	if restaurants.Error != "" {
		return nil, &errPkg.Errors{Text: restaurants.Error}
	}
	return cast.CastRestaurantsTagsProtoToAllRestaurants(restaurants), nil
}

func (r *Wrapper) RecommendedRestaurants() (*restaurant.AllRestaurants, error) {
	restaurants, err := r.Conn.GetRecommendedRestaurants(r.Ctx, &resProto.Empty{})
	if err != nil {
		return nil, err
	}
	if restaurants.Error != "" {
		return nil, &errPkg.Errors{Text: restaurants.Error}
	}
	return cast.CastRecommendedRestaurantsProtoToAllRestaurants(restaurants), nil
}

func (r *Wrapper) GetRestaurant(id int, idClient int) (*restaurant.RestaurantId, error) {
	var restaurantId *resProto.RestaurantId
	restaurantId = &resProto.RestaurantId{}
	restaurantId.Id = int64(id)
	restaurantId.IdClient = int64(idClient)
	rest, err := r.Conn.GetRestaurant(r.Ctx, restaurantId)
	if err != nil {
		return nil, err
	}
	if rest.Error != "" {
		return nil, &errPkg.Errors{Text: rest.Error}
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
		return nil, &errPkg.Errors{Text: dishes.Error}
	}
	return cast.CastDishesProtoToDishes(dishes), nil
}

func (r *Wrapper) CreateReview(id int, review restaurant.NewReview) error {
	createReview, err := r.Conn.CreateReview(r.Ctx, cast.CastNewReviewToNewReviewProto(review, id))
	if err != nil {
		return err
	}
	if createReview.Error != "" {
		return &errPkg.Errors{Text: createReview.Error}
	}
	return nil
}

func (r *Wrapper) GetReview(id int, idClient int) (*restaurant.ResReview, error) {
	var restId *resProto.RestaurantClientId
	restId = &resProto.RestaurantClientId{}
	restId.IdRestaurant = int64(id)
	restId.IdClient = int64(idClient)
	review, err := r.Conn.GetReview(r.Ctx, restId)
	if err != nil {
		return nil, err
	}
	if review.Error != "" {
		return nil, &errPkg.Errors{Text: review.Error}
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
		return nil, &errPkg.Errors{Text: rest.Error}
	}
	return cast.CastRestaurantsProtoToRestaurant(rest), nil
}

func (r *Wrapper) GetFavoriteRestaurants(id int) ([]restaurant.Restaurants, error) {
	restaurants, err := r.Conn.GetFavoriteRestaurants(r.Ctx, &resProto.UserId{Id: int64(id)})
	if err != nil {
		return nil, err
	}
	if restaurants.Error != "" {
		return nil, &errPkg.Errors{Text: restaurants.Error}
	}
	return cast.CastRestaurantsProtoToRestaurants(restaurants.Restaurants), nil
}

func (r *Wrapper) EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error) {
	restaurants, err := r.Conn.EditRestaurantInFavorite(r.Ctx, &resProto.EditRestaurantInFavoriteRequest{IdRestaurant: int64(idRestaurant), IdClient: int64(idClient)})
	if err != nil {
		return false, err
	}
	if restaurants.Error != "" {
		return false, &errPkg.Errors{Text: restaurants.Error}
	}
	return restaurants.Status, nil
}

func (r *Wrapper) DeleteDish(idDish int) error {
	result, err := r.Conn.DeleteDish(r.Ctx, &resProto.DishesId{Id: int64(idDish)})
	if err != nil {
		return err
	}
	if result.Error != "" {
		return &errPkg.Errors{Text: result.Error}
	}
	return nil
}
