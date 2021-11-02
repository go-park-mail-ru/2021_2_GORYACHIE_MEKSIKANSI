package Cart

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type InfoCart struct {
	ConnectionDB  *pgxpool.Pool
	LoggerErrWarn *zap.SugaredLogger
	LoggerInfo    *zap.SugaredLogger
	LoggerTest    *zap.SugaredLogger
}

func (c *InfoCart) GetCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.LoggerErrWarn.Errorf("GetCartHandler: GetId: %s, %v", errors.ErrAtoi, errorConvert)
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
		LoggerErrWarn: c.LoggerErrWarn,
		LoggerInfo:    c.LoggerInfo,
		LoggerTest:    c.LoggerTest,
		RequestId:     &reqId,
	}

	wrapper := Wrapper{Conn: c.ConnectionDB}

	idUrl := ctx.UserValue("id")
	var id int
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.LoggerErrWarn.Errorf("GetCartHandler: error: %s, requestId: %d", errors.ErrAtoi, reqId)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		c.LoggerErrWarn.Errorf("GetCartHandler: error: %s, requestId: %d", errors.ErrNotStringAndInt, reqId)
		return
	}

	result, err := GetCart(&wrapper, id)
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
		c.LoggerErrWarn.Errorf("GetCartHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}

}

func (c *InfoCart) UpdateCartHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	var reqId int
	var errorConvert error
	switch reqIdCtx.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(reqIdCtx.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.LoggerErrWarn.Errorf("UpdateCartHandler: GetId: %s, %v", errors.ErrAtoi, errorConvert)
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
		LoggerErrWarn: c.LoggerErrWarn,
		LoggerInfo:    c.LoggerInfo,
		LoggerTest:    c.LoggerTest,
		RequestId:     &reqId,
	}

	wrapper := Wrapper{Conn: c.ConnectionDB}

	var cartRequest utils.CartRequest
	err := json.Unmarshal(ctx.Request.Body(), &cartRequest)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrUnmarshal))
		c.LoggerErrWarn.Errorf("UpdateCartHandler: error: %s, %v, requestId: %d", errors.ErrUnmarshal, err, reqId)
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
		c.LoggerErrWarn.Errorf("UpdateCartHandler: error: %s, requestId: %d", errors.ErrNotStringAndInt, reqId)
		return
	}

	idUrl := ctx.UserValue("id")
	var id int
	switch idUrl.(type) {
	case string:
		id, errorConvert = strconv.Atoi(idUrl.(string))
		if errorConvert != nil {
			ctx.Response.SetStatusCode(http.StatusInternalServerError)
			ctx.Response.SetBody([]byte(errors.ErrAtoi))
			c.LoggerErrWarn.Errorf("UpdateCartHandler: error: %s, requestId: %d", errors.ErrAtoi, reqId)
			return
		}
	case int:
		id = idUrl.(int)
	default:
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errors.ErrNotStringAndInt))
		c.LoggerErrWarn.Errorf("UpdateCartHandler: error: %s, requestId: %d", errors.ErrNotStringAndInt, reqId)
		return
	}

	result, err := UpdateCart(&wrapper, cartRequest.Cart, id)
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
			c.LoggerErrWarn.Errorf("UpdateCartHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
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
		c.LoggerErrWarn.Errorf("UpdateCartHandler: error: %s, %v, requestId: %d", errors.ErrEncode, err, reqId)
		return
	}
}
