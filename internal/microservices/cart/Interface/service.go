package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/cart/proto"
	"context"
)

type CartManager interface {
	GetCart(ctx context.Context, id *proto.CartId) (*proto.ResponseCartErrors, error)
	UpdateCart(ctx context.Context, dishes *proto.RequestCartDefault) (*proto.ResponseCartErrors, error)
}
