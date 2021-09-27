package Authorization

import (
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type ResultError struct {
	Status	int         `json:"status"`
	Explain	string		`json:"parsedJSON,omitempty"`
}


func checkErrorSignUp(errIn error, ctx *fasthttp.RequestCtx) error {
	if errIn != nil {
		if errIn.Error() == ERRUNIQUE {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRUNIQUE,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRUNIQUE)
			return errors.New("Fatal")
		}

		if errIn.Error() == ERRINFOQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINFOQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINFOQUERY)
			return errors.New("Fatal")
		}
		if errIn.Error() == ERRINFOSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINFOSCAN,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINFOSCAN)
			return errors.New("Fatal")
		}
		if errIn.Error() == ERRINSERTHOSTQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTHOSTQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTHOSTQUERY)
			return errors.New("Fatal")
		}
		if errIn.Error() == ERRINSERTCOOKIEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTCOOKIEQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCOOKIEQUERY)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTCOOKIEQUERY)
			return errors.New("Fatal")
		}
		if errIn.Error() == ERRINSERTCOURIERQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTCOURIERQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
			return errors.New("Fatal")
		}
		if errIn.Error() == ERRINSERTCLIENTQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRINSERTCLIENTQUERY,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCLIENTQUERY)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
			return errors.New("Fatal")
		}
	}
	return nil
}

func checkErrorLogin(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() == ERRNOTLOGINORPASSWORD {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRNOTLOGINORPASSWORD,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRNOTLOGINORPASSWORD)
			return errors.New("Fatal")
		}

		if err.Error() == ERRINSERTLOGINCOOKIEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRINSERTLOGINCOOKIEQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRMAILQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRMAILSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRMAILPASSQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILPASSQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRMAILPASSSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRMAILPASSSCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRPHONEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONEQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRPHONESCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONESCAN)
			return errors.New("Fatal")
		}
		if err.Error() == ERRPHONEPASSQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONEPASSQUERY)
			return errors.New("Fatal")
		}
		if err.Error() == ERRPHONEPASSSCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRPHONEPASSSCAN)
			return errors.New("Fatal")
		}
	}
	return nil
}

func checkErrorLogout(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() == ERRDELETEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRDELETEQUERY)
			return errors.New("Fatal")
		}
	}
	return nil
}

func CheckErrorLoggedIn(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() == mid.ERRCOOKIEQUERY {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return errors.New("Fatal")
		}

		if err.Error() == mid.ERRCOOKIESCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIESCAN)
			return errors.New("Fatal")
		}
		if err.Error() == mid.ERRCOOKIEEXPIRED {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRCOOKIEEXPIRED,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return errors.New("Fatal")
		}

		if err.Error() == mid.ERRSIDNOTFOUND {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return errors.New("Fatal")
		}
	}
	return nil
}

func CheckErrorProfileCookie(err error, ctx *fasthttp.RequestCtx, cookieHTTP *fasthttp.Cookie) error {
	if err != nil {
		if err.Error() == mid.ERRCOOKIEQUERY {
			cookieHTTP.SetKey("session_id")
			cookieHTTP.SetValue(string(ctx.Request.Header.Cookie("session_id")))
			cookieHTTP.SetExpire(time.Now().Add(-72 * time.Hour))
			cookieHTTP.SetHTTPOnly(true)
			cookieHTTP.SetPath("/")
			ctx.Response.Header.SetCookie(cookieHTTP)
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return errors.New("Fatal")
		}

		if err.Error() == mid.ERRCOOKIESCAN {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIESCAN)
			return errors.New("Fatal")
		}
		if err.Error() == mid.ERRCOOKIEEXPIRED {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRCOOKIEEXPIRED,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return errors.New("Fatal")
		}

		if err.Error() == mid.ERRSIDNOTFOUND {
			err := json.NewEncoder(ctx).Encode(&ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("Fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return errors.New("Fatal")
		}
	}
	return nil
}
