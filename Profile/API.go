package Profile

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

type InfoProfile struct {
	ConnectionDB *pgxpool.Pool
}

func (u *InfoProfile) ProfileHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	// TODO(N): add x-csrf-

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	profile, err := GetProfile(&wrapper, id)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfile(err)
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
	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body: &utils.ProfileResponse{
			ProfileUser: profile,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (u *InfoProfile) UpdateUserName(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userName := utils.UpdateName{}
	err := json.Unmarshal(ctx.Request.Body(), &userName)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	_, err = mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
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
	}

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	err = UpdateName(&wrapper, id, userName.Name)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdateName(err)
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
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (u *InfoProfile) UpdateUserEmail(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userEmail := utils.UpdateEmail{}
	err := json.Unmarshal(ctx.Request.Body(), &userEmail)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	_, err = mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
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
	}

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	err = UpdateEmail(&wrapper, id, userEmail.Email)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdateEmail(err)
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
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (u *InfoProfile) UpdateUserPassword(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userPassword := utils.UpdatePassword{}
	err := json.Unmarshal(ctx.Request.Body(), &userPassword)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	_, err = mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
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
	}

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	err = UpdatePassword(&wrapper, id, userPassword.Password)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdatePassword(err)
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
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (u *InfoProfile) UpdateUserPhone(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userPhone := utils.UpdatePhone{}
	err := json.Unmarshal(ctx.Request.Body(), &userPhone)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	_, err = mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
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
	}

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	err = UpdatePhone(&wrapper, id, userPhone.Phone)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdatePhone(err)
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
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (u *InfoProfile) UpdateUserAvatar(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userAvatar := utils.UpdateAvatar{}
	err := json.Unmarshal(ctx.Request.Body(), &userAvatar)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	_, err = mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
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
	}

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	err = UpdateAvatar(&wrapper, id, userAvatar.Avatar)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdateAvatar(err)
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
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (u *InfoProfile) UpdateUserBirthday(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.ConnectionDB}
	userBirthday := utils.UpdateBirthday{}
	err := json.Unmarshal(ctx.Request.Body(), &userBirthday) // TODO(N): birthday проверить
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
		return
	}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	_, err = mid.CheckAccess(u.ConnectionDB, &cookieDB)
	errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
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
	}

	idUrl := ctx.UserValue("id")
	var id int
	var errorConvert error
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			fmt.Printf("Console: %s\n", errors.ErrAtoi)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotString))
		fmt.Printf("Console: %s\n", errors.ErrNotString)
		return
	}

	err = UpdateBirthday(&wrapper, id, userBirthday.Birthday)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdateBirthday(err)
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
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}
}
