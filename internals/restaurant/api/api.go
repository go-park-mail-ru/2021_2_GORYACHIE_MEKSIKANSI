//go:generate mockgen -destination=mocks/api.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/myerror MultiLogger
//go:generate mockgen -destination=mocks/apiApplication.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/application RestaurantApplicationInterface
package api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
)

type RestaurantApiInterface interface {
	RestaurantHandler(ctx *fasthttp.RequestCtx)
	RestaurantIdHandler(ctx *fasthttp.RequestCtx)
	RestaurantDishesHandler(ctx *fasthttp.RequestCtx)
	RecommendedRestaurantsHandler(ctx *fasthttp.RequestCtx)
	UpdateFavouritesHandler(ctx *fasthttp.RequestCtx)
	GetFavouritesHandler(ctx *fasthttp.RequestCtx)
	SearchRestaurantHandler(ctx *fasthttp.RequestCtx)
	GetReviewHandler(ctx *fasthttp.RequestCtx)
	CreateReviewHandler(ctx *fasthttp.RequestCtx)
}

type InfoRestaurant struct {
	Application appPkg.RestaurantApplicationInterface
	Logger      errPkg.MultiLogger
}

func (r *InfoRestaurant) RestaurantHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, errResponse.Error(), reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) RecommendedRestaurantsHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	restaurant, err := r.Application.RecommendedRestaurants()
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) RestaurantIdHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("idRes")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	idCtx = ctx.UserValue("id")
	idClient, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	restaurant, err := r.Application.GetRestaurant(id, idClient)

	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorRestaurantId(err)
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantIdResponse{
			RestaurantsGet: restaurant,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) RestaurantDishesHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("idRes")
	idRes, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	idDishCtx := ctx.UserValue("idDish")
	idDish, errConvert := util.InterfaceConvertInt(idDishCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.DishesResponse{
			DishesGet: dishes,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) CreateReviewHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	newReview := resPkg.NewReview{}
	err := easyjson.Unmarshal(ctx.Request.Body(), &newReview)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		r.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrUnmarshal, err.Error(), reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}
	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := util.InterfaceConvertString(tokenContext)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	err = r.Application.CreateReview(id, newReview)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorCreateReview(err)
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

	response, errResponse := easyjson.Marshal(&util.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) GetReviewHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("idRes")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}
	var idClient = 0
	idCleintCtx := ctx.UserValue("id")
	if idCleintCtx != nil {
		idClient, errConvert = util.InterfaceConvertInt(idCleintCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			r.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
			return
		}
	}
	restaurant, err := r.Application.GetReview(id, idClient)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetReview(err)
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

	if restaurant.Reviews == nil {
		response, errResponse := easyjson.Marshal(&errPkg.ResultErrorMulti{
			Status:  http.StatusNotFound,
			Explain: errPkg.RGetReviewEmpty,
			Body: &resPkg.RestaurantsResponse{
				RestaurantsGet: restaurant,
			},
		})
		ctx.Response.SetBody(response)
		json.NewEncoder(ctx)
		if errResponse != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errPkg.ErrEncode))
			r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
			return
		}
		return
	}

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) SearchRestaurantHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	searchText := string(ctx.FormValue("searchText"))

	restaurant, err := r.Application.SearchRestaurant(searchText)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorSearchRes(err)
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) GetFavouritesHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	restaurant, err := r.Application.GetFavoriteRestaurants(id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetFavorite(err)
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantsResponse{
			RestaurantsGet: restaurant,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.SetStatusCode(http.StatusOK)
}

func (r *InfoRestaurant) UpdateFavouritesHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    r.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		r.Logger.Errorf("%s", errConvert.Error())
	}

	var userFavourite resPkg.ResFavouriteNew
	err := easyjson.Unmarshal(ctx.Request.Body(), &userFavourite)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		r.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrUnmarshal, err.Error(), reqId)
		return
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := util.InterfaceConvertString(tokenContext)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		return
	}

	statusFavourite, err := r.Application.EditRestaurantInFavorite(userFavourite.Id, id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorUpdateFavorite(err)
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &resPkg.RestaurantsResponse{
			RestaurantsGet: resPkg.ResFavouriteStatus{Status: statusFavourite},
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		r.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.SetStatusCode(http.StatusOK)
}
