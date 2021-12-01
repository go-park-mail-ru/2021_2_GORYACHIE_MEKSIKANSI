package api

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/application"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"github.com/valyala/fasthttp"
	"math"
	"net/http"
)

type MiddlewareApiInterface interface {
	LogURL(h fasthttp.RequestHandler) fasthttp.RequestHandler
	GetIdClient(h fasthttp.RequestHandler) fasthttp.RequestHandler
	CheckClient(h fasthttp.RequestHandler) fasthttp.RequestHandler
	CheckWebSocketKey(h fasthttp.RequestHandler) fasthttp.RequestHandler
	GetIdClientForReview(h fasthttp.RequestHandler) fasthttp.RequestHandler
}

type InfoMiddleware struct {
	Application appPkg.MiddlewareApplicationInterface
	Logger      errPkg.MultiLogger
	ReqId       int
}

func (m *InfoMiddleware) LogURL(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		if m.ReqId == math.MaxInt {
			m.ReqId = 0
		}
		m.ReqId++
		m.Logger.Infof("Method: %s, URL: %s, requestId: %d", string(ctx.Method()), ctx.URI(), m.ReqId)
		ctx.SetUserValue("reqId", m.ReqId)
		h(ctx)
	})
}

func (m *InfoMiddleware) GetIdClient(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqIdCtx := ctx.UserValue("reqId")
		reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			m.Logger.Errorf("%s", errConvert.Error())
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errPkg.CheckError{
			Logger:    m.Logger,
			RequestId: reqId,
		}

		cookieDB := util.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		id, err := m.Application.GetIdByCookie(&cookieDB)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorCookie(err)
		if resultOutAccess != nil {
			switch errAccess.Error() {
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
		//ctx.Response.Header.SetContentType("charset=UTF-8")
		//ctx.Response.Header.SetContentType("application/json")
		ctx.SetUserValue("id", id)

		h(ctx)
	})
}

func (m *InfoMiddleware) GetIdClientForReview(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqIdCtx := ctx.UserValue("reqId")
		reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			m.Logger.Errorf("%s", errConvert.Error())
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errPkg.CheckError{
			Logger:    m.Logger,
			RequestId: reqId,
		}

		cookieDB := util.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		id, err := m.Application.GetIdByCookie(&cookieDB)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorCookie(err)
		if resultOutAccess != nil {
			switch errAccess.Error() {
			case errPkg.ErrMarshal:
				ctx.Response.SetStatusCode(codeHTTP)
				ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
				return
			}
		}
		//ctx.Response.Header.SetContentType("charset=UTF-8")
		//ctx.Response.Header.SetContentType("application/json")
		ctx.SetUserValue("id", id)

		h(ctx)
	})
}

func (m *InfoMiddleware) CheckClient(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		reqIdCtx := ctx.UserValue("reqId")
		reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			m.Logger.Errorf("%s", errConvert.Error())
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errPkg.CheckError{
			Logger:    m.Logger,
			RequestId: reqId,
		}

		cookieDB := util.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

		_, err := m.Application.CheckAccess(&cookieDB)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorAccess(err)
		if errAccess != nil {
			switch errAccess.Error() {
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
		ctx.SetUserValue("X-Csrf-Token", cookieDB.CsrfToken)
		//ctx.Response.Header.SetContentType("application/json")

		h(ctx)
	})
}

func (m *InfoMiddleware) CheckWebSocketKey(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		reqIdCtx := ctx.UserValue("reqId")
		reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
		if errConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errConvert.Error()))
			m.Logger.Errorf("%s", errConvert.Error())
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errPkg.CheckError{
			Logger:    m.Logger,
			RequestId: reqId,
		}

		key := string(ctx.FormValue("key"))

		_, err := m.Application.CheckAccessWebsocket(key)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorWsKey(err)
		if errAccess != nil {
			switch errAccess.Error() {
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

		h(ctx)
	})
}
