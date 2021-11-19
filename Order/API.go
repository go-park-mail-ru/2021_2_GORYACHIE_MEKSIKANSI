package Order

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
)

type InfoOrder struct {
	Application interfaces.OrderApplication
	Logger      errPkg.MultiLogger
}

func (u *InfoOrder) CreateOrderHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var createOrder utils.CreateOrder
	err := json.Unmarshal(ctx.Request.Body(), &createOrder)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}
	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := utils.InterfaceConvertString(tokenContext)
	if (errConvert != nil) && (errConvert.Error() == errPkg.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrNotStringAndInt))
		return
	}

	err = u.Application.CreateOrder(id, createOrder)

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

	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *InfoOrder) GetOrdersHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}
	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
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
	err = json.NewEncoder(ctx).Encode(&utils.Result{
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
