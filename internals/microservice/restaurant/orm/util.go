package orm

import "time"

func FormatDate(date time.Time) (string, string) {
	return date.Format("02.01.2006"), date.Format("15:04")
}
