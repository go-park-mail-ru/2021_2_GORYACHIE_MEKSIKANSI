package Authorization

import (
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"math/rand"
	"strings"
	"time"
)

const LENSALT = 5

func randString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func SignUp(db Wrapper, signup Registration) (mid.Defense, error) {
	var cookie mid.Defense
	var err error
	switch signup.TypeIn {
	case "client":
		cookie, err = db.SignupClient(signup)
	case "courier":
		cookie, err = db.SignupCourier(signup)
	case "host":
		cookie, err = db.SignupHost(signup)
	default:
		return mid.Defense{}, err
	}

	if err != nil {
		return cookie, err
	}

	return cookie, nil
}

func Login(db Wrapper, login Authorization) (mid.Defense, error) {
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = db.LoginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = db.LoginByPhone(login.Phone, login.Password)
	default:
		return mid.Defense{}, err
	}

	if err != nil {
		return mid.Defense{}, err
	}

	var cookie mid.Defense
	//TODO Check cookie
	cookie = cookie.GenerateNew()
	err = db.AddCookie(cookie, userId)

	if err != nil {
		return cookie, err
	}
	return cookie, nil
}

func Logout(db Wrapper, cookie mid.Defense) error {
	err := db.DeleteCookie(cookie)
	if err != nil {
		return err
	}
	return nil
}
