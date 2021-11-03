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
		case AGeneralSignUpLoginNotUnique:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusConflict,
				Explain: AGeneralSignUpLoginNotUnique,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", AGeneralSignUpLoginNotUnique)
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", AGeneralSignUpLoginNotUnique, *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case AGeneralSignUpIncorrectPhoneFormat:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: AGeneralSignUpIncorrectPhoneFormat,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", AGeneralSignUpIncorrectPhoneFormat, *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case AGeneralSignUpNotInsert, ASignUpHostHostNotInsert, AAddTransactionCookieNotInsert, ASignUpCourierCourierNotInsert, ASignUpClientClientNotInsert:
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
			c.LoggerErrWarn.Errorf("error: %s, requestId: %d", errIn.Error(), *c.RequestId)
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
		case ALoginOrPasswordIncorrect, ALoginNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ALoginOrPasswordIncorrect,
			})
			if errMarshal != nil {
				c.LoggerErrWarn.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, *c.RequestId)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			c.LoggerErrWarn.Warnf("error: %s, requestId: %d", err.Error(), *c.RequestId)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		case AAddCookieCookieNotInsert, ASaltNotSelect:
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

func (c *CheckError) CheckErrorLogout(err error) (error, []byte, int) {
	if err != nil && err.Error() == ADeleteCookieCookieNotDelete {
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
		c.LoggerErrWarn.Errorf("error: %s, requestId: %d", ADeleteCookieCookieNotDelete, *c.RequestId)
		return &Errors{
				Text: ErrCheck,
				Time: time.Now(),
			},
			result, http.StatusInternalServerError
	}
	return nil, nil, HttpNil
}
