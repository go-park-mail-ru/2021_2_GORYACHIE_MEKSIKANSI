package Middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func CheckAuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
		h(ctx)
	})
}
