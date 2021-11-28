package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	appCart "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/build"
	ormCart "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/service"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
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
	var logger Logger
	logger.Log = NewLogger("./logs.txt")

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
		DB: &cartWrapper,
	}
	cartManager := service.CartManager{Application: &cartApp}
	proto.RegisterCartServiceServer(server, &cartManager)

	logger.Log.Infof("Listen in 127.0.0.1%s", Port)
	errServ := server.Serve(listen)
	if errServ != nil {
		logger.Log.Errorf("%s", errServ.Error())
		os.Exit(1)
	}

}
