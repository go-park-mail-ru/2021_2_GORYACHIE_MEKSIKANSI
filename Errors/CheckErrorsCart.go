package Errors

import (
	"encoding/json"
	"net/http"
	"time"
)

func (c *CheckError) CheckErrorGetCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case GetCartRestaurantNotScan, GetCartCartNotFound, GetCartCartNotScan,
			GetCartDishesNotFound, GetCartDishesNotScan, GetCartRestaurantNotSelect,
			GetCartCheckboxNotScan, GetCartRadiosNotSelect, GetCartRadiosNotScan, GetCartStructRadiosNowScan:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Errorf("error: %s, requestId: %d", err.Error(), *c.RequestId)
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
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", err.Error(), *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		}
	}
	return nil, nil, HttpNil
}

func (c *CheckError) CheckErrorUpdateCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case CartNotDelete, StructureFoodNotDelete, CartRadiosFoodNotDelete, UpdateCartCartNotInsert,
			UpdateCartStructureFoodNotInsert, UpdateCartRadiosNotInsert, GetPriceDeliveryNotFound, GetPriceDeliveryNotScan,
			UpdateCartCartNotFound, UpdateCartCartNotScan, UpdateCartStructureNotSelect, UpdateCartStructRadiosNotSelect:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Errorf("error: %s, requestId: %d", err.Error(), *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}
