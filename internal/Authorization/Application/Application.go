package Application

import (
	Authorization2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
)

const LenSalt = 5

type Authorization struct {
	DB Interface.WrapperAuthorization
}

func (a *Authorization) SignUp(signup *Authorization2.RegistrationRequest) (*Util.Defense, error) {
	var cookie *Util.Defense
	var err error
	newCookie := a.DB.NewDefense()
	switch signup.TypeUser {
	case "client":
		cookie, err = a.DB.SignupClient(signup, newCookie)
	case "courier":
		cookie, err = a.DB.SignupCourier(signup, newCookie)
	case "host":
		cookie, err = a.DB.SignupHost(signup, newCookie)
	default:
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpUnknownType,
		}
	}

	if err != nil {
		return nil, err
	}

	return cookie, nil
}

func (a *Authorization) Login(login *Authorization2.Authorization) (*Util.Defense, error) {
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = a.DB.LoginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = a.DB.LoginByPhone(login.Phone, login.Password)
	default:
		return nil, &errPkg.Errors{
			Alias: errPkg.ALoginVoidLogin,
		}
	}

	if err != nil {
		return nil, err
	}

	cookie := a.DB.NewDefense()
	err = a.DB.AddCookie(cookie, userId)

	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (a *Authorization) Logout(CSRF string) (string, error) {
	return a.DB.DeleteCookie(CSRF)
}