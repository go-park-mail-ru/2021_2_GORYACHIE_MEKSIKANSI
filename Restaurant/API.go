package Restaurant

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

type InfoRestaurant struct {
	Application interfaces.RestaurantApplication
	Logger      errors.MultiLogger
}

func (r *InfoRestaurant) RestaurantHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			r.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    r.Logger,
		RequestId: &reqId,
	}

	restaurant, err := r.Application.AllRestaurants()
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurant(err)
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

	err = json.NewEncoder(ctx).Encode(&utils.Result{
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
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			r.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    r.Logger,
		RequestId: &reqId,
	}

	idCtx := ctx.UserValue("idRes")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			r.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	restaurant, err := r.Application.GetRestaurant(id)

	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurantId(err) // должна появиться новая ошибка +1
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

	err = json.NewEncoder(ctx).Encode(&utils.Result{
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
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			r.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    r.Logger,
		RequestId: &reqId,
	}

	idCtx := ctx.UserValue("idRes")
	idRes, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			r.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	idDishCtx := ctx.UserValue("idDish")
	idDish, errConvert := utils.InterfaceConvertInt(idDishCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			r.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	dishes, err := r.Application.RestaurantDishes(idRes, idDish)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurantDishes(err)
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

	err = json.NewEncoder(ctx).Encode(&utils.Result{
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
