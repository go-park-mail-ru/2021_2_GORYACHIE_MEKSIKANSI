//go:generate mockgen -destination=mocks/api.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/myerror MultiLogger
//go:generate mockgen -destination=mocks/apiApplication.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/cart/application CartApplicationInterface
package api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/cart/application"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
)

type CartApiInterface interface {
	GetCartHandler(ctx *fasthttp.RequestCtx)
	UpdateCartHandler(ctx *fasthttp.RequestCtx)
	AddPromocodeHandler(ctx *fasthttp.RequestCtx)
}

type InfoCart struct {
	Application appPkg.CartApplicationInterface
	Logger      errPkg.MultiLogger
}

func (c *InfoCart) GetCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    c.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvertId := util.InterfaceConvertInt(idCtx)
	if errConvertId != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertId.Error()))
		c.Logger.Errorf("%s, requestId: %d", errConvertId.Error(), reqId)
		return
	}

	result, err := c.Application.GetCart(id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetCart(err)
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
		Body: &cart.ResponseCart{
			Cart: result,
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (c *InfoCart) AddPromocodeHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    c.Logger,
		RequestId: reqId,
	}

	var newPromo cart.CreatePromoCode
	err := easyjson.Unmarshal(ctx.Request.Body(), &newPromo)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrUnmarshal, err, reqId)
		return
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvertToken := util.InterfaceConvertString(tokenContext)
	if errConvertToken != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertToken.Error()))
		c.Logger.Errorf("%s, requestId: %d", errPkg.ErrNotStringAndInt, reqId)
		return
	}
	idCtx := ctx.UserValue("id")
	id, errConvertId := util.InterfaceConvertInt(idCtx)
	if errConvertId != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertId.Error()))
		c.Logger.Errorf("%s, requestId: %d", errConvertId.Error(), reqId)
		return
	}

	errAddPromo := c.Application.AddPromoCode(newPromo.Code, newPromo.RestaurantId, id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorAddPromoCode(errAddPromo)
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
		Body:   newPromo,
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (c *InfoCart) UpdateCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    c.Logger,
		RequestId: reqId,
	}

	var cartRequest cart.CartRequest
	err := easyjson.Unmarshal(ctx.Request.Body(), &cartRequest)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrUnmarshal, err, reqId)
		return
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvertToken := util.InterfaceConvertString(tokenContext)
	if errConvertToken != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertToken.Error()))
		c.Logger.Errorf("%s, requestId: %d", errPkg.ErrNotStringAndInt, reqId)
		return
	}
	idCtx := ctx.UserValue("id")
	id, errConvertId := util.InterfaceConvertInt(idCtx)
	if errConvertId != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertId.Error()))
		c.Logger.Errorf("%s, requestId: %d", errConvertId.Error(), reqId)
		return
	}

	result, err := c.Application.UpdateCart(cartRequest.Cart, id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorUpdateCart(err)
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

	if result != nil {
		response, errResponse := easyjson.Marshal(&authorization.Result{
			Status: http.StatusOK,
			Body: &cart.ResponseCart{
				Cart: result,
			},
		})
		if errResponse != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errPkg.ErrEncode))
			c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
			return
		}

		ctx.Response.SetBody(response)
		json.NewEncoder(ctx)
		return
	}

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body:   cartRequest,
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}
