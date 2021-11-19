package Utils

import "time"

type CreateOrder struct {
	MethodPay string             `json:"methodPay"`
	Address   AddressCoordinates `json:"address"`
	Comment   string             `json:"comment"`
}

type HistoryOrderArray struct {
	Orders []HistoryOrder `json:"orders"`
}

type HistoryOrder struct {
	Status     int                `json:"status"`
	Date       time.Time          `json:"date"`
	Address    AddressCoordinates `json:"address"`
	Restaurant HistoryResOrder    `json:"restaurant"`
	Cart       ResponseCartOrder  `json:"cart"`
}

type HistoryResOrder struct {
	Id      int                `json:"id"`
	Name    string             `json:"name"`
	Img     string             `json:"img"`
	Address AddressCoordinates `json:"address"`
}

type ResponseCartOrder struct {
	Dishes []DishesCartResponse `json:"dishes"`
	Cost   CostCartResponse     `json:"cost"`
}
