package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	orm "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/service"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/build"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	Port    = ":8081"
	Network = "tcp"
)

func main() {
	var logger utils.Logger
	logger.Log = utils.NewLogger("./logs.txt")

	defer func(loggerErrWarn errPkg.MultiLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(logger.Log)

	listen, errListen := net.Listen(Network, Port)
	if errListen != nil {
		logger.Log.Errorf("%s", errListen.Error())
		os.Exit(1)
	}
	server := grpc.NewServer()

	errConf, configRes := build.InitConfig()
	if errConf != nil {
		logger.Log.Errorf("%s", errConf.Error())
		os.Exit(1)
	}
	configDB := configRes[0].(config.DBConfig)

	connectDB, errDb := build.CreateDb(configDB.Db)
	if errDb != nil {
		logger.Log.Errorf("%s", errDb.Error())
		os.Exit(1)
	}

	authWrapper := orm.Wrapper{Conn: connectDB}
	authorizationManager := Application.AuthorizationApplication{DB: &authWrapper}
	authInfo := service.AuthorizationManager{Application: &authorizationManager}

	proto.RegisterAuthorizationServiceServer(server, &authInfo)

	logger.Log.Infof("Listen in 127.0.0.1%s", Port)
	errServ := server.Serve(listen)
	if errServ != nil {
		logger.Log.Errorf("%s", errServ.Error())
		os.Exit(1)
	}

}
