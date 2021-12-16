//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/orm WrapperAuthorizationInterface,ConnectionInterface,TransactionInterface
package orm

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/myerror"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

type WrapperAuthorizationInterface interface {
	SignupClient(signup *authPkg.RegistrationRequest, cookie *authPkg.Defense) (*authPkg.Defense, error)
	SignupCourier(signup *authPkg.RegistrationRequest, cookie *authPkg.Defense) (*authPkg.Defense, error)
	SignupHost(signup *authPkg.RegistrationRequest, cookie *authPkg.Defense) (*authPkg.Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(CSRF string) (string, error)
	NewDefense() *authPkg.Defense
	AddCookie(cookie *authPkg.Defense, id int) error
	CheckAccess(cookie *authPkg.Defense) (bool, error)
	NewCSRF(cookie *authPkg.Defense) (string, error)
	GetIdByCookie(cookie *authPkg.Defense) (int, error)
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

func (db *Wrapper) NewDefense() *authPkg.Defense {
	var tmp authPkg.Defense
	return tmp.GenerateNew()
}

func (db *Wrapper) generalSignUp(signup *authPkg.RegistrationRequest, transaction TransactionInterface, contextTransaction context.Context) (int, error) {
	var userId int

	salt := authPkg.RandString(authPkg.LenSalt)

	Sanitize(signup.Phone)
	if _, err := strconv.Atoi(signup.Phone); err != nil || len(signup.Phone) != PhoneLen {
		return 0, &errPkg.Errors{
			Text: errPkg.AGeneralSignUpIncorrectPhoneFormat,
		}
	}

	s := []rune(signup.Phone)
	s[0] = '8'
	phone := string(s)

	err := transaction.QueryRow(contextTransaction,
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		Sanitize(signup.Name), Sanitize(signup.Email),
		phone, HashPassword(signup.Password, salt), salt).Scan(&userId)

	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "duplicate key") {
			return 0, &errPkg.Errors{
				Text: errPkg.AGeneralSignUpLoginNotUnique,
			}
		}
		return 0, &errPkg.Errors{
			Text: errPkg.AGeneralSignUpNotInsert,
		}
	}
	return userId, nil
}

func (db *Wrapper) SignupHost(signup *authPkg.RegistrationRequest, cookie *authPkg.Defense) (*authPkg.Defense, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.ASignupHostTransactionNotCreate,
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
			Text: errPkg.ASignUpHostHostNotInsert,
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.ASignUpHostNotCommit,
		}
	}
	return cookie, nil
}

func (db *Wrapper) SignupCourier(signup *authPkg.RegistrationRequest, cookie *authPkg.Defense) (*authPkg.Defense, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.ASignupCourierTransactionNotCreate,
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
			Text: errPkg.ASignUpCourierCourierNotInsert,
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.ASignUpCourierNotCommit,
		}
	}

	return cookie, err
}

func (db *Wrapper) SignupClient(signup *authPkg.RegistrationRequest, cookie *authPkg.Defense) (*authPkg.Defense, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.ASignupClientTransactionNotCreate,
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
			Text: errPkg.ASignUpClientClientNotInsert,
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.ASignUpClientNotCommit,
		}
	}

	return cookie, nil
}

func (db *Wrapper) addTransactionCookie(cookie *authPkg.Defense, Transaction TransactionInterface, id int, contextTransaction context.Context) error {
	_, err := Transaction.Exec(contextTransaction,
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.AAddTransactionCookieNotInsert,
		}
	}

	return nil
}

func (db *Wrapper) LoginByEmail(email string, password string) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.ALoginByEmailTransactionNotCreate,
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
				Text: errPkg.ALoginNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Text: errPkg.ASaltNotSelect,
		}
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.ALoginOrPasswordIncorrect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.ALoginByEmailNotCommit,
		}
	}

	return userId, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.ALoginByPhoneTransactionNotCreate,
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
				Text: errPkg.ALoginNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Text: errPkg.ASaltNotSelect,
		}
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.ALoginOrPasswordIncorrect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.ALoginByPhoneNotCommit,
		}
	}

	return userId, nil
}

func (db *Wrapper) DeleteCookie(CSRF string) (string, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.ADeleteCookieTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var sessionId *string
	err = tx.QueryRow(contextTransaction,
		"DELETE FROM cookie WHERE csrf_token = $1 RETURNING session_id", CSRF).Scan(&sessionId)

	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.ADeleteCookieCookieNotDelete,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.ADeleteCookieNotCommit,
		}
	}
	return *sessionId, nil
}

func (db *Wrapper) AddCookie(cookie *authPkg.Defense, id int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.AAddCookieTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.AAddCookieCookieNotInsert,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.AAddCookieNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) CheckAccess(cookie *authPkg.Defense) (bool, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.MCheckAccessTransactionNotCreate,
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
				Text: errPkg.MCheckAccessCookieNotFound,
			}
		}
		return false, &errPkg.Errors{
			Text: errPkg.MCheckAccessCookieNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.MCheckAccessNotCommit,
		}
	}

	if time.Now().Before(timeLiveCookie) {
		return true, nil
	}

	return false, nil
}

func (db *Wrapper) NewCSRF(cookie *authPkg.Defense) (string, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.MNewCSRFCSRFTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	csrfToken := authPkg.RandString(5)
	_, err = tx.Exec(contextTransaction,
		"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
		csrfToken, cookie.SessionId)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.MNewCSRFCSRFNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.MNewCSRFCSRFNotCommit,
		}
	}

	return csrfToken, nil
}

func (db *Wrapper) GetIdByCookie(cookie *authPkg.Defense) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.MGetIdByCookieTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var timeLiveCookie time.Time
	var id int
	err = tx.QueryRow(contextTransaction,
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
		cookie.SessionId).Scan(&id, &timeLiveCookie)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Text: errPkg.MGetIdByCookieCookieNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Text: errPkg.MGetIdByCookieCookieNotScan,
		}
	}

	realTime := time.Now()

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.MGetIdByCookieNotCommit,
		}
	}

	if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, &errPkg.Errors{
		Text: errPkg.MGetIdByCookieCookieExpired,
	}
}
