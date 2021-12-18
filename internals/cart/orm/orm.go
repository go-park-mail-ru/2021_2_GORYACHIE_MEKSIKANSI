//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/cart/orm WrapperCartServerInterface,ConnectCartServiceInterface
package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internals/util/cast"
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
	cartUser, err := db.Conn.GetCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	if cartUser.Error != "" {
		return nil, &errPkg.Errors{Text: cartUser.Error}
	}
	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cartUser), nil
}

func (db *Wrapper) UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error) {
	var a cartProto.RequestCartDefault
	a = *cast.CastRequestCartDefaultToRequestCartDefaultProto(&dishes)
	a.ClientId = int64(clientId)
	cartUser, err := db.Conn.UpdateCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	if cartUser.Error != "" {
		return nil, &errPkg.Errors{Text: cartUser.Error}
	}

	if cartUser.Restaurant == nil {
		return nil, nil
	}

	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cartUser), nil
}
