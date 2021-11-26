package Interface

import (
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/proto"
	"context"
)

type RestaurantManager interface {
	AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error)
	GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error)
	RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error)
	CreateReview(ctx context.Context, rev *resProto.NewReview) (*resProto.Error, error)
	GetReview(ctx context.Context, id *resProto.RestaurantId) (*resProto.ResReview, error)
	SearchRestaurant(ctx context.Context, search *resProto.SearchRestaurantText) (*resProto.Restaurants, error)
}
