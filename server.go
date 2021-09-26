package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	"context"
	_ "encoding/json"
	"fmt"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	_ "net/http"
	"os"
	_ "time"
)

const LOGINDB string = "root"
const PASSWORDDB string = "root"
const DBNAME string = "hot_mexicans_db"

func runServer(port string) {
	// TODO: сделать вернуть connection
	connectionPostgres, err := pgxpool.Connect(context.Background(), "postgres://" + LOGINDB + ":" + PASSWORDDB + "@localhost:5432/" + DBNAME)
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
	myRouter.GET(api+"/", restaurantInfo.ProductsHandler)
	myRouter.POST(api+"/login", userInfo.LoginHandler)
	myRouter.POST(api+"/signup", userInfo.SignUpHandler)

	withCors := cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		AllowedOrigins: []string{"http://127.0.0.1:3000"}, // Only allow example.com to access the resource
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
