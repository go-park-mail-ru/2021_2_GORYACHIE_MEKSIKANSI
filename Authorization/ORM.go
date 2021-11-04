package Authorization

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"strconv"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) GenerateNew() *utils.Defense {
	var tmp utils.Defense
	return tmp.GenerateNew()
}

func GeneralSignUp(signup *utils.RegistrationRequest, transaction pgx.Tx) (int, error) {
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
		signup.Name, signup.Email, signup.Phone, utils.HashPassword(signup.Password, salt), salt).Scan(&userId)

	if err != nil {
		errorText := err.Error()
		if errorText == "ERROR: duplicate key value violates unique constraint \"general_user_info_phone_key\" (SQLSTATE 23505)" ||
			errorText == "0ERROR: duplicate key value violates unique constraint \"general_user_info_email_key\" (SQLSTATE 2355)" {
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
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignupHostTransactionNotCreate,
			Time: time.Now(),
		}
	}

	userId, err := GeneralSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	err = AddTransactionCookie(cookie, tx, userId)
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

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignupCourierTransactionNotCreate,
			Time: time.Now(),
		}
	}

	userId, err := GeneralSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	err = AddTransactionCookie(cookie, tx, userId)
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

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignupClientTransactionNotCreate,
			Time: time.Now(),
		}
	}

	userId, err := GeneralSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	err = AddTransactionCookie(cookie, tx, userId)
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

func AddTransactionCookie(cookie *utils.Defense, Transaction pgx.Tx, id int) error {
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
	var userId int
	var salt string

	err := db.Conn.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE email = $1",
		email).Scan(&salt)
	if err != nil {
		if err.Error() == "no rows in result set" {
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

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, utils.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginOrPasswordIncorrect,
			Time: time.Now(),
		}
	}

	return userId, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	var userId int
	var salt string

	err := db.Conn.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE phone = $1",
		phone).Scan(&salt)
	if err != nil {
		if err.Error() == "no rows in result set" {
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

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, utils.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ALoginOrPasswordIncorrect,
			Time: time.Now(),
		}
	}
	return userId, nil
}

func (db *Wrapper) DeleteCookie(cookie *utils.Defense) error {
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ADeleteCookieCookieNotDelete,
			Time: time.Now(),
		}
	}
	return nil
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
