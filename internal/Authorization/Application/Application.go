package Application

import (
	Authorization2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
)

const LenSalt = 5

type Authorization struct {
	DB Interface.WrapperAuthorization
}

func (a *Authorization) SignUp(signup *Authorization2.RegistrationRequest) (*Util.Defense, error) {
	return a.DB.SignUp(signup)
}

func (a *Authorization) Login(login *Authorization2.Authorization) (*Util.Defense, error) {
	return a.DB.Login(login)
}

func (a *Authorization) Logout(CSRF string) (string, error) {
	return a.DB.Logout(CSRF)
}
