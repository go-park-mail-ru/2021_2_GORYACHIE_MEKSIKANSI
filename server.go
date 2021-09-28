package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	_ "encoding/json"
	"fmt"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	_ "net/http"
	"os"
	_ "time"
)

func runServer(port string) {
	connectionPostgres, err := mid.CreateDb()
	defer connectionPostgres.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	userInfo := auth.UserInfo{ConnectionDB: connectionPostgres}
	restaurantInfo := restaurant.RestaurantInfo{ConnectionDB: connectionPostgres}
	profileInfo := profile.ProfileInfo{ConnectionDB: connectionPostgres}

	myRouter := router.New()
	api := "/api"
/*	cookieDB := mid.Defense{}
	// TODO(Н): Сделать проверку когда надо (проверять не надо только в сигнапе и логине)
	id, err := mid.GetIdByCookie(connectionPostgres, cookieDB)
	if id == 0 {

	}
	// Если id == 0, то сессии или scrf не найден
	if err != nil {
	}*/

	myRouter.GET(api+"/profile", profileInfo.ProfileHandler)
	myRouter.POST(api+"/logout", userInfo.LogoutHandler)
	myRouter.GET(api+"/", restaurantInfo.ProductsHandler)
	myRouter.POST(api+"/login", userInfo.LoginHandler)
	myRouter.POST(api+"/signup", userInfo.SignUpHandler)
	myRouter.GET(api+"/check", userInfo.CheckLoggedInHandler)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:3000"},
		AllowedHeaders: []string{"access-control-allow-origin", "content-type", "x-csrf-token", "access-control-expose-headers"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		ExposedHeaders: []string{"X-Csrf-Token"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})
	err = fasthttp.ListenAndServe(port, withCors.CorsMiddleware(myRouter.Handler))
	if err != nil {
		fmt.Printf("Console: ERROR: fatall lListenAndServe")
		return
	}
}

func main() {
	runServer(":5000")
}
