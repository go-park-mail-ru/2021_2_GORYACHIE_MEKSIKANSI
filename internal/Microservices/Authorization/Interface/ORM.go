package Interface

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
)

type WrapperAuthorization interface {
	SignupClient(signup *proto.RegistrationRequest, cookie *Utils2.Defense) (*Utils2.Defense, error)
	SignupCourier(signup *proto.RegistrationRequest, cookie *Utils2.Defense) (*Utils2.Defense, error)
	SignupHost(signup *proto.RegistrationRequest, cookie *Utils2.Defense) (*Utils2.Defense, error)
	LoginByEmail(email string, password string) (int, error)
	LoginByPhone(phone string, password string) (int, error)
	DeleteCookie(CSRF string) (string, error)
	NewDefense() *Utils2.Defense
	AddCookie(cookie *Utils2.Defense, id int) error
	CheckAccess(cookie *proto.Defense) (bool, error)
	NewCSRF(cookie *proto.Defense) (string, error)
	GetIdByCookie(cookie *proto.Defense) (int, error)
}
