package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"math/rand"
	"os"
	"strings"
	"time"
)

const DAYLIVECOOKIE = 5
const LENSALT = 5

type Restaurant struct {
	Img                 string  `json:"imgUrl"`
	Name                string  `json:"restaurantName"`
	CostForFreeDelivery int     `json:"costForFreeDelivery"`
	MinDelivery         int     `json:"minDelivery"`
	MaxDelivery         int     `json:"maxDelivery"`
	Rating              float32 `json:"rating"`
}

type Wrapper struct {
	Conn        *pgxpool.Pool
	Transaction pgx.Tx
}

type Cookie struct {
	DateLife  time.Time
	SessionId string
}

type Profile struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`
}

func (c Cookie) generateNew() Cookie {
	c.DateLife = time.Now().Add(time.Hour * 24 * DAYLIVECOOKIE)
	c.SessionId = "123"
	return c
}

type Handler struct {
	foobar string
	Conn   *pgxpool.Pool
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password"`
}

type SignUp struct {
	TypeIn   string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
}

// ORM
func (db *Wrapper) getRestaurants() ([]Restaurant, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT avatar, price_delivery, min_delivery_time, max_delivery_time FROM restaurant LIMIT 50")
	if err != nil {
		return nil, err
	}

	p := Restaurant{}
	var result []Restaurant
	for row.Next() {
		p.Rating = rand.Float32() * 5
		err := row.Scan(&p.Img, &p.Name, &p.CostForFreeDelivery, &p.MinDelivery, &p.MaxDelivery)
		//err := row.Scan(&p.img, &p.name, &p.costForFreeDelivery, &p.minDeliveryTime, &p.maxDeliveryTime, &p.rating)
		if err != nil {
			panic(err)
		}
		result = append(result, p)
	}

	return result, nil
}
func (db *Wrapper) getRoleById(id int) (string, error) {
	role := 0

	row, err := db.Conn.Query(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id)
	if err != nil {
		return "", err
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", err
		}
	}
	if role != 0 {
		return "client", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id)
	if err != nil {
		return "", err
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", err
		}
	}
	if role != 0 {
		return "host", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id)
	if err != nil {
		return "", err
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", err
		}
	}
	if role != 0 {
		return "courier", nil
	}

	return "", nil
}

func randString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func hashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt + password))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func (db *Wrapper) generalSignUp(signup SignUp) (int, error) {
	var userId int
	salt := randString(LENSALT)

	err := db.Transaction.QueryRow(context.Background(),
		"INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		signup.Name, signup.Email, signup.Phone, hashPassword(signup.Password, salt), salt).Scan(&userId)
	if err != nil {
		panic(err)
	}
	return userId, nil
}

func (db *Wrapper) signupHost(signup SignUp) (Cookie, error) {
	tx, err := db.Conn.Begin(context.Background())
	db.Transaction = tx
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: add return err
			return
		}
	}(tx, context.Background())

	userId, err := db.generalSignUp(signup)
	if err != nil {
		return Cookie{}, err
	}

	var cookie Cookie
	cookie = cookie.generateNew()
	err = db.addTransactionCookie(cookie, userId)
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

func (db *Wrapper) signupCourier(signup SignUp) (Cookie, error) {
	tx, err := db.Conn.Begin(context.Background())
	db.Transaction = tx
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: add return err
			return
		}
	}(tx, context.Background())

	userId, err := db.generalSignUp(signup)
	if err != nil {
		return Cookie{}, err
	}

	var cookie Cookie
	cookie = cookie.generateNew()
	err = db.addTransactionCookie(cookie, userId)
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

func (db *Wrapper) signupClient(signup SignUp) (Cookie, error) {
	tx, err := db.Conn.Begin(context.Background())
	db.Transaction = tx
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: add return err
			return
		}
	}(tx, context.Background())

	userId, err := db.generalSignUp(signup)
	if err != nil {
		return Cookie{}, err
	}

	var cookie Cookie
	cookie = cookie.generateNew()
	err = db.addTransactionCookie(cookie, userId)
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

func (db *Wrapper) addTransactionCookie(cookie Cookie, id int) error {
	_, err := db.Transaction.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life) VALUES ($1, $2, $3)",
		id, cookie.SessionId, cookie.DateLife)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) addCookie(cookie Cookie, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life) VALUES ($1, $2, $3)",
		id, cookie.SessionId, cookie.DateLife)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) getIdByCookie(cookie Cookie) (int, error) {
	row, err := db.Conn.Query(context.Background(),
		//"SELECT client_id FROM cookie WHERE session_id = $1", cookie.SessionId)
		"SELECT client_id FROM cookie WHERE session_id = $1 AND date_life = $2", cookie.SessionId, cookie.DateLife)
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
	_, err := db.Conn.Query(context.Background(),
		//"DELETE FROM cookie WHERE session_id = $1",
		"DELETE FROM cookie WHERE session_id = $1 AND date_life = $2",
		cookie.SessionId, cookie.DateLife)
	//cookie.SessionId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Wrapper) checkCookie(cookie Cookie) (bool, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT ID FROM cookie WHERE session_id = $1 AND date_life = $2",
		cookie.SessionId, cookie.DateLife)
	if err != nil {
		return false, err
	}

	var id int
	for row.Next() {
		err = row.Scan(&id)
	}
	if id == 0 {
		return false, nil
	}
	return true, nil
}

