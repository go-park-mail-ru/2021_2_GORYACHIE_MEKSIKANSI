package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
	"google.golang.org/grpc"
)

type WrapperAuthorizationInterface interface {
	SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error)
	Login(login *authorization.Authorization) (*Utils2.Defense, error)
	Logout(CSRF string) (string, error)
}

type ConnectAuthServiceInterface interface {
	SignUp(ctx context.Context, in *authProto.RegistrationRequest, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, in *authProto.Authorization, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, in *authProto.CSRF, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
}

type Wrapper struct {
	Conn ConnectAuthServiceInterface
	Ctx  context.Context
}

func (w *Wrapper) SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error) {
	result, err := w.Conn.SignUp(w.Ctx, cast.CastRegistrationRequestToRegistrationRequestProto(signup))
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, &errPkg.Errors{Alias: result.Error}
	}
	return cast.CastDefenseResponseProtoToDefense(result), nil
}

func (w *Wrapper) Login(login *authorization.Authorization) (*Utils2.Defense, error) {
	response, err := w.Conn.Login(w.Ctx, cast.CastAuthorizationToAuthorizationProto(login))
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, &errPkg.Errors{Alias: response.Error}
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
	if logout.Error != "" {
		return "", &errPkg.Errors{Alias: logout.Error}
	}
	return logout.XCsrfToken.XCsrfToken, nil
}
