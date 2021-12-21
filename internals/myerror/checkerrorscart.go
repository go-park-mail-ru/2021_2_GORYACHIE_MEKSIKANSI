package myerror

import (
	"github.com/mailru/easyjson"
	"net/http"
)

func (c *CheckError) CheckErrorGetCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case CGetCartDishesNotFound:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		case RGetRestaurantRestaurantNotFound, CGetCartCartNotFound:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusNotFound,
				Explain: ErrCartNull,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK
		default:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorUpdateCart(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case CUpdateCartCartNotInsert, CUpdateCartStructFoodStructureFoodNotInsert, CUpdateCartRadiosRadiosNotInsert,
			CGetPriceDeliveryPriceNotFound, CGetPriceDeliveryPriceNotScan, CUpdateCartCartNotFound,
			CUpdateCartCartNotScan, CUpdateCartStructureFoodStructureFoodNotSelect,
			CUpdateCartStructRadiosStructRadiosNotSelect, CUpdateCartTransactionNotCreate, CUpdateCartNotCommit:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		default:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorAddPromoCode(err error) (error, []byte, int) {
	if err != nil {
		result, errMarshal := easyjson.Marshal(ResultError{
			Status:  http.StatusInternalServerError,
			Explain: ErrDB,
		})
		if errMarshal != nil {
			c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
			return &Errors{
					Text: ErrMarshal,
				},
				nil, http.StatusInternalServerError
		}
		c.Logger.Errorf("%s, requestId: %d", err.Error(), c.RequestId)
		return &Errors{
				Text: ErrCheck,
			},
			result, http.StatusInternalServerError
	}
	return nil, nil, IntNil
}
