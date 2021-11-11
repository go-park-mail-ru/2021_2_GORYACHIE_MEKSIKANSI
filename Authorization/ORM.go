package Authorization

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) NewDefense() *utils.Defense {
	var tmp utils.Defense
	return tmp.GenerateNew()
}

func (db *Wrapper) generalSignUp(signup *utils.RegistrationRequest, transaction pgx.Tx) (int, error) {
	var userId int

	salt := utils.RandString(LenSalt)

	if _, err := strconv.Atoi(signup.Phone); err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.AGeneralSignUpIncorrectPhoneFormat,
			Time: time.Now(),
		}
	}

	err := transaction.QueryRow(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		utils.Sanitize(signup.Name), utils.Sanitize(signup.Email),
		signup.Phone, utils.HashPassword(signup.Password, salt), salt).Scan(&userId)

	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "duplicate key") {
			return 0, &errorsConst.Errors{
				Text: errorsConst.AGeneralSignUpLoginNotUnique,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.AGeneralSignUpNotInsert,
			Time: time.Now(),
		}
	}
	return userId, nil
}

func (db *Wrapper) SignupHost(signup *utils.RegistrationRequest, cookie *utils.Defense) (*utils.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignupHostTransactionNotCreate,
			Time: time.Now(),
		}
	}

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
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpHostHostNotInsert,
			Time: time.Now(),
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpHostNotCommit,
			Time: time.Now(),
		}
	}
	return cookie, nil
}

func (db *Wrapper) SignupCourier(signup *utils.RegistrationRequest, cookie *utils.Defense) (*utils.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignupCourierTransactionNotCreate,
			Time: time.Now(),
		}
	}

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
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpCourierCourierNotInsert,
			Time: time.Now(),
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpCourierNotCommit,
			Time: time.Now(),
		}
	}

	return cookie, err
}

func (db *Wrapper) SignupClient(signup *utils.RegistrationRequest, cookie *utils.Defense) (*utils.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignupClientTransactionNotCreate,
			Time: time.Now(),
		}
	}

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
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpClientClientNotInsert,
			Time: time.Now(),
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpClientNotCommit,
			Time: time.Now(),
		}
	}

	return cookie, nil
}

func (db *Wrapper) addTransactionCookie(cookie *utils.Defense, Transaction pgx.Tx, id int) error {
	_, err := Transaction.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.AAddTransactionCookieNotInsert,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) LoginByEmail(email string, password string) (int, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginByEmailTransactionNotCreate,
			Time: time.Now(),
		}
	}

	var userId int
	var salt string

	err = tx.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE email = $1",
		email).Scan(&salt)
	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "no rows") {
			return 0, &errorsConst.Errors{
				Text: errorsConst.ALoginNotFound,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.ASaltNotSelect,
			Time: time.Now(),
		}
	}

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, utils.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginOrPasswordIncorrect,
			Time: time.Now(),
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginByEmailNotCommit,
			Time: time.Now(),
		}
	}

	return userId, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginByPhoneTransactionNotCreate,
			Time: time.Now(),
		}
	}

	var userId int
	var salt string

	err = tx.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE phone = $1",
		phone).Scan(&salt)
	if err != nil {
		errorText := err.Error()
		if strings.Contains(errorText, "no rows") {
			return 0, &errorsConst.Errors{
				Text: errorsConst.ALoginNotFound,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.ASaltNotSelect,
			Time: time.Now(),
		}
	}

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, utils.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginOrPasswordIncorrect,
			Time: time.Now(),
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginByPhoneNotCommit,
			Time: time.Now(),
		}
	}

	return userId, nil
}

func (db *Wrapper) DeleteCookie(CSRF string) (string, error) {
	var sessionId string
	var sessionIdScan interface{}
	err := db.Conn.QueryRow(context.Background(),
		"DELETE FROM cookie WHERE csrf_token = $1 RETURNING session_id", CSRF).Scan(&sessionIdScan)
	sessionId = sessionIdScan.(string)
	if err != nil {
		return "", &errorsConst.Errors{
			Text: errorsConst.ADeleteCookieCookieNotDelete,
			Time: time.Now(),
		}
	}
	return sessionId, nil
}

func (db *Wrapper) AddCookie(cookie *utils.Defense, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.AAddCookieCookieNotInsert,
			Time: time.Now(),
		}
	}
	return nil
}
