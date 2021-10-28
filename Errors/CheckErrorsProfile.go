package Errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CheckErrorProfile(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrGetProfileClientScan, ErrGetBirthdayScan, ErrGetProfileCourierScan, ErrGetProfileHostScan,
			ErrClientScan, ErrHostScan, ErrCourierScan:
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

func CheckErrorProfileUpdateName(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdateName:
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
			fmt.Printf("Console: %s\n", ErrUpdateName)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorProfileUpdateEmail(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdateEmail:
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
			fmt.Printf("Console: %s\n", ErrUpdateEmail)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		case ErrUpdateEmailRepeat:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrUpdateEmailRepeat,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrUpdateEmailRepeat)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorProfileUpdatePassword(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdatePassword, ErrSelectSaltInUpdate:
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

func CheckErrorProfileUpdatePhone(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdatePhone:
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
			fmt.Printf("Console: %s\n", ErrUpdatePhone)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		case ErrUpdatePhoneRepeat:
			result, errMarshal := json.Marshal(ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrUpdatePhoneRepeat,
			})
			if errMarshal != nil {
				fmt.Printf("Console: %s\n", ErrMarshal)
				return &Errors{
						Text: ErrMarshal,
						Time: time.Now(),
					},
					nil, http.StatusInternalServerError
			}
			fmt.Printf("Console: %s\n", ErrUpdatePhoneRepeat)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusOK
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorProfileUpdateAvatar(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdateAvatar:
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
			fmt.Printf("Console: %s\n", ErrUpdateAvatar)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorProfileUpdateBirthday(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdateBirthday:
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
			fmt.Printf("Console: %s\n", ErrUpdateBirthday)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}

func CheckErrorProfileUpdateAddress(err error) (error, []byte, int) {
	if err != nil {
		switch err.Error() {
		case ErrUpdateAddress:
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
			fmt.Printf("Console: %s\n", ErrUpdateAddress)
			return &Errors{
					Text: ErrCheck,
					Time: time.Now(),
				},
				result, http.StatusInternalServerError
		}
	}
	return nil, nil, HttpNil
}
