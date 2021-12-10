package orm

import (
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/orm/mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
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
		case **string:
			t := r.row[i].(string)
			*dest[i].(**string) = &t
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
		case **int32:
			t := int32(r.row[i].(int))
			*dest[i].(**int32) = &t
		default:
			dest[i] = nil
		}
	}
	return nil
}

var CheckAccess = []struct {
	testName   string
	input      *util.Defense
	out        bool
	outErr     string
	inputQuery *authProto.Defense
	outQuery   *authProto.CheckAccess
	errQuery   error
}{
	{
		testName: "One",
		input: &util.Defense{
			SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		out:    true,
		outErr: "",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CheckAccess{CheckResult: true},
	},
}

func TestCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionMiddlewareInterface(ctrl)
	for _, tt := range CheckAccess {
		m.
			EXPECT().
			CheckAccessUser(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
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

var NewCSRF = []struct {
	testName   string
	input      *util.Defense
	outErr     string
	inputQuery *authProto.Defense
	outQuery   *authProto.CSRFResponse
	errQuery   error
}{
	{
		testName: "First",
		input: &util.Defense{SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		outErr: "",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CSRFResponse{XCsrfToken: &authProto.CSRF{XCsrfToken: "1"}},
		errQuery: nil,
	},
}

func TestNewCSRF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionMiddlewareInterface(ctrl)
	for _, tt := range NewCSRF {
		m.
			EXPECT().
			NewCSRFUser(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
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

var GetIdByCookie = []struct {
	testName   string
	input      *util.Defense
	out        int
	outErr     string
	inputQuery *authProto.Defense
	outQuery   *authProto.IdClientResponse
	errQuery   error
}{
	{
		testName:   "First",
		input:      &util.Defense{SessionId: "1", DateLife: time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)},
		out:        1,
		outErr:     "",
		inputQuery: &authProto.Defense{SessionId: "1", DateLife: timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local))},
		outQuery:   &authProto.IdClientResponse{IdUser: 1},
		errQuery:   nil,
	},
}

func TestGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionMiddlewareInterface(ctrl)
	for _, tt := range GetIdByCookie {
		m.
			EXPECT().
			GetIdByCookie(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
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

var CheckAccessWebsocket = []struct {
	testName                 string
	input                    string
	out                      bool
	outErr                   string
	inputQuery               string
	outQuery                 Row
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputQuery:               "1",
		outQuery:                 Row{row: []interface{}{1}},
		input:                    "1",
		out:                      true,
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestCheckAccessWebsocket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range CheckAccessWebsocket {
		m.
			EXPECT().
			Begin(gomock.Any()).
			Return(mTx, tt.errBeginTransaction)
		mTx.
			EXPECT().
			Commit(gomock.Any()).
			Return(tt.errCommitTransaction).
			Times(tt.countCommitTransaction)
		mTx.
			EXPECT().
			Rollback(gomock.Any()).
			Return(nil).
			Times(tt.countRollbackTransaction)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id FROM cookie WHERE websocket = $1",
				tt.input).
			Return(&tt.outQuery)
		test := Wrapper{DBConn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CheckAccessWebsocket(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
