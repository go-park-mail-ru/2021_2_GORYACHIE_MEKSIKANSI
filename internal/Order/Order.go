package Order

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Profile"
	"time"
)

type CreateOrder struct {
	MethodPay string                     `json:"methodPay"`
	Address   Profile.AddressCoordinates `json:"address"`
	Comment   string                     `json:"comment"`
}

type HistoryOrderArray struct {
	Orders []HistoryOrder `json:"orders"`
}

type HistoryOrder struct {
	Status     int                        `json:"status"`
	Date       time.Time                  `json:"date"`
	Address    Profile.AddressCoordinates `json:"address"`
	Restaurant HistoryResOrder            `json:"restaurant"`
	Cart       ResponseCartOrder          `json:"cart"`
}

type HistoryResOrder struct {
	Id      int                        `json:"id"`
	Name    string                     `json:"name"`
	Img     string                     `json:"img"`
	Address Profile.AddressCoordinates `json:"address"`
}

type ResponseCartOrder struct {
	Dishes []Cart.DishesCartResponse `json:"dishes"`
	Cost   Cart.CostCartResponse     `json:"cost"`
}
