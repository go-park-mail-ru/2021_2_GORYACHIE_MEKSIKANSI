package Middleware

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
}

func (db *Wrapper) CheckAccess(cookie *utils.Defense) (bool, error) {
	var timeLiveCookie time.Time
	var id int
	err := db.Conn.QueryRow(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken).Scan(&id, &timeLiveCookie)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, &errorsConst.Errors{
				Text: errorsConst.MCheckAccessCookieNotFound,
				Time: time.Now(),
			}
		}
		return false, &errorsConst.Errors{
			Text: errorsConst.MCheckAccessCookieNotScan,
			Time: time.Now(),
		}
	}

	if time.Now().Before(timeLiveCookie) {
		return true, nil
	}

	return false, nil
}

func (db *Wrapper) NewCSRF(cookie *utils.Defense) (string, error) {
	csrfToken := utils.RandString(5)
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
		csrfToken, cookie.SessionId)
	if err != nil {
		return "", &errorsConst.Errors{
			Text: errorsConst.MNewCSRFCSRFNotUpdate,
			Time: time.Now(),
		}
	}

	return csrfToken, nil
}

func (db *Wrapper) GetIdByCookie(cookie *utils.Defense) (int, error) {
	var timeLiveCookie time.Time
	var id int
	err := db.Conn.QueryRow(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
		cookie.SessionId).Scan(&id, &timeLiveCookie)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, &errorsConst.Errors{
				Text: errorsConst.MGetIdByCookieCookieNotFound,
				Time: time.Now(),
			}
		}
		return 0, &errorsConst.Errors{
			Text: errorsConst.MGetIdByCookieCookieNotScan,
			Time: time.Now(),
		}
	}

	realTime := time.Now()

	if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, &errorsConst.Errors{
		Text: errorsConst.MGetIdByCookieCookieExpired,
		Time: time.Now(),
	}
}
