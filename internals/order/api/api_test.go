package api

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/order/api/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"testing"
)

var CreateOrderHandler = []struct {
	testName                      string
	inputValueReqId               interface{}
	inputValueUnmarshal           []byte
	inputCreateOrderHandlerIdCtx  interface{}
	inputCreateOrderHandlerId     int
	inputCreateOrderHandlerDishes order.CreateOrder
	out                           []byte
	inputCreateOrderCSRFCtx       interface{}
	inputErrorfArgs               []interface{}
	inputErrorfFormat             string
	countErrorf                   int
	inputWarnfArgs                []interface{}
	inputWarnfFormat              string
	countWarnf                    int
	outCreateOrder                int
	errCreateOrder                error
	countCreateOrder              int
}{
	{
		testName:                     "Successful CreateOrderHandler handler 1",
		inputValueReqId:              10,
		inputCreateOrderHandlerIdCtx: 1,
		inputCreateOrderHandlerId:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputCreateOrderCSRFCtx:      "token",
		out:                          []byte("{\"status\":200,\"body\":{\"order\":{\"id\":0}}}"),
		countErrorf:                  0,
		countWarnf:                   0,
		errCreateOrder:               nil,
		countCreateOrder:             1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countCreateOrder:  0,
	},
	{
		testName:                     "Error Unmarshall interfaceConvertInt",
		out:                          []byte(errPkg.ErrUnmarshal),
		inputValueReqId:              1,
		inputCreateOrderCSRFCtx:      nil,
		inputCreateOrderHandlerIdCtx: nil,
		inputErrorfArgs:              []interface{}{errPkg.ErrUnmarshal, "unexpected end of JSON input", 1},
		inputErrorfFormat:            "%s, %s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		countCreateOrder:             0,
	},
	{
		testName:                     "Error X-Csrf-Token interfaceConvertInt",
		out:                          []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:              1,
		inputCreateOrderCSRFCtx:      nil,
		inputCreateOrderHandlerIdCtx: nil,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		countCreateOrder:             0,
	},
	{
		testName:                     "Error id interfaceConvertInt",
		out:                          []byte("func Atoi convert string in int"),
		inputValueReqId:              1,
		inputCreateOrderCSRFCtx:      "token",
		inputCreateOrderHandlerIdCtx: "token",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{"func Atoi convert string in int", 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		countCreateOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-404",
		out:                          []byte("{\"status\":404,\"explain\":\"Нечего создавать: корзина пустая\"}"),
		inputCreateOrderCSRFCtx:      "token",
		inputValueReqId:              1,
		inputCreateOrderHandlerIdCtx: 1,
		inputCreateOrderHandlerId:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.CGetCartCartNotFound, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errCreateOrder:               errors.New(errPkg.CGetCartCartNotFound),
		countCreateOrder:             1,
	},
	{
		testName:                     "Error checkError-ErrCheck-500-default",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:              1,
		inputCreateOrderHandlerIdCtx: 1,
		inputCreateOrderHandlerId:    1,
		inputCreateOrderCSRFCtx:      "token",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{"tempError", 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errCreateOrder:               errors.New("tempError"),
		countCreateOrder:             1,
	},
}

func TestCreateOrderHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockOrderApplicationInterface(ctrlApp)
	for _, tt := range CreateOrderHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputCreateOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputCreateOrderHandlerIdCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			CreateOrder(tt.inputCreateOrderHandlerId, tt.inputCreateOrderHandlerDishes).
			Return(tt.outCreateOrder, tt.errCreateOrder).
			Times(tt.countCreateOrder)

		userInfo := InfoOrder{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.CreateOrderHandler(&ctxIn)
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var GetOrdersHandler = []struct {
	testName                   string
	inputValueReqId            interface{}
	inputValueUnmarshal        []byte
	inputGetOrderHandlerIdCtx  interface{}
	inputGetOrderHandlerId     int
	inputGetOrderHandlerDishes *order.HistoryOrderArray
	out                        []byte
	inputGetOrderCSRFCtx       interface{}
	inputErrorfArgs            []interface{}
	inputErrorfFormat          string
	countErrorf                int
	inputWarnfArgs             []interface{}
	inputWarnfFormat           string
	countWarnf                 int
	outGetOrder                *order.HistoryOrderArray
	errGetOrder                error
	countGetOrder              int
}{
	{
		testName:                  "Successful CreateOrderHandler handler",
		inputValueReqId:           10,
		inputGetOrderHandlerIdCtx: 1,
		inputGetOrderHandlerId:    1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:      "token",
		out:                       []byte{0x7b, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x32, 0x30, 0x30, 0x2c, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x6e, 0x75, 0x6c, 0x6c, 0x7d, 0xa},
		countErrorf:               0,
		countWarnf:                0,
		errGetOrder:               nil,
		countGetOrder:             1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:                  "Error id interfaceConvertInt",
		out:                       []byte("func Atoi convert string in int"),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"func Atoi convert string in int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                  "Error checkError-ErrCheck-404",
		out:                       []byte("{\"status\":404,\"explain\":\"Заказов не  найдено\"}"),
		inputGetOrderCSRFCtx:      "token",
		inputValueReqId:           1,
		inputGetOrderHandlerIdCtx: 1,
		inputGetOrderHandlerId:    1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{errPkg.OGetOrdersOrdersIsVoid, 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		errGetOrder:               errors.New(errPkg.OGetOrdersOrdersIsVoid),
		countGetOrder:             1,
	},
	{
		testName:                  "Error checkError-ErrCheck-500-default",
		out:                       []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:           1,
		inputGetOrderHandlerIdCtx: 1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      "token",
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"tempError", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		errGetOrder:               errors.New("tempError"),
		countGetOrder:             1,
	},
}

func TestGetOrdersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockOrderApplicationInterface(ctrlApp)
	for _, tt := range GetOrdersHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			GetOrders(tt.inputGetOrderHandlerId).
			Return(tt.outGetOrder, tt.errGetOrder).
			Times(tt.countGetOrder)

		userInfo := InfoOrder{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.GetOrdersHandler(&ctxIn)
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var GetOrderActiveHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputGetOrderHandlerDishes   *order.HistoryOrderArray
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetOrder                  *order.ActiveOrder
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful GetOrderActiveHandler handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte{0x7b, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x32, 0x30, 0x30, 0x2c, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x6e, 0x75, 0x6c, 0x6c, 0x7d, 0xa},
		countErrorf:                  0,
		countWarnf:                   0,
		errGetOrder:                  nil,
		countGetOrder:                1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:                  "Error id interfaceConvertInt",
		out:                       []byte("func Atoi convert string in int"),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"func Atoi convert string in int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                  "Error idOrd interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-404",
		out:                          []byte("{\"status\":404,\"explain\":\"Заказ не существует\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.OGetOrderNotExist, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.OGetOrderNotExist),
		countGetOrder:                1,
	},
	{
		testName:                     "Error checkError-ErrCheck-500-default",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputGetOrderCSRFCtx:         "token",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{"tempError", 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New("tempError"),
		countGetOrder:                1,
	},
}

func TestGetOrderActiveHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockOrderApplicationInterface(ctrlApp)
	for _, tt := range GetOrderActiveHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.SetUserValue("idOrd", tt.inputGetOrderHandlerIdOrdCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			GetActiveOrder(tt.inputGetOrderHandlerId, tt.inputGetOrderHandlerIdOrd).
			Return(tt.outGetOrder, tt.errGetOrder).
			Times(tt.countGetOrder)

		userInfo := InfoOrder{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.GetOrderActiveHandler(&ctxIn)
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}
