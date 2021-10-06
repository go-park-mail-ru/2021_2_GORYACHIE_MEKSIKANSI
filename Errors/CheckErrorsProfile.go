package Errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func CheckErrorProfile(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ErrGetProfileClientScan, ErrGetBirthdayScan, ErrGetProfileCourierScan, ErrGetProfileHostScan,
		ErrClientScan, ErrHostScan, ErrCourierScan:
			errEncode := json.NewEncoder(ctx).Encode(ResultError{
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


func CheckErrorProfileCookie(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {

		case ErrCookieScan:
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
			fmt.Printf("Console: %s\n", ErrCookieScan)
			return errors.New("fatal")
		case ErrCookieExpired:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ErrCookieExpired,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrCookieExpired)
			return errors.New("fatal")
		case ErrCookieNotFound:
			errEncode := json.NewEncoder(ctx).Encode( ResultError{
				Status:  http.StatusConflict,
				Explain: ErrAuth,
			})
			if errEncode != nil {
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

