package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
)

type AuthorizationApplication interface {
	SignUp(signup *authorization.RegistrationRequest) (*util.Defense, error)
	Login(login *authorization.Authorization) (*util.Defense, error)
	Logout(CSRF string) (string, error)
	CheckAccess(cookie *util.Defense) (bool, error)
	NewCSRF(cookie *util.Defense) (string, error)
	GetIdByCookie(cookie *util.Defense) (int, error)
}
