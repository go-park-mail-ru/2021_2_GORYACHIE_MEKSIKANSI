package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"crypto/sha256"
	"encoding/hex"
	"github.com/microcosm-cc/bluemonday"
	"strings"
)

const (
	PhoneLen       = 11
	LenFirstField  = 8
	LenSecondField = 4
	LenThirdField  = 5
	LenFourthField = 4
	LenFifthField  = 12
)

func Sanitize(str string) string {
	return bluemonday.StrictPolicy().Sanitize(str)
}

func HashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(salt + password))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func RandString(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[util.RandomInteger(0, len(chars))])
	}

	return b.String()
}

func RandPhone() string {
	chars := []rune("0123456789")
	var b strings.Builder

	for i := 0; i < PhoneLen; i++ {
		b.WriteRune(chars[util.RandomInteger(0, len(chars))])
	}

	return b.String()
}

func generateWebsocket() string {
	return RandString(LenFirstField) + "-" + RandString(LenSecondField) + "-" + RandString(LenThirdField) + "-" + RandString(LenFourthField) + "-" + RandString(LenFifthField)
}
