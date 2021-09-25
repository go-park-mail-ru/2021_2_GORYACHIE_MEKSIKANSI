package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"math/rand"
	"os"
	"time"
)

type Restaurant struct {
	img                 string
	name                string
	costForFreeDelivery int
	minDeliveryTime     int
	maxDeliveryTime     int
	rating              float32
}

type Wrapper struct {
	conn *pgxpool.Pool
}

type Cookie struct {
	dateLife  time.Time
	sessionId string
}

func (c Cookie) generateNew() Cookie {
	c.dateLife = time.Now()
	c.sessionId = "123"
	return c
}

type Handler struct {
	foobar string
	conn   *pgxpool.Pool
}

// ORM
func (db *Wrapper) getRestaurants() ([]Restaurant, error) {
	row, err := db.conn.Query(context.Background(),
		"SELECT avatar, price_delivery, min_delivery_time, max_delivery_time FROM restaurant LIMIT 50")
	if err != nil {
		return nil, err
	}

	p := Restaurant{}
	var result []Restaurant
	for row.Next() {
		p.rating = rand.Float32() * 5
		err := row.Scan(&p.img, &p.name, &p.costForFreeDelivery, &p.minDeliveryTime, &p.maxDeliveryTime)
		//err := row.Scan(&p.img, &p.name, &p.costForFreeDelivery, &p.minDeliveryTime, &p.maxDeliveryTime, &p.rating)
		if err != nil {
			panic(err)
		}
		result = append(result, p)
	}

	return result, nil
}

func (db *Wrapper) generalSignUp(name string, email string, phone string, password string) (int, error) {
	var userId int
	err := db.conn.QueryRow(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		name, email, phone, password, "salt").Scan(&userId)
	if err != nil {
		panic(err)
	}
	return userId, nil
}

func (db *Wrapper) signupHost(userId int) error {
	_, err := db.conn.Exec(context.Background(),
		"INSERT INTO host (client_id) VALUES ($1)", userId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) signupCourier(userId int) error {
	_, err := db.conn.Exec(context.Background(),
		"INSERT INTO courier (client_id) VALUES ($1)", userId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) addCookie(cookie Cookie, id int) error {
	_, err := db.conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life) VALUES ($1, $2, $3)",
		id, cookie.sessionId, cookie.dateLife)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) getIdByCookie(cookie Cookie) (int, error) {
	row, err := db.conn.Query(context.Background(),
		"SELECT client_id FROM cookie WHERE session_id = $1 AND date_life = $2", cookie.sessionId, cookie.dateLife)
	if err != nil {
		return 0, err
	}

	var id int
	for row.Next() {
		err = row.Scan(&id)
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *Wrapper) deleteCookie(cookie Cookie) error {
	_, err := db.conn.Query(context.Background(),
		"DELETE FROM cookie WHERE session_id = $1",
		//"DELETE FROM cookie WHERE session_id = $1 AND date_life = $2",
		//cookie.sessionId, cookie.dateLife)
		cookie.sessionId)
	if err != nil {
		return err
	}

	return nil
}

//func (db *Wrapper) checkCookie(cookie Cookie) error {
//	_, err := db.conn.Query(context.Background(),
//		"DELETE FROM cookie WHERE session_id = $1 AND date_life = $2",
//		cookie.sessionId, cookie.dateLife)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (db *Wrapper) signupClient(userId int, birthday time.Time) error {
	_, err := db.conn.Exec(context.Background(),
		"INSERT INTO client (client_id, birthday) VALUES ($1, $2)", userId, birthday)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) login(email string, phone string, password string) (int, error) {
	var user int
	row, err := db.conn.Query(context.Background(),
		"SELECT id FROM general_user_info WHERE email = $1 AND phone = $2 AND password = $3",
		email, phone, password)
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&user)
	}
	if err != nil {
		panic(err)
		return 0, err
	}
	// TODO: check
	if user == 0 {
		panic(err)
		return 0, nil
	}
	return user, nil
}

