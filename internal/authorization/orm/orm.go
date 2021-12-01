package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	cast "2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

type WrapperAuthorizationInterface interface {
	SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error)
	Login(login *authorization.Authorization) (*Utils2.Defense, error)
	Logout(CSRF string) (string, error)
	NewCSRFWebsocket(id int) (string, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type ConnectAuthServiceInterface interface {
	SignUp(ctx context.Context, in *authProto.RegistrationRequest, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, in *authProto.Authorization, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, in *authProto.CSRF, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
}

type Wrapper struct {
	Conn   ConnectAuthServiceInterface
	DBConn ConnectionInterface
	Ctx    context.Context
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

func (w *Wrapper) NewCSRFWebsocket(id int) (string, error) {
	contextTransaction := context.Background()
	tx, err := w.DBConn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.OGetOrderTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	websocket := generateWebsocket()

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO cookie (client_id, websocket) VALUES ($1, $2)", id, websocket)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.OGetOrderNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.OGetOrderNotCommit,
		}
	}

	return websocket, nil
}
