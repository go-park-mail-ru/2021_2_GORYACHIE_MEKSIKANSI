package application

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/application/mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var SignUp = []struct {
	testName   string
	input      *auth.RegistrationRequest
	out        *util.Defense
	outErr     string
	inputQuery *auth.RegistrationRequest
	outQuery   *util.Defense
	errQuery   error
}{
	{
		testName: "Client",
		input: &auth.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		out:    &util.Defense{SessionId: "fdsgdsgf", CsrfToken: "asdasd"},
		outErr: "",
		inputQuery: &auth.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		outQuery: &util.Defense{SessionId: "fdsgdsgf", CsrfToken: "asdasd"},
		errQuery: nil,
	},
}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorizationInterface(ctrl)
	for _, tt := range SignUp {
		m.
			EXPECT().
			SignUp(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Authorization{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignUp(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var Login = []struct {
	testName   string
	input      *auth.Authorization
	out        *util.Defense
	outErr     string
	inputQuery *auth.Authorization
	outQuery   *util.Defense
	errQuery   error
}{
	{
		testName: "Authorization on phone",
		input: &auth.Authorization{
			Email:    "",
			Phone:    "89175554433",
			Password: "root",
		},
		out:    &util.Defense{SessionId: "fdsgdsgf", CsrfToken: "asdasd"},
		outErr: "",
		inputQuery: &auth.Authorization{
			Email:    "",
			Phone:    "89175554433",
			Password: "root",
		},
		outQuery: &util.Defense{SessionId: "fdsgdsgf", CsrfToken: "asdasd"},
		errQuery: nil,
	},
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorizationInterface(ctrl)
	for _, tt := range Login {
		m.
			EXPECT().
			Login(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Authorization{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.Login(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var Logout = []struct {
	testName   string
	input      string
	out        string
	outErr     string
	inputQuery string
	outQuery   string
	errQuery   error
}{
	{
		testName:   "Logout",
		input:      "asdasdadsf",
		out:        "adfgdfgjyihlkgj",
		outErr:     "",
		inputQuery: "asdasdadsf",
		outQuery:   "adfgdfgjyihlkgj",
		errQuery:   nil,
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorizationInterface(ctrl)
	for _, tt := range Logout {
		m.
			EXPECT().
			Logout(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Authorization{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.Logout(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var NewCSRFWebsocket = []struct {
	testName   string
	input      int
	out        string
	outErr     string
	inputQuery int
	outQuery   string
	errQuery   error
}{
	{
		testName:   "New token",
		input:      1,
		out:        "dgiudfgkdfhug",
		outErr:     "",
		inputQuery: 1,
		outQuery:   "dgiudfgkdfhug",
		errQuery:   nil,
	},
}

func TestNewCSRFWebsocket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorizationInterface(ctrl)
	for _, tt := range NewCSRFWebsocket {
		m.
			EXPECT().
			NewCSRFWebsocket(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Authorization{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.NewCSRFWebsocket(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
