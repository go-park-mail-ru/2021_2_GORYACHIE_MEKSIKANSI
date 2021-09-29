package Profile

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	error "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type Profile struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday,omitempty"`
}

type ProfileInfo struct {
	ConnectionDB *pgxpool.Pool
}

func (u *ProfileInfo) ProfileHandler(ctx *fasthttp.RequestCtx) {
	cookieHTTP := fasthttp.Cookie{}
	wrapper := Wrapper{Conn: u.ConnectionDB}
	profile := Profile{}
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}

	id, err := mid.GetIdByCookie(u.ConnectionDB, cookieDB)

	errOut := error.CheckErrorProfileCookie(err, ctx, &cookieHTTP)
	if errOut != nil {
		return
	}

	profile, err = GetProfile(wrapper, id)
	err = error.CheckErrorProfile(err, ctx)
	if err != nil {
		return
	}

	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body:   profile,
	})
	ctx.Response.SetStatusCode(http.StatusOK)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", auth.ERRENCODE)
		return 
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
