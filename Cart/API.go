package Cart

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
)

type InfoCart struct {
	Application interfaces.CartApplication
	Logger      errors.MultiLogger
}

func (c *InfoCart) GetCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    c.Logger,
		RequestId: &reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	result, err := c.Application.GetCart(id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetCart(err)
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
		c.Logger.Errorf("GetCartHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (c *InfoCart) UpdateCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	checkError := &errors.CheckError{
		Logger:    c.Logger,
		RequestId: &reqId,
	}

	var cartRequest utils.CartRequest
	err := json.Unmarshal(ctx.Request.Body(), &cartRequest)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		c.Logger.Errorf("UpdateCartHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
		return
	}

	TokenContext := ctx.UserValue("X-Csrf-Token")
	XCsrfToken, errConvert := utils.InterfaceConvertString(TokenContext)
	if (errConvert != nil) && (errConvert.Error() == errors.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		return
	}
	idCtx := ctx.UserValue("id")
	id, errConvert := utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		switch errConvert.Error() {
		case errors.ErrAtoi:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.Logger.Errorf("SignUpHandler: GetId: %s, %v", errors.ErrAtoi, errConvert)
			return

		case errors.ErrNotStringAndInt:
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
			return
		}
	}

	result, err := c.Application.UpdateCart(cartRequest.Cart, id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorUpdateCart(err)
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
			c.Logger.Errorf("UpdateCartHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
			return
		}
		return
	}

	err = json.NewEncoder(ctx).Encode(&utils.Result{
		Status: http.StatusOK,
		Body:   cartRequest,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrEncode))
		c.Logger.Errorf("UpdateCartHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
}
