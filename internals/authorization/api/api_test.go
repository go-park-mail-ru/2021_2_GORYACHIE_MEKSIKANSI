package api

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/api/mocks"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"testing"
)

var SignUpHandler = []struct {
	testName             string
	inputValueReqId      interface{}
	inputValueUnmarshal  []byte
	out                  []byte
	inputErrorfArgs      []interface{}
	inputErrorfFormat    string
	countErrorf          int
	inputWarnfArgs       []interface{}
	inputWarnfFormat     string
	countWarnf           int
	inputSignUpSignUpAll authorization.RegistrationRequest
	outSignUpCookieDB    *util.Defense
	errSignUp            error
	countSignUp          int
}{
	{
		testName:            "Successful SignUp handler",
		inputValueReqId:     10,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		out:                 []byte("{\"status\":201,\"body\":{\"user\":{\"type\":\"\",\"name\":\"\",\"email\":\"\",\"phone\":\"\"}}}"),
		countErrorf:         0,
		countWarnf:          0,
		outSignUpCookieDB:   &util.Defense{SessionId: "sessionid"},
		errSignUp:           nil,
		countSignUp:         1,
	},
	{
		testName:            "Error reqId interfaceConvertInt",
		out:                 []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:     nil,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat:   "%s",
		countErrorf:         1,
		countWarnf:          0,
		countSignUp:         0,
	},
	{
		testName:            "Error unmarshal",
		out:                 []byte(errPkg.ErrUnmarshal),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"name\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 9 of '1'", 1},
		inputErrorfFormat:   "%s, %s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		countSignUp:         0,
	},
	{
		testName:            "Error checkError-ErrCheck-internal",
		out:                 []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ASignupCourierTransactionNotCreate, 1},
		inputErrorfFormat:   "%s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		outSignUpCookieDB:   nil,
		errSignUp:           errors.New(errPkg.ASignupCourierTransactionNotCreate),
		countSignUp:         1,
	},
	{
		testName:            "Error checkError-ErrCheck-409",
		out:                 []byte("{\"status\":409,\"explain\":\"" + errPkg.AGeneralSignUpLoginNotUnique + "\"}"),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		countErrorf:         0,
		inputWarnfArgs:      []interface{}{errPkg.AGeneralSignUpLoginNotUnique, 1},
		inputWarnfFormat:    "%s, requestId: %d",
		countWarnf:          1,
		outSignUpCookieDB:   nil,
		errSignUp:           errors.New(errPkg.AGeneralSignUpLoginNotUnique),
		countSignUp:         1,
	},
	{
		testName:            "Error checkError-ErrCheck-409",
		out:                 []byte("{\"status\":401,\"explain\":\"" + errPkg.AGeneralSignUpIncorrectPhoneFormat + "\"}"),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		countErrorf:         0,
		inputWarnfArgs:      []interface{}{errPkg.AGeneralSignUpIncorrectPhoneFormat, 1},
		inputWarnfFormat:    "%s, requestId: %d",
		countWarnf:          1,
		outSignUpCookieDB:   nil,
		errSignUp:           errors.New(errPkg.AGeneralSignUpIncorrectPhoneFormat),
		countSignUp:         1,
	},
}

func TestSignUpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockAuthorizationApplicationInterface(ctrlApp)
	for _, tt := range SignUpHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
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
			SignUp(&tt.inputSignUpSignUpAll).
			Return(tt.outSignUpCookieDB, tt.errSignUp).
			Times(tt.countSignUp)

		userInfo := UserInfo{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.SignUpHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var LoginHandler = []struct {
	testName            string
	inputValueReqId     interface{}
	inputValueUnmarshal []byte
	out                 []byte
	inputErrorfArgs     []interface{}
	inputErrorfFormat   string
	countErrorf         int
	inputWarnfArgs      []interface{}
	inputWarnfFormat    string
	countWarnf          int
	inputLoginUserLogin authorization.Authorization
	outSignUpCookieDB   *util.Defense
	errSignUp           error
	countSignUp         int
}{
	{
		testName:            "Successful Login handler",
		inputValueReqId:     10,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		out:                 []byte("{\"status\":200}"),
		countErrorf:         0,
		countWarnf:          0,
		outSignUpCookieDB:   &util.Defense{SessionId: "sessionid"},
		errSignUp:           nil,
		countSignUp:         1,
	},
	{
		testName:            "Error reqId interfaceConvertInt",
		out:                 []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:     nil,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat:   "%s",
		countErrorf:         1,
		countWarnf:          0,
		countSignUp:         0,
	},
	{
		testName:            "Error unmarshal",
		out:                 []byte(errPkg.ErrUnmarshal),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"phone\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 10 of '1'", 1},
		inputErrorfFormat:   "%s, %s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		countSignUp:         0,
	},
	{
		testName:            "Error checkError-ErrCheck-500",
		out:                 []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.AAddCookieCookieNotInsert, 1},
		inputErrorfFormat:   "%s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		outSignUpCookieDB:   nil,
		errSignUp:           errors.New(errPkg.AAddCookieCookieNotInsert),
		countSignUp:         1,
	},
	{
		testName:            "Error checkError-ErrCheck-409",
		out:                 []byte("{\"status\":401,\"explain\":\"" + errPkg.ALoginOrPasswordIncorrect + "\"}"),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		countErrorf:         0,
		inputWarnfArgs:      []interface{}{errPkg.ALoginOrPasswordIncorrect, 1},
		inputWarnfFormat:    "%s, requestId: %d",
		countWarnf:          1,
		outSignUpCookieDB:   nil,
		errSignUp:           errors.New(errPkg.ALoginOrPasswordIncorrect),
		countSignUp:         1,
	},
}

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockAuthorizationApplicationInterface(ctrlApp)
	for _, tt := range LoginHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
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
			Login(&tt.inputLoginUserLogin).
			Return(tt.outSignUpCookieDB, tt.errSignUp).
			Times(tt.countSignUp)

		userInfo := UserInfo{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.LoginHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var LogoutHandler = []struct {
	testName            string
	inputValueReqId     interface{}
	inputValueUnmarshal []byte
	out                 []byte
	inputLogoutCSRF     string
	inputLogoutCSRFCtx  interface{}
	inputErrorfArgs     []interface{}
	inputErrorfFormat   string
	countErrorf         int
	inputWarnfArgs      []interface{}
	inputWarnfFormat    string
	countWarnf          int
	outLogoutXCSRF      string
	errLogout           error
	countSignUp         int
}{
	{
		testName:            "Successful Login handler",
		inputValueReqId:     10,
		inputLogoutCSRF:     "token",
		inputLogoutCSRFCtx:  "token",
		inputValueUnmarshal: []byte("{\"id\":1}"),
		out:                 []byte("{\"status\":200}"),
		countErrorf:         0,
		countWarnf:          0,
		outLogoutXCSRF:      "sessionid",
		errLogout:           nil,
		countSignUp:         1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countSignUp:       0,
	},
	{
		testName:           "Error XSCRF interfaceConvertString",
		out:                []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:    1,
		inputLogoutCSRFCtx: nil,
		inputErrorfArgs:    []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat:  "%s, requestId: %d",
		countErrorf:        1,
		countWarnf:         0,
		countSignUp:        0,
	},
	{
		testName:            "Error checkError-ErrCheck-500",
		out:                 []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:     1,
		inputLogoutCSRF:     "token",
		inputLogoutCSRFCtx:  "token",
		inputValueUnmarshal: []byte("{\"id\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ADeleteCookieCookieNotDelete, 1},
		inputErrorfFormat:   "%s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		outLogoutXCSRF:      "nil",
		errLogout:           errors.New(errPkg.ADeleteCookieCookieNotDelete),
		countSignUp:         1,
	},
}

func TestLogoutHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockAuthorizationApplicationInterface(ctrlApp)
	for _, tt := range LogoutHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputLogoutCSRFCtx)
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
			Logout(tt.inputLogoutCSRF).
			Return(tt.outLogoutXCSRF, tt.errLogout).
			Times(tt.countSignUp)

		userInfo := UserInfo{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.LogoutHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var PayHandler = []struct {
	testName            string
	inputValueReqId     interface{}
	inputValueUnmarshal []byte
	out                 []byte
	inputPayCSRF        string
	inputPayCSRFCtx     interface{}
	inputErrorfArgs     []interface{}
	inputErrorfFormat   string
	countErrorf         int
	inputWarnfArgs      []interface{}
	inputWarnfFormat    string
	countWarnf          int
	outPayXCSRF         string
	errLogout           error
	countLogout         int
}{
	{
		testName:            "Successful Payhandler handler",
		inputValueReqId:     10,
		inputPayCSRF:        "token",
		inputPayCSRFCtx:     "token",
		inputValueUnmarshal: []byte("{\"id\":1}"),
		out:                 []byte("{\"status\":200}"),
		countErrorf:         0,
		countWarnf:          0,
		outPayXCSRF:         "sessionid",
		errLogout:           nil,
		countLogout:         1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countLogout:       0,
	},
	{
		testName:          "Error XSCRF interfaceConvertString",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   1,
		inputPayCSRFCtx:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat: "%s, requestId: %d",
		countErrorf:       1,
		countWarnf:        0,
		countLogout:       0,
	},
}

func TestPayHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockAuthorizationApplicationInterface(ctrlApp)
	for _, tt := range PayHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputPayCSRFCtx)
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

		userInfo := UserInfo{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.PayHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var UserWebSocketNewKey = []struct {
	testName            string
	inputValueReqId     interface{}
	inputValueUnmarshal []byte
	out                 []byte
	inputLogoutCSRF     string
	inputLogoutCSRFCtx  interface{}
	inputErrorfArgs     []interface{}
	inputErrorfFormat   string
	countErrorf         int
	inputWarnfArgs      []interface{}
	inputWarnfFormat    string
	countWarnf          int
	outLogoutXCSRF      string
	errLogout           error
	countSignUp         int
}{
	{
		testName:            "Successful Login handler",
		inputValueReqId:     10,
		inputLogoutCSRF:     "token",
		inputLogoutCSRFCtx:  "token",
		inputValueUnmarshal: []byte("{\"id\":1}"),
		out:                 []byte("{\"status\":200}"),
		countErrorf:         0,
		countWarnf:          0,
		outLogoutXCSRF:      "sessionid",
		errLogout:           nil,
		countSignUp:         1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countSignUp:       0,
	},
	{
		testName:           "Error XSCRF interfaceConvertString",
		out:                []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:    1,
		inputLogoutCSRFCtx: nil,
		inputErrorfArgs:    []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat:  "%s, requestId: %d",
		countErrorf:        1,
		countWarnf:         0,
		countSignUp:        0,
	},
	{
		testName:            "Error checkError-ErrCheck-500",
		out:                 []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputValueReqId:     1,
		inputLogoutCSRF:     "token",
		inputLogoutCSRFCtx:  "token",
		inputValueUnmarshal: []byte("{\"id\":1}"),
		inputErrorfArgs:     []interface{}{errPkg.ADeleteCookieCookieNotDelete, 1},
		inputErrorfFormat:   "%s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		outLogoutXCSRF:      "nil",
		errLogout:           errors.New(errPkg.ADeleteCookieCookieNotDelete),
		countSignUp:         1,
	},
}

func TestUserWebSocketNewKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockAuthorizationApplicationInterface(ctrlApp)
	for _, tt := range UserWebSocketNewKey {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputLogoutCSRFCtx)
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
			NewCSRFWebsocket(tt.inputLogoutCSRF).
			Return(tt.outLogoutXCSRF, tt.errLogout).
			Times(tt.countSignUp)

		userInfo := UserInfo{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.UserWebSocketNewKey(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}
