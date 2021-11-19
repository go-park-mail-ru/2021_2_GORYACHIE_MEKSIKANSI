package Api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyErrors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/internal/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Utils"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
)

type InfoCart struct {
	Application interfaces.CartApplication
	Logger      errPkg.MultiLogger
}

func (c *InfoCart) GetCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    c.Logger,
		RequestId: reqId,
	}

	idCtx := ctx.UserValue("id")
	id, errConvert := Utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
	}

	result, err := c.Application.GetCart(id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorGetCart(err)
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

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
		Body: &Cart.ResponseCart{
			Cart: result,
		},
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
}

func (c *InfoCart) UpdateCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := Utils.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		Logger:    c.Logger,
		RequestId: reqId,
	}

	var cartRequest Cart.CartRequest
	err := json.Unmarshal(ctx.Request.Body(), &cartRequest)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrUnmarshal, err, reqId)
		return
	}

	tokenContext := ctx.UserValue("X-Csrf-Token")
	xCsrfToken, errConvert := Utils.InterfaceConvertString(tokenContext)
	if (errConvert != nil) && (errConvert.Error() == errPkg.ErrNotStringAndInt) {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrNotStringAndInt))
		return
	}
	idCtx := ctx.UserValue("id")
	id, errConvert := Utils.InterfaceConvertInt(idCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		c.Logger.Errorf("%s", errConvert.Error())
	}

	result, err := c.Application.UpdateCart(cartRequest.Cart, id)
	errOut, resultOutAccess, codeHTTP := checkError.CheckErrorUpdateCart(err)
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

	if result != nil {
		err = json.NewEncoder(ctx).Encode(&Authorization.Result{
			Status: http.StatusOK,
			Body: &Cart.ResponseCart{
				Cart: result,
			},
		})
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errPkg.ErrEncode))
			c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
			return
		}
		return
	}

	err = json.NewEncoder(ctx).Encode(&Authorization.Result{
		Status: http.StatusOK,
		Body:   cartRequest,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		c.Logger.Errorf("%s, %v, requestId: %d", errPkg.ErrEncode, err, reqId)
		return
	}
	ctx.Response.Header.Set("X-CSRF-Token", xCsrfToken)
	ctx.Response.SetStatusCode(http.StatusOK)
}
