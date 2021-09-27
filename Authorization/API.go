package Authorization

import (
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
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

type Authorization struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password"`
}

type Result struct {
	Status int         `json:"status,omitempty"`
	Body   interface{} `json:"parsedJSON,omitempty"`
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
	cookieDB := mid.Defense{}
	cookieDB, _ /*err*/ = SignUp(wrapper, signUpAll)

	cookieHttp.SetExpire(cookieDB.DateLife)
	cookieHttp.SetKey("session_id")
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	cookieHttp.SetPath("/")
	ctx.Response.Header.SetCookie(&cookieHttp)

	ctx.Response.SetStatusCode(http.StatusOK)


	json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
		Body: &Registration{
			TypeIn: signUpAll.TypeIn,
			Name: signUpAll.Name,
			Email: signUpAll.Email,
			Phone: signUpAll.Phone,
		},
	})
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
	cookieDB := mid.Defense{}
	cookieDB, _ /*err*/ = Login(wrapper, userLogin) // TODO: проверки на ошибки

	cookieHttp.SetExpire(cookieDB.DateLife)
	cookieHttp.SetKey("session_id")
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	cookieHttp.SetPath("/")
	ctx.Response.Header.SetCookie(&cookieHttp)

	ctx.Response.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}

	cookieHttp := fasthttp.Cookie{}
	cookieDB := mid.Defense{DateLife: cookieHttp.Expire(), SessionId: string(cookieHttp.Value())}
	_ /*err*/ = Logout(wrapper, cookieDB) // TODO: проверки на ошибки

	ctx.Response.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})
	cookieHttp.SetExpire(cookieDB.DateLife) //TODO: уменьшить день
	cookieHttp.SetKey("session_id")
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	cookieHttp.SetPath("/")
	ctx.Response.Header.SetCookie(&cookieHttp)

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func(u *UserInfo) CheckLoggedInHandler(ctx *fasthttp.RequestCtx) {
	// get cookie from request
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	// get cookie from database
	id, err := mid.GetIdByCookie(u.ConnectionDB, cookieDB)
	if id == 0 {
		ctx.Response.SetStatusCode(http.StatusUnauthorized)
		fmt.Printf("Console:  method: %s, url: %s\n Cookie not found. User not authorized", string(ctx.Method()), ctx.URI())
		return // TODO(N): подправить
	}

	if err != nil {
		// TODO: Сделай с этим что-нибудь
	}

	if id == -1 {  // просрочена кука
		ctx.Response.SetStatusCode(http.StatusUnauthorized)
		fmt.Printf("Console:  method: %s, url: %s\n Cookie is expired. User not authorized", string(ctx.Method()), ctx.URI())
		return // TODO(N): подправить
	}
	// -2 - смотреть err

	ctx.Response.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}