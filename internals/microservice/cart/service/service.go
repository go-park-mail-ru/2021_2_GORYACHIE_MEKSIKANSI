package service

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
	"context"
)

type CartManagerInterface interface {
	GetCart(ctx context.Context, id *proto.CartId) (*proto.ResponseCartErrors, error)
	UpdateCart(ctx context.Context, dishes *proto.RequestCartDefault) (*proto.ResponseCartErrors, error)
}

type CartManager struct {
	Application appPkg.CartInterface
}

func (c *CartManager) GetCart(ctx context.Context, id *proto.CartId) (*proto.ResponseCartErrors, error) {
	cart, err := c.Application.GetCart(int(id.Id))
	if err != nil {
		return &proto.ResponseCartErrors{Error: err.Error()}, nil
	}
	sendCart := CastResponseCartErrorsToResponseCartErrorsProto(cart)
	return sendCart, nil
}

func (c *CartManager) UpdateCart(ctx context.Context, dishes *proto.RequestCartDefault) (*proto.ResponseCartErrors, error) {
	cart, err := c.Application.UpdateCart(*CastRequestCartDefaultProtoToRequestCartDefault(dishes), int(dishes.ClientId))
	if err != nil {
		return &proto.ResponseCartErrors{Error: err.Error()}, nil
	}

	if cart == nil {
		return &proto.ResponseCartErrors{}, nil
	}

	sendCart := CastResponseCartErrorsToResponseCartErrorsProto(cart)
	return sendCart, nil
}
