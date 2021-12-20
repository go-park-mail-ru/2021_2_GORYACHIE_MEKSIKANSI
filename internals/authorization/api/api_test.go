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
	inputErrorfAlias     string
	inputErrorfText      string
	inputErrorfArgs      []interface{}
	inputErrorfFormat    string
	countErrorf          int
	inputSignUpSignUpAll authorization.RegistrationRequest
	outSignUpCookieDB    *util.Defense
	errSignUp            error
	countSignUp          int
}{
	{
		testName:            "Successful handler",
		inputValueReqId:     10,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		out:                 []byte("{\"status\":201,\"body\":{\"user\":{\"type\":\"\",\"name\":\"\",\"email\":\"\",\"phone\":\"\"}}}"),
		countErrorf:         0,
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
		countSignUp:         0,
	},
	{
		testName:            "Error checkError-ErrMarshal",
		out:                 []byte(errPkg.ErrMarshal),
		inputValueReqId:     1,
		inputValueUnmarshal: []byte("{\"id\":1}"),
		countErrorf:         0,
		outSignUpCookieDB:   nil,
		errSignUp:           errors.New(errPkg.ASignupCourierTransactionNotCreate),
		countSignUp:         1,
	},
}

func TestSignUpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockInterface := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockAuthorizationApplicationInterface(ctrlApp)
	for _, tt := range SignUpHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockInterface.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockApplication.
			EXPECT().
			SignUp(&tt.inputSignUpSignUpAll).
			Return(tt.outSignUpCookieDB, tt.errSignUp).
			Times(tt.countSignUp)

		userInfo := UserInfo{Logger: mockInterface, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			userInfo.SignUpHandler(&ctxIn)
			//println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}
