package Errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *CheckError) CheckErrorRestaurant(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrRestaurantsNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusNotFound,
				Explain: ErrRestaurantsNotFound,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrRestaurantsNotFound)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case ErrRestaurantsScan:
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
			fmt.Printf("Console: %s\n", ErrRestaurantsScan)
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
		case ErrRestaurantNotFound, ErrCategoryRestaurantScan, ErrRestaurantsDishesNotSelect,
			ErrRestaurantDishesScan, ErrRestaurantDishesNotFound, ErrTagNotFound:
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

func (c *CheckError) CheckErrorRestaurantDishes(err error) (error, []byte, int) {
	if err != nil {

	}
	return nil, nil, HttpNil
}
