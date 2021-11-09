package Middleware

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"github.com/valyala/fasthttp"
	"math"
	"net/http"
)

type InfoMiddleware struct {
	Application interfaces.MiddlewareApplication
	Logger      errors.MultiLogger
}

func (m *InfoMiddleware) PrintURL(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqId := utils.RandomInteger(0, math.MaxInt64)
		m.Logger.Infof("Method: %s, URL: %s, requestId: %d", string(ctx.Method()), ctx.URI(), reqId)
		ctx.SetUserValue("reqId", reqId)
		h(ctx)
	})
}

func (m *InfoMiddleware) GetId(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqIdCtx := ctx.UserValue("reqId")
		reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			m.Logger.Errorf("SignUpHandler: GetId: %s, %v", errConvert.Error(), errConvert)
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errors.CheckError{
			Logger:    m.Logger,
			RequestId: reqId,
		}

		cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		id, err := m.Application.GetIdByCookie(&cookieDB)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorCookie(err)
		if resultOutAccess != nil {
			switch errAccess.Error() {
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
		//ctx.Response.Header.SetContentType("charset=UTF-8")
		//ctx.Response.Header.SetContentType("application/json")
		ctx.SetUserValue("id", id)

		h(ctx)
	})
}

func (m *InfoMiddleware) Check(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqIdCtx := ctx.UserValue("reqId")
		reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			m.Logger.Errorf("SignUpHandler: GetId: %s, %v", errConvert.Error(), errConvert)
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errors.CheckError{
			Logger:    m.Logger,
			RequestId: reqId,
		}

		cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

		_, err := m.Application.CheckAccess(&cookieDB)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorAccess(err)
		if errAccess != nil {
			switch errAccess.Error() {
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
		ctx.SetUserValue("X-Csrf-Token", cookieDB.CsrfToken)
		//ctx.Response.Header.SetContentType("application/json")

		h(ctx)
	})
}
