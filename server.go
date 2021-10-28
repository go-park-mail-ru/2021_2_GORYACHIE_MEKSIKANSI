package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	cart "2021_2_GORYACHIE_MEKSIKANSI/Cart"
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
	cartInfo := cart.InfoCart{ConnectionDB: connectionPostgres}

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	userGroup := versionGroup.Group("/user")
	restaurantGroup := versionGroup.Group("/restaurant")
	cartGroup := userGroup.Group("/cart")

	userGroup.POST("/login", userInfo.LoginHandler)
	userGroup.POST("/signup", userInfo.SignUpHandler)
	userGroup.POST("/logout", userInfo.LogoutHandler)
	userGroup.GET("/", infoMiddleware.GetIdByCookieMiddl(profileInfo.ProfileHandler))
	userGroup.PUT("/name", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserName))
	userGroup.PUT("/email", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserEmail))
	userGroup.PUT("/password", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserPassword))
	userGroup.PUT("/phone", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserPhone))
	userGroup.PUT("/avatar", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserAvatar))
	userGroup.PUT("/birthday", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserBirthday))
	userGroup.PUT("/address", infoMiddleware.GetIdByCookieMiddl(profileInfo.UpdateUserAddress))
	userGroup.POST("/pay", infoMiddleware.CheckAccessMiddl(userInfo.PayHandler))

	restaurantGroup.GET("/", restaurantInfo.RestaurantHandler)
	restaurantGroup.GET("/{idRes}/dish/{idDish}", restaurantInfo.RestaurantDishesHandler)
	restaurantGroup.GET("/{idRes}", restaurantInfo.RestaurantIdHandler)

	cartGroup.GET("/", infoMiddleware.GetIdByCookieMiddl(cartInfo.GetCartHandler))
	cartGroup.PUT("/", infoMiddleware.CheckAccessMiddl(infoMiddleware.GetIdByCookieMiddl(cartInfo.UpdateCartHandler)))

	printURL := infoMiddleware.PrintURLMiddl(myRouter.Handler)

	//logger := logur.NewNoopLogger()

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins: []string{config.AllowedOriginsDomen + ":" + config.AllowedOriginsPort},
		AllowedHeaders: []string{"access-control-allow-origin", "content-type",
			"x-csrf-token", "access-control-expose-headers"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
		ExposedHeaders:   []string{"X-Csrf-Token"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	err = fasthttp.ListenAndServe(port, withCors.CorsMiddleware(printURL))
	if err != nil {
		fmt.Printf("Console: ERROR: fatal ListenAndServe")
		return
	}
}

func main() {
	runServer(":5000")
}
