package Utils

import (
	"time"
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
