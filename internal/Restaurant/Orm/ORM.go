package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectRestaurantService
	Ctx  context.Context
}

func (r *Wrapper) AllRestaurants() ([]Restaurant.Restaurants, error) {
	restaurants, err := r.Conn.AllRestaurants(r.Ctx, &resProto.Empty{})
	if err != nil {
		return nil, err
	}
	if restaurants.Error != "" {
		return nil, &errPkg.Errors{Alias: restaurants.Error}
	}
	return cast.CastRestaurantsProtoToRestaurants(restaurants.Restaurants), nil
}

func (r *Wrapper) GetRestaurant(id int) (*Restaurant.RestaurantId, error) {
	var restaurantId *resProto.RestaurantId
	restaurantId = &resProto.RestaurantId{}
	restaurantId.Id = int64(id)
	restaurant, err := r.Conn.GetRestaurant(r.Ctx, restaurantId)
	if err != nil {
		return nil, err
	}
	if restaurant.Error != "" {
		return nil, &errPkg.Errors{Alias: restaurant.Error}
	}
	return cast.CastRestaurantInfoToRestaurantIdProto(restaurant), nil
}

func (r *Wrapper) RestaurantDishes(restId int, dishId int) (*Restaurant.Dishes, error) {
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

func (r *Wrapper) CreateReview(id int, review Restaurant.NewReview) error {
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

func (r *Wrapper) GetReview(id int) (*Restaurant.ResReview, error) {
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

func (r *Wrapper) SearchRestaurant(search string) ([]Restaurant.Restaurants, error) {
	var searchText *resProto.SearchRestaurantText
	searchText = &resProto.SearchRestaurantText{}
	searchText.Text = search
	restaurant, err := r.Conn.SearchRestaurant(r.Ctx, searchText)
	if err != nil {
		return nil, err
	}
	if restaurant.Error != "" {
		return nil, &errPkg.Errors{Alias: restaurant.Error}
	}
	return cast.CastRestaurantsProtoToRestaurant(restaurant), nil
}
