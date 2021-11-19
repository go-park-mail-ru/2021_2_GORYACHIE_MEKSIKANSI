package Middleware

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/jackc/pgx/v4"
	"strings"
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
		if err == pgx.ErrNoRows {
			return false, &errPkg.Errors{
				Alias: errPkg.MCheckAccessCookieNotFound,
			}
		}
		return false, &errPkg.Errors{
			Alias: errPkg.MCheckAccessCookieNotScan,
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
		return "", &errPkg.Errors{
			Alias: errPkg.MNewCSRFCSRFNotUpdate,
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
		errorText := err.Error()
		if strings.Contains(errorText, "no rows") {
			return 0, &errPkg.Errors{
				Alias: errPkg.MGetIdByCookieCookieNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Alias: errPkg.MGetIdByCookieCookieNotScan,
		}
	}

	realTime := time.Now()

	if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, &errPkg.Errors{
		Alias: errPkg.MGetIdByCookieCookieExpired,
	}
}
