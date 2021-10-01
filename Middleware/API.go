package Middleware

import (
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

const (
	DAYLIVECOOKIE		= 5
	LENSESSINID			= 92
	LENCSRFTOKEN 		= 92
	KEYCOOKIESESSION	= "session_id"
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
	c.DateLife = time.Now().Add(time.Hour * 24 * DAYLIVECOOKIE)
	c.SessionId = randString(LENSESSINID)
	c.CsrfToken = randString(LENCSRFTOKEN)
	return &c
}
