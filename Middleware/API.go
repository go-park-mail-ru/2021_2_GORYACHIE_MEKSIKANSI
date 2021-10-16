package Middleware

import (
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"github.com/valyala/fasthttp"
)

const (
	KeyCookieSessionId	= "session_id"
)

func SetCookieResponse(cookieHTTP *fasthttp.Cookie, cookieDB utils.Defense, sessionId string) {
	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey(sessionId)
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
}
