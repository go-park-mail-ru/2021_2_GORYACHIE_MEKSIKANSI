package Authorization

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"time"
)

const LenSalt = 5

func SignUp(db utils.WrapperAuthorization, signup *utils.RegistrationRequest) (*utils.Defense, error) {
	var cookie *utils.Defense
	var err error
	newCookie := db.GenerateNew()
	switch signup.TypeUser {
	case "client":
		cookie, err = db.SignupClient(signup, newCookie)
	case "courier":
		cookie, err = db.SignupCourier(signup, newCookie)
	case "host":
		cookie, err = db.SignupHost(signup, newCookie)
	default:
		return nil, &errorsConst.Errors{
			Text: errorsConst.ASignUpUnknownType,
			Time: time.Now(),
		}
	}

	if err != nil {
		return nil, err
	}

	return cookie, nil
}

func Login(db utils.WrapperAuthorization, login *Authorization) (*utils.Defense, error) {
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = db.LoginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = db.LoginByPhone(login.Phone, login.Password)
	default:
		return nil, &errorsConst.Errors{
			Text: errorsConst.ALoginVoidLogin,
			Time: time.Now(),
		}
	}

	if err != nil {
		return nil, err
	}

	cookie := db.GenerateNew()
	err = db.AddCookie(cookie, userId)

	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func Logout(db utils.WrapperAuthorization, cookie *utils.Defense) error {
	return db.DeleteCookie(cookie)
}
