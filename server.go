package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	cart "2021_2_GORYACHIE_MEKSIKANSI/Cart"
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"os"
)

func runServer(port string) {
	loggerErrWarn := utils.NewLogger("./loggErrWarn.txt")
	loggerInfo := utils.NewLogger("./loggInfo.txt")
	loggerTest := utils.NewLogger("./loggTest.txt")

	defer func(loggerErrWarn *zap.SugaredLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerErrWarn)

	defer func(loggerInfo *zap.SugaredLogger) {
		errLogger := loggerInfo.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerInfo the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerInfo)

	defer func(loggerTest *zap.SugaredLogger) {
		errLogger := loggerTest.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerTest the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerTest)

	connectionPostgres, err := utils.CreateDb()
	defer connectionPostgres.Close()
	if err != nil {
		loggerErrWarn.Errorf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	startStructure := setUp(connectionPostgres, loggerErrWarn, loggerInfo, loggerTest)

	userInfo := startStructure[0].(auth.UserInfo)
	cartInfo := startStructure[1].(cart.InfoCart)
	profileInfo := startStructure[2].(profile.InfoProfile)
	infoMid := startStructure[3].(mid.InfoMiddleware)
	restaurantInfo := startStructure[4].(restaurant.InfoRestaurant)

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	userGroup := versionGroup.Group("/user")
	restaurantGroup := versionGroup.Group("/restaurant")
	cartGroup := userGroup.Group("/cart")

	userGroup.POST("/login", userInfo.LoginHandler)
	userGroup.POST("/signup", userInfo.SignUpHandler)
	userGroup.POST("/logout", infoMid.Check(infoMid.GetId(userInfo.LogoutHandler)))
	userGroup.GET("/",  infoMid.GetId(profileInfo.ProfileHandler))
	userGroup.PUT("/name",  infoMid.Check(infoMid.GetId(profileInfo.UpdateUserName)))
	userGroup.PUT("/email",  infoMid.Check(infoMid.GetId(profileInfo.UpdateUserEmail)))
	userGroup.PUT("/password", infoMid.Check(infoMid.GetId(profileInfo.UpdateUserPassword)))
	userGroup.PUT("/phone",  infoMid.Check(infoMid.GetId(profileInfo.UpdateUserPhone)))
	userGroup.PUT("/avatar",  infoMid.Check(infoMid.GetId(profileInfo.UpdateUserAvatar)))
	userGroup.PUT("/birthday",  infoMid.Check(infoMid.GetId(profileInfo.UpdateUserBirthday)))
	userGroup.PUT("/address",  infoMid.Check(infoMid.GetId(profileInfo.UpdateUserAddress)))
	userGroup.POST("/pay",  infoMid.Check(infoMid.Check(userInfo.PayHandler)))

	restaurantGroup.GET("/", restaurantInfo.RestaurantHandler)
	restaurantGroup.GET("/{idRes}/dish/{idDish}", restaurantInfo.RestaurantDishesHandler)
	restaurantGroup.GET("/{idRes}", restaurantInfo.RestaurantIdHandler)

	cartGroup.GET("/", infoMid.GetId(cartInfo.GetCartHandler))
	cartGroup.PUT("/", infoMid.Check(infoMid.GetId(cartInfo.UpdateCartHandler)))

	printURL := infoMid.PrintURL(myRouter.Handler)

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

func setUp(connectionDB interfaces.ConnectionInterface, loggerErrWarn *zap.SugaredLogger, loggerInfo *zap.SugaredLogger,
	loggerTest *zap.SugaredLogger) []interface{} {

	authWrapper := auth.Wrapper{Conn: connectionDB}
	authApp := auth.Authorization{DB: &authWrapper}
	userInfo := auth.UserInfo{
		Application:   &authApp,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}

	profileWrapper := profile.Wrapper{Conn: connectionDB}
	profileApp := profile.Profile{DB: &profileWrapper}
	profileInfo := profile.InfoProfile{
		Application:   &profileApp,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}

	midWrapper := mid.Wrapper{Conn: connectionDB}
	midApp := mid.Middleware{DB: &midWrapper}
	infoMiddleware := mid.InfoMiddleware{
		Application:   &midApp,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}

	restWrapper := restaurant.Wrapper{Conn: connectionDB}
	restApp := restaurant.Restaurant{DB: &restWrapper}
	restaurantInfo := restaurant.InfoRestaurant{
		Application:   &restApp,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}

	cartWrapper := cart.Wrapper{Conn: connectionDB}
	cartApp := cart.Cart{DB: &cartWrapper, DBRestaurant: &restWrapper}
	cartInfo := cart.InfoCart{
		Application:   &cartApp,
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
	}

	var result []interface{}
	result = append(result, userInfo)
	result = append(result, cartInfo)
	result = append(result, profileInfo)
	result = append(result, infoMiddleware)
	result = append(result, restaurantInfo)

	return result
}

func main() {
	runServer(":5000")
}
