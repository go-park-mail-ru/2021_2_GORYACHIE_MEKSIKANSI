package application

import (
	Authorization2 "2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
)

const LenSalt = 5

type AuthorizationApplicationInterface interface {
	SignUp(signup *Authorization2.RegistrationRequest) (*util.Defense, error)
	Login(login *Authorization2.Authorization) (*util.Defense, error)
	Logout(CSRF string) (string, error)
}

type Authorization struct {
	DB ormPkg.WrapperAuthorizationInterface
}

func (a *Authorization) SignUp(signup *Authorization2.RegistrationRequest) (*util.Defense, error) {
	return a.DB.SignUp(signup)
}

func (a *Authorization) Login(login *Authorization2.Authorization) (*util.Defense, error) {
	return a.DB.Login(login)
}

func (a *Authorization) Logout(CSRF string) (string, error) {
	return a.DB.Logout(CSRF)
}
