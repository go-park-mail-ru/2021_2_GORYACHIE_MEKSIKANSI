//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internal/cart/orm WrapperCartServerInterface,ConnectCartServiceInterface
package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
	"google.golang.org/grpc"
)

type WrapperCartServerInterface interface {
	GetCart(id int) (*cart.ResponseCartErrors, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error)
}

type ConnectCartServiceInterface interface {
	GetCart(ctx context.Context, in *cartProto.CartId, opts ...grpc.CallOption) (*cartProto.ResponseCartErrors, error)
	UpdateCart(ctx context.Context, in *cartProto.RequestCartDefault, opts ...grpc.CallOption) (*cartProto.ResponseCartErrors, error)
}

type Wrapper struct {
	Conn ConnectCartServiceInterface
	Ctx  context.Context
}

func (db *Wrapper) GetCart(id int) (*cart.ResponseCartErrors, error) {
	var a cartProto.CartId
	a.Id = int64(id)
	cart, err := db.Conn.GetCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	if cart.Error != "" {
		return nil, &errPkg.Errors{Text: cart.Error}
	}
	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cart), nil
}

func (db *Wrapper) UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error) {
	var a cartProto.RequestCartDefault
	a = *cast.CastRequestCartDefaultToRequestCartDefaultProto(&dishes)
	a.ClientId = int64(clientId)
	cart, err := db.Conn.UpdateCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	if cart.Error != "" {
		return nil, &errPkg.Errors{Text: cart.Error}
	}

	if cart.Restaurant == nil {
		return nil, nil
	}

	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cart), nil
}
