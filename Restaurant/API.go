package Restaurant

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

type RestaurantInfo struct {
	ConnectionDB *pgxpool.Pool
}

func (r *RestaurantInfo) RestaurantHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}
	restaurant, err := AllRestaurants(&WrapperDB)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorRestaurant(err)
	if resultOutAccess != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body: &utils.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}
}

func (r *RestaurantInfo) RestaurantDishesHandler(ctx *fasthttp.RequestCtx) {
	/*	WrapperDB := Wrapper{Conn: r.ConnectionDB}
		dishes, err := RestaurantDishes(&WrapperDB)
		errOut, resultOutAccess, codeHTTP  := errors.CheckErrorRestaurantDishes(err)
		if resultOutAccess != nil {
			switch errOut.Error() {
			case errors.ErrMarshal:
				ctx.Response.SetStatusCode(codeHTTP)
				ctx.Response.SetBody([]byte(errors.ErrMarshal))
				return
			case errors.ErrCheck:
				ctx.Response.SetStatusCode(codeHTTP)
				ctx.Response.SetBody(resultOutAccess)
				return
			}
		}*/

	ctx.SetStatusCode(http.StatusOK)

	/*	err = json.NewEncoder(ctx).Encode(&auth.Result{
			Status: http.StatusOK,
			Body: &utils.DishesResponse {
				DishesGet: dishes,
			},
		})
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", errors.ErrEncode)
			return
		}*/
}

func (r *RestaurantInfo) RestaurantIdHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}
	idUrl := ctx.UserValue("id")
	var id int
	var errorAtoi error
	switch idUrl.(type) {
	case string:
		id, errorAtoi = strconv.Atoi(idUrl.(string))
		if errorAtoi != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	restaurant, err := GetRestaurant(&WrapperDB, id)

	errOut, resultOutAccess, codeHTTP  := errors.CheckErrorRestaurantId(err)
	if resultOutAccess != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.SetStatusCode(http.StatusOK)

		err = json.NewEncoder(ctx).Encode(&auth.Result{
			Status: http.StatusOK,
			Body: &utils.RestaurantIdResponse {
				RestaurantsGet: restaurant,
			},
		})
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrEncode))
			fmt.Printf("Console: %s\n", errors.ErrEncode)
			return
		}
}
