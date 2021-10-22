package Cart

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	//mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

type InfoCart struct {
	ConnectionDB *pgxpool.Pool
}

func (c *InfoCart) GetCartHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: c.ConnectionDB}

	cookieDB := utils.Defense{SessionId: string(ctx.Request.Header.Cookie("session_id"))}
	cookieDB.CsrfToken = string(ctx.Request.Header.Peek("X-Csrf-Token"))

	//_, err := mid.CheckAccess(c.ConnectionDB, &cookieDB)
	//errAccess, resultOutAccess, codeHTTP := errors.CheckErrorAccess(err)
	//if errAccess != nil {
	//	switch errAccess.Error() {
	//	case errors.ErrMarshal:
	//		ctx.Response.SetStatusCode(codeHTTP)
	//		ctx.Response.SetBody([]byte(errors.ErrMarshal))
	//		return
	//	case errors.ErrCheck:
	//		ctx.Response.SetStatusCode(codeHTTP)
	//		ctx.Response.SetBody(resultOutAccess)
	//		return
	//	}
	//}

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

	cart, err := GetCart(&wrapper, id)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorGetCart(err)
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

	ctx.Response.Header.Set("X-CSRF-Token", cookieDB.CsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
		Body: cart,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}
