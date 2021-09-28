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

func checkErrorSignUp(errIn error, ctx *fasthttp.RequestCtx) error {
	if errIn != nil {
		switch errIn.Error() {
		case ERRUNIQUE:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRUNIQUE)
			return errors.New("fatal")
		case ERRINFOQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINFOQUERY)
			return errors.New("fatal")
		case ERRINFOSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINFOSCAN)
			return errors.New("fatal")
		case ERRINSERTHOSTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINSERTHOSTQUERY)
			return errors.New("fatal")
		case ERRINSERTCOOKIEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCOOKIEQUERY)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINSERTCOOKIEQUERY)
			return errors.New("fatal")
		case ERRINSERTCOURIERQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
			return errors.New("fatal")
		case ERRINSERTCLIENTQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRINSERTCLIENTQUERY)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINSERTCOURIERQUERY)
			return errors.New("fatal")
		}
	}
	return nil
}

func checkErrorLogin(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		switch err.Error() {
		case ERRNOTLOGINORPASSWORD:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRNOTLOGINORPASSWORD,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRNOTLOGINORPASSWORD)
			return errors.New("fatal")
		case ERRINSERTLOGINCOOKIEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRINSERTLOGINCOOKIEQUERY)
			return errors.New("fatal")
		case ERRMAILQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRMAILQUERY)
			return errors.New("fatal")
		case ERRMAILSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRMAILSCAN)
			return errors.New("fatal")
		case ERRMAILPASSQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRMAILPASSQUERY)
			return errors.New("fatal")
		case ERRMAILPASSSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRMAILPASSSCAN)
			return errors.New("fatal")
		case ERRPHONEQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRPHONEQUERY)
			return errors.New("fatal")
		case ERRPHONESCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRPHONESCAN)
			return errors.New("fatal")
		case ERRPHONEPASSQUERY:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRPHONEPASSQUERY)
			return errors.New("fatal")
		case ERRPHONEPASSSCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", ERRPHONEPASSSCAN)
			return errors.New("fatal")
		}
	}
	return nil
}

func checkErrorLogout(err error, ctx *fasthttp.RequestCtx) error {
	if err != nil {
		if err.Error() == ERRDELETEQUERY {
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusBadRequest,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", ERRDELETEQUERY)
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
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return errors.New("fatal")
		case mid.ERRCOOKIESCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
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
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return errors.New("fatal")
		case mid.ERRSIDNOTFOUND:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
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
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEQUERY)
			return errors.New("fatal")
		case mid.ERRCOOKIESCAN:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: ERRDB,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
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
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRCOOKIEEXPIRED)
			return errors.New("fatal")
		case mid.ERRSIDNOTFOUND:
			err := json.NewEncoder(ctx).Encode(&mid.ResultError{
				Status:  http.StatusUnauthorized,
				Explain: ERRAUTH,
			})
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusOK)
				fmt.Printf("Console: %s\n", ERRENCODE)
				return errors.New("fatal")
			}
			ctx.Response.SetStatusCode(http.StatusOK)
			fmt.Printf("Console: %s\n", mid.ERRSIDNOTFOUND)
			return errors.New("fatal")
		}
	}
	return nil
}
