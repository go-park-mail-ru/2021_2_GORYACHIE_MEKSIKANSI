package orm

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/microcosm-cc/bluemonday"
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
