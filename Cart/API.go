package Cart

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
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
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		fmt.Printf("Console: %s\n", errors.ErrNotStringAndInt)
		return
	}

	result, err := GetCart(&wrapper, id)
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
	ctx.Response.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
		Body: &utils.ResponseCart{
			Cart: result,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return
	}

}

func (c *InfoCart) UpdateCartHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: c.ConnectionDB}

	var cartRequest utils.CartRequest
	err := json.Unmarshal(ctx.Request.Body(), &cartRequest)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		fmt.Printf("Console: %s\n", errors.ErrUnmarshal)
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
		fmt.Printf("Console: %s\n", errors.ErrNotStringAndInt)
		return
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
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		fmt.Printf("Console: %s\n", errors.ErrNotStringAndInt)
		return
	}

	result, err := UpdateCart(&wrapper, cartRequest.Cart, id)
	errOut, resultOutAccess, codeHTTP := errors.CheckErrorUpdateCart(err)
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
	ctx.Response.Header.Set("X-CSRF-Token", XCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)

	if result != nil {
		err = json.NewEncoder(ctx).Encode(&utils.Result{
			Status: http.StatusOK,
			Body: &utils.ResponseCart{
				Cart: result,
			},
		})
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrEncode))
			fmt.Printf("Console: %s\n", errors.ErrEncode)
			return
		}

		return
	}

	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
		Body: cartRequest,

	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		fmt.Printf("Console: %s\n", errors.ErrEncode)
		return

	}
}
