package Authorization

import (
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

//type Identifier []string
//
//type CopyFromSource interface {
//	Next() bool
//	Values() ([]interface{}, error)
//	Err() error
//}
//type batchItem struct {
//	query     string
//	arguments []interface{}
//}
//type Batch struct {
//	items []*batchItem
//}
//type BatchResults interface {
//	Exec() (pgconn.CommandTag, error)
//	Query() (Rows, error)
//	QueryRow() Row
//	Close() error
//}
//type LargeObjects struct {
//	tx Tx
//}
//type Rows interface {
//	Close()
//	Err() error
//	CommandTag() pgconn.CommandTag
//
//	FieldDescriptions() []pgproto3.FieldDescription
//	Next() bool
//	Scan(dest ...interface{}) error
//	Values() ([]interface{}, error)
//	RawValues() [][]byte
//}
//type Row interface {
//	Scan(dest ...interface{}) error
//}
//type QueryFuncRow interface {
//	FieldDescriptions() []pgproto3.FieldDescription
//	RawValues() [][]byte
//}
//type Conn struct {
//	pgConn             *pgconn.PgConn
//	config             *ConnConfig // config used when establishing this connection
//	preparedStatements map[string]*pgconn.StatementDescription
//	stmtcache          stmtcache.Cache
//	logger             Logger
//	logLevel           	LogLevel
//
//	notifications []*pgconn.Notification
//
//	doneChan   chan struct{}
//	closedChan chan error
//
//	connInfo *pgtype.ConnInfo
//
//	wbuf             []byte
//	preallocatedRows []connRows
//	eqb              extendedQueryBuilder
//}
//type TxRow struct {
//
//}
//func (t *TxRow) Begin(ctx context.Context) (Tx, error) {
//	return nil, nil
//}
//
//func (t *TxRow) BeginFunc(ctx context.Context, f func(Tx) error) (err error) {
//	return nil
//}
//
//func (t *TxRow) Commit(ctx context.Context) error {
//	return nil
//}
//
//func (t *TxRow) Rollback(ctx context.Context) error {
//	return nil
//}
//
//func (t *TxRow) CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error) {
//	return 0, nil
//}
//func (t *TxRow) SendBatch(ctx context.Context, b *Batch) BatchResults {
//
//}
//func (t *TxRow) LargeObjects() LargeObjects {
//
//}
//
//func (t *TxRow) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
//
//}
//
//func (t *TxRow) Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
//
//}
//func (t *TxRow) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
//
//}
//func (t *TxRow) QueryRow(ctx context.Context, sql string, args ...interface{}) Row {
//
//}
//func (t *TxRow) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(QueryFuncRow) error) (pgconn.CommandTag, error) {
//
//}
//
//func (t *TxRow) Conn() *Conn {
//	return nil
//}

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
		}
	}
	return nil
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
	testName            string
	input               *utils.Defense
	resultQuerySalt     Row
	resultQueryId       Row
	outErr              string
	errQuerySalt        error
	inputQuerySessionId string
	inputQueryCSRFToken string
}{
	{
		testName:            "One",
		resultQuerySalt:     Row{row: []interface{}{"1"}, errRow: nil},
		resultQueryId:       Row{row: []interface{}{1}, errRow: nil},
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		input:               &utils.Defense{SessionId: "1", CsrfToken: "1"},
		errQuerySalt:        nil,
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
				"DELETE FROM cookie WHERE session_id = $1 AND csrf_token = $2",
				tt.inputQuerySessionId, tt.inputQueryCSRFToken,
			).
			Return(nil, tt.errQuerySalt)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.DeleteCookie(tt.input)
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
	inputCookie         *utils.Defense
	inputId             int
	outErr              string
	errQuerySalt        error
	inputQuerySessionId string
	inputQueryCSRFToken string
	inputQueryClientId  int
	inputQueryDateLife  time.Time
}{
	{
		testName:            "One",
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		inputCookie:         &utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputId:             1,
		errQuerySalt:        nil,
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
			Return(nil, tt.errQuerySalt)
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

var ApplicationSignUp = []struct {
	testName                 string
	out                      *utils.Defense
	outErr                   string
	input                    *utils.RegistrationRequest
	inputSignupClientSignUp  *utils.RegistrationRequest
	resultSignupClient       *utils.Defense
	errSignupClient          error
	countSignupClient        int
	inputSignupCourierSignUp *utils.RegistrationRequest
	resultSignupCourier      *utils.Defense
	errSignupCourier         error
	countSignupCourier       int
	inputSignupHostSignUp    *utils.RegistrationRequest
	resultSignupHost         *utils.Defense
	errSignupHost            error
	countSignupHost          int
	resultGenerateNew        *utils.Defense
	inputSignupClientCookie  *utils.Defense
	inputSignupCourierCookie *utils.Defense
	inputSignupHostCookie    *utils.Defense
}{
	{
		input:                   &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName:                "One",
		outErr:                  "",
		resultSignupClient:      &utils.Defense{},
		inputSignupClientSignUp: &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out:                     &utils.Defense{},
		errSignupClient:         nil,
		countSignupClient:       1,
		resultGenerateNew:       &utils.Defense{},
		inputSignupClientCookie: &utils.Defense{},
	},
	{
		input:                    &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "courier"},
		testName:                 "Two",
		outErr:                   "",
		resultSignupCourier:      &utils.Defense{},
		inputSignupCourierSignUp: &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "courier"},
		out:                      &utils.Defense{},
		errSignupCourier:         nil,
		countSignupCourier:       1,
		resultGenerateNew:        &utils.Defense{},
		inputSignupCourierCookie: &utils.Defense{},
	},
	{
		input:                 &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "host"},
		testName:              "Three",
		outErr:                "",
		resultSignupHost:      &utils.Defense{},
		inputSignupHostSignUp: &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "host"},
		out:                   &utils.Defense{},
		errSignupHost:         nil,
		countSignupHost:       1,
		resultGenerateNew:     &utils.Defense{},
		inputSignupHostCookie: &utils.Defense{},
	},
	{
		input:                   &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName:                "Four",
		outErr:                  "text",
		resultSignupClient:      &utils.Defense{},
		inputSignupClientSignUp: &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out:                     nil,
		errSignupClient:         errors.New("text"),
		countSignupClient:       1,
		resultGenerateNew:       &utils.Defense{},
		inputSignupClientCookie: &utils.Defense{},
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
		t.Run(tt.testName, func(t *testing.T) {
			result, err := SignUp(m, tt.input)
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
	out                  *utils.Defense
	outErr               string
	input                *Authorization
	inputLoginEmail      string
	inputLoginPassword   string
	resultLogin          int
	errLogin             error
	countLoginEmail      int
	countLoginPhone      int
	inputLoginPhone      string
	resultGenerateNew    *utils.Defense
	countGenerateNew     int
	inputAddCookieCookie *utils.Defense
	inputAddCookieId     int
	errAddCookie         error
	countAddCookie       int
}{
	{
		input:                &Authorization{Email: "1", Phone: "", Password: "1"},
		testName:             "One",
		outErr:               "",
		out:                  &utils.Defense{},
		inputLoginEmail:      "1",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      1,
		countLoginPhone:      0,
		resultGenerateNew:    &utils.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &utils.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         nil,
		countAddCookie:       1,
	},
	{
		input:                &Authorization{Email: "", Phone: "1", Password: "1"},
		testName:             "Two",
		outErr:               "",
		out:                  &utils.Defense{},
		inputLoginEmail:      "",
		inputLoginPhone:      "1",
		inputLoginPassword:   "1",
		resultLogin:          1,
		errLogin:             nil,
		countLoginEmail:      0,
		countLoginPhone:      1,
		resultGenerateNew:    &utils.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &utils.Defense{},
		inputAddCookieId:     1,
		errAddCookie:         nil,
		countAddCookie:       1,
	},
	{
		input:                &Authorization{Email: "", Phone: "1", Password: "1"},
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
		resultGenerateNew:    &utils.Defense{},
		countGenerateNew:     1,
		inputAddCookieCookie: &utils.Defense{},
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
		t.Run(tt.testName, func(t *testing.T) {
			result, err := Login(m, tt.input)
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
	testName    string
	out         *utils.Defense
	outErr      string
	input       *utils.Defense
	inputDelete *utils.Defense
	errDelete   error
}{
	{
		input:       &utils.Defense{SessionId: "1", CsrfToken: "1"},
		inputDelete: &utils.Defense{SessionId: "1", CsrfToken: "1"},
		testName:    "One",
		outErr:      "",
		out:         &utils.Defense{},
		errDelete:   nil,
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
			Return(tt.errDelete)
		t.Run(tt.testName, func(t *testing.T) {
			err := Logout(m, tt.input)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
