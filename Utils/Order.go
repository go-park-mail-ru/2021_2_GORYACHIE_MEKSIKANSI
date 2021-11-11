package Utils

import "time"

type CreateOrder struct {
	Address AddressCoordinates `json:"address"`
	Comment string             `json:"comment"`
}

type HistoryOrderArray struct {
	Orders []HistoryOrder `json:"orders"`
}

type HistoryOrder struct {
	Status     int                `json:"status"`
	Date       time.Time          `json:"date"`
	Address    AddressCoordinates `json:"address"`
	Restaurant HistoryResOrder    `json:"restaurant"`
	Cart       ResponseCartErrors `json:"cart"`
}

type HistoryResOrder struct {
	Id      int                `json:"id"`
	Name    string             `json:"name"`
	Img     string             `json:"img"`
	Address AddressCoordinates `json:"address"`
}
