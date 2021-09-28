package Middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/rand"
	"time"
)

const DEBUG = true

const (
	DBLOGIN string = "Captain-matroskin"
	DBPASSWORD string = "74tbr6r54f78"
	DBNAME string = "hot_mexican_db"
	DBHOST = "localhost"
	DBPORT = "5432"
)

const (
	ERRCREATEQUERY = "ERROR: db not created"
	ERRINSERTQUERY = "ERROR: restaurant not insert"
	ERRCOOKIEANDCSRFQUERY = "ERROR: cookie and csrf query"
	ERRCOOKIEANDCSRFSCAN = "ERROR: cookie and csrf scan"
	ERRCOOKIEQUERY = "ERROR: cookie query"
	ERRCOOKIESCAN = "ERROR: cookie scan"
	ERRSIDNOTFOUND = "ERROR: id not found"
	ERRNOTCONNECT      = "ERROR: not connect"
	ERRUPDATECSRFQUERY = "ERROR: csrf not updated"
	ERRCOOKIEEXPIRED   = "ERROR: cookie expired"
	ERRDELETEQUERY = "ERROR: not delete query"
	ERRINSERTROOTQUERY = "ERROR: not create root"
)

func randomInteger(min int, max int) int {
	return rand.Intn(max - min) + min
}

func HashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt + password))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func makeName() string {
	restNames := []string{
	"Cheesecake Factory",
	"Shokolaat",
	"Gordon Biersch",
	"Crepevine",
	"Creamery",
	"Old Pro",
	"House of Bagels",
	"The Prolific Oven",
	"La Strada",
	"Buca di Beppo",
	"Madame Tam",
	"Sprout Cafe",
	"Junoon",
	"Bistro Maxine",
	"Three Seasons",
	"Reposado",
	"Siam Royal",
	"Krung Siam",
	"Thaiphoon",
	"Tamarine",
	"Joya",
	"Jing Jing",
	"Evvia Estiatorio",
	"Cafe 220",
	"Cafe Renaissance",
	"Kan Zeman",
	"Mango Caribbean Cafe",
	"Baklava",
	"Mandarin Gourmet",
	"Bangkok Cuisine",
	"Darbar Indian Cuisine",
	"Mantra",
	"Janta",
	"Hyderabad House",
	"Starbucks",
	"Coupa Cafe",
	"Lytton Coffee Company",
	"Il Fornaio",
	"Lavanda",
	"MacArthur Park",
	"Osteria",
	"Vero",
	"Cafe Renzo",
	"Miyake",
	"Sushi Tomo",
	"Kanpai",
	"Pizza My Heart",
	"New York Pizza",
	"California Pizza Kitchen",
	"Round Table",
	"Loving Hut",
	"Garden Fresh",
	"Cafe Epi",
	"Tai Pan",
	}
	return restNames[randomInteger(0, len(restNames)- 1)]
}

