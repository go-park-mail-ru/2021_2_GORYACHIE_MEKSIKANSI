package Middleware

import (
	config "2021_2_GORYACHIE_MEKSIKANSI/Config"
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/big"
	"time"
)


func randomInteger(min int, max int) int {
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

func makeName() string {
	restNames := []string{
	"Cheesecake Factory", "Shokolaat", "Gordon Biersch", "Crepevine", "Creamery", "Old Pro", "House of Bagels",
	"The Prolific Oven", "La Strada", "Buca di Beppo", "Madame Tam", "Sprout Cafe", "Junoon", "Bistro Maxine",
	"Three Seasons", "Reposado", "Siam Royal", "Krung Siam", "Thaiphoon", "Tamarine", "Joya", "Jing Jing",
	"Evvia Estiatorio", "Cafe 220", "Cafe Renaissance", "Kan Zeman", "Mango Caribbean Cafe", "Baklava",
	"Mandarin Gourmet", "Bangkok Cuisine", "Darbar Indian Cuisine", "Mantra", "Janta", "Hyderabad House",
	"Starbucks", "Coupa Cafe", "Lytton Coffee Company", "Il Fornaio", "Lavanda", "MacArthur Park",
	"Osteria", "Vero", "Cafe Renzo", "Miyake", "Sushi Tomo", "Kanpai", "Pizza My Heart", "New York Pizza",
	"California Pizza Kitchen", "Round Table", "Loving Hut", "Garden Fresh", "Cafe Epi", "Tai Pan",
	}
	return restNames[randomInteger(0, len(restNames)- 1)]
}

func CreateDb() (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(), "postgres://" + config.DBLOGIN + ":" + config.DBPASSWORD + "@" + config.DBHOST + ":" + config.DBPORT + "/" + config.DBNAME)
	if err != nil {
		return nil, errors.New(errorsConst.ErrNotConnect)
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
			return false, errors.New(errorsConst.ErrCheckAccessCookieNotFound)
		}
		return false, errors.New(errorsConst.ErrCookieNotScan)
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
		return "", errors.New(errorsConst.ErrUpdateCSRF)
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
			return 0, errors.New(errorsConst.ErrCookieNotFound)
		}
		return 0, errors.New(errorsConst.ErrCookieScan)
	}

	realTime := time.Now()
    if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, errors.New(errorsConst.ErrCookieExpired)
}
