package Profile

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
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
	cookieHttp := fasthttp.Cookie{}
	wrapper := Wrapper{Conn: u.ConnectionDB}
	profile := Profile{}
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}

	id, err := mid.GetIdByCookie(u.ConnectionDB, cookieDB)
	if id == 0 {
		panic("id not found")
		return // TODO(N): подправить
	}

	if err != nil {
		// TODO: Сделай с этим что-нибудь
	}

	if id == -1 {  // просрочена кука
		cookieHttp.SetKey("session_id")
		cookieHttp.SetValue(string(ctx.Request.Header.Cookie("session_id")))
		cookieHttp.SetExpire(time.Now().Add(-72 * time.Hour))
		cookieHttp.SetHTTPOnly(true)
		cookieHttp.SetPath("/")
		ctx.Response.Header.SetCookie(&cookieHttp)
	}
	// -2 - смотреть err


	profile, _ /*err*/ = GetProfile(wrapper, id) // TODO: проверки на ошибки
	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body:   profile,
	})
	if err != nil {
		return 
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
