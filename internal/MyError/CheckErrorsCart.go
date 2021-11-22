package MyError

import (
	"encoding/json"
	"net/http"
)

func (c *CheckError) CheckErrorGetCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case CGetCartDishesNotFound:
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

		case RGetGeneralInfoRestaurantNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusNotFound,
				Explain: ErrCartNull,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Alias: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Alias: ErrCheck,
				},
				result, http.StatusOK
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorUpdateCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case CDeleteCartCartNotDelete, CDeleteCartStructureFoodNotDelete, CDeleteCartRadiosFoodNotDelete,
			CUpdateCartCartNotInsert, CUpdateCartStructFoodStructureFoodNotInsert, CUpdateCartRadiosRadiosNotInsert,
			CGetPriceDeliveryPriceNotFound, CGetPriceDeliveryPriceNotScan, CUpdateCartCartNotFound,
			CUpdateCartCartNotScan, CUpdateCartStructureFoodStructureFoodNotSelect,
			CUpdateCartStructRadiosStructRadiosNotSelect, CUpdateCartTransactionNotCreate, CUpdateCartNotCommit:
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
