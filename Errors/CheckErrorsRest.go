package Errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func CheckErrorRestaurant(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case  ERRQUERY:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain:  ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n",  ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n",  ERRQUERY)
			return errors.New("fatal")
		case  ERRSCAN:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusInternalServerError,
				Explain:  ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n",  ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n",  ERRSCAN)
			return errors.New("fatal")
		case RESTNULL:
			err := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusBadRequest,
				Explain: RESTNULL,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n",  ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", RESTNULL)
			return errors.New("fatal")
		}
	}
	return nil
}


