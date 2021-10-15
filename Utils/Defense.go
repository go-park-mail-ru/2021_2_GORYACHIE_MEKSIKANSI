package Utils

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

const (
	DayLiveCookie 		= 5
	LenSessionId		= 92
	LenCsrfToken		= 92
)

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}

func RandomInteger(min int, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max - min)))
	if err != nil {
		return 5
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

func (c Defense) GenerateNew() *Defense {
	c.DateLife = time.Now().Add(time.Hour * 24 * DayLiveCookie)
	c.SessionId = RandString(LenSessionId)
	c.CsrfToken = RandString(LenCsrfToken)
	return &c
}
