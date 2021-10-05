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
		case ErrGetProfileClientScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrGetProfileClientScan)
			return errors.New("fatal")
		case ErrGetBirthdayScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrGetBirthdayScan)
			return errors.New("fatal")
		case ErrGetProfileCourierScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrGetProfileCourierScan)
			return errors.New("fatal")
		case ErrGetProfileHostScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrGetProfileHostScan)
			return errors.New("fatal")
		case ErrClientScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrClientScan)
			return errors.New("fatal")
		case ErrHostScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrHostScan)
			return errors.New("fatal")
		case ErrCourierScan:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ErrDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ErrCourierScan)
			return errors.New("fatal")
		}
	}
	return nil
}

