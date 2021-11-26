package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectCartService
	Ctx context.Context
}

func (db *Wrapper) GetCart(id int) (*Cart.ResponseCartErrors, error) {
	var a cartProto.CartId
	a.Id = int64(id)
	cart, err := db.Conn.GetCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cart), nil
}

func (db *Wrapper) UpdateCart(dishes Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, error) {
	var a cartProto.RequestCartDefault
	a.ClientId = int64(clientId)
	cart, err := db.Conn.UpdateCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	return cast.CastResponseCartErrorsProtoToResponseCartErrors(cart), nil
}
