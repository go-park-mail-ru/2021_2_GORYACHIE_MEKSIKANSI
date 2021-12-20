//go:generate mockgen -destination=mocks/api.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/myerror MultiLogger
//go:generate mockgen -destination=mocks/apiApplication.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/application AuthorizationApplicationInterface
package api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/application"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"encoding/json"
	"github.com/fasthttp/websocket"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type AuthorizationApiInterface interface {
	SignUpHandler(ctx *fasthttp.RequestCtx)
	LoginHandler(ctx *fasthttp.RequestCtx)
	LogoutHandler(ctx *fasthttp.RequestCtx)
	PayHandler(ctx *fasthttp.RequestCtx)
	UserWebSocket(ctx *fasthttp.RequestCtx)
	UserWebSocketNewKey(ctx *fasthttp.RequestCtx)
}

type UserInfo struct {
	Application appPkg.AuthorizationApplicationInterface
	Logger      errPkg.MultiLogger
	IntCh       chan authorization.WebSocketOrder
}

func (u *UserInfo) SignUpHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var signUpAll authorization.RegistrationRequest
	err := easyjson.Unmarshal(ctx.Request.Body(), &signUpAll)
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
	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusCreated,
		Body: &authorization.RegistrationResponse{
			User: authorization.UserConvertRegistration(&signUpAll),
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, errResponse.Error(), reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	util.SetCookieResponse(&cookieHTTP, *cookieDB, util.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-Csrf-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
	//ctx.Response.Header.SetContentType("application/json")
}

func (u *UserInfo) LoginHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userLogin authorization.Authorization
	err := easyjson.Unmarshal(ctx.Request.Body(), &userLogin)
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, errResponse.Error(), reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	util.SetCookieResponse(&cookieHTTP, *cookieDB, util.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.Header.Set("X-CSRF-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *UserInfo) LogoutHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var cookieHTTP fasthttp.Cookie
	var cookieDB util.Defense

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvertToken := util.InterfaceConvertString(tokenContext)
	if errConvertToken != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertToken.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvertToken.Error(), reqId)
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

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, errResponse.Error(), reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	cookieDB.DateLife = time.Now().Add(time.Hour * -3)
	util.SetCookieResponse(&cookieHTTP, cookieDB, util.KeyCookieSessionId)
	ctx.Response.Header.SetCookie(&cookieHTTP)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *UserInfo) PayHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
		return
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvertToken := util.InterfaceConvertString(tokenContext)
	if errConvertToken != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvertToken.Error()))
		u.Logger.Errorf("%s, requestId: %d", errConvertToken.Error(), reqId)
		return
	}

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, errResponse, reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *UserInfo) UserWebSocketNewKey(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := util.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}

	newKey, err := u.Application.NewCSRFWebsocket(id)

	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorNewWsKey(err)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	response, errResponse := easyjson.Marshal(&authorization.Result{
		Status: http.StatusOK,
		Body: &authorization.WebSocket{
			Socket: authorization.KeyWebSocket{
				Key: newKey,
			},
		},
	})
	if errResponse != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		u.Logger.Errorf("%s, %s, requestId: %d", errPkg.ErrEncode, errResponse.Error(), reqId)
		return
	}

	ctx.Response.SetBody(response)
	json.NewEncoder(ctx)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (u *UserInfo) UserWebSocket(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		u.Logger.Errorf("%s", errConvert.Error())
	}
	upgrade := websocket.FastHTTPUpgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			return true
		},
	}
	errUpgrade := upgrade.Upgrade(ctx, func(conn *websocket.Conn) {
		for {
			statusOrder := <-u.IntCh

			messageWS := &authorization.Result{
				Status: http.StatusOK,
				Body: authorization.WebSocket{
					Socket: authorization.WebSocketAction{
						Action: "status",
						Order:  statusOrder,
					},
				},
			}

			err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				u.IntCh <- statusOrder
				return
			}

			errWrite := conn.WriteJSON(messageWS)
			if errWrite != nil {
				u.IntCh <- statusOrder
				u.Logger.Errorf("WriteJSON %s, requestId: %d", errWrite.Error(), reqId)
				return
			}
		}
	})
	if errUpgrade != nil {
		u.Logger.Errorf("UpgradeWS %s, requestId: %d", errUpgrade.Error(), reqId)
		return
	}

}
