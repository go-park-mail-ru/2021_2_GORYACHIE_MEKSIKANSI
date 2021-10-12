package Middleware

import (
	config "2021_2_GORYACHIE_MEKSIKANSI/Config"
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

func RandomInteger(min int, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max - min)))
	if err != nil {
		return 5
	}
	n := nBig.Int64()
	return int(n) + min
}

func HashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt + password))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}


func CreateDb() (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://" + config.DBLogin+ ":" + config.DBPassword+
		"@" + config.DBHost+ ":" + config.DBPort+ "/" + config.DBName)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrNotConnect,
			Time: time.Now(),
		}
	}

	if config.DEBUG {
		file, err := ioutil.ReadFile("PostgreSQL/DeleteTables.sql")
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrDeleteFileNotFound,
				Time: time.Now(),
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errorsConst.Errors{
					Text: errorsConst.ErrNotDeleteTables,
					Time: time.Now(),
				}
			}
		}
	}

	file, err := ioutil.ReadFile("PostgreSQL/CreateTables.sql")
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrFileNotFound,
			Time: time.Now(),
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrNotCreateTables,
				Time: time.Now(),
			}
		}
	}

	if config.DEBUG {
		file, err := ioutil.ReadFile("PostgreSQL/Fill.sql")
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrFillFileNotFound,
				Time: time.Now(),
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errorsConst.Errors{
					Text: errorsConst.ErrNotFillTables,
					Time: time.Now(),
				}
			}
		}
	}
	return conn, nil
}

func CheckAccess(conn *pgxpool.Pool, cookie *Defense) (bool, error) {
	var timeLiveCookie time.Time
	var id int
	err := conn.QueryRow(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken).Scan(&id, &timeLiveCookie)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, &errorsConst.Errors{
				Text: errorsConst.ErrCheckAccessCookieNotFound,
				Time: time.Now(),
			}
		}
		return false, &errorsConst.Errors{
			Text: errorsConst.ErrCookieNotScan,
			Time: time.Now(),
		}
	}

	if time.Now().Before(timeLiveCookie) {
		return true, nil
	}

	return false, nil
}

func NewCsrf(conn *pgxpool.Pool, cookie *Defense) (string, error) {
	csrfToken := randString(5)
	_, err := conn.Exec(context.Background(),
		"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
		csrfToken, cookie.SessionId)
	if err != nil {
		return "", &errorsConst.Errors{
			Text: errorsConst.ErrUpdateCSRF,
			Time: time.Now(),
		}
	}

	return csrfToken, nil
}

func GetIdByCookie(conn *pgxpool.Pool, cookie *Defense) (int, error) {
	var timeLiveCookie time.Time
	var id int
	err := conn.QueryRow(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
		cookie.SessionId).Scan(&id, &timeLiveCookie)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, &errorsConst.Errors{
				Text: errorsConst.ErrCookieNotFound,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.ErrCookieScan,
			Time: time.Now(),
		}
	}

	realTime := time.Now()
    if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, &errorsConst.Errors{
		Text: errorsConst.ErrCookieExpired,
		Time: time.Now(),
	}
}
