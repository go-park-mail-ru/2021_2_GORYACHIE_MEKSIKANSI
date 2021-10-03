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
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusConflict,
				Explain: ErrGeneralInfoUnique,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrGeneralInfoUnique)
			return errors.New("fatal")
		case ErrGeneralInfoScan:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrGeneralInfoScan)
			return errors.New("fatal")
		case ErrInsertHost:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrInsertHost)
			return errors.New("fatal")
		case ErrInsertTransactionCookie:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrInsertTransactionCookie)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrInsertTransactionCookie)
			return errors.New("fatal")
		case ErrInsertCourier:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrInsertCourier)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrInsertCourier)
			return errors.New("fatal")
		case ErrInsertClient:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrInsertCourier)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogin(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ErrLoginOrPasswordIncorrect:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrLoginOrPasswordIncorrect,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrLoginOrPasswordIncorrect)
			return errors.New("fatal")
		case ErrInsertCookie:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrInsertCookie)
			return errors.New("fatal")
		case ErrSelectSalt:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrSelectSalt)
			return errors.New("fatal")

		}
	}
	return nil
}

func CheckErrorLogout(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() != ErrDeleteCookie {
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusBadRequest,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrDeleteCookie)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogoutAccess(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ErrCookieNotScan:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrCookieNotScan)
			return errors.New("fatal")
		case ErrCheckAccessCookieNotFound:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrCheckAccessCookieNotFound,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusUnauthorized)
			fmt.Printf("Console: %s\n", ErrCheckAccessCookieNotFound)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLoggedIn(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ErrCookieScan:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrCookieScan)
			return errors.New("fatal")
		case ErrCookieExpired:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrCookieExpired,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCookieExpired)
			return errors.New("fatal")
		case ErrCheckAccessCookieNotFound:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrAuth,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCheckAccessCookieNotFound)
			return errors.New("fatal")
		case ErrCookieNotFound:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrAuth,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCookieNotFound)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorProfileCookie(err error, ctx *fasthttp.RequestCtx, cookieHTTP *fasthttp.Cookie) error {
	if err != nil {
		switch err.Error() {

		case ErrCookieScan:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrCookieScan)
			return errors.New("fatal")
		case ErrCookieExpired:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrCookieExpired,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCookieExpired)
			return errors.New("fatal")
		case ErrCookieNotFound:
			err := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusConflict,
				Explain: ErrAuth,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCookieNotFound)
			return errors.New("fatal")
		}
	}
	return nil
}
