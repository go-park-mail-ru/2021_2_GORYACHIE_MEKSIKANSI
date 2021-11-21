package Api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"

	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type UserInfo struct {
	Application interfaces.AuthorizationApplication
	Logger      errPkg.MultiLogger
}

func (u *UserInfo) SignUpHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var signUpAll Authorization.RegistrationRequest
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrUnmarshal, err.Error(), reqId)
		return
	}

	var cookieHTTP fasthttp.Cookie
	cookieDB, errIn := u.Application.SignUp(&signUpAll)

	errOut, resultOut, codeHTTP := checkError.CheckErrorSignUp(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusCreated,
		Body: &Authorization.RegistrationResponse{
			User: Authorization.UserConvertRegistration(&signUpAll),
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, err.Error(), reqId)
		return
	}

	Util.SetCookieResponse(&cookieHTTP, *cookieDB, Util.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-Csrf-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
	//ctx.Response.Header.SetContentType("application/json")
}

func (u *UserInfo) LoginHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userLogin Authorization.Authorization
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrUnmarshal, err.Error(), reqId)
		return
	}
	var cookieHTTP fasthttp.Cookie
	cookieDB, err := u.Application.Login(&userLogin)

	errOut, resultOut, codeHTTP := checkError.CheckErrorLogin(err)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, err.Error(), reqId)
		return
	}

	Util.SetCookieResponse(&cookieHTTP, *cookieDB, Util.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-CSRF-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var cookieHTTP fasthttp.Cookie
	var cookieDB Util.Defense

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := Util.InterfaceConvertString(tokenContext)
	if (errConvert != nil) && (errConvert.Error() == errPkg.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrNotStringAndInt))
		return
	}

	var err error
	cookieDB.SessionId, err = u.Application.Logout(xCsrfToken)
	errOut, resultOut, codeHTTP := checkError.CheckErrorLogout(err)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	cookieDB.DateLife = time.Now().Add(time.Hour * -3)
	Util.SetCookieResponse(&cookieHTTP, cookieDB, Util.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, err.Error(), reqId)
		return
	}

}

func (u *UserInfo) PayHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := Util.InterfaceConvertString(tokenContext)
	if (errConvert != nil) && (errConvert.Error() == errPkg.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrNotStringAndInt))
		return
	}

	err := json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}