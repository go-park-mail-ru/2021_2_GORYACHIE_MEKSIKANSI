package Middleware

import "github.com/valyala/fasthttp"

func SetHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
}
