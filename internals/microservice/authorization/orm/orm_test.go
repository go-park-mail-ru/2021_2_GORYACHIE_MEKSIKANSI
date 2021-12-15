package orm

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/orm/mocks"
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

func TestGenerateNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	testUser := &Wrapper{Conn: m}
	t.Run("First", func(t *testing.T) {
		result := testUser.NewDefense()
		require.NotEqual(t, time.Time{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", time.Time{}, result.DateLife))
		require.NotEqual(t, "", result, fmt.Sprintf("Expected: %v\nbut got: %v", "", result.SessionId))
		require.NotEqual(t, "", result, fmt.Sprintf("Expected: %v\nbut got: %v", "", result.CsrfToken))
	})
}

var GeneralSignUp = []struct {
	testName                 string
	inputSignup              *authPkg.RegistrationRequest
	inputTransaction         pgx.Tx
	inputQueryPhone          string
	inputQueryEmail          string
	inputQueryName           string
	resultQuery              Row
	out                      int
	outErr                   string
	errCommitTransaction     error
	countCommitTransaction   int
	countRollbackTransaction int
}{
	{
		testName: "First",
		out:      1,
		inputSignup: &authPkg.RegistrationRequest{
			Phone:    "89165554433",
			Email:    "1",
			Password: "1",
			Name:     "1",
		},
		resultQuery:              Row{row: []interface{}{1}, errRow: nil},
		inputQueryPhone:          "89165554433",
		inputQueryEmail:          "1",
		inputQueryName:           "1",
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		countRollbackTransaction: 1,
	},
}

func TestGeneralSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GeneralSignUp {
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
				tt.inputQueryName, tt.inputQueryEmail, tt.inputQueryPhone, gomock.Any(), gomock.Any(),
			).
			Return(&tt.resultQuery)
		testUser := Wrapper{}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.generalSignUp(tt.inputSignup, mTx, context.Background())
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var LoginByEmail = []struct {
	testName                 string
	inputEmail               string
	inputPassword            string
	resultQuerySalt          Row
	resultQueryId            Row
	out                      int
	outErr                   string
	inputQuerySalt           string
	inputQueryPassword       string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		out:                      1,
		resultQuerySalt:          Row{row: []interface{}{"1"}, errRow: nil},
		resultQueryId:            Row{row: []interface{}{1}, errRow: nil},
		inputQuerySalt:           "1",
		inputQueryPassword:       "4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8",
		inputEmail:               "1",
		inputPassword:            "1",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestLoginByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range LoginByEmail {
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
				"SELECT salt FROM general_user_info WHERE email = $1",
				tt.inputQuerySalt,
			).
			Return(&tt.resultQuerySalt)
		mTx.
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

var LoginByPhone = []struct {
	testName                 string
	inputPhone               string
	inputPassword            string
	resultQuerySalt          Row
	resultQueryId            Row
	out                      int
	outErr                   string
	inputQuerySalt           string
	inputQueryPassword       string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		out:                      1,
		resultQuerySalt:          Row{row: []interface{}{"1"}, errRow: nil},
		resultQueryId:            Row{row: []interface{}{1}, errRow: nil},
		inputQuerySalt:           "1",
		inputQueryPassword:       "4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8",
		inputPhone:               "1",
		inputPassword:            "1",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestLoginByPhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range LoginByPhone {
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
				"SELECT salt FROM general_user_info WHERE phone = $1",
				tt.inputQuerySalt,
			).
			Return(&tt.resultQuerySalt)
		mTx.
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

var DeleteCookie = []struct {
	testName                 string
	input                    string
	out                      string
	outErr                   string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
	inputDelete              string
	outDelete                Row
	countDelete              int
}{
	{
		testName:                 "First",
		input:                    "1",
		out:                      "1",
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
		inputDelete:              "1",
		outDelete:                Row{row: []interface{}{"1"}},
		countDelete:              1,
	},
}

func TestDeleteCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range DeleteCookie {
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
				"DELETE FROM cookie WHERE csrf_token = $1 RETURNING session_id",
				tt.inputDelete,
			).
			Return(&tt.outDelete).
			Times(tt.countDelete)
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

var AddCookie = []struct {
	testName                 string
	inputCookie              *authPkg.Defense
	inputId                  int
	outErr                   string
	errQuery                 error
	inputQuerySessionId      string
	inputQueryCSRFToken      string
	inputQueryClientId       int
	inputQueryDateLife       time.Time
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputQuerySessionId:      "1",
		inputQueryCSRFToken:      "1",
		inputCookie:              &authPkg.Defense{SessionId: "1", CsrfToken: "1"},
		inputId:                  1,
		errQuery:                 nil,
		inputQueryClientId:       1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestAddCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range AddCookie {
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

var AddTransactionCookie = []struct {
	testName            string
	inputCookie         *authPkg.Defense
	inputId             int
	outErr              string
	errQuery            error
	inputQuerySessionId string
	inputQueryCSRFToken string
	inputQueryClientId  int
	inputQueryDateLife  time.Time
}{
	{
		testName:            "First",
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		inputCookie:         &authPkg.Defense{SessionId: "1", CsrfToken: "1"},
		inputId:             1,
		errQuery:            nil,
		inputQueryClientId:  1,
	},
}

func TestAddTransactionCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range AddTransactionCookie {
		m.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
				tt.inputQueryClientId, tt.inputQuerySessionId, tt.inputQueryDateLife, tt.inputQueryCSRFToken,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.addTransactionCookie(tt.inputCookie, m, tt.inputId, context.Background())
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var SignupClient = []struct {
	testName                  string
	inputCookie               *authPkg.Defense
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
	errBeginTransaction       error
	errCommitTransaction      error
	countCommitTransaction    int
	errRollbackTransaction    error
	countRollbackTransaction  int
}{
	{
		testName:                  "First",
		inputQueryCookieSessionId: "1",
		inputQueryCookieCSRFToken: "1",
		inputQueryCookieDateLife:  time.Time{},
		inputCookie:               &authPkg.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp: &authPkg.RegistrationRequest{
			Phone:    "89175554433",
			Email:    "1",
			Password: "1",
			Name:     "1",
		},
		errQueryCookie:           nil,
		inputQueryCookieClientId: 1,
		resultQueryInfo:          Row{row: []interface{}{1}, errRow: nil},
		inputQueryInfoPhone:      "89175554433",
		inputQueryInfoEmail:      "1",
		inputQueryInfoName:       "1",
		inputInsert:              1,
		ErrInsert:                nil,
		countInsert:              1,
		countQueryCookie:         1,
		countQueryInfo:           1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
		outErr:                   "",
	},
}

func TestSignupClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range SignupClient {
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
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupClient(tt.inputSignUp, tt.inputCookie)
			require.NotEqual(t, &authPkg.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &authPkg.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var SignupCourier = []struct {
	testName                  string
	inputCookie               *authPkg.Defense
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
	errBeginTransaction       error
	errCommitTransaction      error
	countCommitTransaction    int
	errRollbackTransaction    error
	countRollbackTransaction  int
}{
	{
		testName:                  "First",
		inputQueryCookieSessionId: "1",
		inputQueryCookieCSRFToken: "1",
		inputQueryCookieDateLife:  time.Time{},
		inputCookie:               &authPkg.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp: &authPkg.RegistrationRequest{Phone: "89175554433",
			Email:    "1",
			Password: "1",
			Name:     "1",
		},
		errQueryCookie:           nil,
		inputQueryCookieClientId: 1,
		resultQueryInfo:          Row{row: []interface{}{1}, errRow: nil},
		inputQueryInfoPhone:      "89175554433",
		inputQueryInfoEmail:      "1",
		inputQueryInfoName:       "1",
		inputInsert:              1,
		ErrInsert:                nil,
		countInsert:              1,
		countQueryCookie:         1,
		countQueryInfo:           1,
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestSignupCourier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range SignupCourier {
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
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupCourier(tt.inputSignUp, tt.inputCookie)
			require.NotEqual(t, &authPkg.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &authPkg.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var SignupHost = []struct {
	testName                  string
	inputCookie               *authPkg.Defense
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
	errBeginTransaction       error
	errCommitTransaction      error
	countCommitTransaction    int
	errRollbackTransaction    error
	countRollbackTransaction  int
}{
	{
		testName:                  "First",
		inputQueryCookieSessionId: "1",
		inputQueryCookieCSRFToken: "1",
		inputQueryCookieDateLife:  time.Time{},
		inputCookie:               &authPkg.Defense{SessionId: "1", CsrfToken: "1"},
		inputSignUp: &authPkg.RegistrationRequest{
			Phone:    "89175554433",
			Email:    "1",
			Password: "1",
			Name:     "1",
		},
		errQueryCookie:           nil,
		inputQueryCookieClientId: 1,
		resultQueryInfo:          Row{row: []interface{}{1}, errRow: nil},
		inputQueryInfoPhone:      "89175554433",
		inputQueryInfoEmail:      "1",
		inputQueryInfoName:       "1",
		inputInsert:              1,
		ErrInsert:                nil,
		countInsert:              1,
		countQueryCookie:         1,
		countQueryInfo:           1,
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestSignupHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range SignupHost {
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
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupHost(tt.inputSignUp, tt.inputCookie)
			require.NotEqual(t, &authPkg.Defense{}, result, fmt.Sprintf("Expected: %v\nbut got: %v", &authPkg.Defense{}, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var CheckAccess = []struct {
	testName                  string
	input                     *authPkg.Defense
	out                       bool
	outErr                    string
	inputCheckAccessSessionId string
	inputCheckAccessCSRFToken string
	outCheckAccess            Row
	countCheckAccess          int
	errBeginTransaction       error
	errCommitTransaction      error
	countCommitTransaction    int
	errRollbackTransaction    error
	countRollbackTransaction  int
}{
	{
		testName: "First",
		input: &authPkg.Defense{
			SessionId: "1",
			CsrfToken: "1",
		},
		out:                       false,
		inputCheckAccessSessionId: "1",
		inputCheckAccessCSRFToken: "1",
		outCheckAccess:            Row{row: []interface{}{1, time.Now()}},
		countCheckAccess:          1,
		outErr:                    "",
		errBeginTransaction:       nil,
		errCommitTransaction:      nil,
		countCommitTransaction:    1,
		errRollbackTransaction:    nil,
		countRollbackTransaction:  1,
	},
}

func TestCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range CheckAccess {
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
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
				tt.inputCheckAccessSessionId, tt.inputCheckAccessCSRFToken,
			).
			Return(&tt.outCheckAccess).
			Times(tt.countCheckAccess)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.CheckAccess(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var NewCSRF = []struct {
	testName                 string
	input                    *authPkg.Defense
	out                      string
	outErr                   string
	inputNewCSRF             string
	errNewCSRF               error
	countNewCSRF             int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName: "First",
		input: &authPkg.Defense{
			SessionId: "1",
			CsrfToken: "1",
		},
		out:                      "text",
		inputNewCSRF:             "1",
		errNewCSRF:               nil,
		countNewCSRF:             1,
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestNewCSRF(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range NewCSRF {
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
			Exec(context.Background(),
				"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
				gomock.Any(), tt.inputNewCSRF,
			).
			Return(nil, tt.errNewCSRF).
			Times(tt.countNewCSRF)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.NewCSRF(tt.input)
			require.NotNil(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetIdByCookie = []struct {
	testName                 string
	input                    *authPkg.Defense
	out                      int
	outErr                   string
	inputGetIdByCookie       string
	outGetIdByCookie         Row
	countGetIdByCookie       int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName: "First",
		input: &authPkg.Defense{
			SessionId: "1",
			CsrfToken: "1",
		},
		out:                0,
		outErr:             errPkg.MGetIdByCookieCookieExpired,
		inputGetIdByCookie: "1",
		outGetIdByCookie: Row{row: []interface{}{
			1,
			time.Now().Add(1 * time.Second),
		},
		},
		countGetIdByCookie:       1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetIdByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetIdByCookie {
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
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
				tt.inputGetIdByCookie,
			).
			Return(&tt.outGetIdByCookie).
			Times(tt.countGetIdByCookie)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetIdByCookie(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
