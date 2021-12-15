package service

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/application"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	"context"
)

type AuthorizationManagerInterface interface {
	CheckAccessUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CheckAccess, error)
	NewCSRFUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CSRFResponse, error)
	GetIdByCookie(ctx context.Context, cookie *authProto.Defense) (*authProto.IdClientResponse, error)
	SignUp(ctx context.Context, signup *authProto.RegistrationRequest) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, login *authProto.Authorization) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, CSRF *authProto.CSRF) (*authProto.CSRFResponse, error)
}

type AuthorizationManager struct {
	Application appPkg.AuthorizationInterface
}

func (am *AuthorizationManager) CheckAccessUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CheckAccess, error) {
	status, err := am.Application.CheckAccess(CastDefenseProtoToDefense(cookie))
	if err != nil {
		return &authProto.CheckAccess{
			Error: err.Error(),
		}, nil
	}

	return &authProto.CheckAccess{
		CheckResult: status,
	}, nil
}

func (am *AuthorizationManager) NewCSRFUser(ctx context.Context, cookie *authProto.Defense) (*authProto.CSRFResponse, error) {
	csrf, err := am.Application.NewCSRF(CastDefenseProtoToDefense(cookie))
	if err != nil {
		return &authProto.CSRFResponse{
			Error: err.Error(),
		}, nil
	}

	return &authProto.CSRFResponse{
		XCsrfToken: &authProto.CSRF{
			XCsrfToken: csrf,
		},
	}, nil
}

func (am *AuthorizationManager) GetIdByCookie(ctx context.Context, cookie *authProto.Defense) (*authProto.IdClientResponse, error) {
	csrf, err := am.Application.GetIdByCookie(CastDefenseProtoToDefense(cookie))
	if err != nil {
		return &authProto.IdClientResponse{
			Error: err.Error(),
		}, nil
	}
	return &authProto.IdClientResponse{
		IdUser: int64(csrf),
	}, nil
}

func (am *AuthorizationManager) SignUp(ctx context.Context, signup *authProto.RegistrationRequest) (*authProto.DefenseResponse, error) {
	cookie, err := am.Application.SignUp(CastRegistrationRequestProtoToRegistrationRequest(signup))
	if err != nil {
		return &authProto.DefenseResponse{
			Error: err.Error(),
		}, nil
	}

	return &authProto.DefenseResponse{
		Defense: CastDefenseToDefenseProto(cookie),
	}, nil
}

func (am *AuthorizationManager) Login(ctx context.Context, login *authProto.Authorization) (*authProto.DefenseResponse, error) {
	csrf, err := am.Application.Login(CastAuthorizationProtoToAuthorization(login))
	if err != nil {
		return &authProto.DefenseResponse{
			Error: err.Error(),
		}, nil
	}

	return &authProto.DefenseResponse{
		Defense: CastDefenseToDefenseProto(csrf),
	}, nil
}

func (am *AuthorizationManager) Logout(ctx context.Context, CSRF *authProto.CSRF) (*authProto.CSRFResponse, error) {
	cookie, err := am.Application.Logout(CSRF.XCsrfToken)
	if err != nil {
		return &authProto.CSRFResponse{
			Error: err.Error(),
		}, nil
	}

	return &authProto.CSRFResponse{XCsrfToken: &authProto.CSRF{
		XCsrfToken: cookie,
	},
	}, nil
}
