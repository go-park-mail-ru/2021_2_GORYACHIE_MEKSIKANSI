package MyError

import (
	"encoding/json"
	"net/http"
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
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Alias: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("%s, requestId: %d", RGetGeneralInfoRestaurantNotFound, c.RequestId)
			return &Errors{
					Alias: ErrCheck,
				},
				result, http.StatusOK

		case RGetRestaurantsRestaurantsNotScan, RGetRestaurantsRestaurantsNotSelect:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Alias: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Alias: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
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
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Alias: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Alias: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorRestaurantDishes(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case RGetDishesDishesNotFound, RGetStructDishesStructDishesNotSelect, RGetStructDishesStructDishesNotScan,
			RGetRadiosRadiosNotScan, RGetTagsCategoryNotSelect, RGetRadiosRadiosNotSelect, RGetTagsTagsNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Alias: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Alias: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}
