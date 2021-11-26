package Service

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/Interface"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"
	"context"
)

type AuthorizationManager struct {
	Application Interface.AuthorizationApplication

}

func (am *AuthorizationManager) CheckAccessUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CheckAccess, error) {
	status, err := am.Application.CheckAccess(cast.CastDefenseProtoToDefense(cookie))
	if err != nil {
		return nil, err
	}

	return &authProto.CheckAccess{
		CheckResult: status,
	}, nil
}

func (am *AuthorizationManager) NewCSRFUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CSRFResponse, error) {
	csrf, err := am.Application.NewCSRF(cast.CastDefenseProtoToDefense(cookie))
	if err != nil {
		return nil, err
	}

	return &authProto.CSRFResponse{
		XCsrfToken: &authProto.CSRF{
			XCsrfToken: csrf,
		},
	}, nil
}

func (am *AuthorizationManager) GetIdByCookie(ctx context.Context, cookie *authProto.Defense) (*authProto.IdClientResponse, error) {
	csrf, err := am.Application.GetIdByCookie(cast.CastDefenseProtoToDefense(cookie))
	if err != nil {
		return nil, err
	}
	return &authProto.IdClientResponse{
		IdUser: int64(csrf),
	}, nil
}

func (am *AuthorizationManager) SignUp(ctx context.Context, signup *authProto.RegistrationRequest) (*authProto.DefenseResponse, error) {
	cookie, err := am.Application.SignUp(cast.CastRegistrationRequestProtoToRegistrationRequest(signup))
	if err != nil {
		return nil, err
	}

	return &authProto.DefenseResponse{
		Defense: cast.CastDefenseToDefenseProto(cookie),
	}, nil
}

func (am *AuthorizationManager) Login(ctx context.Context, login *authProto.Authorization) (*authProto.DefenseResponse, error) {
	csrf, err := am.Application.Login(cast.CastAuthorizationProtoToAuthorization(login))
	if err != nil {
		return nil, err
	}

	return &authProto.DefenseResponse{
		Defense: cast.CastDefenseToDefenseProto(csrf),
	}, nil
}

func (am *AuthorizationManager) Logout(ctx context.Context, CSRF *authProto.CSRF) (*authProto.CSRFResponse, error) {
	cookie, err := am.Application.Logout(CSRF.XCsrfToken)
	if err != nil {
		return nil, err
	}

	return &authProto.CSRFResponse{XCsrfToken: &authProto.CSRF{XCsrfToken: cookie}}, nil
}
