package Errors

import (
	"encoding/json"
	"net/http"
	"time"
)

func (c *CheckError) CheckErrorRestaurant(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case RGetRestaurantsRestaurantsNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusNotFound,
				Explain: RGetRestaurantsRestaurantsNotFound,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", RGetGeneralInfoRestaurantNotFound, *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case RGetRestaurantsRestaurantsNotScan:
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
			c.LoggerErrWarn.Errorf("error: %s, requestId: %d", RGetRestaurantsRestaurantsNotScan, *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func (c *CheckError) CheckErrorRestaurantId(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case RGetGeneralInfoRestaurantNotFound, RGetTagsCategoryRestaurantNotScan, RGetMenuDishesNotSelect,
			RGetDishesRestaurantDishesNotScan, RGetMenuDishesNotFound, RGetTagsTagsNotFound:
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

func (c *CheckError) CheckErrorRestaurantDishes(err error) (error, []byte, int) {
	if err != nil {

	}
	return nil, nil, HttpNil
}
