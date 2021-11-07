package Authorization

import (
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Tx struct {
	allGood bool
}

func (tx *Tx) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, nil
}

func (tx *Tx) BeginFunc(ctx context.Context, f func(pgx.Tx) error) error {
	return nil
}

func (tx *Tx) Commit(ctx context.Context) error {
	return nil
}

func (tx *Tx) Rollback(ctx context.Context) error {
	return nil
}

func (tx *Tx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return 0, nil
}

func (tx *Tx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}

func (tx *Tx) LargeObjects() pgx.LargeObjects {
	return pgx.LargeObjects{}
}

func (tx *Tx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return nil, nil
}

func (tx *Tx) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return nil, nil
}

func (tx *Tx) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return nil, nil
}

func (tx *Tx) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return Row{}
}

func (tx *Tx) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return nil, nil
}

func (tx *Tx) Conn() *pgx.Conn {
	return nil
}


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

var OrmSignupClient = []struct {
	testName            string
	inputCookie         *utils.Defense
	inputSignUp         *utils.RegistrationRequest
	out                 *utils.Defense
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
		inputSignUp:         &utils.RegistrationRequest{},
		errQuerySalt:        nil,
		inputQueryClientId:  1,
	},
}

func TestOrmSignupClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmSignupClient {
		m.
			EXPECT().
			Begin(context.Background()).
			Return(&Tx{}, nil)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignupClient(tt.inputSignUp, tt.inputCookie)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
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
		test := Authorization{DB: m}
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
	out                  *utils.Defense
	outErr               string
	input                *utils.Authorization
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
		input:                &utils.Authorization{Email: "1", Phone: "", Password: "1"},
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
		input:                &utils.Authorization{Email: "", Phone: "1", Password: "1"},
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
		input:                &utils.Authorization{Email: "", Phone: "1", Password: "1"},
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
		test := Authorization{DB: m}
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
		test := Authorization{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.Logout(tt.input)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
