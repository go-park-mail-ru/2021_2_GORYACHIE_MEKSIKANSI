package service

import(
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/Interface"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/proto"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
)

type RestaurantManager struct {
	Application Interface.RestaurantApplication
}

func (r *RestaurantManager) AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error) {
	restaurants, err := r.Application.AllRestaurants()
	if err != nil {
		return nil, err
	}

	return cast.CastRestaurantsToRestaurantsProto(restaurants), nil
}

func (r *RestaurantManager) GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error) {
	restaurant, err := r.Application.GetRestaurant(int(id.Id))
	if err != nil {
		return nil, err
	}
	return cast.CastRestaurantIdToRestaurantInfoProto(restaurant), nil
}

func (r *RestaurantManager) RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error) {
	dishes, err := r.Application.RestaurantDishes(int(id.RestaurantId), int(id.DishId))
	if err != nil {
		return nil, err
	}
	dish := cast.CastDishesToDishesProto(dishes)
	dish.Ingredients = cast.CastIngredientsToIngredientsProto(dishes.Ingredient)
	dish.Radios = cast.CastRadiosToRadiosProto(dishes.Radios)
	return dish, nil
}

func (r *RestaurantManager) CreateReview(ctx context.Context, rev *resProto.NewReview) (*resProto.Error, error) {
	var rest restaurant.RestaurantId
	rest.Id = int(rev.Restaurant.Id)
	err := r.Application.CreateReview(int(rev.Id), restaurant.NewReview{
		Rate: int(rev.Rate),
		Restaurant: rest,
		Text: rev.Text,
	})
	if err != nil {
		return nil, err
	}
	return &resProto.Error{}, nil
}

func (r *RestaurantManager) GetReview(ctx context.Context, id *resProto.RestaurantId) (*resProto.ResReview, error) {
	review, err := r.Application.GetReview(int(id.Id))
	if err != nil {
		return nil, err
	}

	return cast.CastResReviewToResReviewProto(review), nil
}

func (r *RestaurantManager) SearchRestaurant(ctx context.Context, search *resProto.SearchRestaurantText) (*resProto.Restaurants, error) {
	searchRestaurant, err := r.Application.SearchRestaurant(search.Text)
	if err != nil {
		return nil, err
	}
	var searchResult *resProto.Restaurants
	for _, restaurantInfo := range searchRestaurant {
		searchResult.Restaurants = append(searchResult.Restaurants, cast.CastRestaurantsToRestaurantProto(&restaurantInfo))
	}
	return searchResult, nil
}

