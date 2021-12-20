//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm WrapperAuthorizationInterface,ConnectAuthServiceInterface
package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util/cast"
	"context"
	"google.golang.org/grpc"
)

type WrapperAuthorizationInterface interface {
	SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error)
	Login(login *authorization.Authorization) (*Utils2.Defense, error)
	Logout(CSRF string) (string, error)
	NewCSRFWebsocket(id int) (string, error)
}

type ConnectAuthServiceInterface interface {
	SignUp(ctx context.Context, in *authProto.RegistrationRequest, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, in *authProto.Authorization, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, in *authProto.CSRF, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
	NewCSRFWebsocket(ctx context.Context, client *authProto.IdClient, opts ...grpc.CallOption) (*authProto.WebsocketResponse, error)
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
		return nil, &errPkg.Errors{Text: result.Error}
	}
	return cast.CastDefenseResponseProtoToDefense(result), nil
}

func (w *Wrapper) Login(login *authorization.Authorization) (*Utils2.Defense, error) {
	response, err := w.Conn.Login(w.Ctx, cast.CastAuthorizationToAuthorizationProto(login))
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, &errPkg.Errors{Text: response.Error}
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
		return "", &errPkg.Errors{Text: logout.Error}
	}
	return logout.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) NewCSRFWebsocket(id int) (string, error) {
	var idClient authProto.IdClient
	idClient.ClientId = int64(id)
	websocket, err := w.Conn.NewCSRFWebsocket(w.Ctx, &idClient)
	if err != nil {
		return "", err
	}
	if websocket.Error != "" {
		return "", &errPkg.Errors{
			Text: websocket.Error,
		}
	}
	return websocket.Websocket, nil
}
