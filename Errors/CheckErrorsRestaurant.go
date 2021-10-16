package Errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CheckErrorRestaurant(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrRestaurantsNotFound:
			result, errMarshal:= json.Marshal(ResultError{
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
		case ErrRestaurantScan:
			result, errMarshal:= json.Marshal(ResultError{
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

func CheckErrorRestaurantDishes(err error) (error, []byte, int) {
	if err != nil {

	}
	return nil, nil, HttpNil
}


