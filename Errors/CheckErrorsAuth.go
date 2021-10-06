package Errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func CheckErrorSignUp(errIn error, ctx *fasthttp.RequestCtx) error {
	if errIn != nil {
		switch errIn.Error() {
		case ErrGeneralInfoUnique:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusConflict,
				Explain: ErrGeneralInfoUnique,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrGeneralInfoUnique)
			return errors.New("fatal")
		case ErrPhoneFormat:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrPhoneFormat,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrPhoneFormat)
			return errors.New("fatal")
		case ErrGeneralInfoScan, ErrInsertHost, ErrInsertTransactionCookie, ErrInsertCourier, ErrInsertClient:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", errIn.Error())
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogin(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ErrLoginOrPasswordIncorrect, ErrUserNotFoundLogin:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrLoginOrPasswordIncorrect,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", err.Error())
			return errors.New("fatal")
		case ErrInsertCookie, ErrSelectSaltInLogin:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", err.Error())
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogout(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil && err.Error() != ErrDeleteCookie {
		errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrDeleteCookie)
			return errors.New("fatal")
	}
	return nil
}

func CheckErrorLogoutAccess(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ErrCookieNotScan:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrCookieNotScan)
			return errors.New("fatal")
		case ErrCheckAccessCookieNotFound:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrCheckAccessCookieNotFound,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCheckAccessCookieNotFound)
			return errors.New("fatal")
		}
	}
	return nil
}
