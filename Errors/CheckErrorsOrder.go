package Errors

import (
	"encoding/json"
	"net/http"
	"time"
)

func (c *CheckError) CheckErrorCreateOrder(err error) (error, []byte, int) {
	if err != nil {
		result, errMarshal := json.Marshal(ResultError{
			Status:  http.StatusInternalServerError,
			Explain: ErrDB,
		})
		if errMarshal != nil {
			c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
			return &Errors{
					Text: ErrMarshal,
					Time: time.Now(),
				},
				nil, http.StatusInternalServerError
		}
		c.Logger.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
		return &Errors{
				Text: ErrCheck,
				Time: time.Now(),
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
			c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
			return &Errors{
					Text: ErrMarshal,
					Time: time.Now(),
				},
				nil, http.StatusInternalServerError
		}
		c.Logger.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
		return &Errors{
				Text: ErrCheck,
				Time: time.Now(),
			},
			result, http.StatusInternalServerError
	}
	return nil, nil, IntNil
}
