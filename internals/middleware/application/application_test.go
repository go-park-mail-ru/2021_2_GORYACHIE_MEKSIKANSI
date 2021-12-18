package application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/application/mocks"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var CheckAccess = []struct {
	testName         string
	input            *utils.Defense
	out              bool
	outErr           string
	inputCheckAccess *utils.Defense
	outCheckAccess   bool
	errCheckAccess   error
}{
	{
		testName:         "Check access",
		input:            &utils.Defense{},
		out:              true,
		outErr:           "",
		inputCheckAccess: &utils.Defense{},
		outCheckAccess:   true,
		errCheckAccess:   nil,
	},
	{
		testName:         "Error check",
		input:            &utils.Defense{},
		out:              false,
		outErr:           "text",
		inputCheckAccess: &utils.Defense{},
		outCheckAccess:   false,
		errCheckAccess:   errors.New("text"),
	},
}

func TestCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddlewareInterface(ctrl)
	for _, tt := range CheckAccess {
		m.
			EXPECT().
			CheckAccess(tt.inputCheckAccess).
			Return(tt.outCheckAccess, tt.errCheckAccess)
		test := Middleware{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CheckAccess(tt.input)
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

var NewCSRF = []struct {
	testName     string
	input        *utils.Defense
	out          string
	outErr       string
	inputNewCSRF *utils.Defense
	outNewCSRF   string
	errNewCSRF   error
}{
	{
		testName:     "Generate new csrf",
		input:        &utils.Defense{},
		out:          "CSRF-token",
		outErr:       "",
		inputNewCSRF: &utils.Defense{},
		outNewCSRF:   "CSRF-token",
		errNewCSRF:   nil,
	},
	{
		testName:     "Error generate new csrf",
		input:        &utils.Defense{},
		out:          "",
		outErr:       "text",
		inputNewCSRF: &utils.Defense{},
		outNewCSRF:   "",
		errNewCSRF:   errors.New("text"),
	},
}

func TestNewCSRF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddlewareInterface(ctrl)
	for _, tt := range NewCSRF {
		m.
			EXPECT().
			NewCSRF(tt.inputNewCSRF).
			Return(tt.outNewCSRF, tt.errNewCSRF)
		test := Middleware{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.NewCSRF(tt.input)
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
	testName           string
	input              *utils.Defense
	out                int
	outErr             string
	inputGetIdByCookie *utils.Defense
	outGetIdByCookie   int
	errGetIdByCookie   error
}{
	{
		testName:           "Get id",
		input:              &utils.Defense{},
		out:                1,
		outErr:             "",
		inputGetIdByCookie: &utils.Defense{},
		outGetIdByCookie:   1,
		errGetIdByCookie:   nil,
	},
	{
		testName:           "Error get id",
		input:              &utils.Defense{},
		out:                0,
		outErr:             "text",
		inputGetIdByCookie: &utils.Defense{},
		outGetIdByCookie:   0,
		errGetIdByCookie:   errors.New("text"),
	},
}

func TestGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddlewareInterface(ctrl)
	for _, tt := range GetIdByCookie {
		m.
			EXPECT().
			GetIdByCookie(tt.inputGetIdByCookie).
			Return(tt.outGetIdByCookie, tt.errGetIdByCookie)
		test := Middleware{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetIdByCookie(tt.input)
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

var CheckAccessWebsocket = []struct {
	testName                  string
	input                     string
	out                       bool
	outErr                    string
	inputCheckAccessWebsocket string
	outCheckAccessWebsocket   bool
	errCheckAccessWebsocket   error
}{
	{
		testName:                  "Check websocket",
		input:                     "cookie",
		out:                       true,
		outErr:                    "",
		inputCheckAccessWebsocket: "cookie",
		outCheckAccessWebsocket:   true,
		errCheckAccessWebsocket:   nil,
	},
	{
		testName:                  "Error check access",
		input:                     "cookie",
		out:                       false,
		outErr:                    "text",
		inputCheckAccessWebsocket: "cookie",
		outCheckAccessWebsocket:   false,
		errCheckAccessWebsocket:   errors.New("text"),
	},
}

func TestCheckAccessWebsocket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddlewareInterface(ctrl)
	for _, tt := range CheckAccessWebsocket {
		m.
			EXPECT().
			CheckAccessWebsocket(tt.inputCheckAccessWebsocket).
			Return(tt.outCheckAccessWebsocket, tt.errCheckAccessWebsocket)
		test := Middleware{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CheckAccessWebsocket(tt.input)
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
