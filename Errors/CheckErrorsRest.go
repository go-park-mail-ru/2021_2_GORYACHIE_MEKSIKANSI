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
		case ErrRestaurantsNotFound:
			errEncode := json.NewEncoder(ctx).Encode(ResultError{
				Status:  http.StatusNoContent,
				Explain: ErrRestaurantsNotFound,
			})
			if errEncode != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ErrEncode)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ErrRestaurantsNotFound)
			return errors.New("fatal")
		case ErrRestaurantScan:
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
			fmt.Printf("Console: %s\n", ErrRestaurantScan)
			return errors.New("fatal")
		}
	}
	return nil
}