func (db *Wrapper) updateName(id int, name string) error {
	_, err := db.conn.Query(context.Background(),
		"UPDATE general_user_info SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updateEmail(id int, email string) error {
	_, err := db.conn.Query(context.Background(),
		"UPDATE general_user_info SET email = $1 WHERE id = $2", email, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updatePassword(id int, password string) error {
	_, err := db.conn.Query(context.Background(),
		"UPDATE general_user_info SET password = $1 WHERE id = $2", password, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updateAdditionalInfo(id int, phone string) error {
	_, err := db.conn.Query(context.Background(),
		"UPDATE general_user_info SET phone = $1 WHERE id = $2", phone, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updateAvatar(id int, avatar string) error {
	_, err := db.conn.Query(context.Background(),
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2", avatar, id)
	if err != nil {
		return err
	}
	return nil
}

// Application
func allRestaurants(db Wrapper) ([]Restaurant, error) {
	result, err := db.getRestaurants()
	if err != nil {
		return nil, err
	}
	if result != nil {
		return nil, err
	}
	// TODO: add text in error
	return result, nil
}

func signUp(db Wrapper, typeUser string, name string, email string, phone string, password string) (Cookie, error) {
	var id int
	id, err := db.generalSignUp(name, email, phone, password)
	if err != nil {
		var cookie Cookie
		return cookie, err
	}
	switch typeUser {
	case "host":
		err = db.signupHost(id)
		// TODO: if err then delete from general
	case "courier":
		err = db.signupCourier(id)
	}

	if err != nil {
		var cookie Cookie
		return cookie, nil
	}

	var cookie Cookie
	cookie = cookie.generateNew()
	err = db.addCookie(cookie, id)
	if err != nil {
		return cookie, err
	}

	return cookie, nil
}

func signUpClient(db Wrapper, name string, email string, phone string, password string, birthday time.Time) (int, error) {
	var id int
	id, err := db.generalSignUp(name, email, phone, password)
	if err != nil {
		return 0, nil
	}
	err = db.signupClient(id, birthday)

	if err != nil {
		return 0, err
	}

	var cookie Cookie
	cookie = cookie.generateNew()
	err = db.addCookie(cookie, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func updateName(db Wrapper, cookie Cookie, id int, name string) error {
	err := db.updateName(id, name)
	// TODO: get id from cookie
	if err != nil {
		return err
	}
	return nil
}

func updateEmail(db Wrapper, cookie Cookie, id int, email string) error {
	err := db.updateEmail(id, email)
	if err != nil {
		return err
	}
	return nil
}

func updatePassword(db Wrapper, cookie Cookie, id int, password string) error {
	err := db.updatePassword(id, password)
	if err != nil {
		return err
	}
	return nil
}

func updateAdditionalInfo(db Wrapper, cookie Cookie, id int, phone string) error {
	err := db.updateAdditionalInfo(id, phone)
	if err != nil {
		return err
	}
	return nil
}

func updateAvatar(db Wrapper, cookie Cookie, id int, avatar string) error {
	err := db.updateAvatar(id, avatar)
	if err != nil {
		return err
	}
	return nil
}

func login(db Wrapper, email string, phone string, password string) (Cookie, error) {
	userId, err := db.login(email, phone, password)
	if err != nil {
		return Cookie{}, err
	}

	var cookie Cookie
	cookie = cookie.generateNew()
	err = db.addCookie(cookie, userId)
	if err != nil {
		return cookie, err
	}
	return cookie, nil
}

func logout(db Wrapper, cookie Cookie) error {
	err := db.deleteCookie(cookie)
	if err != nil {
		return err
	}
	return nil
}

// API
func productsHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	restaurants, _ := allRestaurants(db)
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func loginHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	cookie, _ := login(db, "1", "1", "1")
	fmt.Fprintf(ctx, "Hi there! RequestURI is %s", cookie.sessionId)
}

func logoutHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	var cookie Cookie
	cookie = cookie.generateNew()
	err := logout(db, cookie)
	if err != nil {
		return
	}
}

func signupHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	cookie, _ := signUp(db, "host", "2", "2", "2", "2")
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", cookie.sessionId)
}

// Routing(transportation)

func (h *Handler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/api/restaurant":
		productsHandler(ctx, h.conn)
	case "/login":
		loginHandler(ctx, h.conn)
	case "/logout":
		logoutHandler(ctx, h.conn)
	case "/signup":
		signupHandler(ctx, h.conn)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func main() {
	conn, err := pgxpool.Connect(context.Background(), "postgres://root:root@localhost:5432/hot_mexicans_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	handler := &Handler{foobar: "foobar", conn: conn}
	err = fasthttp.ListenAndServe(":8080", handler.HandleFastHTTP)
	if err != nil {
		return
	}
}
