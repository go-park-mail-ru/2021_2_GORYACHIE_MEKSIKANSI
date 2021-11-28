package orm

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/application"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

const (
	PhoneLen = 11
)

type WrapperAuthorization interface {
	SignupClient(signup *auth.RegistrationRequest, cookie *util.Defense) (*util.Defense, error)
	SignupCourier(signup *auth.RegistrationRequest, cookie *util.Defense) (*util.Defense, error)
	SignupHost(signup *auth.RegistrationRequest, cookie *util.Defense) (*util.Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(CSRF string) (string, error)
	NewDefense() *util.Defense
	AddCookie(cookie *util.Defense, id int) error
	CheckAccess(cookie *util.Defense) (bool, error)
	NewCSRF(cookie *util.Defense) (string, error)
	GetIdByCookie(cookie *util.Defense) (int, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
}

type Wrapper struct {
	Conn ConnectionInterface
}

func (db *Wrapper) NewDefense() *util.Defense {
	var tmp util.Defense
	return tmp.GenerateNew()
}

func (db *Wrapper) generalSignUp(signup *auth.RegistrationRequest, transaction TransactionInterface, contextTransaction context.Context) (int, error) {
	var userId int

	salt := util.RandString(application.LenSalt)

	util.Sanitize(signup.Phone)
	if _, err := strconv.Atoi(signup.Phone); err != nil || len(signup.Phone) != PhoneLen {
		return 0, &errPkg.Errors{
			Alias: errPkg.AGeneralSignUpIncorrectPhoneFormat,
		}
	}

	s := []rune(signup.Phone)
	s[0] = '8'
	phone := string(s)

	err := transaction.QueryRow(contextTransaction,
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		util.Sanitize(signup.Name), util.Sanitize(signup.Email),
		phone, util.HashPassword(signup.Password, salt), salt).Scan(&userId)

	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "duplicate key") {
			return 0, &errPkg.Errors{
				Alias: errPkg.AGeneralSignUpLoginNotUnique,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.AGeneralSignUpNotInsert,
		}
	}
	return userId, nil
}

func (db *Wrapper) SignupHost(signup *auth.RegistrationRequest, cookie *util.Defense) (*util.Defense, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignupHostTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	userId, err := db.generalSignUp(signup, tx, contextTransaction)
	if err != nil {
		return nil, err
	}

	err = db.addTransactionCookie(cookie, tx, userId, contextTransaction)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO host (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpHostHostNotInsert,
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpHostNotCommit,
		}
	}
	return cookie, nil
}

func (db *Wrapper) SignupCourier(signup *auth.RegistrationRequest, cookie *util.Defense) (*util.Defense, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignupCourierTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	userId, err := db.generalSignUp(signup, tx, contextTransaction)
	if err != nil {
		return nil, err
	}

	err = db.addTransactionCookie(cookie, tx, userId, contextTransaction)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO courier (client_id) VALUES ($1)", userId, contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpCourierCourierNotInsert,
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpCourierNotCommit,
		}
	}

	return cookie, err
}

func (db *Wrapper) SignupClient(signup *auth.RegistrationRequest, cookie *util.Defense) (*util.Defense, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignupClientTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	userId, err := db.generalSignUp(signup, tx, contextTransaction)
	if err != nil {
		return nil, err
	}

	err = db.addTransactionCookie(cookie, tx, userId, contextTransaction)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO client (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpClientClientNotInsert,
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpClientNotCommit,
		}
	}

	return cookie, nil
}

func (db *Wrapper) addTransactionCookie(cookie *util.Defense, Transaction TransactionInterface, id int, contextTransaction context.Context) error {
	_, err := Transaction.Exec(contextTransaction,
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.AAddTransactionCookieNotInsert,
		}
	}

	return nil
}

func (db *Wrapper) LoginByEmail(email string, password string) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByEmailTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var userId int
	var salt string

	err = tx.QueryRow(contextTransaction,
		"SELECT salt FROM general_user_info WHERE email = $1",
		email).Scan(&salt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Alias: errPkg.ALoginNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.ASaltNotSelect,
		}
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, util.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginOrPasswordIncorrect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByEmailNotCommit,
		}
	}

	return userId, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByPhoneTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var userId int
	var salt string

	err = tx.QueryRow(contextTransaction,
		"SELECT salt FROM general_user_info WHERE phone = $1",
		phone).Scan(&salt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Alias: errPkg.ALoginNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.ASaltNotSelect,
		}
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, util.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginOrPasswordIncorrect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByPhoneNotCommit,
		}
	}

	return userId, nil
}

func (db *Wrapper) DeleteCookie(CSRF string) (string, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.ADeleteCookieTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var sessionId *string
	err = tx.QueryRow(contextTransaction,
		"DELETE FROM cookie WHERE csrf_token = $1 RETURNING session_id", CSRF).Scan(&sessionId)

	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.ADeleteCookieCookieNotDelete,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.ADeleteCookieNotCommit,
		}
	}
	return *sessionId, nil
}

func (db *Wrapper) AddCookie(cookie *util.Defense, id int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.AAddCookieTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.AAddCookieCookieNotInsert,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.AAddCookieNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) CheckAccess(cookie *util.Defense) (bool, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.MCheckAccessTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var timeLiveCookie time.Time
	var id int
	err = tx.QueryRow(contextTransaction,
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken).Scan(&id, &timeLiveCookie)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, &errPkg.Errors{
				Alias: errPkg.MCheckAccessCookieNotFound,
			}
		}
		return false, &errPkg.Errors{
			Alias: errPkg.MCheckAccessCookieNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.MCheckAccessNotCommit,
		}
	}

	if time.Now().Before(timeLiveCookie) {
		return true, nil
	}

	return false, nil
}

func (db *Wrapper) NewCSRF(cookie *util.Defense) (string, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.MNewCSRFCSRFTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	csrfToken := util.RandString(5)
	_, err = tx.Exec(contextTransaction,
		"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
		csrfToken, cookie.SessionId)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.MNewCSRFCSRFNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.MNewCSRFCSRFNotCommit,
		}
	}

	return csrfToken, nil
}

func (db *Wrapper) GetIdByCookie(cookie *util.Defense) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.MGetIdByCookieTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var timeLiveCookie time.Time
	var id int
	err = db.Conn.QueryRow(contextTransaction,
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
		cookie.SessionId).Scan(&id, &timeLiveCookie)
	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "no rows") {
			return 0, &errPkg.Errors{
				Alias: errPkg.MGetIdByCookieCookieNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.MGetIdByCookieCookieNotScan,
		}
	}

	realTime := time.Now()

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.MGetIdByCookieNotCommit,
		}
	}

	if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, &errPkg.Errors{
		Alias: errPkg.MGetIdByCookieCookieExpired,
	}
}
