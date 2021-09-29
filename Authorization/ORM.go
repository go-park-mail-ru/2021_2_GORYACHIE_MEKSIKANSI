package Authorization

import (
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ERRPHONEQUERY = "ERROR: phone error query"
	ERRPHONESCAN  = "ERROR: phone not scan"
	ERRPHONEPASSQUERY = "ERROR: password and phone error query"
	ERRPHONEPASSSCAN  = "ERROR: phone and password not scan"
	ERRMAILQUERY = "ERROR: email error query"
	ERRMAILSCAN = "ERROR: email not scan"
	ERRMAILPASSQUERY = "ERROR: password or email error query"
	ERRMAILPASSSCAN  = "ERROR: email and pass not scan"
	ERRNOTLOGINORPASSWORD = "Неправильный логин или пароль"
	ERRINFOQUERY = "ERROR: not insert info query"
	ERRINFOSCAN = "ERROR: info not scan"
	ERRINSERTHOSTQUERY = "ERROR: not insert host query"
	ERRINSERTCOURIERQUERY = "ERROR: not insert courier query"
	ERRINSERTCLIENTQUERY = "ERROR: not insert client query"
	ERRINSERTCOOKIEQUERY = "ERROR: not insert cookie transact query"
	ERRDELETEQUERY = "ERROR: cookie not delete query"
	ERRINSERTLOGINCOOKIEQUERY = "ERROR: not insert cookie query"
	ERRUNIQUE = "Телефон или email уже зарегистрирован"
)

type Wrapper struct {
	Conn        *pgxpool.Pool
	Transaction pgx.Tx
}

func (db *Wrapper) GeneralSignUp(signup Registration) (int, error) {
	var userId int
	var err error

	salt := randString(LENSALT)

	row := db.Transaction.QueryRow(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		signup.Name, signup.Email, signup.Phone, mid.HashPassword(signup.Password, salt), salt)
	if err != nil {
		return 0, errors.New(ERRINFOQUERY)
	}

		err = row.Scan(&userId)
		if err != nil {
			errorText := err.Error()
            if errorText == "ERROR: duplicate key value violates unique constraint \"general_user_info_phone_key\" (SQLSTATE 23505)" {
            	return 0, errors.New(ERRUNIQUE)
			}
            return 0, errors.New(ERRINFOSCAN)
		}

	return userId, nil
}

func (db *Wrapper) SignupHost(signup Registration) (mid.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())
	db.Transaction = tx

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())

	userId, err := db.GeneralSignUp(signup)
	if err != nil {
		return mid.Defense{}, err
	}

	var cookie mid.Defense
	cookie = cookie.GenerateNew()
	err = db.AddTransactionCookie(cookie, userId)
	if err != nil {
		return cookie, err
	}

	err = tx.Commit(context.Background())

	_, err = db.Conn.Exec(context.Background(),
		"INSERT INTO host (client_id) VALUES ($1)", userId)
	if err != nil {
		_, err = db.Conn.Exec(context.Background(),
			"DELETE FROM host WHERE client_id = $1", userId)
		return cookie, errors.New(ERRINSERTHOSTQUERY)
	}
	return cookie, nil
}

func (db *Wrapper) SignupCourier(signup Registration) (mid.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())
	db.Transaction = tx

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())

	userId, err := db.GeneralSignUp(signup)
	if err != nil {
		return mid.Defense{}, err
	}

	var cookie mid.Defense
	cookie = cookie.GenerateNew()
	err = db.AddTransactionCookie(cookie, userId)
	if err != nil {
		return cookie, err
	}

	err = tx.Commit(context.Background())

	_, err = db.Conn.Exec(context.Background(),
		"INSERT INTO courier (client_id) VALUES ($1)", userId)
	if err != nil {
		_, err = db.Conn.Exec(context.Background(),
			"DELETE FROM courier WHERE client_id = $1", userId)
		return cookie, errors.New(ERRINSERTCOURIERQUERY)
	}
	return cookie, err
}

func (db *Wrapper) SignupClient(signup Registration) (mid.Defense, error) {
	tx, err := db.Conn.Begin(context.Background())
	db.Transaction = tx

	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, context.Background())

	userId, err := db.GeneralSignUp(signup)
	if err != nil {
		return mid.Defense{}, err
	}

	var cookie mid.Defense
	cookie = cookie.GenerateNew()
	err = db.AddTransactionCookie(cookie, userId)
	if err != nil {
		return cookie, err
	}

	err = tx.Commit(context.Background())
	_, err = db.Conn.Exec(context.Background(),
		"INSERT INTO client (client_id, date_birthday) VALUES ($1, $2)", userId, signup.Birthday)
	if err != nil {
		_, err = db.Conn.Exec(context.Background(),
			"DELETE FROM general_user_info WHERE client_id = $1", userId)
		return cookie, errors.New(ERRINSERTCLIENTQUERY)
	}

	return cookie, nil
}

func (db *Wrapper) AddTransactionCookie(cookie mid.Defense, id int) error {
	_, err := db.Transaction.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return errors.New(ERRINSERTCOOKIEQUERY)
	}

	return nil
}

func (db *Wrapper) LoginByEmail(email string, password string) (int, error) {
	var userId int
	var salt string

	row, err := db.Conn.Query(context.Background(),
		"SELECT salt FROM general_user_info WHERE email = $1",
		email)
	if err != nil {
		return 0, errors.New(ERRMAILQUERY)
	}

	for row.Next() {
		err = row.Scan(&salt)

		if err != nil {
			return 0, errors.New(ERRMAILSCAN)
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, mid.HashPassword(password, salt))
	if err != nil {
		return 0, errors.New(ERRMAILPASSQUERY)
	}

	for row.Next() {
		err = row.Scan(&userId)
		if err != nil {
			return 0, errors.New(ERRMAILPASSSCAN)
		}
	}

	if userId == 0 {
		return 0, errors.New(ERRNOTLOGINORPASSWORD)
	}
	return userId, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	var user int
	var salt string

	row, err := db.Conn.Query(context.Background(),
		"SELECT salt FROM general_user_info WHERE phone = $1",
		phone)
	if err != nil {
		return 0, errors.New(ERRPHONEQUERY)
	}

	for row.Next() {
		err = row.Scan(&salt)
		if err != nil {
			return 0, errors.New(ERRPHONESCAN)
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, mid.HashPassword(password, salt))
	if err != nil {
		return 0, errors.New(ERRPHONEPASSQUERY)
	}

	for row.Next() {
		err = row.Scan(&user)
		if err != nil {
			return 0, errors.New(ERRPHONEPASSSCAN)
		}
	}

	if user == 0 {
		return 0, errors.New(ERRNOTLOGINORPASSWORD)
	}
	return user, nil
}

func (db *Wrapper) DeleteCookie(cookie mid.Defense) error {
	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken)
	if err != nil {
		return errors.New(ERRDELETEQUERY)
	}

	return nil
}

func (db *Wrapper) AddCookie(cookie mid.Defense, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return errors.New(ERRINSERTLOGINCOOKIEQUERY)
	}

	return nil
}
