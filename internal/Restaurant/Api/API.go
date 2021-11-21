package Api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
)

type InfoRestaurant struct {
	Application interfaces.RestaurantApplication
	Logger      errPkg.MultiLogger
}

func (r *InfoRestaurant) RestaurantHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	restaurant, err := r.Application.AllRestaurants()
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurant(err)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
		Body: &Restaurant.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) RestaurantIdHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("idRes")
	id, errConvert := Util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	restaurant, err := r.Application.GetRestaurant(id)

	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurantId(err) // должна появиться новая ошибка +1
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
		Body: &Restaurant.RestaurantIdResponse{
			RestaurantsGet: restaurant,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) RestaurantDishesHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("idRes")
	idRes, errConvert := Util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	idDishCtx := ctx.UserValue("idDish")
	idDish, errConvert := Util.InterfaceConvertInt(idDishCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	dishes, err := r.Application.RestaurantDishes(idRes, idDish)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurantDishes(err)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
		Body: &Restaurant.DishesResponse{
			DishesGet: dishes,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}
