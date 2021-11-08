package Utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Log *zap.SugaredLogger
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.Debugf(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.Infof(template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.Warnf(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.Errorf(template, args)
}

func (l *Logger) Sync() error {
	err := l.Sync()
	if err != nil {
		return err
	}
	return nil
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
	sugarLogger := logger.Sugar()
	return sugarLogger
}
