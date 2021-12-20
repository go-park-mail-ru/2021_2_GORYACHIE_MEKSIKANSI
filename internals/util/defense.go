package util

import (
	"time"
)

const (
	DayLiveCookie = 5
	LenSessionId  = 92
	LenCsrfToken  = 92
)

type Defense struct {
	DateLife  time.Time
	SessionId string
	CsrfToken string
}
