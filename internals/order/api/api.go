//go:generate mockgen -destination=mocks/api.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/myerror MultiLogger
//go:generate mockgen -destination=mocks/apiApplication.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/order/application OrderApplicationInterface

package api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/order/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
)

type OrderApiInterface interface {
	CreateOrderHandler(ctx *fasthttp.RequestCtx)
	GetOrdersHandler(ctx *fasthttp.RequestCtx)
}

type InfoOrder struct {
	Application appPkg.OrderApplicationInterface
	Logger      errPkg.MultiLogger
}

func (u *InfoOrder) CreateOrderHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var createOrder order.CreateOrder
	err := json.Unmarshal(ctx.Request.Body(), &createOrder)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrUnmarshal, err.Error(), reqId)
		return
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := util.InterfaceConvertString(tokenContext)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	idOrder, err := u.Application.CreateOrder(id, createOrder)

	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorCreateOrder(err)
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

	request, errRequest := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body:   order.ConvertCreateOrderIdToOrderResponse(idOrder),
	})
	if errRequest != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errRequest, reqId)
		return
	}

	ctx.Response.SetBody(request)
	json.NewEncoder(ctx)
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *InfoOrder) GetOrdersHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}
	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	historyOrders, err := u.Application.GetOrders(id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetOrders(err)
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
	err = json.NewEncoder(ctx).Encode(&authorization.Result{
		Status: http.StatusOK,
		Body:   historyOrders,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *InfoOrder) GetOrderActiveHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	idCtxOrd := ctx.UserValue("idOrd")
	idOrd, errConvert := util.InterfaceConvertInt(idCtxOrd)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvert.Error(), reqId)
		return
	}

	activeOrder, err := u.Application.GetActiveOrder(id, idOrd)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetOrderActive(err)
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
	err = json.NewEncoder(ctx).Encode(&authorization.Result{
		Status: http.StatusOK,
		Body:   activeOrder,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
}
