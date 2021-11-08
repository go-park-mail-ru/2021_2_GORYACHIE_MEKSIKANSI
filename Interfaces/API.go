package Interfaces

import "github.com/valyala/fasthttp"

type AuthorizationAPI interface {
	SignUpHandler(ctx *fasthttp.RequestCtx)
	LoginHandler(ctx *fasthttp.RequestCtx)
	LogoutHandler(ctx *fasthttp.RequestCtx)
	PayHandler(ctx *fasthttp.RequestCtx)
}

type CartApi interface {
	GetCartHandler(ctx *fasthttp.RequestCtx)
	UpdateCartHandler(ctx *fasthttp.RequestCtx)
}

type ProfileAPI interface {
	ProfileHandler(ctx *fasthttp.RequestCtx)
	UpdateUserName(ctx *fasthttp.RequestCtx)
	UpdateUserEmail(ctx *fasthttp.RequestCtx)
	UpdateUserPassword(ctx *fasthttp.RequestCtx)
	UpdateUserPhone(ctx *fasthttp.RequestCtx)
	UpdateUserAvatar(ctx *fasthttp.RequestCtx)
	UpdateUserBirthday(ctx *fasthttp.RequestCtx)
	UpdateUserAddress(ctx *fasthttp.RequestCtx)
}

type MiddlewareAPI interface {
	PrintURL(h fasthttp.RequestHandler) fasthttp.RequestHandler
	GetId(h fasthttp.RequestHandler) fasthttp.RequestHandler
	Check(h fasthttp.RequestHandler) fasthttp.RequestHandler
}

type RestaurantAPI interface {
	RestaurantHandler(ctx *fasthttp.RequestCtx)
	RestaurantIdHandler(ctx *fasthttp.RequestCtx)
	RestaurantDishesHandler(ctx *fasthttp.RequestCtx)
}