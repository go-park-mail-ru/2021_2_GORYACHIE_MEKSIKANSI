package orm

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/orm/mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Row struct {
	row    []interface{}
	errRow error
}

func (r Row) Scan(dest ...interface{}) error {
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
		}
	}
	return nil
}

func TestOrmGenerateNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	testUser := &Wrapper{Conn: m}
	t.Run("One", func(t *testing.T) {
		result := testUser.NewDefense()
		require.NotEqual(t, time.Time{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", time.Time{}, result.DateLife))
		require.NotEqual(t, "", result, fmt.Sprintf("Expected: %v\nbut got: %v", "", result.SessionId))
		require.NotEqual(t, "", result, fmt.Sprintf("Expected: %v\nbut got: %v", "", result.CsrfToken))
	})
}

var OrmGeneralSignUp = []struct {
	testName         string
	inputSignup      *authPkg.RegistrationRequest
	inputTransaction pgx.Tx
	inputQueryPhone  string
	inputQueryEmail  string
	inputQueryName   string
	resultQuery      Row
	out              int
	outErr           string
}{
	{
		testName:        "One",
		out:             1,
		inputSignup:     &authPkg.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
		resultQuery:     Row{row: []interface{}{1}, errRow: nil},
		inputQueryPhone: "1",
		inputQueryEmail: "1",
		inputQueryName:  "1",
	},
}

func TestOrmGeneralSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGeneralSignUp {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				tt.inputQueryName, tt.inputQueryEmail, tt.inputQueryPhone, gomock.Any(), gomock.Any(),
			).
			Return(&tt.resultQuery)
		testUser := &Wrapper{}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.generalSignUp(tt.inputSignup, m)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmLoginByEmail = []struct {
	testName           string
	inputEmail         string
	inputPassword      string
	resultQuerySalt    Row
	resultQueryId      Row
	out                int
	outErr             string
	inputQuerySalt     string
	inputQueryPassword string
}{
	{
		testName:           "One",
		out:                1,
		resultQuerySalt:    Row{row: []interface{}{"1"}, errRow: nil},
		resultQueryId:      Row{row: []interface{}{1}, errRow: nil},
		inputQuerySalt:     "1",
		inputQueryPassword: "4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8",
		inputEmail:         "1",
		inputPassword:      "1",
	},
}

func TestOrmLoginByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmLoginByEmail {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT salt FROM general_user_info WHERE email = $1",
				tt.inputQuerySalt,
			).
			Return(&tt.resultQuerySalt)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
				tt.inputQuerySalt, tt.inputQueryPassword,
			).
			Return(&tt.resultQueryId)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.LoginByEmail(tt.inputEmail, tt.inputPassword)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmLoginByPhone = []struct {
	testName           string
	inputPhone         string
	inputPassword      string
	resultQuerySalt    Row
	resultQueryId      Row
	out                int
	outErr             string
	inputQuerySalt     string
	inputQueryPassword string
}{
	{
		testName:           "One",
		out:                1,
		resultQuerySalt:    Row{row: []interface{}{"1"}, errRow: nil},
		resultQueryId:      Row{row: []interface{}{1}, errRow: nil},
		inputQuerySalt:     "1",
		inputQueryPassword: "4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8",
		inputPhone:         "1",
		inputPassword:      "1",
	},
}

func TestOrmLoginByPhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmLoginByPhone {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT salt FROM general_user_info WHERE phone = $1",
				tt.inputQuerySalt,
			).
			Return(&tt.resultQuerySalt)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
				tt.inputQuerySalt, tt.inputQueryPassword,
			).
			Return(&tt.resultQueryId)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.LoginByPhone(tt.inputPhone, tt.inputPassword)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmDeleteCookie = []struct {
	testName    string
	input       string
	out         string
	outErr      string
	inputDelete string
	errDelete   error
	countDelete int
	inputQuery  string
	errQuery    error
	resultQuery Row
}{
	{
		testName:    "One",
		input:       "1",
		out:         "1",
		outErr:      "",
		inputDelete: "1",
		errDelete:   nil,
		countDelete: 1,
		inputQuery:  "1",
		errQuery:    nil,
		resultQuery: Row{row: []interface{}{"1"}},
	},
}

func TestOrmDeleteCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmDeleteCookie {
		m.
			EXPECT().
			Exec(context.Background(),
				"DELETE FROM cookie WHERE csrf_token = $1",
				tt.inputDelete,
			).
			Return(nil, tt.errDelete).
			Times(tt.countDelete)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT session_id FROM cookie WHERE csrf_token = $1",
				tt.inputQuery,
			).
			Return(tt.resultQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.DeleteCookie(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmAddCookie = []struct {
	testName            string
	inputCookie         *util.Defense
	inputId             int
	outErr              string
	errQuery            error
	inputQuerySessionId string
	inputQueryCSRFToken string
	inputQueryClientId  int
	inputQueryDateLife  time.Time
}{
	{
		testName:            "One",
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		inputCookie:         &util.Defense{SessionId: "1", CsrfToken: "1"},
		inputId:             1,
		errQuery:            nil,
		inputQueryClientId:  1,
	},
}

func TestOrmAddCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmAddCookie {
		m.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
				tt.inputQueryClientId, tt.inputQuerySessionId, tt.inputQueryDateLife, tt.inputQueryCSRFToken,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.AddCookie(tt.inputCookie, tt.inputId)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmAddTransactionCookie = []struct {
	testName            string
	inputCookie         *util.Defense
	inputId             int
	outErr              string
	errQuery            error
	inputQuerySessionId string
	inputQueryCSRFToken string
	inputQueryClientId  int
	inputQueryDateLife  time.Time
}{
	{
		testName:            "One",
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		inputCookie:         &util.Defense{SessionId: "1", CsrfToken: "1"},
		inputId:             1,
		errQuery:            nil,
		inputQueryClientId:  1,
	},
}

func TestOrmAddTransactionCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmAddTransactionCookie {
		m.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
				tt.inputQueryClientId, tt.inputQuerySessionId, tt.inputQueryDateLife, tt.inputQueryCSRFToken,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.addTransactionCookie(tt.inputCookie, m, tt.inputId)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmSignupClient = []struct {
	testName                  string
	inputCookie               *util.Defense
	inputSignUp               *authPkg.RegistrationRequest
	outErr                    string
	errQueryCookie            error
	inputQueryCookieSessionId string
	inputQueryCookieCSRFToken string
	inputQueryCookieClientId  int
	inputQueryCookieDateLife  time.Time
	inputQueryInfoPhone       string
	inputQueryInfoEmail       string
	inputQueryInfoName        string
	resultQueryInfo           Row
	inputInsert               int
	ErrInsert                 error
	countInsert               int
	countQueryCookie          int
	countQueryInfo            int
	countRollback             int
	countCommit               int
	errRollback               error
	errCommit                 error
}{
	{
		testName:                  "One",
		inputQueryCookieSessionId: "1",
		inputQueryCookieCSRFToken: "1",
		inputQueryCookieDateLife:  time.Time{},
		inputCookie:               &util.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp:               &authPkg.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
		errQueryCookie:            nil,
		inputQueryCookieClientId:  1,
		resultQueryInfo:           Row{row: []interface{}{1}, errRow: nil},
		inputQueryInfoPhone:       "1",
		inputQueryInfoEmail:       "1",
		inputQueryInfoName:        "1",
		inputInsert:               1,
		ErrInsert:                 nil,
		countInsert:               1,
		countQueryCookie:          1,
		countQueryInfo:            1,
		countRollback:             0,
		countCommit:               1,
		errRollback:               nil,
		errCommit:                 nil,
		outErr:                    "",
	},
}

func TestOrmSignupClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmSignupClient {
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO client (client_id) VALUES ($1)",
				tt.inputInsert).
			Return(nil, tt.ErrInsert).
			Times(tt.countInsert)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
				tt.inputQueryCookieClientId, tt.inputQueryCookieSessionId, tt.inputQueryCookieDateLife, tt.inputQueryCookieCSRFToken,
			).
			Return(nil, tt.errQueryCookie).
			Times(tt.countQueryCookie)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				tt.inputQueryInfoName, tt.inputQueryInfoEmail, tt.inputQueryInfoPhone, gomock.Any(), gomock.Any(),
			).
			Return(&tt.resultQueryInfo).
			Times(tt.countQueryInfo)
		mTx.
			EXPECT().
			Rollback(context.Background()).
			Return(tt.errRollback).
			Times(tt.countRollback)
		mTx.
			EXPECT().
			Commit(context.Background()).
			Return(tt.errCommit).
			Times(tt.countCommit)
		m.
			EXPECT().
			Begin(context.Background()).
			Return(mTx, nil)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupClient(tt.inputSignUp, tt.inputCookie)
			require.NotEqual(t, &util.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &util.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmSignupCourier = []struct {
	testName                  string
	inputCookie               *util.Defense
	inputSignUp               *authPkg.RegistrationRequest
	outErr                    string
	errQueryCookie            error
	inputQueryCookieSessionId string
	inputQueryCookieCSRFToken string
	inputQueryCookieClientId  int
	inputQueryCookieDateLife  time.Time
	inputQueryInfoPhone       string
	inputQueryInfoEmail       string
	inputQueryInfoName        string
	resultQueryInfo           Row
	inputInsert               int
	ErrInsert                 error
	countInsert               int
	countQueryCookie          int
	countQueryInfo            int
	countRollback             int
	countCommit               int
	errRollback               error
	errCommit                 error
}{
	{
		testName:                  "One",
		inputQueryCookieSessionId: "1",
		inputQueryCookieCSRFToken: "1",
		inputQueryCookieDateLife:  time.Time{},
		inputCookie:               &util.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp:               &authPkg.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
		errQueryCookie:            nil,
		inputQueryCookieClientId:  1,
		resultQueryInfo:           Row{row: []interface{}{1}, errRow: nil},
		inputQueryInfoPhone:       "1",
		inputQueryInfoEmail:       "1",
		inputQueryInfoName:        "1",
		inputInsert:               1,
		ErrInsert:                 nil,
		countInsert:               1,
		countQueryCookie:          1,
		countQueryInfo:            1,
		countRollback:             0,
		countCommit:               1,
		errRollback:               nil,
		errCommit:                 nil,
		outErr:                    "",
	},
}

func TestOrmSignupCourier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmSignupCourier {
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO courier (client_id) VALUES ($1)",
				tt.inputInsert).
			Return(nil, tt.ErrInsert).
			Times(tt.countInsert)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
				tt.inputQueryCookieClientId, tt.inputQueryCookieSessionId, tt.inputQueryCookieDateLife, tt.inputQueryCookieCSRFToken,
			).
			Return(nil, tt.errQueryCookie).
			Times(tt.countQueryCookie)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				tt.inputQueryInfoName, tt.inputQueryInfoEmail, tt.inputQueryInfoPhone, gomock.Any(), gomock.Any(),
			).
			Return(&tt.resultQueryInfo).
			Times(tt.countQueryInfo)
		mTx.
			EXPECT().
			Rollback(context.Background()).
			Return(tt.errRollback).
			Times(tt.countRollback)
		mTx.
			EXPECT().
			Commit(context.Background()).
			Return(tt.errCommit).
			Times(tt.countCommit)
		m.
			EXPECT().
			Begin(context.Background()).
			Return(mTx, nil)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupCourier(tt.inputSignUp, tt.inputCookie)
			require.NotEqual(t, &util.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &util.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmSignupHost = []struct {
	testName                  string
	inputCookie               *util.Defense
	inputSignUp               *authPkg.RegistrationRequest
	outErr                    string
	errQueryCookie            error
	inputQueryCookieSessionId string
	inputQueryCookieCSRFToken string
	inputQueryCookieClientId  int
	inputQueryCookieDateLife  time.Time
	inputQueryInfoPhone       string
	inputQueryInfoEmail       string
	inputQueryInfoName        string
	resultQueryInfo           Row
	inputInsert               int
	ErrInsert                 error
	countInsert               int
	countQueryCookie          int
	countQueryInfo            int
	countRollback             int
	countCommit               int
	errRollback               error
	errCommit                 error
}{
	{
		testName:                  "One",
		inputQueryCookieSessionId: "1",
		inputQueryCookieCSRFToken: "1",
		inputQueryCookieDateLife:  time.Time{},
		inputCookie:               &util.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp:               &authPkg.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
		errQueryCookie:            nil,
		inputQueryCookieClientId:  1,
		resultQueryInfo:           Row{row: []interface{}{1}, errRow: nil},
		inputQueryInfoPhone:       "1",
		inputQueryInfoEmail:       "1",
		inputQueryInfoName:        "1",
		inputInsert:               1,
		ErrInsert:                 nil,
		countInsert:               1,
		countQueryCookie:          1,
		countQueryInfo:            1,
		countRollback:             0,
		countCommit:               1,
		errRollback:               nil,
		errCommit:                 nil,
		outErr:                    "",
	},
}

func TestOrmSignupHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmSignupHost {
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO host (client_id) VALUES ($1)",
				tt.inputInsert).
			Return(nil, tt.ErrInsert).
			Times(tt.countInsert)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
				tt.inputQueryCookieClientId, tt.inputQueryCookieSessionId, tt.inputQueryCookieDateLife, tt.inputQueryCookieCSRFToken,
			).
			Return(nil, tt.errQueryCookie).
			Times(tt.countQueryCookie)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				tt.inputQueryInfoName, tt.inputQueryInfoEmail, tt.inputQueryInfoPhone, gomock.Any(), gomock.Any(),
			).
			Return(&tt.resultQueryInfo).
			Times(tt.countQueryInfo)
		mTx.
			EXPECT().
			Rollback(context.Background()).
			Return(tt.errRollback).
			Times(tt.countRollback)
		mTx.
			EXPECT().
			Commit(context.Background()).
			Return(tt.errCommit).
			Times(tt.countCommit)
		m.
			EXPECT().
			Begin(context.Background()).
			Return(mTx, nil)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupHost(tt.inputSignUp, tt.inputCookie)
			require.NotEqual(t, &util.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &util.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
