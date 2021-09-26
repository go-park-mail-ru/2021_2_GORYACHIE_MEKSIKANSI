package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"time"
)

const PASSWORDDB string = ""
const LOGINDB string = "Captain-matroskin"

type SignUp struct {
	TypeIn   	string 		`json:"type"`
	Name     	string 		`json:"name"`
	Email    	string 		`json:"email"`
	Phone    	string 		`json:"phone"`
	Password 	string 		`json:"password"`
	Birthday 	time.Time 	`json:"birthday,omitempty"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password"`
}

type Restaurant struct {
	Img         		string  `json:"imgUrl"`
	Name 				string  `json:"restaurantName"`
	CostForFreeDelivery int  	`json:"costForFreeDelivery"`
	MinDelivery    		int  	`json:"minDelivery"`
	MaxDelivery    		int  	`json:"maxDelivery"`
	Rating         		float32 `json:"rating"`
}

type Profile struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday,omitempty"`
}

type UserInfo struct {
	connectionDB *pgxpool.Pool
}
type RestaurantInfo struct {
	connectionDB *pgxpool.Pool
}

func (r *RestaurantInfo) productsHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.connectionDB}
	restaurant, _/*err*/ := allRestaurants(WrapperDB)  // TODO: проверки на ошибки
	if restaurant!= nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)  // TODO: какой код?
	}
	fmt.Printf("Console:  method: %s, url: %s\n",string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) signUpHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.connectionDB}
	signUpAll := SignUp{}
	err := json.Unmarshal(ctx.Request.Body(), &signUpAll)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	cookieHttp := fasthttp.Cookie{}
	cookieDB := Cookie{}
	cookieDB, _/*err*/ = signUp(wrapper, signUpAll)

	cookieHttp.SetExpire(cookieDB.DateLife)
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус
	fmt.Printf("Console:  method: %s, url: %s\n",string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) loginHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.connectionDB}
	userLogin := Login{}
	err := json.Unmarshal(ctx.Request.Body(), &userLogin)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	cookieHttp := fasthttp.Cookie{}
	cookieDB := Cookie{}
	cookieDB, _/*err*/ = login(wrapper, userLogin)  // TODO: проверки на ошибки

	cookieHttp.SetExpire(cookieDB.DateLife)
	cookieHttp.SetValue(cookieDB.SessionId)
	cookieHttp.SetHTTPOnly(true)
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус
	fmt.Printf("Console:  method: %s, url: %s\n",string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) logoutHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.connectionDB}

	cookieHttp := fasthttp.Cookie{}  // TODO: считать куки
	cookieDB := Cookie{DateLife: cookieHttp.Expire(), SessionId: string(cookieHttp.Value())}
	_/*err*/ = logout(wrapper, cookieDB) // TODO: проверки на ошибки

	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус
	fmt.Printf("Console:  method: %s, url: %s\n",string(ctx.Method()), ctx.URI())
}

func (u *UserInfo) profileHandler(ctx *fasthttp.RequestCtx) {
	wrapper := Wrapper{Conn: u.connectionDB}
	profile := Profile{}

	cookieHttp := fasthttp.Cookie{}  // TODO: считать куки
	cookieDB := Cookie{DateLife: cookieHttp.Expire(), SessionId: string(cookieHttp.Value())}
	profile, _/*err*/ = getProfile(wrapper, cookieDB)  // TODO: проверки на ошибки
	if profile.Email != "" {  // TODO: заглушка на unused
		ctx.Response.SetStatusCode(http.StatusBadRequest)
	}
	ctx.Response.SetStatusCode(http.StatusOK)
	// TODO: записать в json статус
	fmt.Printf("Console:  method: %s, url: %s\n",string(ctx.Method()), ctx.URI())
}

func runServer(port string) {
	connectionPostgres, err := pgxpool.Connect(context.Background(), "postgres://"+LOGINDB+":"+PASSWORDDB+"@localhost:5432/hot_mexican")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer connectionPostgres.Close()

	userInfo := UserInfo{connectionDB: connectionPostgres}
	restaurantInfo := RestaurantInfo{connectionDB: connectionPostgres}

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
