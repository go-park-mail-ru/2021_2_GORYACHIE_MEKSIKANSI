package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/build"
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	auth "2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/api"
	cart "2021_2_GORYACHIE_MEKSIKANSI/internal/cart/api"
	mid "2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/api"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	order "2021_2_GORYACHIE_MEKSIKANSI/internal/order/api"
	profile "2021_2_GORYACHIE_MEKSIKANSI/internal/profile/api"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/api"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"os"
)

func runServer(port string) {
	var logger utils.Logger
	logger.Log = utils.NewLogger("./logs.txt")

	defer func(loggerErrWarn errPkg.MultiLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(logger.Log)

	errConfig, configStructure := build.InitConfig()
	if errConfig != nil {
		logger.Log.Errorf("%s", errConfig.Error())
		return
	}
	appConfig := configStructure[0].(config.AppConfig)
	dbConfig := configStructure[1].(config.DBConfig)
	awsConfig := configStructure[2].(config.AwsConfig)

	connectionPostgres, err := build.CreateDb(dbConfig.Db, appConfig.Primary.Debug)
	defer connectionPostgres.Close()
	if err != nil {
		logger.Log.Errorf("Unable to connect to database: %s", err.Error())
		os.Exit(1)
	}

	errAws, sess := build.ConnectAws(awsConfig.Aws)
	if errAws != nil {
		logger.Log.Errorf("AWS: %s", errAws.Error())
		return
	}
	uploader := s3manager.NewUploader(sess)
	nameBucket := awsConfig.Aws.Name

	startStructure := build.SetUp(connectionPostgres, logger.Log, uploader, nameBucket)

	userInfo := startStructure[0].(auth.UserInfo)
	cartInfo := startStructure[1].(cart.InfoCart)
	profileInfo := startStructure[2].(profile.InfoProfile)
	infoMid := startStructure[3].(mid.InfoMiddleware)
	restaurantInfo := startStructure[4].(restaurant.InfoRestaurant)
	orderInfo := startStructure[5].(order.InfoOrder)

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	userGroup := versionGroup.Group("/user")
	restaurantGroup := versionGroup.Group("/restaurant")
	cartGroup := userGroup.Group("/cart")
	orderGroup := userGroup.Group("/order")

	userGroup.POST("/login", userInfo.LoginHandler)
	userGroup.POST("/signup", userInfo.SignUpHandler)
	userGroup.POST("/logout", infoMid.CheckClient(infoMid.GetIdClient(userInfo.LogoutHandler)))
	userGroup.GET("/", infoMid.GetIdClient(profileInfo.ProfileHandler))
	userGroup.PUT("/name", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserName)))
	userGroup.PUT("/email", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserEmail)))
	userGroup.PUT("/password", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserPassword)))
	userGroup.PUT("/phone", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserPhone)))
	userGroup.PUT("/avatar", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserAvatar)))
	userGroup.PUT("/birthday", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserBirthday)))
	userGroup.PUT("/address", infoMid.CheckClient(infoMid.GetIdClient(profileInfo.UpdateUserAddress)))
	userGroup.POST("/pay", infoMid.CheckClient(infoMid.GetIdClient(userInfo.PayHandler)))
	userGroup.POST("/review", infoMid.CheckClient(infoMid.GetIdClient(restaurantInfo.CreateReviewHandler)))

	restaurantGroup.GET("/", restaurantInfo.RestaurantHandler)
	restaurantGroup.GET("/{idRes}/dish/{idDish}", restaurantInfo.RestaurantDishesHandler)
	restaurantGroup.GET("/{idRes}", restaurantInfo.RestaurantIdHandler)
	restaurantGroup.GET("/{idRes}/review", restaurantInfo.GetReviewHandler)
	restaurantGroup.GET("/search", restaurantInfo.SearchRestaurantHandler)

	cartGroup.GET("/", infoMid.GetIdClient(cartInfo.GetCartHandler))
	cartGroup.PUT("/", infoMid.CheckClient(infoMid.GetIdClient(cartInfo.UpdateCartHandler)))

	orderGroup.GET("/", infoMid.GetIdClient(orderInfo.GetOrdersHandler))
	orderGroup.POST("/", infoMid.CheckClient(infoMid.GetIdClient(orderInfo.CreateOrderHandler)))
	orderGroup.GET("/{idOrd}/active", infoMid.GetIdClient(orderInfo.GetOrderActiveHandler))

	printURL := infoMid.LogURL(myRouter.Handler)

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins: []string{appConfig.Cors.Host + ":" + appConfig.Cors.Port},
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
		logger.Log.Errorf("Listen and server error: %v", err)
		os.Exit(1)
	}
}

func main() {
	runServer(":5000")
}
