package Middleware

import (
	"math/rand"
	"strings"
	"time"
)

const DAYLIVECOOKIE = 5
const LENSESSINID = 92
const LENCSRFTOKEN = 92

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}

func randString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func (c Defense) GenerateNew() Defense {
	c.DateLife = time.Now().Add(time.Hour * 24 * DAYLIVECOOKIE)
	c.SessionId = randString(LENSESSINID)
	c.CsrfToken = randString(LENCSRFTOKEN)
	return c
}