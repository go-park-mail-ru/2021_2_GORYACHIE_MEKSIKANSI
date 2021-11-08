package Utils

import (
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/microcosm-cc/bluemonday"
	"github.com/valyala/fasthttp"
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

func InterfaceConvertInt(value interface{}) (int, error) {
	var intConvert int
	var errorConvert error
	switch value.(type) {
	case string:
		intConvert, errorConvert = strconv.Atoi(value.(string))
		if errorConvert != nil {
			return errors.IntNil, &errors.Errors{
				Text: errors.ErrAtoi,
			}
		}
		return intConvert, nil
	case int:
		intConvert = value.(int)
		return intConvert, nil
	default:
		return errors.IntNil, &errors.Errors{
			Text: errors.ErrNotStringAndInt,
		}
	}

}

func InterfaceConvertString(value interface{}) (string, error) {
	var StringConvert string
	switch value.(type) {
	case string:
		StringConvert = value.(string)
		return StringConvert, nil
	case int:
		StringConvert = strconv.Itoa(value.(int))
		return StringConvert, nil
	default:
		return "", &errors.Errors{
			Text: errors.ErrNotStringAndInt,
		}
	}

}

func Sanitize(str string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(str)
}
