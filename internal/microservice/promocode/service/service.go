package service

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/promocode/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/promocode/proto"
	"context"
)

type PromocodeManagerInterface interface {
	GetTypePromoCode(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.TypePromoCodeResponse, error)
	ActiveCostForFreeDelivery(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.NewCostResponse, error)
	ActiveCostForFreeDish(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.FreeDishResponse, error)
	ActiveCostForSale(ctx context.Context, promoCode *proto.PromoCodeWithAmount) (*proto.NewCostResponse, error)
	ActiveTimeForSale(ctx context.Context, promoCode *proto.PromoCodeWithAmount) (*proto.NewCostResponse, error)
}

type PromocodeManager struct {
	Application appPkg.PromocodeApplicationInterface
}

func (pm *PromocodeManager) GetTypePromoCode(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.TypePromoCodeResponse, error) {
	result, err := pm.Application.GetTypePromoCode(promoCode.PromoCode, int(promoCode.Restaurant))
	if err != nil {
		return &proto.TypePromoCodeResponse{Error: err.Error()}, nil
	}
	return &proto.TypePromoCodeResponse{Type: int64(result)}, nil
}

func (pm *PromocodeManager) ActiveCostForFreeDelivery(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.NewCostResponse, error) {
	result, err := pm.Application.ActiveCostForFreeDelivery(promoCode.PromoCode, int(promoCode.Restaurant))
	if err != nil {
		return &proto.NewCostResponse{Error: err.Error()}, nil
	}
	return &proto.NewCostResponse{Cost: int64(result)}, nil
}

func (pm *PromocodeManager) ActiveCostForFreeDish(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.FreeDishResponse, error) {
	cost, dishId, err := pm.Application.ActiveCostForFreeDish(promoCode.PromoCode, int(promoCode.Restaurant))
	if err != nil {
		return &proto.FreeDishResponse{Error: err.Error()}, nil
	}
	return &proto.FreeDishResponse{Cost: int64(cost), DishId: int64(dishId)}, nil
}

func (pm *PromocodeManager) ActiveCostForSale(ctx context.Context, promoCode *proto.PromoCodeWithAmount) (*proto.NewCostResponse, error) {
	result, err := pm.Application.ActiveCostForSale(promoCode.PromoCode, int(promoCode.Amount), int(promoCode.Restaurant))
	if err != nil {
		return &proto.NewCostResponse{Error: err.Error()}, nil
	}
	return &proto.NewCostResponse{Cost: int64(result)}, nil
}

func (pm *PromocodeManager) ActiveTimeForSale(ctx context.Context, promoCode *proto.PromoCodeWithAmount) (*proto.NewCostResponse, error) {
	result, err := pm.Application.ActiveTimeForSale(promoCode.PromoCode, int(promoCode.Amount), int(promoCode.Restaurant))
	if err != nil {
		return &proto.NewCostResponse{Error: err.Error()}, nil
	}
	return &proto.NewCostResponse{Cost: int64(result)}, nil
}
