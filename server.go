package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	cart "2021_2_GORYACHIE_MEKSIKANSI/Cart"
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"os"
)

func runServer(port string) {
	loggerErrWarn := utils.NewLogger("./loggErrWarn.txt")
	loggerInfo := utils.NewLogger("./loggInfo.txt")
	loggerTest := utils.NewLogger("./loggTest.txt")

/*	defer func(sugarLogger *zap.SugaredLogger) {
		errLogger := sugarLogger.Sync()
		if errLogger != nil {
			zap.S().Errorf("Logger the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerErrWarn)*/

	connectionPostgres, err := utils.CreateDb()
	defer connectionPostgres.Close()
	if err != nil {
		if config.TEST {
			loggerTest.Errorf("Unable to connect to database: %v", err)
			os.Exit(1)
		}
		loggerErrWarn.Errorf("Unable to connect to database: %v", err)
		os.Exit(1)
	}
	userInfo := auth.UserInfo{
		ConnectionDB:  connectionPostgres,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}
	restaurantInfo := restaurant.InfoRestaurant{
		ConnectionDB:  connectionPostgres,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}
	profileInfo := profile.InfoProfile{
		ConnectionDB:  connectionPostgres,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}
	infoMiddleware := mid.InfoMiddleware{
		ConnectionDB:  connectionPostgres,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}
	cartInfo := cart.InfoCart{
		ConnectionDB:  connectionPostgres,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	userGroup := versionGroup.Group("/user")
	restaurantGroup := versionGroup.Group("/restaurant")
	cartGroup := userGroup.Group("/cart")

	userGroup.POST("/login", userInfo.LoginHandler)
	userGroup.POST("/signup", userInfo.SignUpHandler)
	userGroup.POST("/logout", userInfo.LogoutHandler)
	userGroup.GET("/", infoMiddleware.GetId(profileInfo.ProfileHandler))
	userGroup.PUT("/name", infoMiddleware.GetId(profileInfo.UpdateUserName))
	userGroup.PUT("/email", infoMiddleware.GetId(profileInfo.UpdateUserEmail))
	userGroup.PUT("/password", infoMiddleware.GetId(profileInfo.UpdateUserPassword))
	userGroup.PUT("/phone", infoMiddleware.GetId(profileInfo.UpdateUserPhone))
	userGroup.PUT("/avatar", infoMiddleware.GetId(profileInfo.UpdateUserAvatar))
	userGroup.PUT("/birthday", infoMiddleware.GetId(profileInfo.UpdateUserBirthday))
	userGroup.PUT("/address", infoMiddleware.GetId(profileInfo.UpdateUserAddress))
	userGroup.POST("/pay", infoMiddleware.Check(userInfo.PayHandler))

	restaurantGroup.GET("/", restaurantInfo.RestaurantHandler)
	restaurantGroup.GET("/{idRes}/dish/{idDish}", restaurantInfo.RestaurantDishesHandler)
	restaurantGroup.GET("/{idRes}", restaurantInfo.RestaurantIdHandler)

	cartGroup.GET("/", infoMiddleware.GetId(cartInfo.GetCartHandler))
	cartGroup.PUT("/", infoMiddleware.Check(infoMiddleware.GetId(cartInfo.UpdateCartHandler)))

	printURL := infoMiddleware.PrintURL(myRouter.Handler)

	/*	for i := 0; i < 1000; i++ {
		reqId := utils.RandomInteger(0, math.MaxInt64)
		loggerErrWarn.Infof("URL = %s, request_id = %d", "https://github.com/uber-go/zap:", reqId)
	}*/

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins: []string{config.AllowedOriginsDomain + ":" + config.AllowedOriginsPort},
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
		if config.TEST {
			loggerTest.Errorf("Unable to connect to database: %v", err)
			os.Exit(1)
		}
		loggerErrWarn.Errorf("Listen and server error: %v", err)
		os.Exit(1)
	}
}

func main() {
	runServer(":5000")
}
