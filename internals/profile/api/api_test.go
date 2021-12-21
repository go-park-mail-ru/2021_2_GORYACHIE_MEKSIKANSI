package api

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/profile/api/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"testing"
)

var GetOrderActiveHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful ProfileHandler handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200,\"body\":{\"user\":null}}"),
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
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.PGetProfileClientClientNotScan, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.PGetProfileClientClientNotScan),
		countGetOrder:                1,
	},
}

func TestProfileHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
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
			GetProfile(tt.inputGetOrderHandlerId).
			Return(tt.outGetProfile, tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.ProfileHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UpdateUserName = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputUpdateName:              "",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200}"),
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
		testName:                  "Error Unmarshall",
		out:                       []byte(errPkg.ErrUnmarshal),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"name\":1}"),
		inputErrorfArgs:           []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 9 of '1'", 1},
		inputErrorfFormat:         "%s, %s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	}, {
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
		testName:                  "Error X-CSRF interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      nil,
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.PUpdateNameNameNotUpdate, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.PUpdateNameNameNotUpdate),
		countGetOrder:                1,
	},
}

func TestUpdateUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
	for _, tt := range UpdateUserName {
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
			UpdateName(tt.inputGetOrderHandlerId, tt.inputUpdateName).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.UpdateUserName(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UpdateUserEmail = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputUpdateName:              "",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200}"),
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
		testName:                  "Error Unmarshall",
		out:                       []byte(errPkg.ErrUnmarshal),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"email\":1}"),
		inputErrorfArgs:           []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 10 of '1'", 1},
		inputErrorfFormat:         "%s, %s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	}, {
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
		testName:                  "Error X-CSRF interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      nil,
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.PUpdateEmailEmailNotUpdate, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.PUpdateEmailEmailNotUpdate),
		countGetOrder:                1,
	},
	{
		testName:                     "Error checkError-ErrCheck-401",
		out:                          []byte("{\"status\":401,\"explain\":\"email already exist\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputWarnfArgs:               []interface{}{errPkg.PUpdateEmailEmailRepeat, 1},
		inputWarnfFormat:             "%s, requestId: %d",
		countErrorf:                  0,
		countWarnf:                   1,
		errGetOrder:                  errors.New(errPkg.PUpdateEmailEmailRepeat),
		countGetOrder:                1,
	},
}

func TestUpdateUserEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
	for _, tt := range UpdateUserEmail {
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
			UpdateEmail(tt.inputGetOrderHandlerId, tt.inputUpdateName).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.UpdateUserEmail(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UpdateUserPassword = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputUpdateName:              "",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200}"),
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
		testName:                  "Error Unmarshall",
		out:                       []byte(errPkg.ErrUnmarshal),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"password\":1}"),
		inputErrorfArgs:           []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 13 of '1'", 1},
		inputErrorfFormat:         "%s, %s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	}, {
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
		testName:                  "Error X-CSRF interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      nil,
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.PUpdatePasswordPasswordNotUpdate, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.PUpdatePasswordPasswordNotUpdate),
		countGetOrder:                1,
	},
}

func TestUpdateUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
	for _, tt := range UpdateUserPassword {
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
			UpdatePassword(tt.inputGetOrderHandlerId, tt.inputUpdateName).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.UpdateUserPassword(&ctxIn)
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UpdateUserPhone = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputUpdateName:              "",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200}"),
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
		testName:                  "Error Unmarshall",
		out:                       []byte(errPkg.ErrUnmarshal),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"Phone\":1}"),
		inputErrorfArgs:           []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 10 of '1'", 1},
		inputErrorfFormat:         "%s, %s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	}, {
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
		testName:                  "Error X-CSRF interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      nil,
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-401",
		out:                          []byte("{\"status\":401,\"explain\":\"phone already exist\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputWarnfArgs:               []interface{}{errPkg.PUpdatePhonePhoneRepeat, 1},
		inputWarnfFormat:             "%s, requestId: %d",
		countErrorf:                  0,
		countWarnf:                   1,
		errGetOrder:                  errors.New(errPkg.PUpdatePhonePhoneRepeat),
		countGetOrder:                1,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.PUpdatePhonePhoneNotUpdate, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.PUpdatePhonePhoneNotUpdate),
		countGetOrder:                1,
	},
}

func TestUpdateUserPhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
	for _, tt := range UpdateUserPhone {
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
			UpdatePhone(tt.inputGetOrderHandlerId, tt.inputUpdateName).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.UpdateUserPhone(&ctxIn)
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UpdateUserAvatar = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputAvatarCtx               interface{}
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
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
		testName:                  "Error X-CSRF interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      nil,
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
}

func TestUpdateUserAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
	for _, tt := range UpdateUserAvatar {
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
			UpdateAvatar(tt.inputGetOrderHandlerId, tt.inputUpdateName).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.UpdateUserPhone(&ctxIn)
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UpdateUserAddress = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              profile.AddressCoordinates
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *profile.Profile
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200}"),
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
		testName:                  "Error Unmarshall",
		out:                       []byte(errPkg.ErrUnmarshal),
		inputValueReqId:           1,
		inputGetOrderCSRFCtx:      "token",
		inputGetOrderHandlerIdCtx: "token",
		inputValueUnmarshal:       []byte("{\"address\"}:1}"),
		inputErrorfArgs:           []interface{}{errPkg.ErrUnmarshal, "parse error: syntax error near offset 10 of '{\"address\"}:1}'", 1},
		inputErrorfFormat:         "%s, %s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	}, {
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
		testName:                  "Error X-CSRF interfaceConvertInt",
		out:                       []byte("expected type string or int"),
		inputValueReqId:           1,
		inputGetOrderHandlerId:    1,
		inputGetOrderCSRFCtx:      nil,
		inputGetOrderHandlerIdCtx: 1,
		inputValueUnmarshal:       []byte("{\"id\":1}"),
		inputErrorfArgs:           []interface{}{"expected type string or int", 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.PUpdateAddressAddressNotUpdate, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.PUpdateAddressAddressNotUpdate),
		countGetOrder:                1,
	},
}

func TestUpdateUserAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockProfileApplicationInterface(ctrlApp)
	for _, tt := range UpdateUserAddress {
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
			UpdateAddress(tt.inputGetOrderHandlerId, tt.inputUpdateName).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoProfile{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.UpdateUserAddress(&ctxIn)
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}
