package Middleware

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

func CheckAccess(conn *pgxpool.Pool, cookie *utils.Defense) (bool, error) {
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

func NewCsrf(conn *pgxpool.Pool, cookie *utils.Defense) (string, error) {
	csrfToken := utils.RandString(5)
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

func GetIdByCookie(conn *pgxpool.Pool, cookie *utils.Defense) (int, error) {
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
