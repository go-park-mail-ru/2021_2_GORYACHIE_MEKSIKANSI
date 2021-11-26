package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	ormCart "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/Orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/Service"
	appCart "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/build"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

const (
	Port    = ":8082"
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

	cartWrapper := ormCart.Wrapper{Conn: connectDB}
	cartApp := appCart.Cart{
		DB:    &cartWrapper,
	}
	cartManager := Service.CartManager{Application: &cartApp}
	proto.RegisterCartServiceServer(server, &cartManager)

	logger.Log.Infof("Listen in 127.0.0.1%s", Port)
	errServ := server.Serve(listen)
	if errServ != nil {
		logger.Log.Errorf("%s", errServ.Error())
		os.Exit(1)
	}

}
