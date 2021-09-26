package Middleware

import (
	"github.com/valyala/fasthttp"
	"math/rand"
	"strings"
	"time"
)

const DAYLIVECOOKIE = 5
const LENSESSINID = 92
const LENCSRFTOKEN = 92

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}

func randString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func (c Defense) GenerateNew() Defense {
	c.DateLife = time.Now().Add(time.Hour * 24 * DAYLIVECOOKIE)
	c.SessionId = randString(LENSESSINID)
	c.CsrfToken = randString(LENCSRFTOKEN)
	return c
}

func SetHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
}
