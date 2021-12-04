package Application

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/application/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var SignUp = []struct {
	testName                 string
	out                      *authPkg.Defense
	outErr                   string
	input                    *authPkg.RegistrationRequest
	inputSignupClientSignUp  *authPkg.RegistrationRequest
	resultSignupClient       *authPkg.Defense
	errSignupClient          error
	countSignupClient        int
	inputSignupCourierSignUp *authPkg.RegistrationRequest
	resultSignupCourier      *authPkg.Defense
	errSignupCourier         error
	countSignupCourier       int
	inputSignupHostSignUp    *authPkg.RegistrationRequest
	resultSignupHost         *authPkg.Defense
	errSignupHost            error
	countSignupHost          int
	resultGenerateNew        *authPkg.Defense
	inputSignupClientCookie  *authPkg.Defense
	inputSignupCourierCookie *authPkg.Defense
	inputSignupHostCookie    *authPkg.Defense
}{
	{
		input:                   &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName:                "One",
		outErr:                  "",
		resultSignupClient:      &authPkg.Defense{},
		inputSignupClientSignUp: &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out:                     &authPkg.Defense{},
		errSignupClient:         nil,
		countSignupClient:       1,
		resultGenerateNew:       &authPkg.Defense{},
		inputSignupClientCookie: &authPkg.Defense{},
	},
	{
		input:                    &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "courier"},
		testName:                 "Two",
		outErr:                   "",
		resultSignupCourier:      &authPkg.Defense{},
		inputSignupCourierSignUp: &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "courier"},
		out:                      &authPkg.Defense{},
		errSignupCourier:         nil,
		countSignupCourier:       1,
		resultGenerateNew:        &authPkg.Defense{},
		inputSignupCourierCookie: &authPkg.Defense{},
	},
	{
		input:                 &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "host"},
		testName:              "Three",
		outErr:                "",
		resultSignupHost:      &authPkg.Defense{},
		inputSignupHostSignUp: &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "host"},
		out:                   &authPkg.Defense{},
		errSignupHost:         nil,
		countSignupHost:       1,
		resultGenerateNew:     &authPkg.Defense{},
		inputSignupHostCookie: &authPkg.Defense{},
	},
	{
		input:                   &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName:                "Four",
		outErr:                  "text",
		resultSignupClient:      &authPkg.Defense{},
		inputSignupClientSignUp: &authPkg.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out:                     nil,
		errSignupClient:         errors.New("text"),
		countSignupClient:       1,
		resultGenerateNew:       &authPkg.Defense{},
		inputSignupClientCookie: &authPkg.Defense{},
	},
}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range SignUp {
		m.
			EXPECT().
			SignupClient(tt.inputSignupClientSignUp, tt.inputSignupClientCookie).
			Return(tt.resultSignupClient, tt.errSignupClient).
			Times(tt.countSignupClient)
		m.
			EXPECT().
			SignupCourier(tt.inputSignupCourierSignUp, tt.inputSignupCourierCookie).
			Return(tt.resultSignupCourier, tt.errSignupCourier).
			Times(tt.countSignupCourier)
		m.
			EXPECT().
			SignupHost(tt.inputSignupHostSignUp, tt.inputSignupHostCookie).
			Return(tt.resultSignupHost, tt.errSignupHost).
			Times(tt.countSignupHost)
		m.
			EXPECT().
			NewDefense().
			Return(tt.resultGenerateNew)
		test := AuthorizationApplication{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.SignUp(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var Login = []struct {
	testName             string
	out                  *authPkg.Defense
	outErr               string
	input                *authPkg.Authorization
	inputLoginEmail      string
	inputLoginPassword   string
	resultLogin          int
	errLogin             error
	countLoginEmail      int
	countLoginPhone      int
	inputLoginPhone      string
	resultGenerateNew    *authPkg.Defense
	countGenerateNew     int
	inputAddCookieCookie *authPkg.Defense
	inputAddCookieId     int
	errAddCookie         error
	countAddCookie       int
}{
	{
		input:                &authPkg.Authorization{Email: "1", Phone: "", Password: "1"},
		testName:             "One",
		outErr:               "",
		out:                  &authPkg.Defense{},
		inputLoginEmail:      "1",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      1,
		countLoginPhone:      0,
		resultGenerateNew:    &authPkg.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &authPkg.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         nil,
		countAddCookie:       1,
	},
	{
		input:                &authPkg.Authorization{Email: "", Phone: "1", Password: "1"},
		testName:             "Two",
		outErr:               "",
		out:                  &authPkg.Defense{},
		inputLoginEmail:      "",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      0,
		countLoginPhone:      1,
		resultGenerateNew:    &authPkg.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &authPkg.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         nil,
		countAddCookie:       1,
	},
	{
		input:                &authPkg.Authorization{Email: "", Phone: "1", Password: "1"},
		testName:             "Three",
		outErr:               "text",
		out:                  nil,
		inputLoginEmail:      "",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      0,
		countLoginPhone:      1,
		resultGenerateNew:    &authPkg.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &authPkg.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         errors.New("text"),
		countAddCookie:       1,
	},
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range Login {
		m.
			EXPECT().
			LoginByEmail(tt.inputLoginEmail, tt.inputLoginPassword).
			Return(tt.resultLogin, tt.errLogin).
			Times(tt.countLoginEmail)
		m.
			EXPECT().
			LoginByPhone(tt.inputLoginPhone, tt.inputLoginPassword).
			Return(tt.resultLogin, tt.errLogin).
			Times(tt.countLoginPhone)
		m.
			EXPECT().
			NewDefense().
			Return(tt.resultGenerateNew).
			Times(tt.countGenerateNew)
		m.
			EXPECT().
			AddCookie(tt.inputAddCookieCookie, tt.inputAddCookieId).
			Return(tt.errAddCookie).
			Times(tt.countAddCookie)
		test := AuthorizationApplication{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.Login(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var Logout = []struct {
	testName     string
	outErr       string
	out          string
	input        string
	inputDelete  string
	resultDelete string
	errDelete    error
}{
	{
		testName:     "One",
		out:          "1",
		outErr:       "",
		input:        "1",
		inputDelete:  "1",
		resultDelete: "1",
		errDelete:    nil,
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range Logout {
		m.
			EXPECT().
			DeleteCookie(tt.inputDelete).
			Return(tt.resultDelete, tt.errDelete)
		test := AuthorizationApplication{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.Logout(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
