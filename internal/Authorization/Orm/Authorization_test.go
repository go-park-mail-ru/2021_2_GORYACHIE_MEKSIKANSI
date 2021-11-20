package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Utils"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/test/Mocks"
	"context"
	"errors"
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
	inputSignup      *Authorization.RegistrationRequest
	inputTransaction pgx.Tx
	inputQueryPhone  string
	inputQueryEmail  string
	inputQueryName string
	resultQuery    Row
	out            int
	outErr           string
}{
	{
		testName:        "One",
		out:             1,
		inputSignup:     &Authorization.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
		resultQuery:     Row{row: []interface{}{1}, errRow: nil},
		inputQueryPhone: "1",
		inputQueryEmail: "1",
		inputQueryName:  "1",
	},
}

func TestOrmGeneralSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockTransactionInterface(ctrl)
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
	inputPassword   string
	resultQuerySalt Row
	resultQueryId   Row
	out             int
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
	inputPassword   string
	resultQuerySalt Row
	resultQueryId   Row
	out             int
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
	inputCookie         *Utils.Defense
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
		inputCookie:         &Utils.Defense{SessionId: "1", CsrfToken: "1"},
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
	inputCookie         *Utils.Defense
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
		inputCookie:         &Utils.Defense{SessionId: "1", CsrfToken: "1"},
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
	inputCookie               *Utils.Defense
	inputSignUp               *Authorization.RegistrationRequest
	outErr                    string
	errQueryCookie            error
	inputQueryCookieSessionId string
	inputQueryCookieCSRFToken string
	inputQueryCookieClientId  int
	inputQueryCookieDateLife  time.Time
	inputQueryInfoPhone       string
	inputQueryInfoEmail       string
	inputQueryInfoName string
	resultQueryInfo    Row
	inputInsert        int
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
		inputCookie:               &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp:               &Authorization.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
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
			require.NotEqual(t, &Utils.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &Utils.Defense{}, result))
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
	inputCookie               *Utils.Defense
	inputSignUp               *Authorization.RegistrationRequest
	outErr                    string
	errQueryCookie            error
	inputQueryCookieSessionId string
	inputQueryCookieCSRFToken string
	inputQueryCookieClientId  int
	inputQueryCookieDateLife  time.Time
	inputQueryInfoPhone       string
	inputQueryInfoEmail       string
	inputQueryInfoName string
	resultQueryInfo    Row
	inputInsert        int
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
		inputCookie:               &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp:               &Authorization.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
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
			require.NotEqual(t, &Utils.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &Utils.Defense{}, result))
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
	inputCookie               *Utils.Defense
	inputSignUp               *Authorization.RegistrationRequest
	outErr                    string
	errQueryCookie            error
	inputQueryCookieSessionId string
	inputQueryCookieCSRFToken string
	inputQueryCookieClientId  int
	inputQueryCookieDateLife  time.Time
	inputQueryInfoPhone       string
	inputQueryInfoEmail       string
	inputQueryInfoName string
	resultQueryInfo    Row
	inputInsert        int
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
		inputCookie:               &Utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp:               &Authorization.RegistrationRequest{Phone: "1", Email: "1", Password: "1", Name: "1"},
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
			require.NotEqual(t, &Utils.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &Utils.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationSignUp = []struct {
	testName                 string
	out                      *Utils.Defense
	outErr                   string
	input                    *Authorization.RegistrationRequest
	inputSignupClientSignUp  *Authorization.RegistrationRequest
	resultSignupClient       *Utils.Defense
	errSignupClient          error
	countSignupClient        int
	inputSignupCourierSignUp *Authorization.RegistrationRequest
	resultSignupCourier      *Utils.Defense
	errSignupCourier         error
	countSignupCourier       int
	inputSignupHostSignUp    *Authorization.RegistrationRequest
	resultSignupHost         *Utils.Defense
	errSignupHost            error
	countSignupHost          int
	resultGenerateNew        *Utils.Defense
	inputSignupClientCookie  *Utils.Defense
	inputSignupCourierCookie *Utils.Defense
	inputSignupHostCookie    *Utils.Defense
}{
	{
		input:                   &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName:                "One",
		outErr:                  "",
		resultSignupClient:      &Utils.Defense{},
		inputSignupClientSignUp: &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out:                     &Utils.Defense{},
		errSignupClient:         nil,
		countSignupClient:       1,
		resultGenerateNew:       &Utils.Defense{},
		inputSignupClientCookie: &Utils.Defense{},
	},
	{
		input:                    &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "courier"},
		testName:                 "Two",
		outErr:                   "",
		resultSignupCourier:      &Utils.Defense{},
		inputSignupCourierSignUp: &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "courier"},
		out:                      &Utils.Defense{},
		errSignupCourier:         nil,
		countSignupCourier:       1,
		resultGenerateNew:        &Utils.Defense{},
		inputSignupCourierCookie: &Utils.Defense{},
	},
	{
		input:                 &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "host"},
		testName:              "Three",
		outErr:                "",
		resultSignupHost:      &Utils.Defense{},
		inputSignupHostSignUp: &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "host"},
		out:                   &Utils.Defense{},
		errSignupHost:         nil,
		countSignupHost:       1,
		resultGenerateNew:     &Utils.Defense{},
		inputSignupHostCookie: &Utils.Defense{},
	},
	{
		input:                   &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName:                "Four",
		outErr:                  "text",
		resultSignupClient:      &Utils.Defense{},
		inputSignupClientSignUp: &Authorization.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out:                     nil,
		errSignupClient:         errors.New("text"),
		countSignupClient:       1,
		resultGenerateNew:       &Utils.Defense{},
		inputSignupClientCookie: &Utils.Defense{},
	},
}

func TestApplicationSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range ApplicationSignUp {
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
			GenerateNew().
			Return(tt.resultGenerateNew)
		test := Application.Authorization{DB: m}
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

var ApplicationLogin = []struct {
	testName             string
	out                  *Utils.Defense
	outErr               string
	input                *Authorization.Authorization
	inputLoginEmail      string
	inputLoginPassword   string
	resultLogin          int
	errLogin             error
	countLoginEmail      int
	countLoginPhone      int
	inputLoginPhone      string
	resultGenerateNew    *Utils.Defense
	countGenerateNew     int
	inputAddCookieCookie *Utils.Defense
	inputAddCookieId     int
	errAddCookie         error
	countAddCookie       int
}{
	{
		input:                &Authorization.Authorization{Email: "1", Phone: "", Password: "1"},
		testName:             "One",
		outErr:               "",
		out:                  &Utils.Defense{},
		inputLoginEmail:      "1",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      1,
		countLoginPhone:      0,
		resultGenerateNew:    &Utils.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &Utils.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         nil,
		countAddCookie:       1,
	},
	{
		input:                &Authorization.Authorization{Email: "", Phone: "1", Password: "1"},
		testName:             "Two",
		outErr:               "",
		out:                  &Utils.Defense{},
		inputLoginEmail:      "",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      0,
		countLoginPhone:      1,
		resultGenerateNew:    &Utils.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &Utils.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         nil,
		countAddCookie:       1,
	},
	{
		input:                &Authorization.Authorization{Email: "", Phone: "1", Password: "1"},
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
		resultGenerateNew:    &Utils.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &Utils.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         errors.New("text"),
		countAddCookie:       1,
	},
}

func TestApplicationLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range ApplicationLogin {
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
			GenerateNew().
			Return(tt.resultGenerateNew).
			Times(tt.countGenerateNew)
		m.
			EXPECT().
			AddCookie(tt.inputAddCookieCookie, tt.inputAddCookieId).
			Return(tt.errAddCookie).
			Times(tt.countAddCookie)
		test := Application.Authorization{DB: m}
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

var ApplicationLogout = []struct {
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

func TestApplicationLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range ApplicationLogout {
		m.
			EXPECT().
			DeleteCookie(tt.inputDelete).
			Return(tt.resultDelete, tt.errDelete)
		test := Application.Authorization{DB: m}
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
