package Utils

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/natefinch/lumberjack"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/big"
	"strconv"
	"strings"
)

const (
	KeyCookieSessionId = "session_id"
)

type ResponseStatus struct {
	StatusHTTP int `json:"status"`
}

func SetCookieResponse(cookieHTTP *fasthttp.Cookie, cookieDB Defense, sessionId string) {
	cookieHTTP.SetExpire(cookieDB.DateLife)
	cookieHTTP.SetKey(sessionId)
	cookieHTTP.SetValue(cookieDB.SessionId)
	cookieHTTP.SetHTTPOnly(true)
	cookieHTTP.SetPath("/")
}

func RandomInteger(min int, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return max - min
	}
	n := nBig.Int64()
	return int(n) + min
}

func RandString(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[RandomInteger(0, len(chars))])
	}

	return b.String()
}

func HashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt + password))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
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

func interfaceConvertInt(value interface{}) (int, error) {
	var reqId int
	var errorConvert error
	switch value.(type) {
	case string:
		reqId, errorConvert = strconv.Atoi(value.(string))
		if errorConvert != nil {
			return errors.IntNil, &errors.Errors{
				Text: errors.ErrAtoi,
			}
		}
		return reqId, nil
	case int:
		reqId = value.(int)
		return reqId, nil
	default:
		return errors.IntNil, &errors.Errors{
			Text: errors.ErrNotStringAndInt,
		}
	}

}
