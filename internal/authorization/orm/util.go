package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"strings"
)

const (
	LenFirstField = 8
	LenSecondField = 4
	LenThirdField = 5
	LenFourthField = 4
	LenFifthField = 12
)

func RandString(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[util.RandomInteger(0, len(chars))])
	}

	return b.String()
}

func generateWebsocket() string {
	return RandString(LenFirstField) + "-" + RandString(LenSecondField) + "-" + RandString(LenThirdField) + "-" + RandString(LenFourthField) + "-" + RandString(LenFifthField)
}
