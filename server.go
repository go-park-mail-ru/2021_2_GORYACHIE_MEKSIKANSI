package main

import (
	"context"
	"fmt"
	"github.com/fasthttp/router"
	_ "github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

const PASSWORD_DB string = ""
const LOGIN_DB string = "Captain-Matroskin"

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

type SignUp struct {
	TypeIn   string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password"`
}

type Restaurants struct {
	ImgUrl         string  `json:"imgUrl"`
	RestaurantName string  `json:"restaurantName"`
	PriceDelivery  string  `json:"costForFreeDelivery"`
	MinDelivery    string  `json:"minDelivery"`
	MaxDelivery    string  `json:"maxDelivery"`
	Rating         float32 `json:"rating"`
}

type Profile struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`
}

type UserInfo struct {
	connectionDB *pgxpool.Pool
}
type RestaurantInfo struct {
	connectionDB *pgxpool.Pool
}

func (db *Wrapper) getRestaurants() ([]Restaurant, error) {
	row, _ := db.conn.Query(context.Background(), "SELECT id FROM general_user_info")
	p := Restaurant{}
	var result []Restaurant
	for row.Next() {
		err := row.Scan(&p.costForFreeDelivery)
		if err != nil {
			panic(err)
		}
		result = append(result, p)
	}

	return result, nil
}

func allRestaurants(db Wrapper) []Restaurant {
	result, _ := db.getRestaurants()
	return result
}

func (r *RestaurantInfo) productsHandler(ctx *fasthttp.RequestCtx) {

	// TODO: make response
	//ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.WriteString("Welcome!")
	fmt.Fprintf(ctx, "Hi there! RequestURI is ")

}

func (u *UserInfo) signUpHandler(ctx *fasthttp.RequestCtx) {
	var db = Wrapper{u.connectionDB}
	restaurants := allRestaurants(db)
	//signUp := SignUp{}
	/*	switch signUp.TypeIn {
		case "client":
			signUpClient(db, name, email, phone, password, birthday)
		case "courier":
			signUpCourier(db, name, email, phone, password)
		case "host":
			signUpHost(db, name, email, phone, password)
		default:
			fmt.Printf("error")


		}*/
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func (u *UserInfo) loginHandler(ctx *fasthttp.RequestCtx) {
	var db = Wrapper{u.connectionDB}
	restaurants := allRestaurants(db)
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func (u *UserInfo) logoutHandler(ctx *fasthttp.RequestCtx) {
	var db = Wrapper{u.connectionDB}
	restaurants := allRestaurants(db)
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func (u *UserInfo) profileHandler(ctx *fasthttp.RequestCtx) {
	var db = Wrapper{u.connectionDB}
	restaurants := allRestaurants(db)
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func runServer(port string) {
	conn, err := pgxpool.Connect(context.Background(), "postgres://"+LOGIN_DB+":"+PASSWORD_DB+"@localhost:5432/hot_mexican")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	userInfo := UserInfo{connectionDB: conn}
	restaurantInfo := RestaurantInfo{connectionDB: conn}

	myRouter := router.New()
	api := "/api"
	myRouter.GET(api+"/restaurants", restaurantInfo.productsHandler)
	myRouter.POST(api+"/login", userInfo.loginHandler)
	myRouter.POST(api+"/logout", userInfo.logoutHandler)
	myRouter.POST(api+"/signup", userInfo.signUpHandler)
	myRouter.GET(api+"/profile", userInfo.profileHandler)
	
	err = fasthttp.ListenAndServe(port, myRouter.Handler)
	if err != nil {
		return
	}
}

func main() {
	runServer(":8080")
}
