package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
)

type WrapperAuthorization interface {
	SignupClient(signup *Authorization.RegistrationRequest, cookie *Utils2.Defense) (*Utils2.Defense, error)
	SignupCourier(signup *Authorization.RegistrationRequest, cookie *Utils2.Defense) (*Utils2.Defense, error)
	SignupHost(signup *Authorization.RegistrationRequest, cookie *Utils2.Defense) (*Utils2.Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(CSRF string) (string, error)
	NewDefense() *Utils2.Defense
	AddCookie(cookie *Utils2.Defense, id int) error
	CheckAccess(cookie *Utils2.Defense) (bool, error)
	NewCSRF(cookie *Utils2.Defense) (string, error)
	GetIdByCookie(cookie *Utils2.Defense) (int, error)
}
