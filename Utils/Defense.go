package Utils

import (
	"time"
)

const (
	DayLiveCookie      = 5
	LenSessionId       = 92
	LenCsrfToken       = 92
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
