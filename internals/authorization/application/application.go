//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm WrapperAuthorizationInterface
package application

import (
	Authorization2 "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
)

type AuthorizationApplicationInterface interface {
	SignUp(signup *Authorization2.RegistrationRequest) (*util.Defense, error)
	Login(login *Authorization2.Authorization) (*util.Defense, error)
	Logout(CSRF string) (string, error)
	NewCSRFWebsocket(id int) (string, error)
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

func (a *Authorization) NewCSRFWebsocket(id int) (string, error) {
	return a.DB.NewCSRFWebsocket(id)
}
