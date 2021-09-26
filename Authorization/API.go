package Authorization

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	middleware "project/Middleware"
	"time"
)

type UserInfo struct {
	ConnectionDB *pgxpool.Pool
}

type Registration struct {
	TypeIn   string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday,omitempty"`
}

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}

type Authorization struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password"`
}

func (u *UserInfo) SignUpHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	signUpAll := Registration{}
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	cookieHttp := fasthttp.Cookie{}
	cookieDB := Defense{}
	cookieDB, _ /*err*/ = SignUp(wrapper, signUpAll)

	cookieHttp.SetExpire(cookieDB.DateLife)
	cookieHttp.SetKey("session_id")
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус
	// json.NewEncoder()
	middleware.SetHeaders(ctx)
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) LoginHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userLogin := Authorization{}
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	cookieHttp := fasthttp.Cookie{}
	cookieDB := Defense{}
	cookieDB, _ /*err*/ = Login(wrapper, userLogin) // TODO: проверки на ошибки

	cookieHttp.SetExpire(cookieDB.DateLife)
	cookieHttp.SetKey("session_id")
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус

	middleware.SetHeaders(ctx)
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}

	cookieHttp := fasthttp.Cookie{}
	cookieDB := Defense{DateLife: cookieHttp.Expire(), SessionId: string(cookieHttp.Value())}
	_ /*err*/ = Logout(wrapper, cookieDB) // TODO: проверки на ошибки

	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус
	middleware.SetHeaders(ctx)
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
	// TODO: отдать просроченную куку, key=value, sessionId=432423
}
