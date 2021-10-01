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
		case ERRGETPROFILECLIENTSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILECLIENTSCAN)
			return errors.New("fatal")
		case ERRGETBIRTHDAYSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETBIRTHDAYSCAN)
			return errors.New("fatal")
		case ERRGETPROFILECOURIERSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILECOURIERSCAN)
			return errors.New("fatal")
		case ERRGETPROFILEHOSTSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILEHOSTSCAN)
			return errors.New("fatal")
		case ERRCLIENTSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRCLIENTSCAN)
			return errors.New("fatal")
		case ERRHOSTSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRHOSTSCAN)
			return errors.New("fatal")
		case ERRCORIERSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRCORIERSCAN)
			return errors.New("fatal")
		}
	}
	return nil
}

