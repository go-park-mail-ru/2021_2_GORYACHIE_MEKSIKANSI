package Middleware

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type InfoMiddleware struct {
	ConnectionDB *pgxpool.Pool
}

func (m *InfoMiddleware) PrintURLMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
		h(ctx)
	})
}

func (m *InfoMiddleware) GetIdByCookieMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
		id, err := GetIdByCookie(m.ConnectionDB, &cookieDB)
		errAccess, resultOutAccess, codeHTTP := errors.CheckErrorCookie(err)
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

		ctx.SetUserValue("id", id)
		h(ctx)
	})
}
