//go:generate mockgen -destination=mocks/service.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/application PromocodeApplicationInterface
package service

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
	"context"
)

type PromocodeManagerInterface interface {
	GetTypePromoCode(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.TypePromoCodeResponse, error)
	ActiveFreeDelivery(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.FreeDeliveryResponse, error)
	ActiveCostForFreeDish(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.FreeDishResponse, error)
	ActiveCostForSale(ctx context.Context, promoCode *proto.PromoCodeWithAmount) (*proto.NewCostResponse, error)
	ActiveTimeForSale(ctx context.Context, promoCode *proto.PromoCodeWithAmount) (*proto.NewCostResponse, error)
	AddPromoCode(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantIdAndClient) (*proto.Error, error)
	GetPromoCode(ctx context.Context, promoCode *proto.ClientId) (*proto.PromoCodeText, error)
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

func (pm *PromocodeManager) ActiveFreeDelivery(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantId) (*proto.FreeDeliveryResponse, error) {
	result, err := pm.Application.ActiveFreeDelivery(promoCode.PromoCode, int(promoCode.Restaurant))
	if err != nil {
		return &proto.FreeDeliveryResponse{Error: err.Error()}, nil
	}
	return &proto.FreeDeliveryResponse{Have: result}, nil
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

func (pm *PromocodeManager) AddPromoCode(ctx context.Context, promoCode *proto.PromoCodeWithRestaurantIdAndClient) (*proto.Error, error) {
	err := pm.Application.AddPromoCode(promoCode.PromoCode, int(promoCode.Restaurant), int(promoCode.Client))
	if err != nil {
		return &proto.Error{Error: err.Error()}, nil
	}
	return &proto.Error{}, nil
}

func (pm *PromocodeManager) GetPromoCode(ctx context.Context, promoCode *proto.ClientId) (*proto.PromoCodeText, error) {
	promoCodeText, err := pm.Application.GetPromoCode(int(promoCode.ClientId))
	if err != nil {
		return &proto.PromoCodeText{Error: err.Error()}, nil
	}
	return &proto.PromoCodeText{
		PromoCodeText: promoCodeText,
	}, nil
}
