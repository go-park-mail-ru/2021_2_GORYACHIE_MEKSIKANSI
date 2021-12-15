package main

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/build"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	address := configApp.Primary.Host + ":" + configApp.Primary.Port

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

	authInfo := build.SetUp(connectDB)

	proto.RegisterAuthorizationServiceServer(server, &authInfo)

	logger.Log.Infof("Listen in %s", address)
	errServ := server.Serve(listen)
	if errServ != nil {
		logger.Log.Errorf("%s", errServ.Error())
		os.Exit(1)
	}

}

type Logger struct {
	Log errPkg.MultiLogger
}

func NewLogger(filePath string) *zap.SugaredLogger {
	configLog := zap.NewProductionEncoderConfig()
	configLog.TimeKey = "time_stamp"
	configLog.LevelKey = "level"
	configLog.MessageKey = "note"
	configLog.EncodeTime = zapcore.ISO8601TimeEncoder
	configLog.EncodeLevel = zapcore.CapitalLevelEncoder

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     60,
		Compress:   false,
	}
	writerSyncer := zapcore.AddSync(lumberJackLogger)
	encoder := zapcore.NewConsoleEncoder(configLog)

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	zapLogger := logger.Sugar()
	return zapLogger
}
