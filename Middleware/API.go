package Middleware

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"math"
	"net/http"
	"strconv"
)

type InfoMiddleware struct {
	ConnectionDB  *pgxpool.Pool
	LoggerErrWarn *zap.SugaredLogger
	LoggerInfo    *zap.SugaredLogger
	LoggerTest    *zap.SugaredLogger
}

func (m *InfoMiddleware) PrintURL(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqId := utils.RandomInteger(0, math.MaxInt64)
		m.LoggerInfo.Infof("Method: %s, URL: %s, requestId: %d", string(ctx.Method()), ctx.URI(), reqId)
		ctx.SetUserValue("reqId", reqId)
		h(ctx)
	})
}

func (m *InfoMiddleware) GetId(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		reqIdCtx := ctx.UserValue("reqId")
		var reqId int
		var errorConvert error
		switch reqIdCtx.(type) {
		case string:
			reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
			if errorConvert != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				ctx.Response.SetBody([]byte(errors.ErrAtoi))
				m.LoggerErrWarn.Errorf("Middleware GetId: %s, %v", errors.ErrAtoi, errorConvert)
				return
			}
		case int:
			reqId = reqIdCtx.(int)
		default:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errors.CheckError{
			LoggerErrWarn: m.LoggerErrWarn,
			LoggerInfo:    m.LoggerInfo,
			LoggerTest:    m.LoggerTest,
			RequestId:     &reqId,
		}

		cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		id, err := GetIdByCookie(m.ConnectionDB, &cookieDB)
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
		var reqId int
		var errorConvert error
		switch reqIdCtx.(type) {
		case string:
			reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
			if errorConvert != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				ctx.Response.SetBody([]byte(errors.ErrAtoi))
				m.LoggerErrWarn.Errorf("Middleware Check: %s, %v", errors.ErrAtoi, errorConvert)
				return
			}
		case int:
			reqId = reqIdCtx.(int)
		default:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
		ctx.SetUserValue("reqId", reqId)

		checkError := &errors.CheckError{
			LoggerErrWarn: m.LoggerErrWarn,
			LoggerInfo:    m.LoggerInfo,
			LoggerTest:    m.LoggerTest,
			RequestId:     &reqId,
		}

		cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

		_, err := CheckAccess(m.ConnectionDB, &cookieDB)
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
