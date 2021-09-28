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

const(
	ERRDB		=	"ERROR: database is not responding"
	ERRENCODE	=	"ERROR: Encode"
	ERRUNMARSHAL=	"ERROR: unmarshal json"
	ERRAUTH 	= 	"ERROR: authorization failed"
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
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", ERRUNMARSHAL)
		return
	}
	cookieHTTP := fasthttp.Cookie{}
	cookieDB := mid.Defense{}
	cookieDB, errIn := SignUp(wrapper, signUpAll)

	errOut := checkErrorSignUp(errIn, ctx)
	if errOut!= nil{
		return
	}

	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey("session_id")
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
	cookieHTTP.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(&cookieHTTP)

	ctx.Response.Header.Set("X-Csrf-Token", cookieDB.CsrfToken)

	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
		Body: &Registration{
			TypeIn: signUpAll.TypeIn,
			Name:   signUpAll.Name,
			Email:  signUpAll.Email,
			Phone:  signUpAll.Phone,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", ERRENCODE)
		return
	}
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) LoginHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userLogin := Authorization{}
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", ERRUNMARSHAL)
		return
	}
	cookieHTTP := fasthttp.Cookie{}
	cookieDB := mid.Defense{}
	cookieDB, err = Login(wrapper, userLogin)

	errOut := checkErrorLogin(err, ctx)
	if errOut != nil {
		return
	}

	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey("session_id")
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
	ctx.Response.Header.SetCookie(&cookieHTTP)

	ctx.Response.Header.Set("X-CSRF-Token", cookieDB.CsrfToken)

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", ERRENCODE)
		return
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	cookieHTTP := fasthttp.Cookie{}
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	err := Logout(wrapper, cookieDB)
	errOut := checkErrorLogout(err, ctx)
	if errOut != nil {
		return
	}

	cookieHTTP.SetExpire(time.Now().Add(time.Hour * -3))
	cookieHTTP.SetKey("session_id")
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
	cookieHTTP.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(&cookieHTTP)

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})
	if err != nil {
		return 
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func(u *UserInfo) CheckLoggedInHandler(ctx *fasthttp.RequestCtx) {
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	_ , err := mid.GetIdByCookie(u.ConnectionDB, cookieDB)

	errOut := CheckErrorLoggedIn(err, ctx)
	if errOut != nil {
		return
	}
	
	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", ERRENCODE)
		return
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
