package orm

import (
	"github.com/microcosm-cc/bluemonday"
	"time"
)

func FormatDate(date time.Time) (string, string) {
	return date.Format("02.01.2006"), date.Format("15:04")
}

func Sanitize(str string) string {
	return bluemonday.StrictPolicy().Sanitize(str)
}
