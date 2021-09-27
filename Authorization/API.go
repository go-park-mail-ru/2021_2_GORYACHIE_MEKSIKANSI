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
	ERRAUTH = "ERROR: authorization failed"
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

type ResultError struct {
	Status	int         `json:"status"`
	Explain	string		`json:"parsedJSON,omitempty"`
}

/*func checkErrorSignUp(errIn error, ctx *fasthttp.RequestCtx) error {

}*/
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
	if errIn != nil {
		if errIn.Error() == ERRUNIQUE {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRUNIQUE,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRUNIQUE)
			return
		}

		if errIn.Error() == ERRINFOQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINFOQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINFOQUERY)
			return
		}
		if errIn.Error() == ERRINFOSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINFOSCAN,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINFOSCAN)
			return
		}
		if errIn.Error() == ERRINSERTHOSTQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTHOSTQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTHOSTQUERY)
			return
		}
		if errIn.Error() == ERRINSERTCOOKIEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTCOOKIEQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCOOKIEQUERY)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTCOOKIEQUERY)
			return
		}
		if errIn.Error() == ERRINSERTCOURIERQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTCOURIERQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
			return
		}
		if errIn.Error() == ERRINSERTCLIENTQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTCLIENTQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCLIENTQUERY)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
			return
		}
	}

	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey("session_id")
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
	cookieHTTP.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(&cookieHTTP)

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
	cookieDB, err = Login(wrapper, userLogin) // TODO: проверки на ошибки
	if err != nil {
		if err.Error() == ERRNOTLOGINORPASSWORD {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRNOTLOGINORPASSWORD,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRNOTLOGINORPASSWORD)
			return
		}

		if err.Error() == ERRINSERTLOGINCOOKIEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTLOGINCOOKIEQUERY)
			return
		}
		if err.Error() == ERRMAILQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILQUERY)
			return
		}
		if err.Error() == ERRMAILSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILSCAN)
			return
		}
		if err.Error() == ERRMAILPASSQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILPASSQUERY)
			return
		}
		if err.Error() == ERRMAILPASSSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILPASSSCAN)
			return
		}
		if err.Error() == ERRPHONEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONEQUERY)
			return
		}
		if err.Error() == ERRPHONESCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONESCAN)
			return
		}
		if err.Error() == ERRPHONEPASSQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONEPASSQUERY)
			return
		}
		if err.Error() == ERRPHONEPASSSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONEPASSSCAN)
			return
		}
	}

	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey("session_id")
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
	ctx.Response.Header.SetCookie(&cookieHTTP)

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
	ctx.Response.Header.Peek("X-CSRF-Token")

	err := Logout(wrapper, cookieDB)
	if err != nil {
		if err.Error() == ERRDELETEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRDELETEQUERY)
		}
	}

	cookieHTTP.SetExpire(time.Now().Add(time.Hour * -3)) //TODO: уменьшить день
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
	// get cookie from request
	cookieDB := mid.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}

	_ , err := mid.GetIdByCookie(u.ConnectionDB, cookieDB)
	if err != nil {
		if err.Error() == mid.ERRCOOKIEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return
		}

		if err.Error() == mid.ERRCOOKIESCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIESCAN)
			return
		}
		if err.Error() == mid.ERRCOOKIEEXPIRED {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRCOOKIEEXPIRED,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return
		}

		if err.Error() == mid.ERRSIDNOTFOUND {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return
		}
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
