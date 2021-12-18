package orm

import (
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/orm/mocks"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
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
		if r.row[i] == nil {
			dest[i] = nil
			continue
		}
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
		case *time.Time:
			*dest[i].(*time.Time) = r.row[i].(time.Time)
		case *bool:
			*dest[i].(*bool) = r.row[i].(bool)
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
		testName: "Check access",
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
		outQuery: &authProto.CheckAccess{
			CheckResult: true,
			Error:       "",
		},
	},
	{
		testName: "Error check access",
		input: &util.Defense{
			SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		out:    false,
		outErr: "text",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CheckAccess{
			CheckResult: true,
			Error:       "text",
		},
	},
	{
		testName: "Error microservice",
		input: &util.Defense{
			SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		out:    false,
		outErr: "text",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CheckAccess{
			CheckResult: true,
			Error:       "",
		},
		errQuery: errors.New("text"),
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
	testName   string
	input      *util.Defense
	out        string
	outErr     string
	inputQuery *authProto.Defense
	outQuery   *authProto.CSRFResponse
	errQuery   error
}{
	{
		testName: "Generate new csrf",
		input: &util.Defense{
			SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		out:    "",
		outErr: "",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CSRFResponse{
			XCsrfToken: &authProto.CSRF{
				XCsrfToken: "1",
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName: "Error generate new csrf",
		input: &util.Defense{
			SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		out:    "text",
		outErr: "text",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CSRFResponse{
			XCsrfToken: &authProto.CSRF{
				XCsrfToken: "1",
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName: "Error microservice",
		input: &util.Defense{
			SessionId: "1",
			CsrfToken: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
		},
		out:    "text",
		outErr: "text",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			XCsrfToken: "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
		},
		outQuery: &authProto.CSRFResponse{
			XCsrfToken: &authProto.CSRF{
				XCsrfToken: "1",
			},
			Error: "",
		},
		errQuery: errors.New("text"),
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
			require.NotEqual(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
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
	input      *util.Defense
	out        int
	outErr     string
	inputQuery *authProto.Defense
	outQuery   *authProto.IdClientResponse
	errQuery   error
}{
	{
		testName: "Get id",
		input: &util.Defense{
			SessionId: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
			CsrfToken: "",
		},
		out:    1,
		outErr: "",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
			XCsrfToken: "",
		},
		outQuery: &authProto.IdClientResponse{
			IdUser: 1,
			Error:  "",
		},
		errQuery: nil,
	},
	{
		testName: "Error get id",
		input: &util.Defense{
			SessionId: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
			CsrfToken: "",
		},
		out:    0,
		outErr: "text",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
			XCsrfToken: "",
		},
		outQuery: &authProto.IdClientResponse{
			IdUser: 1,
			Error:  "text",
		},
		errQuery: nil,
	},
	{
		testName: "Error microservice",
		input: &util.Defense{
			SessionId: "1",
			DateLife:  time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local),
			CsrfToken: "",
		},
		out:    0,
		outErr: "text",
		inputQuery: &authProto.Defense{
			SessionId:  "1",
			DateLife:   timestamp.New(time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)),
			XCsrfToken: "",
		},
		outQuery: &authProto.IdClientResponse{
			IdUser: 1,
			Error:  "",
		},
		errQuery: errors.New("text"),
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
	testName                 string
	input                    string
	out                      bool
	outErr                   string
	inputQuery               string
	outQuery                 Row
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "Check websocket",
		input:                    "1",
		out:                      true,
		outErr:                   "",
		inputQuery:               "1",
		outQuery:                 Row{row: []interface{}{1}},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:                 "Error begin transaction",
		input:                    "1",
		out:                      false,
		outErr:                   errPkg.MCheckAccessWebsocketTransactionNotCreate,
		inputQuery:               "1",
		outQuery:                 Row{row: []interface{}{1}},
		countQuery:               0,
		errBeginTransaction:      errors.New(errPkg.MCheckAccessWebsocketTransactionNotCreate),
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 0,
	},
	{
		testName:                 "Error select",
		input:                    "1",
		out:                      false,
		outErr:                   errPkg.MCheckAccessWebsocketNotSelect,
		inputQuery:               "1",
		outQuery:                 Row{errRow: errors.New("text"), row: []interface{}{1}},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:                 "Error commit",
		input:                    "1",
		out:                      false,
		outErr:                   errPkg.MCheckAccessWebsocketNotCommit,
		inputQuery:               "1",
		outQuery:                 Row{row: []interface{}{1}},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     errors.New("text"),
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:                 "Check websocket permission denied, error commit",
		input:                    "1",
		out:                      false,
		outErr:                   errPkg.MCheckAccessWebsocketNotCommit,
		inputQuery:               "1",
		outQuery:                 Row{row: []interface{}{1}, errRow: pgx.ErrNoRows},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     errors.New("text"),
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:                 "Check websocket permission denied",
		input:                    "1",
		out:                      false,
		outErr:                   "",
		inputQuery:               "1",
		outQuery:                 Row{row: []interface{}{1}, errRow: pgx.ErrNoRows},
		countQuery:               1,
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
			Return(&tt.outQuery).
			Times(tt.countQuery)
		test := Wrapper{DBConn: m}
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
