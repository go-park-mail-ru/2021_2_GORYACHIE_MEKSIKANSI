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
	RESTAURANTNULL = "ERROR: restaurants not found"
)

func CheckErrorRestaurant(err error, ctx *fasthttp.RequestCtx, restaurant []Restaurant) error {
	if err != nil {
		if err.Error() == ERRQUERY {
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
			fmt.Printf("Console: %s\n", ERRQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRSCAN {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRCOOKIEEXPIRED,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRSCAN)
			return errors.New("Fatal")
		}
	} else {
		if restaurant == nil {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusBadRequest,
				Explain: RESTAURANTNULL,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", RESTAURANTNULL)
			return errors.New("Fatal")
		}
	}
	return nil
}

