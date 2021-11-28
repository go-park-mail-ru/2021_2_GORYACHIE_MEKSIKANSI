package main

import (
	appRes "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/build"
	confPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/config"
	ormRes "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/orm"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/service"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	var logger Logger
	logger.Log = NewLogger("./logs.txt")

	defer func(loggerErrWarn errPkg.MultiLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(logger.Log)

	errConf, configRes := build.InitConfig()
	if errConf != nil {
		logger.Log.Errorf("%s", errConf.Error())
		os.Exit(1)
	}
	configDB := configRes[0].(confPkg.DBConfig)
	configApp := configRes[1].(confPkg.AppConfig)

	address := configApp.Primary.Host  + ":" + configApp.Primary.Port

	listen, errListen := net.Listen(configApp.Primary.Network, address)
	if errListen != nil {
		logger.Log.Errorf("%s", errListen.Error())
		os.Exit(1)
	}
	server := grpc.NewServer()

	connectDB, errDb := build.CreateDb(configDB.Db)
	if errDb != nil {
		logger.Log.Errorf("%s", errDb.Error())
		os.Exit(1)
	}

	resWrapper := ormRes.Wrapper{Conn: connectDB}
	restaurantApp := appRes.Restaurant{
		DB: &resWrapper,
	}
	restaurantManager := service.RestaurantManager{Application: &restaurantApp}

	resProto.RegisterRestaurantServiceServer(server, &restaurantManager)

	logger.Log.Infof("Listen in %s", address)
	errServ := server.Serve(listen)
	if errServ != nil {
		logger.Log.Errorf("%s", errServ.Error())
		os.Exit(1)
	}

}
