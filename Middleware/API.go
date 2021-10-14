package Middleware

import (
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

const (
	DayLiveCookie 		= 5
	LenSessionId		= 92
	LenCsrfToken		= 92
	KeyCookieSessionId	= "session_id"
)

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}

func SetCookieResponse(cookieHTTP *fasthttp.Cookie, cookieDB Defense, sessionId string) {
	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey(sessionId)
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
}

func randString(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[RandomInteger(0, len(chars))])
	}

	return b.String()
}

func (c Defense) GenerateNew() *Defense {
	c.DateLife = time.Now().Add(time.Hour * 24 * DayLiveCookie)
	c.SessionId = randString(LenSessionId)
	c.CsrfToken = randString(LenCsrfToken)
	return &c
}
