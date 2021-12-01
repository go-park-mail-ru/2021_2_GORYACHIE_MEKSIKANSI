package service

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant"
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/application"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/proto"
	"context"
)

type RestaurantManagerInterface interface {
	AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error)
	GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error)
	RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error)
	CreateReview(ctx context.Context, rev *resProto.NewReview) (*resProto.Error, error)
	GetReview(ctx context.Context, id *resProto.RestaurantClientId) (*resProto.ResReview, error)
	SearchRestaurant(ctx context.Context, search *resProto.SearchRestaurantText) (*resProto.Restaurants, error)
	GetFavoriteRestaurants(ctx context.Context, id *resProto.UserId) (*resProto.Restaurants, error)
	EditRestaurantInFavorite(ctx context.Context, ids *resProto.EditRestaurantInFavoriteRequest) (*resProto.ResponseEditRestaurantInFavorite, error)
}

type RestaurantManager struct {
	Application appPkg.RestaurantApplicationInterface
}

func (r *RestaurantManager) AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error) {
	restaurants, err := r.Application.AllRestaurants()
	if err != nil {
		return &resProto.Restaurants{Error: err.Error()}, nil
	}
	sendRestaurant := CastRestaurantsToRestaurantsProto(restaurants)
	return sendRestaurant, nil
}

func (r *RestaurantManager) GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error) {
	restaurant, err := r.Application.GetRestaurant(int(id.Id))
	if err != nil {
		return &resProto.RestaurantInfo{Error: err.Error()}, nil
	}
	sendRestaurant := CastRestaurantIdToRestaurantInfoProto(restaurant)
	return sendRestaurant, nil
}

func (r *RestaurantManager) RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error) {
	dishes, err := r.Application.RestaurantDishes(int(id.RestaurantId), int(id.DishId))
	if err != nil {
		return &resProto.Dishes{Error: err.Error()}, nil
	}
	dish := CastDishesToDishesProto(dishes)
	dish.Ingredients = CastIngredientsToIngredientsProto(dishes.Ingredient)
	dish.Radios = CastRadiosToRadiosProto(dishes.Radios)
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

func (r *RestaurantManager) GetReview(ctx context.Context, id *resProto.RestaurantClientId) (*resProto.ResReview, error) {
	review, err := r.Application.GetReview(int(id.IdRestaurant), int(id.IdClient))
	if err != nil {
		return &resProto.ResReview{Error: err.Error()}, nil
	}

	sendReview := CastResReviewToResReviewProto(review)
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
		searchResult.Restaurants = append(searchResult.Restaurants, CastRestaurantsToRestaurantProto(&restaurantInfo))
	}
	return searchResult, nil
}

func (r *RestaurantManager) GetFavoriteRestaurants(ctx context.Context, clientId *resProto.UserId) (*resProto.Restaurants, error) {
	restaurants, err := r.Application.GetFavoriteRestaurants(int(clientId.Id))
	if err != nil {
		return &resProto.Restaurants{Error: err.Error()}, nil
	}
	return CastRestaurantsToRestaurantsProto(restaurants), nil
}

func (r *RestaurantManager) EditRestaurantInFavorite(ctx context.Context, restaurant *resProto.EditRestaurantInFavoriteRequest) (*resProto.ResponseEditRestaurantInFavorite, error) {
	status, err := r.Application.EditRestaurantInFavorite(int(restaurant.IdRestaurant), int(restaurant.IdClient))
	if err != nil {
		return &resProto.ResponseEditRestaurantInFavorite{Error: err.Error()}, nil
	}
	var result resProto.ResponseEditRestaurantInFavorite
	result.Status = status
	return &result, nil
}
