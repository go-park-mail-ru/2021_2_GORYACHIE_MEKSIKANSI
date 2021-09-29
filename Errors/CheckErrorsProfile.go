package Errors

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func CheckErrorProfile(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case profile.ERRGETPROFILECLIENTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETPROFILECLIENTQUERY)
			return errors.New("fatal")
		case profile.ERRGETPROFILECLIENTSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETPROFILECLIENTSCAN)
			return errors.New("fatal")
		case profile.ERRGETBIRTHDAYQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETBIRTHDAYQUERY)
			return errors.New("fatal")
		case profile.ERRGETBIRTHDAYSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETBIRTHDAYSCAN)
			return errors.New("fatal")
		case profile.ERRGETPROFILECOURIERQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETPROFILECOURIERQUERY)
			return errors.New("fatal")
		case profile.ERRGETPROFILECOURIERSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETPROFILECOURIERSCAN)
			return errors.New("fatal")
		case profile.ERRGETPROFILEHOSTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETPROFILEHOSTQUERY)
			return errors.New("fatal")
		case profile.ERRGETPROFILEHOSTSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRGETPROFILEHOSTSCAN)
			return errors.New("fatal")
		case profile.ERRCLIENTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRCLIENTQUERY)
			return errors.New("fatal")
		case profile.ERRCLIENTSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRCLIENTSCAN)
			return errors.New("fatal")
		case profile.ERRHOSTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRHOSTQUERY)
			return errors.New("fatal")
		case profile.ERRHOSTSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRHOSTSCAN)
			return errors.New("fatal")
		case profile.ERRCORIERQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRCORIERQUERY)
			return errors.New("fatal")
		case profile.ERRCORIERSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", profile.ERRCORIERSCAN)
			return errors.New("fatal")
		}
	}
	return nil
}

