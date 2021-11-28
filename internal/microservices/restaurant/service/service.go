package service

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant/application"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant/proto"
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	castPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
)

type RestaurantManagerInterface interface {
	AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error)
	GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error)
	RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error)
	CreateReview(ctx context.Context, rev *resProto.NewReview) (*resProto.Error, error)
	GetReview(ctx context.Context, id *resProto.RestaurantId) (*resProto.ResReview, error)
	SearchRestaurant(ctx context.Context, search *resProto.SearchRestaurantText) (*resProto.Restaurants, error)
}

type RestaurantManager struct {
	Application appPkg.RestaurantApplicationInterface
}

func (r *RestaurantManager) AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error) {
	restaurants, err := r.Application.AllRestaurants()
	if err != nil {
		return &resProto.Restaurants{Error: err.Error()}, nil
	}
	sendRestaurant := castPkg.CastRestaurantsToRestaurantsProto(restaurants)
	return sendRestaurant, nil
}

func (r *RestaurantManager) GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error) {
	restaurant, err := r.Application.GetRestaurant(int(id.Id))
	if err != nil {
		return &resProto.RestaurantInfo{Error: err.Error()}, nil
	}
	sendRestaurant := castPkg.CastRestaurantIdToRestaurantInfoProto(restaurant)
	return sendRestaurant, nil
}

func (r *RestaurantManager) RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error) {
	dishes, err := r.Application.RestaurantDishes(int(id.RestaurantId), int(id.DishId))
	if err != nil {
		return &resProto.Dishes{Error: err.Error()}, nil
	}
	dish := castPkg.CastDishesToDishesProto(dishes)
	dish.Ingredients = castPkg.CastIngredientsToIngredientsProto(dishes.Ingredient)
	dish.Radios = castPkg.CastRadiosToRadiosProto(dishes.Radios)
	return dish, nil
}

func (r *RestaurantManager) CreateReview(ctx context.Context, rev *resProto.NewReview) (*resProto.Error, error) {
	var rest resPkg.RestaurantId
	rest.Id = int(rev.Restaurant.Id)
	err := r.Application.CreateReview(int(rev.Id), resPkg.NewReview{
		Rate:       int(rev.Rate),
		Restaurant: rest,
		Text:       rev.Text,
	})
	if err != nil {
		return &resProto.Error{Error: err.Error()}, nil
	}
	return &resProto.Error{}, nil
}

func (r *RestaurantManager) GetReview(ctx context.Context, id *resProto.RestaurantId) (*resProto.ResReview, error) {
	review, err := r.Application.GetReview(int(id.Id))
	if err != nil {
		return &resProto.ResReview{Error: err.Error()}, nil
	}

	sendReview := castPkg.CastResReviewToResReviewProto(review)
	return sendReview, nil
}

func (r *RestaurantManager) SearchRestaurant(ctx context.Context, search *resProto.SearchRestaurantText) (*resProto.Restaurants, error) {
	searchRestaurant, err := r.Application.SearchRestaurant(search.Text)
	if err != nil {
		return &resProto.Restaurants{Error: err.Error()}, nil
	}
	var searchResult *resProto.Restaurants
	searchResult = &resProto.Restaurants{}
	for _, restaurantInfo := range searchRestaurant {
		searchResult.Restaurants = append(searchResult.Restaurants, castPkg.CastRestaurantsToRestaurantProto(&restaurantInfo))
	}
	return searchResult, nil
}
