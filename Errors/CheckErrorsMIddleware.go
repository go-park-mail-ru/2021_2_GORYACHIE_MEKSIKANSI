package Errors

import (
	"encoding/json"
	"net/http"
)

func (c *CheckError) CheckErrorCookie(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case MGetIdByCookieCookieNotScan:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("error: %s, requestId: %d", MGetIdByCookieCookieNotScan, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError

		case MGetIdByCookieCookieExpired, MGetIdByCookieCookieNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: err.Error(),
			})
			if errMarshal != nil {
				c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorAccess(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case MCheckAccessCookieNotScan:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Errorf("error: %s, requestId: %d", AGeneralSignUpLoginNotUnique, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError

		case MCheckAccessCookieNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: MCheckAccessCookieNotFound,
			})
			if errMarshal != nil {
				c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("error: %s, requestId: %d", MCheckAccessCookieNotFound, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK
		}
	}
	return nil, nil, IntNil
}
