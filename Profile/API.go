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

	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorProfileCookie(err)
	if resultOutAccess != nil {
		switch errAccess.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	profile, err := GetProfile(wrapper, id)
	errOut, resultOutAccess, codeHTTP  := errors.CheckErrorProfile(err)
	if resultOutAccess != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body:   &ProfileResponse{
			profile,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}
