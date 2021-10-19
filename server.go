package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"fmt"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"os"
)

func runServer(port string) {
	connectionPostgres, err := utils.CreateDb()
	defer connectionPostgres.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	userInfo := auth.UserInfo{ConnectionDB: connectionPostgres}
	restaurantInfo := restaurant.InfoRestaurant{ConnectionDB: connectionPostgres}
	profileInfo := profile.InfoProfile{ConnectionDB: connectionPostgres}
	infoMiddleware := mid.InfoMiddleware{ConnectionDB: connectionPostgres}

	myRouter := router.New()
	api := myRouter.Group("/api")
	version := api.Group("/v1")
	user := version.Group("/user")
	restaurants := version.Group("/restaurant")

	user.POST("/login", userInfo.LoginHandler)
	user.POST("/signup", userInfo.SignUpHandler)
	user.POST("/logout", userInfo.LogoutHandler)

	restaurants.GET("/", restaurantInfo.RestaurantHandler)
	restaurants.GET("/{idRes}/dishes/{idDish}", restaurantInfo.RestaurantDishesHandler)
	restaurants.GET("/{idRes}", restaurantInfo.RestaurantIdHandler)

	user.GET("/", infoMiddleware.GetIdByCookieMiddleware(profileInfo.ProfileHandler))
	user.PUT("/name", profileInfo.UpdateUserName)

	printURL := infoMiddleware.PrintURLMiddleware(myRouter.Handler)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   []string{config.AllowedOriginsDomen + ":" + config.AllowedOriginsPort},
		AllowedHeaders:   []string{"access-control-allow-origin", "content-type", "x-csrf-token", "access-control-expose-headers"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		ExposedHeaders:   []string{"X-Csrf-Token"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	err = fasthttp.ListenAndServe(port, withCors.CorsMiddleware(printURL))
	if err != nil {
		fmt.Printf("Console: ERROR: fatall lListenAndServe")
		return
	}
}

func main() {
	runServer(":5000")
}
