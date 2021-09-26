package Profile

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	middleware "project/Middleware"
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
	Id           int
	ConnectionDB *pgxpool.Pool
}

func (u *ProfileInfo) ProfileHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	profile := Profile{}

	profile, _ /*err*/ = GetProfile(wrapper, u.Id) // TODO: проверки на ошибки
	if profile.Email != "" {                       // TODO: заглушка на unused
		ctx.Response.SetStatusCode(http.StatusBadRequest)
	}
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус

	middleware.SetHeaders(ctx)
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
