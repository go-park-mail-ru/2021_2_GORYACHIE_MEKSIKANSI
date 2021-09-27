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
		// TODO(Н): подправить
	}
	// Если id == 0, то сессии или scrf не найден
	if err != nil {
		// TODO: Сделай с этим что-нибудь
	}*/

	
	myRouter.GET(api+"/profile", profileInfo.ProfileHandler)
	myRouter.POST(api+"/logout", userInfo.LogoutHandler)
	myRouter.GET(api+"/", restaurantInfo.ProductsHandler)
	myRouter.POST(api+"/login", userInfo.LoginHandler)
	myRouter.POST(api+"/signup", userInfo.SignUpHandler)
	myRouter.GET(api+"/check", userInfo.CheckLoggedInHandler)

	withCors := cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		AllowedOrigins: []string{"http://167.172.179.1:3000"}, // Only allow example.com to access the resource
		// if you leave allowedHeaders empty then fasthttpcors will accept any non-simple headers
		AllowedHeaders: []string{"access-control-allow-origin", "content-type"}, // only allow x-something-client and Content-Type in actual request
		// if you leave this empty, only simple method will be accepted
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, // only allow get or post to resource
		AllowCredentials: true,                   // resource doesn't support credentials
		AllowMaxAge:      5600,                    // cache the preflight result
		Debug:            true,
	})
	err = fasthttp.ListenAndServe(port, withCors.CorsMiddleware(myRouter.Handler))
	if err != nil {
		return
	}
}

func main() {

	runServer(":5000")
}

