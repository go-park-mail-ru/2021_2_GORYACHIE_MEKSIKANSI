package Authorization

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserInfo struct {
	Application   interfaces.AuthorizationApplication
	LoggerErrWarn *zap.SugaredLogger
	LoggerInfo    *zap.SugaredLogger
	LoggerTest    *zap.SugaredLogger
}

func (u *UserInfo) SignUpHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		LoggerErrWarn: u.LoggerErrWarn,
		LoggerInfo:    u.LoggerInfo,
		LoggerTest:    u.LoggerTest,
		RequestId:     &reqId,
	}

	var signUpAll utils.RegistrationRequest
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.LoggerErrWarn.Errorf("SignUpHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
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
			User: &utils.User{
				TypeUser: signUpAll.TypeUser,
				Name:     signUpAll.Name,
				Email:    signUpAll.Email,
				Phone:    signUpAll.Phone,
			},
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.LoggerErrWarn.Errorf("SignUpHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
	//ctx.Response.Header.SetContentType("application/json")
}

func (u *UserInfo) LoginHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	intConvert, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		LoggerErrWarn: u.LoggerErrWarn,
		LoggerInfo:    u.LoggerInfo,
		LoggerTest:    u.LoggerTest,
		RequestId:     &intConvert,
	}

	var userLogin utils.Authorization
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.LoggerErrWarn.Errorf("LoginHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, intConvert)
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
		u.LoggerErrWarn.Errorf("LoginHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, intConvert)
		return
	}

}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		LoggerErrWarn: u.LoggerErrWarn,
		LoggerInfo:    u.LoggerInfo,
		LoggerTest:    u.LoggerTest,
		RequestId:     &reqId,
	}

	var cookieHTTP fasthttp.Cookie
	var cookieDB utils.Defense

	TokenContext := ctx.UserValue("X-Csrf-Token")
	XCsrfToken, errConvert := utils.InterfaceConvertString(TokenContext)
	if (errConvert != nil) && (errConvert.Error() == errors.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	var err error
	cookieDB.SessionId, err = u.Application.Logout(XCsrfToken)
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
		u.LoggerErrWarn.Errorf("LogoutHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *UserInfo) PayHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	TokenContext := ctx.UserValue("X-Csrf-Token")
	XCsrfToken, errConvert := utils.InterfaceConvertString(TokenContext)
	if (errConvert != nil) && (errConvert.Error() == errors.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	ctx.Response.Header.Set("X-CSRF-Token", XCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
	err := json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.LoggerErrWarn.Errorf("PayHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
}
