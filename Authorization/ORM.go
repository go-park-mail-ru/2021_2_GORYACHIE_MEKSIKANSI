package Authorization

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Wrapper struct {
	Conn        *pgxpool.Pool
}

func (db *Wrapper) GeneralSignUp(signup *Registration, 	transaction pgx.Tx) (int, error) {
	var userId int

	salt := randString(LENSALT)

	err := transaction.QueryRow(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",

		signup.Name, signup.Email, signup.Phone, mid.HashPassword(signup.Password, salt), salt).Scan(&userId)
        
		if err != nil {
			errorText := err.Error()
            if errorText == "ERROR: duplicate key value violates unique constraint \"general_user_info_phone_key\" (SQLSTATE 23505)" ||
				errorText == "ERROR: duplicate key value violates unique constraint \"general_user_info_email_key\" (SQLSTATE 23505)" {
            	return 0, &errorsConst.Errors{
					Text: errorsConst.ERRUNIQUE,
					Time: time.Now(),
				}
			}
            return 0, &errorsConst.Errors {
				Text: errorsConst.ErrGeneralInfoUnique,
				Time: time.Now(),
			}

		}
		return 0, errors.New(errorsConst.ErrGeneralInfoScan)
	}

	return userId, nil
}

func (db *Wrapper) SignupHost(signup *Registration) (*mid.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())

	userId, err := db.GeneralSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	var temp mid.Defense
	cookie := temp.GenerateNew()
	err = db.AddTransactionCookie(cookie, tx, userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO host (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrInsertHost,
			Time: time.Now(),
		}
	}
	err = tx.Commit(context.Background())
	return cookie, nil
}

func (db *Wrapper) SignupCourier(signup *Registration) (*mid.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())

	userId, err := db.GeneralSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	var tmp mid.Defense
	cookie := tmp.GenerateNew()
	err = db.AddTransactionCookie(cookie, tx, userId)
	if err != nil {
		return nil, err
	}


	_, err = tx.Exec(context.Background(),
		"INSERT INTO courier (client_id) VALUES ($1)", userId)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrInsertCourier,
			Time: time.Now(),
		}
	}
	err = tx.Commit(context.Background())

	return cookie, err
}

func (db *Wrapper) SignupClient(signup *Registration) (*mid.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())

	userId, err := db.GeneralSignUp(signup, tx)
	if err != nil {
		return nil, err
	}

	var tmp mid.Defense
	cookie := tmp.GenerateNew()
	err = db.AddTransactionCookie(cookie, tx, userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO client (client_id, date_birthday) VALUES ($1, $2)", userId, signup.Birthday)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrInsertClient,
			Time: time.Now(),
		}
	}
	err = tx.Commit(context.Background())

	return cookie, nil
}

func (db *Wrapper) AddTransactionCookie(cookie *mid.Defense, Transaction pgx.Tx, id int) error {
	_, err := Transaction.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrInsertTransactionCookie,
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
		return 0, &errorsConst.Errors{
			Text: errorsConst.ErrSelectSalt,
			Time: time.Now(),
		}
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, mid.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ErrLoginOrPasswordIncorrect,
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
		return 0, &errorsConst.Errors{
			Text: errorsConst.ErrSelectSalt,
			Time: time.Now(),
		}
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, mid.HashPassword(password, salt)).Scan(&userId)
	if err != nil {
		return 0, &errorsConst.Errors{
			Text: errorsConst.ErrLoginOrPasswordIncorrect,
			Time: time.Now(),
		}
	}
	return userId, nil
}

func (db *Wrapper) DeleteCookie(cookie *mid.Defense) error {
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrDeleteCookie,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) AddCookie(cookie *mid.Defense, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrInsertCookie,
			Time: time.Now(),
		}
	}

	return nil
}
