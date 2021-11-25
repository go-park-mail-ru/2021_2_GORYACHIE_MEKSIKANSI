package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectCartService
	Ctx context.Context
}

func (db *Wrapper) CalculatePriceDelivery(id int) (int, error) {
	// TODO: add convert func
	var a cartProto.CalculatePriceDeliveryId
	a.Id = int64(id)
	delivery, err := db.Conn.CalculatePriceDelivery(db.Ctx, &a)
	if err != nil {
		return 0, err
	}
	return int(delivery.Id), nil
}

func (db *Wrapper) GetCart(id int) (*Cart.ResponseCartErrors, error) {
	// TODO: add convert func
	var a cartProto.CartId
	a.Id = int64(id)
	_, err := db.Conn.GetCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (db *Wrapper) UpdateCart(dishes Cart.RequestCartDefault, clientId int) (*Cart.ResponseCartErrors, error) {
	// TODO: add convert func
	var a cartProto.RequestCartDefault
	a.ClientId = int64(clientId)
	_, err := db.Conn.UpdateCart(db.Ctx, &a)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (db *Wrapper) DeleteCart(id int) error  {
	// TODO: add convert func
	var a cartProto.DeleteCartId
	a.Id = int64(id)
	_, err := db.Conn.DeleteCart(db.Ctx, &a)
	if err != nil {
		return err
	}
	return nil
}
