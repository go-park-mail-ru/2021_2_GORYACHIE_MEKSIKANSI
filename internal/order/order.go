package order

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/profile"
)

type CreateOrder struct {
	MethodPay string                     `json:"methodPay"`
	Address   profile.AddressCoordinates `json:"address"`
	Comment   string                     `json:"comment"`
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
	Id         int                        `json:"id"`
	Status     int                        `json:"status"`
	Date       string                     `json:"date"`
	Time       string                     `json:"time"`
	Address    profile.AddressCoordinates `json:"address"`
	Restaurant HistoryResOrder            `json:"restaurant"`
	Cart       ResponseCartOrder          `json:"cart"`
}
