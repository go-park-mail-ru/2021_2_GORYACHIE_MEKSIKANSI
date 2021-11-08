package Profile

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	_ "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

type InfoProfile struct {
	Application interfaces.ProfileApplication
	Logger      errors.MultiLogger
}

func (u *InfoProfile) ProfileHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	profile, err := u.Application.GetProfile(id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfile(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
		Body: &utils.ProfileResponse{
			ProfileUser: profile,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("ProfileHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *InfoProfile) UpdateUserName(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userName utils.UpdateName
	err := json.Unmarshal(ctx.Request.Body(), &userName)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("UpdateUserName: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err = u.Application.UpdateName(id, userName.Name)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdateName(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserName: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *InfoProfile) UpdateUserEmail(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userEmail utils.UpdateEmail
	err := json.Unmarshal(ctx.Request.Body(), &userEmail)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("UpdateUserEmail: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err = u.Application.UpdateEmail(id, userEmail.Email)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdateEmail(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserEmail: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *InfoProfile) UpdateUserPassword(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("UpdateUserPassword: GetId: %s, %v", errors.ErrAtoi, errorConvert)
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
		Logger:    u.Logger,
		RequestId: reqId,
	}

	userPassword := utils.UpdatePassword{}
	err := json.Unmarshal(ctx.Request.Body(), &userPassword)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("UpdateUserPassword: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err = u.Application.UpdatePassword(id, userPassword.Password)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdatePassword(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserPassword: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *InfoProfile) UpdateUserPhone(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userPhone utils.UpdatePhone
	err := json.Unmarshal(ctx.Request.Body(), &userPhone)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("UpdateUserPhone: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err = u.Application.UpdatePhone(id, userPhone.Phone)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdatePhone(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserPhone: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *InfoProfile) UpdateUserAvatar(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}
	headerAvatar, errFile := ctx.FormFile("avatar")
	if errFile != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotSearchAvatar))
		return
	}
	var userAvatar utils.UpdateAvatar
	userAvatar.FileHeader = headerAvatar

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err := u.Application.UpdateAvatar(id, &userAvatar)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdateAvatar(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
		Body: &utils.UpdateAvatarRequest{
			PathImg: userAvatar.Avatar,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserAvatar: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (u *InfoProfile) UpdateUserBirthday(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	var userBirthday utils.UpdateBirthday
	err := json.Unmarshal(ctx.Request.Body(), &userBirthday)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("UpdateUserBirthday: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err = u.Application.UpdateBirthday(id, userBirthday.Birthday)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdateBirthday(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserBirthday: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
}

func (u *InfoProfile) UpdateUserAddress(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    u.Logger,
		RequestId: reqId,
	}

	userAddress := utils.UpdateAddress{}
	err := json.Unmarshal(ctx.Request.Body(), &userAddress)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		u.Logger.Errorf("UpdateUserAddress: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			u.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
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

	err = u.Application.UpdateAddress(id, userAddress.Address)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorProfileUpdateAddress(err)
	if errOut != nil {
		switch errOut.Error() {
		case errors.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errors.ErrMarshal))
			return
		case errors.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOutAccess)
			return
		}
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.ResponseStatus{
		StatusHTTP: http.StatusOK,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		u.Logger.Errorf("UpdateUserAddress: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
}
