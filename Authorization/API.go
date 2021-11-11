package Authorization

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"

	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type UserInfo struct {
	Application interfaces.AuthorizationApplication
	Logger      errors.MultiLogger
}

func (u *UserInfo) SignUpHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errConvert.Error(), errConvert)
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var signUpAll utils.RegistrationRequest
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("SignUpHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	var cookieHTTP fasthttp.Cookie
	cookieDB, errIn := u.Application.SignUp(&signUpAll)

	errOut, resultOut, codeHTTP := checkError.CheckErrorSignUp(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	utils.SetCookieResponse(&cookieHTTP, *cookieDB, utils.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-Csrf-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusCreated,
		Body: &utils.RegistrationResponse{
			User: utils.UserConvertRegistration(&signUpAll),
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("SignUpHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
	//ctx.Response.Header.SetContentType("application/json")
}

func (u *UserInfo) LoginHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errConvert.Error(), errConvert)
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userLogin utils.Authorization
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("LoginHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}
	var cookieHTTP fasthttp.Cookie
	cookieDB, err := u.Application.Login(&userLogin)

	errOut, resultOut, codeHTTP := checkError.CheckErrorLogin(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	utils.SetCookieResponse(&cookieHTTP, *cookieDB, utils.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-CSRF-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("LoginHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errConvert.Error(), errConvert)
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var cookieHTTP fasthttp.Cookie
	var cookieDB utils.Defense

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := utils.InterfaceConvertString(tokenContext)
	if (errConvert != nil) && (errConvert.Error() == errors.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	var err error
	cookieDB.SessionId, err = u.Application.Logout(xCsrfToken)
	errOut, resultOut, codeHTTP := checkError.CheckErrorLogout(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	cookieDB.DateLife = time.Now().Add(time.Hour * -3)
	utils.SetCookieResponse(&cookieHTTP, cookieDB, utils.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.SetStatusCode(http.StatusOK)

	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("LogoutHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *UserInfo) PayHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errConvert.Error(), errConvert)
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := utils.InterfaceConvertString(tokenContext)
	if (errConvert != nil) && (errConvert.Error() == errors.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
	err := json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("PayHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
}
