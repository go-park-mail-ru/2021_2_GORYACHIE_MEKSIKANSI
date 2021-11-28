package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/build"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/config"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/service"
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
	configDB := configRes[0].(config.DBConfig)
	configApp := configRes[1].(config.AppConfig)

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

	authWrapper := ormPkg.Wrapper{Conn: connectDB}
	authorizationManager := Application.AuthorizationApplication{DB: &authWrapper}
	authInfo := service.AuthorizationManager{Application: &authorizationManager}

	proto.RegisterAuthorizationServiceServer(server, &authInfo)

	logger.Log.Infof("Listen in %s", address)
	errServ := server.Serve(listen)
	if errServ != nil {
		logger.Log.Errorf("%s", errServ.Error())
		os.Exit(1)
	}

}
