package Errors

import (
	"encoding/json"
	"net/http"
)

func (c *CheckError) CheckErrorProfile(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PGetProfileClientClientNotScan, PGetProfileClientBirthdayNotScan, PGetProfileCourierCourierNotScan, PGetProfileHostHostNotScan,
			PGetRoleByIdClientNotScan, PGetRoleByIdHostNotScan, PGetRoleByIdCourierNotScan, PGetProfileUnknownRole:
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
			c.Logger.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdateName(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdateNameNameNotUpdate:
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
			c.Logger.Errorf("error: %s, requestId: %d", PUpdateNameNameNotUpdate, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdateEmail(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdateEmailEmailNotUpdate:
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
			c.Logger.Errorf("error: %s, requestId: %d", PUpdateEmailEmailNotUpdate, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError

		case PUpdateEmailEmailRepeat:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: PUpdateEmailEmailRepeat,
			})
			if errMarshal != nil {
				c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("error: %s, requestId: %d", PUpdateEmailEmailRepeat, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdatePassword(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdatePasswordPasswordNotUpdate, PUpdatePasswordSaltNotSelect:
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
			c.Logger.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdatePhone(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdatePhonePhoneNotUpdate, PUpdatePhoneIncorrectPhoneFormat:
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
			c.Logger.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError

		case PUpdatePhonePhoneRepeat:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: PUpdatePhonePhoneRepeat,
			})
			if errMarshal != nil {
				c.Logger.Errorf("error: %s, %v, requestId: %d", ErrMarshal, errMarshal, c.RequestId)
				return &Errors{
						Text: ErrMarshal,
					},
					nil, http.StatusInternalServerError
			}
			c.Logger.Warnf("error: %s, requestId: %d", PUpdatePhonePhoneRepeat, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusOK
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdateAvatar(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdateAvatarAvatarNotUpdate, PUpdateAvatarAvatarNotOpen, PUpdateAvatarAvatarNotUpload,
			PUpdateAvatarFileNameEmpty, PUpdateAvatarFileWithoutExtension:
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
			c.Logger.Errorf("error: %s, requestId: %d", err.Error(), c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdateBirthday(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdateBirthdayBirthdayNotUpdate:
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
			c.Logger.Errorf("error: %s, requestId: %d", PUpdateBirthdayBirthdayNotUpdate, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}

func (c *CheckError) CheckErrorProfileUpdateAddress(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case PUpdateAddressAddressNotUpdate:
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
			c.Logger.Errorf("error: %s, requestId: %d", PUpdateAddressAddressNotUpdate, c.RequestId)
			return &Errors{
					Text: ErrCheck,
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, IntNil
}
