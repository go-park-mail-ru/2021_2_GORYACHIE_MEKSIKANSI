package Authorization

import (
	"math/rand"
	"strings"
	"time"
)

const DAYLIVECOOKIE = 5
const LENSALT = 5
const LENSESSINID = 92
const LENCSRFTOKEN = 92

func (c Defense) generateNew() Defense {
	c.DateLife = time.Now().Add(time.Hour * 24 * DAYLIVECOOKIE)
	c.SessionId = randString(LENSESSINID)
	c.CsrfToken = randString(LENCSRFTOKEN)
	return c
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

func SignUp(db Wrapper, signup Registration) (Defense, error) {
	var cookie Defense
	var err error
	switch signup.TypeIn {
	case "client":
		cookie, err = db.SignupClient(signup)
	case "courier":
		cookie, err = db.SignupCourier(signup)
	case "host":
		cookie, err = db.SignupHost(signup)
	default:
		return Defense{}, err
	}

	if err != nil {
		return cookie, err
	}

	return cookie, nil
}

func Login(db Wrapper, login Authorization) (Defense, error) {
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = db.LoginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = db.LoginByPhone(login.Phone, login.Password)
	default:
		return Defense{}, err
	}

	if err != nil {
		return Defense{}, err
	}

	var cookie Defense
	cookie = cookie.generateNew()
	err = db.AddCookie(cookie, userId)

	if err != nil {
		return cookie, err
	}
	return cookie, nil
}

func Logout(db Wrapper, cookie Defense) error {
	err := db.DeleteCookie(cookie)
	if err != nil {
		return err
	}
	return nil
}
