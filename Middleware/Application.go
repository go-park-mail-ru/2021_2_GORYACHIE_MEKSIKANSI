package Middleware

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/rand"
	"time"
)

const LOGINDB string = "constantil"
const PASSWORDDB string = "293456QwErty"
const DBNAME string = "hot_mexicans_db"

const (
	ERRQUERY = "Error query"
	ERRSCAN = "Error scan"
	ERRSIDNOTFOUND = "Error id not found"
	ERRNOTCONNECT = "Error not connect"
	IDNOTFOUND = -1
)

func randomInteger(min int, max int) int {
	return rand.Intn(max - min) + min
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
	conn, err := pgxpool.Connect(context.Background(), "postgres://" + LOGINDB + ":" + PASSWORDDB + "@localhost:5432/" + DBNAME)
	if err != nil {
		return nil, errors.New(ERRNOTCONNECT)
	}

	tableGeneralUserInfo:= "CREATE TABLE IF NOT EXISTS general_user_info (id SERIAL PRIMARY KEY, name text NOT NULL, password varchar(64) NOT NULL, salt varchar(5) NOT NULL, phone varchar(15) UNIQUE NOT NULL, email text UNIQUE, avatar text DEFAULT '/uploads/', date_registration timestamp DEFAULT NOW() NOT NULL, deleted boolean DEFAULT false);"
	tableRestaurant := "CREATE TABLE IF NOT EXISTS restaurant (id serial PRIMARY KEY, owner INTEGER, FOREIGN KEY (owner) REFERENCES general_user_info (id) On DELETE CASCADE, name text NOT NULL, description text NOT NULL, created timestamp DEFAULT NOW() NOT NULL, deleted boolean DEFAULT false, avatar text DEFAULT '/uploads/', min_price int DEFAULT 0, price_delivery int NOT NULL, min_delivery_time int, max_delivery_time int, city text NOT NULL, street text NOT NULL, house text NOT NULL, floor int, rating float, location text);"
	tableCookie := "CREATE TABLE IF NOT EXISTS cookie (id serial PRIMARY KEY, client_id INTEGER, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE, session_id text NOT NULL, date_life timestamp NOT NULL, csrf_token varchar(256) NOT NULL);"
	tableHost := "CREATE TABLE IF NOT EXISTS host (id serial PRIMARY KEY, client_id INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE);"
	tableClient := "CREATE TABLE IF NOT EXISTS client (id serial PRIMARY KEY, client_id INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE, date_birthday timestamp NOT NULL);"
	tableCourier := "CREATE TABLE IF NOT EXISTS courier (id serial PRIMARY KEY, client_id  INTEGER UNIQUE, FOREIGN KEY (client_id) REFERENCES general_user_info (id) On DELETE CASCADE);"
	_, err = conn.Exec(context.Background(), tableGeneralUserInfo + tableRestaurant + tableCookie + tableHost + tableClient + tableCourier)
	if err != nil {
		return nil, errors.New(ERRQUERY)
	}

	for  i := 0; i < 100; i++ {
		_, err := conn.Exec(context.Background(), "INSERT INTO restaurant (name, description, price_delivery, city, street, house) VALUES ($1, $2, $3, $4, $5, $6)", makeName(), makeName(), randomInteger(10, 15), "city", "street", "house", 0)
		if err != nil {
			return nil, errors.New(ERRQUERY)
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
		return false, errors.New(ERRQUERY)
	}

	for row.Next() {
		err = row.Scan(&id, &timeLiveCookie)
		if err != nil {
			return false, errors.New(ERRSCAN)
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

func GetIdByCookie(conn *pgxpool.Pool, cookie Defense) (int, error) {
	var timeLiveCookie time.Time
	var id int
	row, err := conn.Query(context.Background(),
		"SELECT client_id, date_life FROM cookie WHERE session_id = $1",
		cookie.SessionId)
	if err != nil {
		return 0, errors.New(ERRQUERY)
	}

	for row.Next() {
		err = row.Scan(&id, &timeLiveCookie)
		if err != nil {
			return 0, errors.New(ERRSCAN)
		}
	}

	if id == 0 {
		return 0, errors.New(ERRSIDNOTFOUND)
	}

	realTime := time.Now()
    if realTime.Before(timeLiveCookie) {
		return id, nil
	}

	return IDNOTFOUND, nil
}
