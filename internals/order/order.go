//go:generate easyjson -no_std_marshalers order.go
package order

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
)

type ResponseOrder struct {
	Order interface{} `json:"order"`
}

//easyjson:json
type CreateOrder struct {
	MethodPay string                     `json:"methodPay"`
	Address   profile.AddressCoordinates `json:"address"`
	Comment   string                     `json:"comment"`
}

type CreateOrderId struct {
	Id int `json:"id"`
}

type HistoryOrderArray struct {
	Orders []HistoryOrder `json:"orders"`
}

type HistoryOrder struct {
	Id         int                        `json:"id"`
	Status     int                        `json:"status"`
	Date       string                     `json:"date"`
	Time       string                     `json:"time"`
	Address    profile.AddressCoordinates `json:"address"`
	Restaurant HistoryResOrder            `json:"restaurant"`
	Cart       ResponseCartOrder          `json:"cart"`
}

type HistoryResOrder struct {
	Id      int                        `json:"id"`
	Name    string                     `json:"name"`
	Img     string                     `json:"img"`
	Address profile.AddressCoordinates `json:"address"`
}

type ResponseCartOrder struct {
	Dishes []cart.DishesCartResponse `json:"dishes"`
	Cost   cart.CostCartResponse     `json:"cost"`
}

type ActiveOrder struct {
	Id           int                        `json:"id"`
	Status       int                        `json:"status"`
	Date         string                     `json:"date"`
	Time         string                     `json:"time"`
	TimeDelivery string                     `json:"time_delivery"`
	TextCancel   string                     `json:"text_cancel"`
	Address      profile.AddressCoordinates `json:"address"`
	Restaurant   HistoryResOrder            `json:"restaurant"`
	Cart         ResponseCartOrder          `json:"cart"`
}

type CancelOrderHost struct {
	TextCancel string `json:"text_cancel"`
}

func ConvertCreateOrderIdToOrderResponse(id int) ResponseOrder {
	return ResponseOrder{Order: CreateOrderId{Id: id}}
}
