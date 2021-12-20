package myerror

import (
	"github.com/mailru/easyjson"
	"net/http"
)

func (c CheckError) CheckErrorSignUp(errIn error) (error, []byte, int) {
	if errIn != nil {
		switch errIn.Error() {
		case AGeneralSignUpLoginNotUnique:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusConflict,
				Explain: AGeneralSignUpLoginNotUnique,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("%s, requestId: %d", AGeneralSignUpLoginNotUnique, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK

		case AGeneralSignUpIncorrectPhoneFormat:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: AGeneralSignUpIncorrectPhoneFormat,
			})
			if errMarshal != nil {
				c.Logger.Errorf("%s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("%s, requestId: %d", AGeneralSignUpIncorrectPhoneFormat, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK

		case AGeneralSignUpNotInsert, ASignUpHostHostNotInsert, AAddTransactionCookieNotInsert,
			ASignUpCourierCourierNotInsert, ASignUpClientClientNotInsert, ASignupHostTransactionNotCreate,
			ASignupCourierTransactionNotCreate, ASignupClientTransactionNotCreate, ASignUpUnknownType,
			ASignUpHostNotCommit, ASignUpCourierNotCommit, ASignUpClientNotCommit:
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
			c.Logger.Errorf("%s, requestId: %d", errIn.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorLogin(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ALoginOrPasswordIncorrect, ALoginNotFound:
			result, errMarshal := easyjson.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ALoginOrPasswordIncorrect,
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

		case AAddCookieCookieNotInsert, ASaltNotSelect, ALoginVoidLogin, ALoginByEmailTransactionNotCreate,
			ALoginByEmailNotCommit, ALoginByPhoneTransactionNotCreate, ALoginByPhoneNotCommit:
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

func (c *CheckError) CheckErrorLogout(err error) (error, []byte, int) {
	if err != nil && err.Error() == ADeleteCookieCookieNotDelete {
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
		c.Logger.Errorf("%s, requestId: %d", ADeleteCookieCookieNotDelete, c.RequestId)
		return &Errors{
				Text: ErrCheck,
			},
			result, http.StatusInternalServerError
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorNewWsKey(err error) (error, []byte, int) {
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
