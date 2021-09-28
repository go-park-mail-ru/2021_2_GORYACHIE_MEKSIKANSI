package Restaurant

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

const(
	RESTNULL = "ERROR: restaurants not found"
)

func CheckErrorRestaurant(err error, ctx *fasthttp.RequestCtx, restaurant []Restaurant) error {
	if err != nil {
		switch err.Error() {
		case ERRQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRQUERY)
			return errors.New("fatal")
		case ERRSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRSCAN)
			return errors.New("fatal")

		}
	} else {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusBadRequest,
				Explain: RESTNULL,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", RESTNULL)
			return errors.New("fatal")
	}
	return nil
}

