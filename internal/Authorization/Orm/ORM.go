package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Application"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyErrors"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) NewDefense() *Utils.Defense {
	var tmp Utils.Defense
	return tmp.GenerateNew()
}

func (db *Wrapper) generalSignUp(signup *Authorization.RegistrationRequest, transaction pgx.Tx) (int, error) {
	var userId int

	salt := Utils.RandString(Application.LenSalt)

	if _, err := strconv.Atoi(signup.Phone); err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.AGeneralSignUpIncorrectPhoneFormat,
		}
	}

	err := transaction.QueryRow(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		Utils.Sanitize(signup.Name), Utils.Sanitize(signup.Email),
		signup.Phone, Utils.HashPassword(signup.Password, salt), salt).Scan(&userId)

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

func (db *Wrapper) SignupHost(signup *Authorization.RegistrationRequest, cookie *Utils.Defense) (*Utils.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignupHostTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	userId, err := db.generalSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	err = db.addTransactionCookie(cookie, tx, userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO host (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpHostHostNotInsert,
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpHostNotCommit,
		}
	}
	return cookie, nil
}

func (db *Wrapper) SignupCourier(signup *Authorization.RegistrationRequest, cookie *Utils.Defense) (*Utils.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignupCourierTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	userId, err := db.generalSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	err = db.addTransactionCookie(cookie, tx, userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO courier (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpCourierCourierNotInsert,
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpCourierNotCommit,
		}
	}

	return cookie, err
}

func (db *Wrapper) SignupClient(signup *Authorization.RegistrationRequest, cookie *Utils.Defense) (*Utils.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignupClientTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	userId, err := db.generalSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	err = db.addTransactionCookie(cookie, tx, userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO client (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpClientClientNotInsert,
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpClientNotCommit,
		}
	}

	return cookie, nil
}

func (db *Wrapper) addTransactionCookie(cookie *Utils.Defense, Transaction pgx.Tx, id int) error {
	_, err := Transaction.Exec(context.Background(),
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
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByEmailTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var userId int
	var salt string

	err = tx.QueryRow(context.Background(),
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

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, Utils.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginOrPasswordIncorrect,
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByEmailNotCommit,
		}
	}

	return userId, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByPhoneTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var userId int
	var salt string

	err = tx.QueryRow(context.Background(),
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

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, Utils.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginOrPasswordIncorrect,
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.ALoginByPhoneNotCommit,
		}
	}

	return userId, nil
}

func (db *Wrapper) DeleteCookie(CSRF string) (string, error) {
	var sessionId string
	var sessionIdScan interface{}
	err := db.Conn.QueryRow(context.Background(),
		"DELETE FROM cookie WHERE csrf_token = $1 RETURNING session_id", CSRF).Scan(&sessionIdScan)
	if sessionIdScan != nil {
		sessionId = sessionIdScan.(string)
	}
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.ADeleteCookieCookieNotDelete,
		}
	}
	return sessionId, nil
}

func (db *Wrapper) AddCookie(cookie *Utils.Defense, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.AAddCookieCookieNotInsert,
		}
	}
	return nil
}
