package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/orm/mocks"
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
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

var CheckAccess = []struct {
	testName            string
	input               *util.Defense
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
		input:               &util.Defense{SessionId: "1", CsrfToken: "1"},
	},
}

func TestCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range CheckAccess {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
				tt.inputQuerySessionId, tt.inputQueryCSRFToken).
			Return(&tt.outQuery)
		test := Wrapper{DBConn: m}
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

var NewCSRF = []struct {
	testName            string
	input               *util.Defense
	outErr              string
	inputQuerySessionId string
	errQuery            error
}{
	{
		testName:            "One",
		outErr:              "",
		inputQuerySessionId: "1",
		errQuery:            nil,
		input:               &util.Defense{SessionId: "1", CsrfToken: "1"},
	},
}

func TestNewCSRF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range NewCSRF {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
				gomock.Any(), tt.inputQuerySessionId).
			Return(nil, tt.errQuery)
		test := Wrapper{DBConn: m}
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

var GetIdByCookie = []struct {
	testName            string
	input               *util.Defense
	out                 int
	outErr              string
	inputQuerySessionId string
	outQuery            Row
}{
	{
		testName:            "One",
		inputQuerySessionId: "1",
		outQuery:            Row{row: []interface{}{1, time.Now()}},
		input:               &util.Defense{SessionId: "1"},
		out:                 0,
		outErr:              errorsConst.MGetIdByCookieCookieExpired,
	},
}

func TestGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetIdByCookie {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
				tt.inputQuerySessionId).
			Return(&tt.outQuery)
		test := Wrapper{DBConn: m}
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
