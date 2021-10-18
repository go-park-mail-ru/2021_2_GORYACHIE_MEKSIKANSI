package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	"fmt"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"os"
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
	api := myRouter.Group("/api")
	//version := api.Group("/v1")
	restaurants := api.Group("/restaurant")
	api.GET("/profile", profileInfo.ProfileHandler)
	api.POST("/logout", userInfo.LogoutHandler)
	api.POST("/login", userInfo.LoginHandler)
	api.POST("/signup", userInfo.SignUpHandler)

	api.GET("/", restaurantInfo.RestaurantHandler)
	restaurants.GET("/", restaurantInfo.RestaurantDishesHandler)

	siteHandler := mid.CheckAuthMiddleware(myRouter.Handler)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   []string{config.AllowedOriginsDomen + ":" + config.AllowedOriginsPort},
		AllowedHeaders:   []string{"access-control-allow-origin", "content-type", "x-csrf-token", "access-control-expose-headers"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		ExposedHeaders:   []string{"X-Csrf-Token"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	err = fasthttp.ListenAndServe(port, withCors.CorsMiddleware(siteHandler))
	if err != nil {
		fmt.Printf("Console: ERROR: fatall lListenAndServe")
		return
	}
}

func main() {
	runServer(":5000")
}
