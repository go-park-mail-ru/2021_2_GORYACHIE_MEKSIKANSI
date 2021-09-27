package Profile

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func CheckErrorProfile(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() == ERRGETPROFILECLIENTQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILECLIENTQUERY)
			return errors.New("Fatal")
		}

		if err.Error() == ERRGETPROFILECLIENTSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILECLIENTSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRGETBIRTHDAYQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETBIRTHDAYQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRGETBIRTHDAYSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETBIRTHDAYSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRGETPROFILECOURIERQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILECOURIERQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRGETPROFILECOURIERSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILECOURIERSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRGETPROFILEHOSTQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILEHOSTQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRGETPROFILEHOSTSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRGETPROFILEHOSTSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRCLIENTQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRCLIENTQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRCLIENTSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRCLIENTSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRHOSTQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRHOSTQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRHOSTSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRHOSTSCAN)
			return errors.New("Fatal")
		}

		if err.Error() == ERRHOSTSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRHOSTSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRCORIERQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRCORIERQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRCORIERSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRCORIERSCAN)
			return errors.New("Fatal")
		}
	}
	return nil
}
