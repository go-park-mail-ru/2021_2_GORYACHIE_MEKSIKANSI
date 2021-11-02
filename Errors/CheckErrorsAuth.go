package Errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (c *CheckError) CheckErrorSignUp(errIn error) (error, []byte, int) {
	if errIn != nil {
		switch errIn.Error() {
		case ErrGeneralInfoUnique:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusConflict,
				Explain: ErrGeneralInfoUnique,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrGeneralInfoUnique)
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", ErrGeneralInfoUnique, c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case ErrPhoneFormat:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrPhoneFormat,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", ErrPhoneFormat, c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case ErrGeneralInfoScan, ErrInsertHost, ErrInsertTransactionCookie, ErrInsertCourier, ErrInsertClient:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Errorf("error: %s, requestId: %d", errIn.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func (c *CheckError) CheckErrorLogin(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrLoginOrPasswordIncorrect, ErrUserNotFoundLogin:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrLoginOrPasswordIncorrect,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case ErrInsertCookie, ErrSelectSaltInLogin:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func (c *CheckError) CheckErrorLogout(err error) (error, []byte, int) {
	if err != nil && err.Error() == ErrDeleteCookie {
		result, errMarshal := json.Marshal(ResultError{
			Status:  http.StatusInternalServerError,
			Explain: ErrDB,
		})
		if errMarshal != nil {
			c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
			return &Errors{
					Text: ErrMarshal,
					Time: time.Now(),
				},
				nil, http.StatusInternalServerError
		}
		c.LoggerErrWarn.Errorf("error: %s, requestId: %d", ErrDeleteCookie, c.RequestId)
		return &Errors{
				Text: ErrCheck,
				Time: time.Now(),
			},
			result, http.StatusInternalServerError
	}
	return nil, nil, HttpNil
}
