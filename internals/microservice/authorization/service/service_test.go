package service

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/service/mocks"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

var CheckAccessUser = []struct {
	testName   string
	input      *authProto.Defense
	out        *authProto.CheckAccess
	outErr     string
	inputQuery *authPkg.Defense
	outQuery   bool
	errQuery   error
	countQuery int
}{
	{
		testName: "Check access",
		input: &authProto.Defense{
			SessionId:  "zffdsfsdf",
			XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
		},
		out: &authProto.CheckAccess{
			CheckResult: true,
			Error:       "",
		},
		outErr: "",
		inputQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outQuery:   true,
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error check access",
		input: &authProto.Defense{
			SessionId:  "zffdsfsdf",
			XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
		},
		out: &authProto.CheckAccess{
			CheckResult: false,
			Error:       "text",
		},
		outErr: "",
		inputQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outQuery:   false,
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestCheckAccessUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range CheckAccessUser {
		m.
			EXPECT().
			CheckAccess(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CheckAccessUser(context.Background(), tt.input)
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

var NewCSRFUser = []struct {
	testName   string
	input      *authProto.Defense
	out        *authProto.CSRFResponse
	outErr     string
	inputQuery *authPkg.Defense
	outQuery   string
	errQuery   error
	countQuery int
}{
	{
		testName: "Generate new csrf",
		input: &authProto.Defense{
			SessionId:  "zffdsfsdf",
			XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
		},
		out: &authProto.CSRFResponse{
			XCsrfToken: &authProto.CSRF{
				XCsrfToken: "afdakjf-sadfkjs-sdfsd-sadf",
			},
			Error: "",
		},
		outErr: "",
		inputQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outQuery:   "afdakjf-sadfkjs-sdfsd-sadf",
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error generate csrf",
		input: &authProto.Defense{
			SessionId:  "zffdsfsdf",
			XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
		},
		out: &authProto.CSRFResponse{
			XCsrfToken: nil,
			Error:      "text",
		},
		outErr: "",
		inputQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outQuery:   "afdakjf-sadfkjs-sdfsd-sadf",
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestNewCSRFUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range NewCSRFUser {
		m.
			EXPECT().
			NewCSRF(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.NewCSRFUser(context.Background(), tt.input)
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

var GetIdByCookie = []struct {
	testName   string
	input      *authProto.Defense
	out        *authProto.IdClientResponse
	outErr     string
	inputQuery *authPkg.Defense
	outQuery   int
	errQuery   error
	countQuery int
}{
	{
		testName: "Get id",
		input: &authProto.Defense{
			SessionId:  "zffdsfsdf",
			XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
		},
		out: &authProto.IdClientResponse{
			IdUser: 1,
			Error:  "",
		},
		outErr: "",
		inputQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outQuery:   1,
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get id",
		input: &authProto.Defense{
			SessionId:  "zffdsfsdf",
			XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
		},
		out: &authProto.IdClientResponse{
			IdUser: 0,
			Error:  "text",
		},
		outErr: "",
		inputQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outQuery:   1,
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range GetIdByCookie {
		m.
			EXPECT().
			GetIdByCookie(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetIdByCookie(context.Background(), tt.input)
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

var SignUp = []struct {
	testName   string
	input      *authProto.RegistrationRequest
	out        *authProto.DefenseResponse
	outErr     string
	inputQuery *authPkg.RegistrationRequest
	outQuery   *authPkg.Defense
	errQuery   error
	countQuery int
}{
	{
		testName: "Sign up",
		input: &authProto.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иван Иванов",
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		out: &authProto.DefenseResponse{
			Defense: &authProto.Defense{
				SessionId:  "zffdsfsdf",
				XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
				DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			Error: "",
		},
		outErr: "",
		inputQuery: &authPkg.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иван Иванов",
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		outQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error sign up",
		input: &authProto.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иван Иванов",
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		out: &authProto.DefenseResponse{
			Defense: nil,
			Error:   "text",
		},
		outErr: "",
		inputQuery: &authPkg.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иван Иванов",
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		outQuery:   nil,
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range SignUp {
		m.
			EXPECT().
			SignUp(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.SignUp(context.Background(), tt.input)
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
	input      *authProto.Authorization
	out        *authProto.DefenseResponse
	outErr     string
	inputQuery *authPkg.Authorization
	outQuery   *authPkg.Defense
	errQuery   error
	countQuery int
}{
	{
		testName: "Log in",
		input: &authProto.Authorization{
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		out: &authProto.DefenseResponse{
			Defense: &authProto.Defense{
				SessionId:  "zffdsfsdf",
				XCsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
				DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			Error: "",
		},
		outErr: "",
		inputQuery: &authPkg.Authorization{
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		outQuery: &authPkg.Defense{
			SessionId: "zffdsfsdf",
			CsrfToken: "dfhklkjle-sdad-dsfdsf-fdsf",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error log in",
		input: &authProto.Authorization{
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		out: &authProto.DefenseResponse{
			Defense: nil,
			Error:   "text",
		},
		outErr: "",
		inputQuery: &authPkg.Authorization{
			Email:    "root@root",
			Phone:    "89175554433",
			Password: "password",
		},
		outQuery:   nil,
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range Login {
		m.
			EXPECT().
			Login(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.Login(context.Background(), tt.input)
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
	input      *authProto.CSRF
	out        *authProto.CSRFResponse
	outErr     string
	inputQuery string
	outQuery   string
	errQuery   error
	countQuery int
}{
	{
		testName: "Log out",
		input: &authProto.CSRF{
			XCsrfToken: "asd-asdfa-sadsa-dsa",
		},
		out: &authProto.CSRFResponse{
			XCsrfToken: &authProto.CSRF{
				XCsrfToken: "gfhgf-fghfgh-dhghfh-fgdf",
			},
			Error: "",
		},
		outErr:     "",
		inputQuery: "asd-asdfa-sadsa-dsa",
		outQuery:   "gfhgf-fghfgh-dhghfh-fgdf",
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error log out",
		input: &authProto.CSRF{
			XCsrfToken: "asd-asdfa-sadsa-dsa",
		},
		out: &authProto.CSRFResponse{
			XCsrfToken: nil,
			Error:      "text",
		},
		outErr:     "",
		inputQuery: "asd-asdfa-sadsa-dsa",
		outQuery:   "gfhgf-fghfgh-dhghfh-fgdf",
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range Logout {
		m.
			EXPECT().
			Logout(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.Logout(context.Background(), tt.input)
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
	input      *authProto.IdClient
	out        *authProto.WebsocketResponse
	outErr     string
	inputQuery int
	outQuery   string
	errQuery   error
	countQuery int
}{
	{
		testName: "New CSRF websocket",
		input: &authProto.IdClient{
			ClientId: 1,
		},
		out: &authProto.WebsocketResponse{
			Websocket: "gfhgf-fghfgh-dhghfh-fgdf",
			Error:     "",
		},
		outErr:     "",
		inputQuery: 1,
		outQuery:   "gfhgf-fghfgh-dhghfh-fgdf",
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error new CSRF websocket",
		input: &authProto.IdClient{
			ClientId: 1,
		},
		out: &authProto.WebsocketResponse{
			Websocket: "",
			Error:     "text",
		},
		outErr:     "",
		inputQuery: 1,
		outQuery:   "gfhgf-fghfgh-dhghfh-fgdf",
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestNewCSRFWebsocket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockAuthorizationInterface(ctrl)
	for _, tt := range NewCSRFWebsocket {
		m.
			EXPECT().
			NewCSRFWebsocket(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := AuthorizationManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.NewCSRFWebsocket(context.Background(), tt.input)
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
