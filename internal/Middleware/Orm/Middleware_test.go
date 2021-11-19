package Orm

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/internal/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Utils"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/test/Mocks"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Row struct {
	row    []interface{}
	errRow error
}

func (r *Row) Scan(dest ...interface{}) error {
	if r.errRow != nil {
		return r.errRow
	}
	for i := range dest {
		switch dest[i].(type) {
		case *int:
			*dest[i].(*int) = r.row[i].(int)
		case *string:
			*dest[i].(*string) = r.row[i].(string)
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
		case *time.Time:
			*dest[i].(*time.Time) = r.row[i].(time.Time)
		}
	}
	return nil
}

var OrmCheckAccess = []struct {
	testName            string
	input               *Utils.Defense
	out                 bool
	outErr              string
	inputQuerySessionId string
	inputQueryCSRFToken string
	outQuery            Row
}{
	{
		testName:            "One",
		outErr:              "",
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		outQuery:            Row{row: []interface{}{1, time.Now()}},
		input:               &Utils.Defense{SessionId: "1", CsrfToken: "1"},
	},
}

func TestOrmCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmCheckAccess {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
				tt.inputQuerySessionId, tt.inputQueryCSRFToken).
			Return(&tt.outQuery)
		test := Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CheckAccess(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmNewCSRF = []struct {
	testName            string
	input               *Utils.Defense
	outErr              string
	inputQuerySessionId string
	errQuery            error
}{
	{
		testName:            "One",
		outErr:              "",
		inputQuerySessionId: "1",
		errQuery:            nil,
		input:               &Utils.Defense{SessionId: "1", CsrfToken: "1"},
	},
}

func TestOrmNewCSRF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmNewCSRF {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
				gomock.Any(), tt.inputQuerySessionId).
			Return(nil, tt.errQuery)
		test := Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.NewCSRF(tt.input)
			require.NotEqual(t, "", result, fmt.Sprintf("Expected: %v\nbut got: %v", "", result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetIdByCookie = []struct {
	testName            string
	input               *Utils.Defense
	out                 int
	outErr              string
	inputQuerySessionId string
	outQuery            Row
}{
	{
		testName:            "One",
		inputQuerySessionId: "1",
		outQuery:            Row{row: []interface{}{1, time.Now()}},
		input:               &Utils.Defense{SessionId: "1"},
		out:                 0,
		outErr:              errorsConst.MGetIdByCookieCookieExpired,
	},
}

func TestOrmGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetIdByCookie {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
				tt.inputQuerySessionId).
			Return(&tt.outQuery)
		test := Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetIdByCookie(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationCheckAccess = []struct {
	testName         string
	input            *Utils.Defense
	inputCheckAccess *Utils.Defense
	out              bool
	outCheckAccess   bool
	outErr           string
	errCheckAccess   error
}{
	{
		testName:         "One",
		outErr:           "",
		out:              false,
		outCheckAccess:   false,
		input:            &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputCheckAccess: &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		errCheckAccess:   nil,
	},
}

func TestApplicationCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddleware(ctrl)
	for _, tt := range ApplicationCheckAccess {
		m.
			EXPECT().
			CheckAccess(tt.inputCheckAccess).
			Return(tt.outCheckAccess, tt.errCheckAccess)
		test := Application.Middleware{m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CheckAccess(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationNewCSRF = []struct {
	testName     string
	input        *Utils.Defense
	inputNewCSRF *Utils.Defense
	out          string
	outErr       string
	errNewCSRF   error
	outNewCSRF   string
}{
	{
		testName:     "One",
		outErr:       "",
		errNewCSRF:   nil,
		input:        &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputNewCSRF: &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		out:          "1",
		outNewCSRF:   "1",
	},
}

func TestApplicationNewCSRF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddleware(ctrl)
	for _, tt := range ApplicationNewCSRF {
		m.
			EXPECT().
			NewCSRF(tt.inputNewCSRF).
			Return(tt.outNewCSRF, tt.errNewCSRF)
		test := Application.Middleware{m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.NewCSRF(tt.input)
			require.NotEqual(t, "", result, fmt.Sprintf("Expected: %v\nbut got: %v", "", result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationGetIdByCookie = []struct {
	testName    string
	input       *Utils.Defense
	inputGetId  *Utils.Defense
	out         int
	outErr      string
	resultGetId int
	errGetId    error
}{
	{
		testName:    "One",
		input:       &Utils.Defense{SessionId: "1"},
		inputGetId:  &Utils.Defense{SessionId: "1"},
		out:         1,
		outErr:      "",
		resultGetId: 1,
		errGetId:    nil,
	},
}

func TestApplicationGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperMiddleware(ctrl)
	for _, tt := range ApplicationGetIdByCookie {
		m.
			EXPECT().
			GetIdByCookie(tt.inputGetId).
			Return(tt.resultGetId, tt.errGetId)
		test := Application.Middleware{m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetIdByCookie(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
