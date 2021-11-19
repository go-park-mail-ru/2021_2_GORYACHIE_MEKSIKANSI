package Errors

import (
	"encoding/json"
	"net/http"
)

func (c *CheckError) CheckErrorCreateOrder(err error) (error, []byte, int) {
	if err != nil {
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
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorGetOrders(err error) (error, []byte, int) {
	if err != nil {
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
	return nil, nil, IntNil
}
