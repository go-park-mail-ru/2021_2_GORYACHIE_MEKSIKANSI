package Authorization

import (
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Wrapper struct {
	Conn        *pgxpool.Pool
	Transaction pgx.Tx
}

func hashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt + password))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func (db *Wrapper) GeneralSignUp(signup Registration) (int, error) {
	var userId int
	var err error

	salt := randString(LENSALT)

	row, err := db.Transaction.Query(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		signup.Name, signup.Email, signup.Phone, hashPassword(signup.Password, salt), salt)
	if err != nil {
		panic(err)
		return 0, err
	}

	for row.Next() {
		err := row.Scan(&userId)
		if err != nil {
			panic(err)
            return 0, err
		}
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
		return cookie, err
	}

	return cookie, err
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
		return cookie, err
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
		return cookie, err
	}

	return cookie, nil
}

func (db *Wrapper) AddTransactionCookie(cookie mid.Defense, id int) error {
	_, err := db.Transaction.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func (db *Wrapper) LoginByEmail(email string, password string) (int, error) {
	var user int
	var salt string

	row, err := db.Conn.Query(context.Background(),
		"SELECT salt FROM general_user_info WHERE email = $1",
		email)
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&salt)

		if err != nil {
			return 0, err
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND password = $2",
		email, hashPassword(password, salt))
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&user)
		if err != nil {
			return 0, err
		}
	}

	if user == 0 {
		return 0, nil
	}
	return user, nil
}

func (db *Wrapper) LoginByPhone(phone string, password string) (int, error) {
	var user int
	var salt string

	row, err := db.Conn.Query(context.Background(),
		"SELECT salt FROM general_user_info WHERE phone = $1",
		phone)
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&salt)
		if err != nil {
			return 0, err
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, hashPassword(password, salt))
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&user)
		if err != nil {
			return 0, err
		}
	}

	if user == 0 {
		return 0, nil
	}
	return user, nil
}

func (db *Wrapper) DeleteCookie(cookie mid.Defense) error {
	_, err := db.Conn.Query(context.Background(),
		"DELETE FROM cookie WHERE session_id = $1 AND date_life = $2",
		cookie.SessionId, cookie.DateLife)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) AddCookie(cookie mid.Defense, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES ($1, $2, $3, $4)",
		id, cookie.SessionId, cookie.DateLife, "s")
	if err != nil {
		return err
	}

	return nil
}
