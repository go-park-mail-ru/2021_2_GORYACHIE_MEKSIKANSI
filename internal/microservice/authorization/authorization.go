package authorization

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

const (
	DayLiveCookie = 5
	LenSessionId  = 92
	LenCsrfToken  = 92
	LenSalt = 5
)

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}

func (c Defense) GenerateNew() *Defense {
	c.DateLife = time.Now().Add(time.Hour * 24 * DayLiveCookie)
	c.SessionId = RandString(LenSessionId)
	c.CsrfToken = RandString(LenCsrfToken)
	return &c
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

type RegistrationRequest struct {
	TypeUser string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type User struct {
	TypeUser string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Authorization struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
