package api

//import (
//	"2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/api/mocks"
//	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
//	"errors"
//	"fmt"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/require"
//	"github.com/valyala/fasthttp"
//	"testing"
//)
//
//var LogURL = []struct {
//	testName                     string
//	inputValueReqId              interface{}
//	inputValueUnmarshal          []byte
//	inputUpdateCartHandlerIdCtx  interface{}
//	inputUpdateCartHandlerId     int
//	inputUpdateCartHandlerDishes cart.RequestCartDefault
//	out                          []byte
//	inputUpdateCartCSRFCtx       interface{}
//	inputErrorfArgs              []interface{}
//	inputErrorfFormat            string
//	countErrorf                  int
//	inputWarnfArgs               []interface{}
//	inputWarnfFormat             string
//	countWarnf                   int
//	outUpdateCart                *cart.ResponseCartErrors
//	errGetCart                   error
//	countGetCart                 int
//}{
//	{
//		testName:                    "Successful LogURL handler 1",
//		inputValueReqId:             10,
//		inputUpdateCartHandlerIdCtx: 1,
//		inputUpdateCartHandlerId:    1,
//		inputValueUnmarshal:         []byte("{\"id\":1}"),
//		inputUpdateCartCSRFCtx:      "token",
//		out:                         []byte("{\"status\":200,\"body\":{\"cart\":{\"restaurant\":{\"id\":0},\"dishes\":null,\"promo_code\":\"\"}}}"),
//		countErrorf:                 0,
//		countWarnf:                  0,
//		errGetCart:                  nil,
//		countGetCart:                1,
//	},
//	{
//		testName:                    "Successful LogURL handler 2",
//		inputValueReqId:             10,
//		inputUpdateCartHandlerIdCtx: 1,
//		inputUpdateCartHandlerId:    1,
//		outUpdateCart:               nil,
//		inputValueUnmarshal:         []byte("{\"id\":1}"),
//		inputUpdateCartCSRFCtx:      "token",
//		out:                         []byte("{\"status\":200,\"body\":{\"cart\":{\"restaurant\":{\"id\":0},\"dishes\":null,\"promo_code\":\"\"}}}"),
//		countErrorf:                 0,
//		countWarnf:                  0,
//		errGetCart:                  nil,
//		countGetCart:                1,
//	},
//	{
//		testName:          "Error reqId interfaceConvertInt",
//		out:               []byte(errPkg.ErrNotStringAndInt),
//		inputValueReqId:   nil,
//		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
//		inputErrorfFormat: "%s",
//		countErrorf:       1,
//		countWarnf:        0,
//		countGetCart:      0,
//	},
//	{
//		testName:                    "Error Unmarshall interfaceConvertInt",
//		out:                         []byte(errPkg.ErrUnmarshal),
//		inputValueReqId:             1,
//		inputUpdateCartCSRFCtx:      nil,
//		inputUpdateCartHandlerIdCtx: nil,
//		inputErrorfArgs:             []interface{}{errPkg.ErrUnmarshal, errors.New("EOF"), 1},
//		inputErrorfFormat:           "%s, %v, requestId: %d",
//		countErrorf:                 1,
//		countWarnf:                  0,
//		countGetCart:                0,
//	},
//	{
//		testName:                    "Error X-Csrf-Token interfaceConvertInt",
//		out:                         []byte(errPkg.ErrNotStringAndInt),
//		inputValueReqId:             1,
//		inputUpdateCartCSRFCtx:      nil,
//		inputUpdateCartHandlerIdCtx: nil,
//		inputValueUnmarshal:         []byte("{\"id\":1}"),
//		inputErrorfArgs:             []interface{}{errPkg.ErrNotStringAndInt, 1},
//		inputErrorfFormat:           "%s, requestId: %d",
//		countErrorf:                 1,
//		countWarnf:                  0,
//		countGetCart:                0,
//	},
//	{
//		testName:                    "Error id interfaceConvertInt",
//		out:                         []byte("func Atoi convert string in int"),
//		inputValueReqId:             1,
//		inputUpdateCartCSRFCtx:      "token",
//		inputUpdateCartHandlerIdCtx: "token",
//		inputValueUnmarshal:         []byte("{\"id\":1}"),
//		inputErrorfArgs:             []interface{}{"func Atoi convert string in int", 1},
//		inputErrorfFormat:           "%s, requestId: %d",
//		countErrorf:                 1,
//		countWarnf:                  0,
//		countGetCart:                0,
//	},
//	{
//		testName:                    "Error checkError-ErrCheck-500",
//		out:                         []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
//		inputUpdateCartCSRFCtx:      "token",
//		inputValueReqId:             1,
//		inputUpdateCartHandlerIdCtx: 1,
//		inputUpdateCartHandlerId:    1,
//		inputValueUnmarshal:         []byte("{\"id\":1}"),
//		inputErrorfArgs:             []interface{}{errPkg.CUpdateCartCartNotInsert, 1},
//		inputErrorfFormat:           "%s, requestId: %d",
//		countErrorf:                 1,
//		countWarnf:                  0,
//		errGetCart:                  errors.New(errPkg.CUpdateCartCartNotInsert),
//		countGetCart:                1,
//	},
//	{
//		testName:                    "Error checkError-ErrCheck-500-default",
//		out:                         []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
//		inputValueReqId:             1,
//		inputUpdateCartHandlerIdCtx: 1,
//		inputUpdateCartHandlerId:    1,
//		inputUpdateCartCSRFCtx:      "token",
//		inputValueUnmarshal:         []byte("{\"id\":1}"),
//		inputErrorfArgs:             []interface{}{"tempError", 1},
//		inputErrorfFormat:           "%s, requestId: %d",
//		countErrorf:                 1,
//		countWarnf:                  0,
//		errGetCart:                  errors.New("tempError"),
//		countGetCart:                1,
//	},
//}
//
//func TestUpdateCartHandler(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	ctrlApp := gomock.NewController(t)
//	ctrlCounter := gomock.NewController(t)
//	ctrlCounterVec := gomock.NewController(t)
//	defer ctrl.Finish()
//	defer ctrlApp.Finish()
//	defer ctrlCounter.Finish()
//	defer ctrlCounterVec.Finish()
//
//	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
//	mockApplication := mocks.NewMockMiddlewareApplicationInterface(ctrlApp)
//	mockCounterMetric := mocks.NewMockCounterMetricInterface(ctrlCounter)
//	mockCounterVecMetric := mocks.NewMockCounterVecMetricInterface(ctrlCounterVec)
//	for _, tt := range LogURL {
//		ctxIn := fasthttp.RequestCtx{}
//		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
//		ctxIn.SetUserValue("X-Csrf-Token", tt.inputUpdateCartCSRFCtx)
//		ctxIn.SetUserValue("id", tt.inputUpdateCartHandlerIdCtx)
//		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
//		ctxExpected := fasthttp.RequestCtx{}
//		ctxExpected.Response.SetBody(tt.out)
//		mockMultilogger.
//			EXPECT().
//			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
//			Times(tt.countErrorf)
//
//		mockMultilogger.
//			EXPECT().
//			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
//			Times(tt.countWarnf)
//
//		//mockApplication.
//		//	EXPECT().
//		//	GetIdByCookie(tt.inputUpdateCartHandlerDishes, tt.inputUpdateCartHandlerId).
//		//	Return(tt.outUpdateCart, tt.errGetCart).
//		//	Times(tt.countGetCart)
//		userInfo := InfoMiddleware{Logger: mockMultilogger, Application: mockApplication}
//		t.Run(tt.testName, func(t *testing.T) {
//			userInfo.LogURL(fasthttp.RequestHandler(&ctxIn))
//			println(string(ctxIn.Response.Body()))
//			//println(string(ctxExpected.Response.Body()))
//			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))
//
//		})
//	}
//
//}
