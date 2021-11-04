package Authorization

import (
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errorConvert)
			return
		}
	case int:
		reqId = reqIdCtx.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	checkError := &errors.CheckError{
		LoggerErrWarn: u.LoggerErrWarn,
		LoggerInfo:    u.LoggerInfo,
		LoggerTest:    u.LoggerTest,
		RequestId:     &reqId,
	}

	signUpAll := utils.RegistrationRequest{}
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		if config.TEST {
			u.LoggerTest.Errorf("SignUpHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
			return
		}
		u.LoggerErrWarn.Errorf("SignUpHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	cookieHTTP := fasthttp.Cookie{}
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
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("LoginHandler: GetId: %s, %v", errors.ErrAtoi, errorConvert)
			return
		}
	case int:
		reqId = reqIdCtx.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	checkError := &errors.CheckError{
		LoggerErrWarn: u.LoggerErrWarn,
		LoggerInfo:    u.LoggerInfo,
		LoggerTest:    u.LoggerTest,
		RequestId:     &reqId,
	}

	userLogin := utils.Authorization{}
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.LoggerErrWarn.Errorf("LoginHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}
	cookieHTTP := fasthttp.Cookie{}
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
		u.LoggerErrWarn.Errorf("LoginHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("LogoutHandler: GetId: %s, %v", errors.ErrAtoi, errorConvert)
			return
		}
	case int:
		reqId = reqIdCtx.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	checkError := &errors.CheckError{
		LoggerErrWarn: u.LoggerErrWarn,
		LoggerInfo:    u.LoggerInfo,
		LoggerTest:    u.LoggerTest,
		RequestId:     &reqId,
	}

	cookieHTTP := fasthttp.Cookie{}
	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))
	/*	_, err := mid.CheckAccess(u.Application, &cookieDB)
		errAccess, resultOutAccess, codeHTTP := checkError.CheckErrorAccess(err)
		if errAccess != nil {
			switch errAccess.Error() {
			case errors.ErrMarshal:
				ctx.Response.SetStatusCode(codeHTTP)
				ctx.Response.SetBody([]byte(errors.ErrMarshal))

				return
			case errors.ErrCheck:
				ctx.Response.SetStatusCode(codeHTTP)
				ctx.Response.SetBody(resultOutAccess)
				return
			}
		}*/

	err := u.Application.Logout(&cookieDB)
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
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.LoggerErrWarn.Errorf("PayHandler: GetId: %s, %v", errors.ErrAtoi, errorConvert)
			return
		}
	case int:
		reqId = reqIdCtx.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}

	TokenContext := ctx.UserValue("X-Csrf-Token")
	var XCsrfToken string
	switch TokenContext.(type) {
	case string:
		XCsrfToken = TokenContext.(string)
	case int:
		XCsrfToken = strconv.Itoa(TokenContext.(int))
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		u.LoggerErrWarn.Errorf("PayHandler: error: %s, requestId: %d", errors.ErrNotStringAndInt, reqId)
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
