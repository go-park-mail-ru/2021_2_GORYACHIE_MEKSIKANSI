package Service

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
)

type CartManager struct {
	Application Interface.CartApplication
}

func (c *CartManager) GetCart(ctx context.Context, id *proto.CartId) (*proto.ResponseCartErrors, error) {
	cart, err := c.Application.GetCart(int(id.Id))
	if err != nil {
		return &proto.ResponseCartErrors{Error: err.Error()}, nil
	}
	sendCart := cast.CastResponseCartErrorsToResponseCartErrorsProto(cart)
	return sendCart, nil
}

func (c *CartManager) UpdateCart(ctx context.Context, dishes *proto.RequestCartDefault) (*proto.ResponseCartErrors, error) {
	cart, err := c.Application.UpdateCart(*cast.CastRequestCartDefaultProtoToRequestCartDefault(dishes), int(dishes.ClientId))
	if err != nil {
		return &proto.ResponseCartErrors{Error: err.Error()}, nil
	}

	if cart == nil {
		return &proto.ResponseCartErrors{}, nil
	}

	sendCart := cast.CastResponseCartErrorsToResponseCartErrorsProto(cart)
	sendCart.Error = err.Error()
	return sendCart, nil
}
