package Orm

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

const (
	PhoneLen = 11
)

type Wrapper struct {
	Conn Interface.ConnectionInterface
}

func (db *Wrapper) NewDefense() *Util.Defense {
	var tmp Util.Defense
	return tmp.GenerateNew()
}

func (db *Wrapper) generalSignUp(signup *auth.RegistrationRequest, transaction Interface.TransactionInterface, contextTransaction context.Context) (int, error) {
	var userId int

	salt := Util.RandString(Application.LenSalt)

	Util.Sanitize(signup.Phone)
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
		Util.Sanitize(signup.Name), Util.Sanitize(signup.Email),
		phone, Util.HashPassword(signup.Password, salt), salt).Scan(&userId)

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

func (db *Wrapper) SignupHost(signup *auth.RegistrationRequest, cookie *Util.Defense) (*Util.Defense, error) {
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

func (db *Wrapper) SignupCourier(signup *auth.RegistrationRequest, cookie *Util.Defense) (*Util.Defense, error) {
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

func (db *Wrapper) SignupClient(signup *auth.RegistrationRequest, cookie *Util.Defense) (*Util.Defense, error) {
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

func (db *Wrapper) addTransactionCookie(cookie *Util.Defense, Transaction Interface.TransactionInterface, id int, contextTransaction context.Context) error {
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
		email, Util.HashPassword(password, salt)).Scan(&userId)
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
		phone, Util.HashPassword(password, salt)).Scan(&userId)
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

func (db *Wrapper) AddCookie(cookie *Util.Defense, id int) error {
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

func (db *Wrapper) CheckAccess(cookie *Util.Defense) (bool, error) {
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

func (db *Wrapper) NewCSRF(cookie *Util.Defense) (string, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.MNewCSRFCSRFTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	csrfToken := Util.RandString(5)
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

func (db *Wrapper) GetIdByCookie(cookie *Util.Defense) (int, error) {
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
