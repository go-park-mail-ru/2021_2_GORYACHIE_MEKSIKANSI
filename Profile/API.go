package Profile

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"time"
)

type Profile struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`
}

type ProfileInfo struct {
	ConnectionDB *pgxpool.Pool
}

func (u *ProfileInfo) ProfileHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	profile := Profile{}
	cookieHttp := fasthttp.Cookie{}
	cookieDB := mid.Defense{DateLife: cookieHttp.Expire(), SessionId: string(cookieHttp.Value())}
	id, _ := mid.GetIdByCookie(u.ConnectionDB, cookieDB)

	profile, _ /*err*/ = GetProfile(wrapper, id) // TODO: проверки на ошибки
	if profile.Email != "" {                     // TODO: заглушка на unused
		ctx.Response.SetStatusCode(http.StatusBadRequest)
	}
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус

	mid.SetHeaders(ctx)
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