func (db *Wrapper) loginByEmail(email string, password string) (int, error) {
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
			panic(err)
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
			panic(err)
			return 0, err
		}
	}

	if user == 0 {
		panic(err)
		return 0, nil
	}
	return user, nil
}

func (db *Wrapper) loginByPhone(phone string, password string) (int, error) {
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
			panic(err)
			return 0, err
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM general_user_info WHERE phone = $1 AND password = $2",
		phone, password)
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&user)
		if err != nil {
			panic(err)
			return 0, err
		}
	}

	if user == 0 {
		panic(err)
		return 0, nil
	}
	return user, nil
}

func (db *Wrapper) getProfileHost(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, err
	}
	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			panic(err)
			return Profile{}, err
		}
	}
	return profile, err
}

func (db *Wrapper) getProfileClient(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, err
	}
	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			panic(err)
			return Profile{}, err
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT date_birthday FROM client WHERE client_id = $1", id)
	if err != nil {
		return Profile{}, err
	}
	for row.Next() {
		err = row.Scan(&profile.Birthday)
		if err != nil {
			panic(err)
			return Profile{}, err
		}
	}
	return profile, err
}

func (db *Wrapper) getProfileCourier(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, err
	}
	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			panic(err)
			return Profile{}, err
		}
	}
	return profile, err
}

func (db *Wrapper) updateName(id int, name string) error {
	_, err := db.Conn.Query(context.Background(),
		"UPDATE general_user_info SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updateEmail(id int, email string) error {
	_, err := db.Conn.Query(context.Background(),
		"UPDATE general_user_info SET email = $1 WHERE id = $2", email, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updatePassword(id int, password string) error {
	_, err := db.Conn.Query(context.Background(),
		"UPDATE general_user_info SET password = $1 WHERE id = $2", password, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updateAdditionalInfo(id int, phone string) error {
	_, err := db.Conn.Query(context.Background(),
		"UPDATE general_user_info SET phone = $1 WHERE id = $2", phone, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Wrapper) updateAvatar(id int, avatar string) error {
	_, err := db.Conn.Query(context.Background(),
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2", avatar, id)
	if err != nil {
		return err
	}
	return nil
}

// Application
func getProfile(db Wrapper, cookie Cookie) (Profile, error) {
	id, err := db.getIdByCookie(cookie)
	if err != nil {
		return Profile{}, err
	}
	role, err := db.getRoleById(id)
	if err != nil {
		return Profile{}, err
	}
	var result Profile
	switch role {
	case "client":
		result, err = db.getProfileClient(id)
	case "courier":
		result, err = db.getProfileCourier(id)
	case "host":
		result, err = db.getProfileHost(id)
	}
	if err != nil {
		return Profile{}, err
	}
	return result, nil
}

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

func signUp(db Wrapper, signup SignUp) (Cookie, error) {
	var cookie Cookie
	var err error
	switch signup.TypeIn {
	case "client":
		cookie, err = db.signupClient(signup)
	case "courier":
		cookie, err = db.signupCourier(signup)
	case "host":
		cookie, err = db.signupHost(signup)
	default:
		return Cookie{}, err

	}
	if err != nil {
		return cookie, err
	}

	return cookie, nil
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

func login(db Wrapper, login Login) (Cookie, error) {
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = db.loginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = db.loginByPhone(login.Phone, login.Password)
	default:
		return Cookie{}, err
	}
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
	var db = Wrapper{Conn: conn}
	restaurants, _ := allRestaurants(db)
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].CostForFreeDelivery)
}

func loginHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{Conn: conn}
	var loginUser = Login{"email", "phone", "password"}
	cookie, _ := login(db, loginUser)
	fmt.Fprintf(ctx, "Hi there! RequestURI is %s", cookie.SessionId)
}

func logoutHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{Conn: conn}
	var cookie Cookie
	cookie = cookie.generateNew()
	err := logout(db, cookie)
	if err != nil {
		return
	}
}

func signupHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{Conn: conn}
	var signup = SignUp{"client", "name", "email", "phone", "password", time.Now()}
	_, err := signUp(db, signup)
	if err != nil {
		return
	}
}

func profileHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{Conn: conn}
	var cookie = Cookie{}
	_, err := getProfile(db, cookie.generateNew())
	if err != nil {
		return
	}
}

// Routing(transportation)

func (h *Handler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/api/restaurant":
		productsHandler(ctx, h.Conn)
	case "/login":
		loginHandler(ctx, h.Conn)
	case "/logout":
		logoutHandler(ctx, h.Conn)
	case "/signup":
		signupHandler(ctx, h.Conn)
	case "/profile":
		profileHandler(ctx, h.Conn)
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

	handler := &Handler{foobar: "foobar", Conn: conn}
	err = fasthttp.ListenAndServe(":8080", handler.HandleFastHTTP)
	if err != nil {
		return
	}
}
