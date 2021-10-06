package Profile

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type Profile struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday,omitempty"`
}

type ProfileResponse struct {
	ProfileUser	interface{}	`json:"profile"`
}


type ProfileInfo struct {
	ConnectionDB *pgxpool.Pool
}

func (u *ProfileInfo) ProfileHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}

	id, err := mid.GetIdByCookie(u.ConnectionDB, &cookieDB)

	errOut := errors.CheckErrorProfileCookie(err, ctx)
	if errOut != nil {
		return
	}

	profile, err := GetProfile(wrapper, id)
	err = errors.CheckErrorProfile(err, ctx)
	if err != nil {
		return
	}

	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body:   &ProfileResponse{
			profile,
		},
	})
	ctx.Response.SetStatusCode(http.StatusOK)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
