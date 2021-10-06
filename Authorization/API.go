package Authorization

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
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

type RegistrationRequest struct {
	TypeUser string		`json:"type"`
	Name     string 	`json:"name"`
	Email    string		`json:"email"`
	Phone    string		`json:"phone"`
	Password string		`json:"password"`
}

type User struct {
	TypeUser string		`json:"type"`
	Name     string    	`json:"name"`
	Email    string		`json:"email"`
	Phone    string		`json:"phone"`
}

type RegistrationResponse struct {
	User	interface{}	`json:"user"`
}

type Authorization struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Result struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}

func (u *UserInfo) SignUpHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	signUpAll := RegistrationRequest{}
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieHTTP := fasthttp.Cookie{}
	cookieDB, errIn := SignUp(wrapper, &signUpAll)

	errOut := errors.CheckErrorSignUp(errIn, ctx)
	if errOut != nil{
		return
	}
	mid.SetCookieResponse(&cookieHTTP, *cookieDB, mid.KEYCOOKIESESSION)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-Csrf-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
		Body: &RegistrationResponse{
			User: &User{
				TypeUser: signUpAll.TypeUser,
				Name:     signUpAll.Name,
				Email:    signUpAll.Email,
				Phone:    signUpAll.Phone,
			},
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", errors.ErrEncode)
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
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}
	cookieHTTP := fasthttp.Cookie{}
	cookieDB, err := Login(wrapper, &userLogin)

	errOut := errors.CheckErrorLogin(err, ctx)
	if errOut != nil {
		return
	}

	mid.SetCookieResponse(&cookieHTTP, *cookieDB, mid.KEYCOOKIESESSION)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-CSRF-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	cookieHTTP := fasthttp.Cookie{}
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))
	_, err := mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess := errors.CheckErrorLogoutAccess(err, ctx)
	if errAccess != nil {
		return
	}

	err = Logout(wrapper, &cookieDB)
	errOut := errors.CheckErrorLogout(err, ctx)
	if errOut != nil {
		return
	}

	cookieDB.DateLife = time.Now().Add(time.Hour * -3)
	mid.SetCookieResponse(&cookieHTTP, cookieDB, mid.KEYCOOKIESESSION)
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
