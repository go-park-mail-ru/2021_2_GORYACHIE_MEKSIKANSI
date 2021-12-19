package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/build"
	configPkg "2021_2_GORYACHIE_MEKSIKANSI/config"
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	metrics "github.com/w1ck3dg0ph3r/fastprometrics"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	runServer()
}

func runServer() {
	var logger utils.Logger
	logger.Log = utils.NewLogger("./logs.txt")

	defer func(loggerErrWarn errPkg.MultiLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(2)
		}
	}(logger.Log)

	errConfig, configStructure := build.InitConfig()
	if errConfig != nil {
		logger.Log.Errorf("%s", errConfig.Error())
		os.Exit(2)
	}
	appConfig := configStructure[0].(configPkg.AppConfig)
	dbConfig := configStructure[1].(configPkg.DBConfig)
	awsConfig := configStructure[2].(configPkg.AwsConfig)
	microserviceConfig := configStructure[3].(configPkg.MicroserviceConfig)

	connectionPostgres, err := build.CreateDb(dbConfig.Db, appConfig.Primary.Debug)
	defer connectionPostgres.Close()
	if err != nil {
		logger.Log.Errorf("Unable to connect to database: %s", err.Error())
		os.Exit(2)
	}

	errAws, sess := build.ConnectAws(awsConfig.Aws)
	if errAws != nil {
		logger.Log.Errorf("AWS: %s", errAws.Error())
		os.Exit(2)
	}
	uploader := s3manager.NewUploader(sess)
	nameBucket := awsConfig.Aws.Name

	intCh := make(chan authPkg.WebSocketOrder, 10)

	startStructure := build.SetUp(connectionPostgres, logger.Log, uploader, nameBucket, microserviceConfig, intCh)

	userInfo := startStructure.User
	cartInfo := startStructure.Cart
	profileInfo := startStructure.Profile
	infoMid := startStructure.Midle
	restaurantInfo := startStructure.Restaraunt
	orderInfo := startStructure.Order

	userInfo.IntCh = intCh

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	userGroup := versionGroup.Group("/user")
	restaurantGroup := versionGroup.Group("/restaurant")
	cartGroup := userGroup.Group("/cart")
	orderGroup := userGroup.Group("/order")
	webSocketGroup := versionGroup.Group("/ws")
	userWSGroup := userGroup.Group("/ws")
	favouriteGroup := userGroup.Group("/restaurant/favourite")

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
	favouriteGroup.GET("/", infoMid.GetIdClient(restaurantInfo.GetFavouritesHandler))
	favouriteGroup.PUT("/", infoMid.CheckClient(infoMid.GetIdClient(restaurantInfo.UpdateFavouritesHandler)))

	restaurantGroup.GET("/", restaurantInfo.RestaurantHandler)
	restaurantGroup.GET("/{idRes}/dish/{idDish}", restaurantInfo.RestaurantDishesHandler)
	restaurantGroup.GET("/{idRes}", infoMid.GetIdClientIgnoreErr(restaurantInfo.RestaurantIdHandler))
	restaurantGroup.GET("/{idRes}/review", infoMid.GetIdClientIgnoreErr(restaurantInfo.GetReviewHandler))
	restaurantGroup.GET("/search", restaurantInfo.SearchRestaurantHandler)
	restaurantGroup.GET("/recommend", infoMid.GetIdClient(restaurantInfo.RecommendedRestaurantsHandler))

	cartGroup.GET("/", infoMid.GetIdClient(cartInfo.GetCartHandler))
	cartGroup.PUT("/", infoMid.CheckClient(infoMid.GetIdClient(cartInfo.UpdateCartHandler)))

	orderGroup.GET("/", infoMid.GetIdClient(orderInfo.GetOrdersHandler))
	orderGroup.POST("/", infoMid.CheckClient(infoMid.GetIdClient(orderInfo.CreateOrderHandler)))
	orderGroup.GET("/{idOrd}/active", infoMid.GetIdClient(orderInfo.GetOrderActiveHandler))

	webSocketGroup.GET("/", infoMid.CheckWebSocketKey(userInfo.UserWebSocket))
	userWSGroup.GET("/key", infoMid.GetIdClient(userInfo.UserWebSocketNewKey))

	metricsHandler := metrics.Add(func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(http.StatusOK)
	}, metrics.WithPath("/metrics"), metrics.WithSubsystem("http"))

	myRouter.GET("/metrics", metricsHandler)
	myRouter.GET("/internal", infoMid.MetricsInternal(func(ctx *fasthttp.RequestCtx) { // TODO(): delete test handler
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	}))

	printURL := infoMid.LogURL(infoMid.MetricsHits(myRouter.Handler))

	addressAllowedCors := appConfig.Cors.Host + ":" + appConfig.Cors.Port
	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins: []string{addressAllowedCors},
		AllowedHeaders: []string{"access-control-allow-origin", "content-type",
			"x-csrf-token", "access-control-expose-headers"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
		ExposedHeaders:   []string{"X-Csrf-Token"},
		AllowCredentials: true,
		AllowMaxAge:      5600,
		Debug:            true,
	})

	port := ":" + appConfig.Port
	logger.Log.Infof("Listen in 127:0.0.1%s", port)
	err = fasthttp.ListenAndServe(port, withCors.CorsMiddleware(infoMid.MetricsTiming(printURL)))
	if err != nil {
		logger.Log.Errorf("Listen and server error: %v", err)
		os.Exit(2)
	}
}
