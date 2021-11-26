package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectAuthService
	Ctx context.Context
}

func (w *Wrapper) SignUp(signup *Authorization.RegistrationRequest) (*Utils2.Defense, error) {
	result, err := w.Conn.SignUp(w.Ctx, cast.CastRegistrationRequestToRegistrationRequestProto(signup))
	if err != nil {
		return nil, err
	}
	return cast.CastDefenseResponseProtoToDefense(result), nil
}

func (w *Wrapper) Login(login *Authorization.Authorization) (*Utils2.Defense, error) {
	response, err := w.Conn.Login(w.Ctx, cast.CastAuthorizationToAuthorizationProto(login))
	if err != nil {
		return nil, err
	}
	return cast.CastDefenseResponseProtoToDefense(response), nil
}

func (w *Wrapper) Logout(CSRF string) (string, error) {
	var csrfToken authProto.CSRF
	csrfToken.XCsrfToken = CSRF
	logout, err := w.Conn.Logout(w.Ctx, &csrfToken)
	if err != nil {
		return "", err
	}
	return logout.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) CheckAccess(cookie *Utils2.Defense) (bool, error) {
	user, err := w.Conn.CheckAccessUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return false, err
	}
	return user.CheckResult, nil
}

func (w *Wrapper) NewCSRF(cookie *Utils2.Defense) (string, error){
	user, err := w.Conn.NewCSRFUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return "", err
	}
	return user.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) GetIdByCookie(cookie *Utils2.Defense) (int, error){
	byCookie, err := w.Conn.GetIdByCookie(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return 0, err
	}
	return int(byCookie.IdUser), nil
}
