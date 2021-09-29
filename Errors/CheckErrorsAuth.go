package Errors


import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func CheckErrorSignUp(errIn error, ctx *fasthttp.RequestCtx) error {
	if errIn != nil {
		switch errIn.Error() {
		case auth.ERRUNIQUE:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusConflict,
				Explain: auth.ERRUNIQUE,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", auth.ERRUNIQUE)
			return errors.New("fatal")
		case auth.ERRINFOQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINFOQUERY)
			return errors.New("fatal")
		case auth.ERRINFOSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINFOSCAN)
			return errors.New("fatal")
		case auth.ERRINSERTHOSTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINSERTHOSTQUERY)
			return errors.New("fatal")
		case auth.ERRINSERTCOOKIEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRINSERTCOOKIEQUERY)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINSERTCOOKIEQUERY)
			return errors.New("fatal")
		case auth.ERRINSERTCOURIERQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRINSERTCOURIERQUERY)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINSERTCOURIERQUERY)
			return errors.New("fatal")
		case auth.ERRINSERTCLIENTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINSERTCOURIERQUERY)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogin(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case auth.ERRNOTLOGINORPASSWORD:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: auth.ERRNOTLOGINORPASSWORD,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", auth.ERRNOTLOGINORPASSWORD)
			return errors.New("fatal")
		case auth.ERRINSERTLOGINCOOKIEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRINSERTLOGINCOOKIEQUERY)
			return errors.New("fatal")
		case auth.ERRMAILQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRMAILQUERY)
			return errors.New("fatal")
		case auth.ERRMAILSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRMAILSCAN)
			return errors.New("fatal")
		case auth.ERRMAILPASSQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRMAILPASSQUERY)
			return errors.New("fatal")
		case auth.ERRMAILPASSSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRMAILPASSSCAN)
			return errors.New("fatal")
		case auth.ERRPHONEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRPHONEQUERY)
			return errors.New("fatal")
		case auth.ERRPHONESCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRPHONESCAN)
			return errors.New("fatal")
		case auth.ERRPHONEPASSQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRPHONEPASSQUERY)
			return errors.New("fatal")
		case auth.ERRPHONEPASSSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", auth.ERRPHONEPASSSCAN)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogout(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() != auth.ERRDELETEQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusBadRequest,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", auth.ERRDELETEQUERY)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLogoutAccess(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case mid.ERRCOOKIEANDCSRFQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEANDCSRFQUERY)
			return errors.New("fatal")
		case mid.ERRCOOKIEANDCSRFSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEANDCSRFSCAN)
			return errors.New("fatal")
		case mid.ERRSIDNOTFOUND:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRSIDNOTFOUND,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusUnauthorized)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorLoggedIn(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case mid.ERRCOOKIEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return errors.New("fatal")
		case mid.ERRCOOKIESCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIESCAN)
			return errors.New("fatal")
		case mid.ERRCOOKIEEXPIRED:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRCOOKIEEXPIRED,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return errors.New("fatal")
		case mid.ERRSIDNOTFOUND:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: auth.ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return errors.New("fatal")
		}
	}
	return nil
}

func CheckErrorProfileCookie(err error, ctx *fasthttp.RequestCtx, cookieHTTP *fasthttp.Cookie) error {
	if err != nil {
		switch err.Error() {
		case mid.ERRCOOKIEQUERY:
			cookieHTTP.SetKey("session_id")
			cookieHTTP.SetValue(string(ctx.Request.Header.Cookie("session_id")))
			cookieHTTP.SetExpire(time.Now().Add(-72 * time.Hour))
			cookieHTTP.SetHTTPOnly(true)
			cookieHTTP.SetPath("/")
			ctx.Response.Header.SetCookie(cookieHTTP)
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return errors.New("fatal")
		case mid.ERRCOOKIESCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: auth.ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIESCAN)
			return errors.New("fatal")
		case mid.ERRCOOKIEEXPIRED:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: mid.ERRCOOKIEEXPIRED,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return errors.New("fatal")
		case mid.ERRSIDNOTFOUND:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: auth.ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusInternalServerError)
				fmt.Printf("Console: %s\n", auth.ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return errors.New("fatal")
		}
	}
	return nil
}
