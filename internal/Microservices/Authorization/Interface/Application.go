package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
)

type AuthorizationApplication interface {
	SignUp(signup *Authorization.RegistrationRequest) (*Util.Defense, error)
	Login(login *Authorization.Authorization) (*Util.Defense, error)
	Logout(CSRF string) (string, error)
	CheckAccess(cookie *Util.Defense) (bool, error)
	NewCSRF(cookie *Util.Defense) (string, error)
	GetIdByCookie(cookie *Util.Defense) (int, error)
}
