package Errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CheckErrorGetCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error(){
		case GetCartRestaurantNotScan, GetCartCartNotFound, GetCartCartNotScan,
			GetCartDishesNotFound, GetCartDishesNotScan, GetCartRestaurantNotSelect,
			GetCartCheckboxNotScan, GetCartRadiosNotSelect, GetCartRadiosNotScan, GetCartStructRadiosNowScan:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", err.Error())
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		case GetCartRestaurantNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusNotFound,
				Explain: ErrCartNull,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", err.Error())
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorUpdateCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error(){
		case CartNotDelete, StructureFoodNotDelete, CartRadiosFoodNotDelete, UpdateCartCartNotInsert,
		UpdateCartStructureFoodNotInsert, UpdateCartRadiosNotInsert, GetPriceDeliveryNotFound, GetPriceDeliveryNotScan,
			UpdateCartCartNotSelect:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", err.Error())
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}
