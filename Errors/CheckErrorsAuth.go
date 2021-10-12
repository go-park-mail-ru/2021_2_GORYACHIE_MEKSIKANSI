package Errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CheckErrorSignUp(errIn error) (error, []byte, int) {
	if errIn != nil {
		switch errIn.Error() {
		case ErrGeneralInfoUnique:
			result, errMarshal:= json.Marshal(ResultError{
				Status:  http.StatusConflict,
				Explain: ErrGeneralInfoUnique,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
					Text: ErrMarshal,
					Time: time.Now(),
				},
				nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrGeneralInfoUnique)
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
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrPhoneFormat)
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
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrPhoneFormat)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorLogin(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrLoginOrPasswordIncorrect, ErrUserNotFoundLogin:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrLoginOrPasswordIncorrect,
			})
			if errMarshal != nil {
				//ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			//ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", err.Error())
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
				//ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			//ctx.Response.SetStatusCode(http.StatusInternalServerError)
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

func CheckErrorLogout(err error) (error, []byte, int) {
	if err != nil && err.Error() != ErrDeleteCookie {
		result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errMarshal != nil {
				//ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			//ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrDeleteCookie)
		return &Errors{
				Text: ErrCheck,
				Time: time.Now(),
			},
			result, http.StatusInternalServerError
	}
	return nil, nil, HttpNil
}

func CheckErrorLogoutAccess(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrCookieNotScan:
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
			fmt.Printf("Console: %s\n", ErrCookieNotScan)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		case ErrCheckAccessCookieNotFound:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrCheckAccessCookieNotFound,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrCheckAccessCookieNotFound)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		}
	}
	return nil, nil, HttpNil
}