func CreateDb() (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(), "postgres://" + DBLOGIN + ":" + DBPASSWORD + "@" + DBHOST + ":" + DBPORT + "/" + DBNAME)
	if err != nil {
		return nil, errors.New(ERRNOTCONNECT)
	}
	if DEBUG {
		_, err = conn.Exec(context.Background(), "DROP TABLE IF EXISTS restaurant, general_user_info, host, client, courier CASCADE")
		if err != nil {
			return nil, errors.New(ERRDELETEQUERY)
		}
	}

	tableGeneralUserInfo:= "CREATE TABLE IF NOT EXISTS general_user_info (id SERIAL PRIMARY KEY, name text NOT NULL, password varchar(64) NOT NULL, salt varchar(5) NOT NULL, phone varchar(15) UNIQUE NOT NULL, email text UNIQUE, avatar text DEFAULT '/uploads/', date_registration timestamp DEFAULT NOW() NOT NULL, deleted boolean DEFAULT false);"
	tableRestaurant := "CREATE TABLE IF NOT EXISTS restaurant (id serial PRIMARY KEY, owner INTEGER, FOREIGN KEY (owner) REFERENCES general_user_info (id) On DELETE CASCADE, name text NOT NULL, description text NOT NULL, created timestamp DEFAULT NOW() NOT NULL, deleted boolean DEFAULT false, avatar text DEFAULT '/uploads/', min_price int DEFAULT 0, price_delivery int NOT NULL, min_delivery_time int, max_delivery_time int, city text NOT NULL, street text NOT NULL, house text NOT NULL, floor int, rating double precision, location text);"
	tableCookie := "CREATE TABLE IF NOT EXISTS cookie (id serial PRIMARY KEY, client_id INTEGER, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE, session_id text NOT NULL, date_life timestamp NOT NULL, csrf_token varchar(256) NOT NULL);"
	tableHost := "CREATE TABLE IF NOT EXISTS host (id serial PRIMARY KEY, client_id INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE);"
	tableClient := "CREATE TABLE IF NOT EXISTS client (id serial PRIMARY KEY, client_id INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE, date_birthday timestamp NOT NULL);"
	tableCourier := "CREATE TABLE IF NOT EXISTS courier (id serial PRIMARY KEY, client_id  INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE);"
	_, err = conn.Exec(context.Background(), tableGeneralUserInfo + tableRestaurant + tableCookie + tableHost + tableClient + tableCourier)
	if err != nil {
		return nil, errors.New(ERRCREATEQUERY)
	}

	if DEBUG {
		_, err = conn.Exec(context.Background(),
			"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5)",
			"root", "root@root", "88888888888", HashPassword("root", "salt"), "salt")

		if err != nil {
			return nil, errors.New(ERRINSERTROOTQUERY)
		}

		for i := 0; i < 500; i++ {
			_, err := conn.Exec(context.Background(),
				"INSERT INTO restaurant (name, description, owner, price_delivery, city, street, house, rating, min_delivery_time, max_delivery_time, avatar, floor, location) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
				makeName(), makeName(), 1, randomInteger(1, 15000), "city", "street", "house", randomInteger(1, 5), randomInteger(1, 300), randomInteger(1, 300), "https://avatars.mds.yandex.net/get-mpic/4944925/img_id5013960435158963405.jpeg/13hq", randomInteger(1, 163), "location")
			if err != nil {
				return nil, errors.New(ERRINSERTQUERY)
			}
		}
	}

	return conn, nil
}

func CheckAccess(conn *pgxpool.Pool, cookie Defense) (bool, error) {
	var timeLiveCookie time.Time
	var id int
	row, err := conn.Query(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
		cookie.SessionId, cookie.CsrfToken)
	if err != nil {
		return false, errors.New(ERRCOOKIEANDCSRFQUERY)
	}

	for row.Next() {
		err = row.Scan(&id, &timeLiveCookie)
		if err != nil {
			return false, errors.New(ERRCOOKIEANDCSRFSCAN)
		}
	}

	if id == 0 {
		return false, errors.New(ERRSIDNOTFOUND)
	}

	realTime := time.Now()
	if realTime.Before(timeLiveCookie) {
		return true, nil
	}

	return false, nil
}

func NewCsrf(conn *pgxpool.Pool, cookie Defense) (string, error) {
	csrfToken := randString(5)
	err := conn.QueryRow(context.Background(),
		"UPDATE cookie SET csrf_token = $1 WHERE session_id = $2",
		csrfToken, cookie.SessionId)
	if err != nil {
		return "", errors.New(ERRUPDATECSRFQUERY)
	}

	return csrfToken, nil
}

func GetIdByCookie(conn *pgxpool.Pool, cookie Defense) (int, error) {
	var timeLiveCookie time.Time
	var id int
	row, err := conn.Query(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
		cookie.SessionId)
	if err != nil {
		return 0, errors.New(ERRCOOKIEQUERY)
	}

	for row.Next() {
		err = row.Scan(&id, &timeLiveCookie)
		if err != nil {
			return 0, errors.New(ERRCOOKIESCAN)
		}
	}

	if id == 0 {
		return 0, errors.New(ERRSIDNOTFOUND)
	}

	realTime := time.Now()
    if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return 0, errors.New(ERRCOOKIEEXPIRED)
}
