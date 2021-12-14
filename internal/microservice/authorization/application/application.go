//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/orm WrapperAuthorizationInterface
package Application

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/myerror"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/orm"
)

type AuthorizationInterface interface {
	SignUp(signup *authPkg.RegistrationRequest) (*authPkg.Defense, error)
	Login(login *authPkg.Authorization) (*authPkg.Defense, error)
	Logout(CSRF string) (string, error)
	CheckAccess(cookie *authPkg.Defense) (bool, error)
	NewCSRF(cookie *authPkg.Defense) (string, error)
	GetIdByCookie(cookie *authPkg.Defense) (int, error)
}

type AuthorizationApplication struct {
	DB ormPkg.WrapperAuthorizationInterface
}

func (ap *AuthorizationApplication) SignUp(signup *authPkg.RegistrationRequest) (*authPkg.Defense, error) {
	var cookie *authPkg.Defense
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

func (ap *AuthorizationApplication) Login(login *authPkg.Authorization) (*authPkg.Defense, error) {
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

func (ap *AuthorizationApplication) CheckAccess(cookie *authPkg.Defense) (bool, error) {
	return ap.DB.CheckAccess(cookie)
}

func (ap *AuthorizationApplication) NewCSRF(cookie *authPkg.Defense) (string, error) {
	return ap.DB.NewCSRF(cookie)
}

func (ap *AuthorizationApplication) GetIdByCookie(cookie *authPkg.Defense) (int, error) {
	return ap.DB.GetIdByCookie(cookie)
}
