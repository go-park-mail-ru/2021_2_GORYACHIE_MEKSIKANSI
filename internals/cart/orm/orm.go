//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/cart/orm WrapperCartServerInterface,ConnectCartServiceInterface,ConnectPromocodeServiceInterface
package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
	promoProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internals/util/cast"
	"context"
	"google.golang.org/grpc"
)

type WrapperCartServerInterface interface {
	GetCart(id int) (*cart.ResponseCartErrors, error)
	UpdateCart(dishes cart.RequestCartDefault, clientId int) (*cart.ResponseCartErrors, error)
	AddPromoCode(promoCode string, restaurantId int, clientId int) error
}

type ConnectCartServiceInterface interface {
	GetCart(ctx context.Context, in *cartProto.CartId, opts ...grpc.CallOption) (*cartProto.ResponseCartErrors, error)
	UpdateCart(ctx context.Context, in *cartProto.RequestCartDefault, opts ...grpc.CallOption) (*cartProto.ResponseCartErrors, error)
}

type ConnectPromocodeServiceInterface interface {
	AddPromoCode(ctx context.Context, in *promoProto.PromoCodeWithRestaurantIdAndClient, opts ...grpc.CallOption) (*promoProto.Error, error)
}

type Wrapper struct {
	ConnCart  ConnectCartServiceInterface
	ConnPromo ConnectPromocodeServiceInterface
	CtxCart   context.Context
	CtxPromo  context.Context
}

func (db *Wrapper) GetCart(id int) (*cart.ResponseCartErrors, error) {
	var a cartProto.CartId
	a.Id = int64(id)
	cartUser, err := db.ConnCart.GetCart(db.CtxCart, &a)
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
	cartUser, err := db.ConnCart.UpdateCart(db.CtxCart, &a)
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

func (db *Wrapper) AddPromoCode(promoCode string, restaurantId int, clientId int) error {
	result, err := db.ConnPromo.AddPromoCode(db.CtxPromo, &promoProto.PromoCodeWithRestaurantIdAndClient{
		PromoCode:  promoCode,
		Restaurant: int64(restaurantId),
		Client:     int64(clientId),
	})
	if err != nil {
		return err
	}
	if result.Error != "" {
		return &errPkg.Errors{Text: result.Error}
	}

	return nil
}
