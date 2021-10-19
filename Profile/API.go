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

	id, errorConvert:= strconv.Atoi(ctx.UserValue("id").(string))
	if errorConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrAtoi))
		fmt.Printf("Console: %s\n", errors.ErrAtoi)
		return
	}

	profile, err := GetProfile(&wrapper, id)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfile(err)
	if resultOutAccess != nil {
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
	if resultOutAccess != nil {
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

	id, errorConvert:= strconv.Atoi(ctx.UserValue("id").(string))
	if errorConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrAtoi))
		fmt.Printf("Console: %s\n", errors.ErrAtoi)
		return
	}

	err = UpdateName(&wrapper, id, userName.Name)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorProfileUpdateName(err)  // work in progress on CheckError
	if resultOutAccess != nil {
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
