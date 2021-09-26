package Middleware

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	auth "project/Authorization"
)

func CreateDb(conn *pgxpool.Pool) error {
	sql := "CREATE TABLE IF NOT EXISTS restaurant (id serial PRIMARY KEY, owner INTEGER, FOREIGN KEY (owner) REFERENCES general_user_info (id) On DELETE CASCADE, name text NOT NULL, description text NOT NULL, created timestamp DEFAULT NOW() NOT NULL, deleted boolean DEFAULT false, avatar text DEFAULT '/uploads/', min_price int DEFAULT 0, price_delivery int NOT NULL, min_delivery_time timestamp, max_delivery_time timestamp, city text NOT NULL, street text NOT NULL, house text NOT NULL, floor int, rating int, location text); CREATE TABLE IF NOT EXISTS cookie (id serial PRIMARY KEY, client_id INTEGER, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE, session_id text NOT NULL, date_life timestamp NOT NULL, csrf_token varchar(64) NOT NULL); CREATE TABLE IF NOT EXISTS host (id serial PRIMARY KEY, client_id INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE); CREATE TABLE IF NOT EXISTS client (id serial PRIMARY KEY, client_id INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE, date_birthday timestamp NOT NULL); CREATE TABLE IF NOT EXISTS courier (id serial PRIMARY KEY, client_id  INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE);"
	_, err := conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	return nil
}

// TODO: merge check and get
func checkDefence(conn *pgxpool.Pool, cookie auth.Defense) (bool, error) {
	row, err := conn.Query(context.Background(),
		"SELECT ID FROM cookie WHERE session_id = $1 AND date_life = $2 AND csrf_token = $3",
		cookie.SessionId, cookie.DateLife, cookie.CsrfToken)
	if err != nil {
		return false, err
	}

	var id int
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return false, err
		}
	}
	if id == 0 {
		return false, nil
	}
	return true, nil
}

func CheckAccess(conn *pgxpool.Pool, cookie auth.Defense) (bool, error) {
	result, err := checkDefence(conn, cookie)
	if err != nil {
		return false, err
	}
	return result, nil
}

func GetIdByCookie(conn *pgxpool.Pool, cookie auth.Defense) (int, error) {
	row, err := conn.Query(context.Background(),
		"SELECT client_id FROM cookie WHERE session_id = $1 AND date_life = $2", cookie.SessionId, cookie.DateLife)
	if err != nil {
		return 0, err
	}

	var id int
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}
