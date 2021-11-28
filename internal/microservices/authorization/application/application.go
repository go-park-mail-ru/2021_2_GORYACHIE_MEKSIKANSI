package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/orm"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
)

type AuthorizationInterface interface {
	SignUp(signup *authorization.RegistrationRequest) (*util.Defense, error)
	Login(login *authorization.Authorization) (*util.Defense, error)
	Logout(CSRF string) (string, error)
	CheckAccess(cookie *util.Defense) (bool, error)
	NewCSRF(cookie *util.Defense) (string, error)
	GetIdByCookie(cookie *util.Defense) (int, error)
}

type AuthorizationApplication struct {
	DB ormPkg.WrapperAuthorization
}

func (ap *AuthorizationApplication) SignUp(signup *authorization.RegistrationRequest) (*util.Defense, error) {
	var cookie *util.Defense
	var err error
	newCookie := ap.DB.NewDefense()
	switch signup.TypeUser {
	case "client":
		cookie, err = ap.DB.SignupClient(signup, newCookie)
	case "courier":
		cookie, err = ap.DB.SignupCourier(signup, newCookie)
	case "host":
		cookie, err = ap.DB.SignupHost(signup, newCookie)
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

func (ap *AuthorizationApplication) Login(login *authorization.Authorization) (*util.Defense, error) {
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = ap.DB.LoginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = ap.DB.LoginByPhone(login.Phone, login.Password)
	default:
		return nil, &errPkg.Errors{
			Alias: errPkg.ALoginVoidLogin,
		}
	}

	if err != nil {
		return nil, err
	}

	cookie := ap.DB.NewDefense()
	err = ap.DB.AddCookie(cookie, userId)

	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (ap *AuthorizationApplication) Logout(CSRF string) (string, error) {
	cookie, err := ap.DB.DeleteCookie(CSRF)
	if err != nil {
		return "", err
	}
	return cookie, nil
}

func (ap *AuthorizationApplication) CheckAccess(cookie *util.Defense) (bool, error) {
	status, err := ap.DB.CheckAccess(cookie)
	if err != nil {
		return false, err
	}
	return status, nil
}

func (ap *AuthorizationApplication) NewCSRF(cookie *util.Defense) (string, error) {
	csrf, err := ap.DB.NewCSRF(cookie)
	if err != nil {
		return "", err
	}
	return csrf, nil
}

func (ap *AuthorizationApplication) GetIdByCookie(cookie *util.Defense) (int, error) {
	byCookie, err := ap.DB.GetIdByCookie(cookie)
	if err != nil {
		return 0, err
	}
	return byCookie, nil
}
