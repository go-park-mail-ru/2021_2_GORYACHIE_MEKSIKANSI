package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/cart/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectCartService
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
		return nil, &errPkg.Errors{Alias: cart.Error}
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
		return nil, &errPkg.Errors{Alias: cart.Error}
	}

	if cart.Restaurant == nil {
		return nil, nil
	}

	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cart), nil
}
