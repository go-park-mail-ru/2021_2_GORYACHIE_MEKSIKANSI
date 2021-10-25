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

type InfoRestaurant struct {
	ConnectionDB *pgxpool.Pool
}

func (r *InfoRestaurant) RestaurantHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}

	restaurant, err := AllRestaurants(&WrapperDB)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorRestaurant(err)
	if errOut != nil {
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
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}
}

func (r *InfoRestaurant) RestaurantIdHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}

	idUrl := ctx.UserValue("idRes")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		fmt.Printf("Console: %s\n", errors.ErrNotStringAndInt)
		return
	}

	restaurant, err := GetRestaurant(&WrapperDB, id)

	errOut, resultOutAccess, codeHTTP := errors.CheckErrorRestaurantId(err) // должна появиться новая ошибка +1
	if errOut != nil {
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
		Body: &utils.RestaurantIdResponse{
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

func (r *InfoRestaurant) RestaurantDishesHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}

	idResIn := ctx.UserValue("idRes")
	var idRes int
	var errorConvert error
	switch idResIn.(type) {
	case string:
		idRes, errorConvert = strconv.Atoi(idResIn.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		idRes = idResIn.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		fmt.Printf("Console: %s\n", errors.ErrNotStringAndInt)
		return
	}

	idDishIn := ctx.UserValue("idDish")
	var idDish int
	switch idDishIn.(type) {
	case string:
		idDish, errorConvert = strconv.Atoi(idDishIn.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		idDish = idDishIn.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		fmt.Printf("Console: %s\n", errors.ErrNotStringAndInt)
		return
	}

	dishes, err := RestaurantDishes(&WrapperDB, idRes, idDish)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorRestaurantDishes(err)
	if errOut != nil {
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
		Body: &utils.DishesResponse{
			DishesGet: dishes,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}
}
