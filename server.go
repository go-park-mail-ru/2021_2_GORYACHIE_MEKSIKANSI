package main

import (
	"context"
	_ "encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	_ "net/http"
	"os"
	auth "project/Authorization"
	mid "project/Middleware"
	profile "project/Profile"
	restaurant "project/Restaurant"
	_ "time"
)

const PASSWORDDB string = "root"
const LOGINDB string = "root"

type Result struct {
	Status int         `json:"status,omitempty"`
	Body   interface{} `json:"parsedJSON,omitempty"`
}

func runServer(port string) {
	// TODO: сделать вернуть connection

	connectionPostgres, err := pgxpool.Connect(context.Background(), "postgres://"+LOGINDB+":"+PASSWORDDB+"@localhost:5432/hot_mexicans_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer connectionPostgres.Close()
	err = mid.CreateDb(connectionPostgres)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	userInfo := auth.UserInfo{ConnectionDB: connectionPostgres}
	restaurantInfo := restaurant.RestaurantInfo{ConnectionDB: connectionPostgres}
	profileInfo := profile.ProfileInfo{ConnectionDB: connectionPostgres}

	myRouter := router.New()
	api := "/api"
	cookieDB := mid.Defense{}
	mid.CheckAccess(connectionPostgres, cookieDB) // TODO: проверка кук для этих профиля и логаута

	myRouter.GET(api+"/profile", profileInfo.ProfileHandler)
	myRouter.POST(api+"/logout", userInfo.LogoutHandler)

	myRouter.GET(api+"/restaurants", restaurantInfo.ProductsHandler)
	myRouter.POST(api+"/login", userInfo.LoginHandler)
	myRouter.POST(api+"/signup", userInfo.SignUpHandler)

	err = fasthttp.ListenAndServe(port, myRouter.Handler)
	if err != nil {
		return
	}
}

func main() {
	runServer(":8080")
}
