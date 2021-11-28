package orm

import (
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"google.golang.org/grpc"

	"context"
)

type WrapperMiddlewareInterface interface {
	CheckAccess(cookie *util.Defense) (bool, error)
	NewCSRF(cookie *util.Defense) (string, error)
	GetIdByCookie(cookie *util.Defense) (int, error)
}

type ConnectionMiddlewareInterface interface {
	CheckAccessUser(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.CheckAccess, error)
	NewCSRFUser(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
	GetIdByCookie(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.IdClientResponse, error)
}

type Wrapper struct {
	Conn ConnectionMiddlewareInterface
	Ctx  context.Context
}

func (w *Wrapper) CheckAccess(cookie *util.Defense) (bool, error) {
	user, err := w.Conn.CheckAccessUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return false, err
	}
	if user.Error != "" {
		return false, &errPkg.Errors{Alias: user.Error}
	}
	return user.CheckResult, nil
}

func (w *Wrapper) NewCSRF(cookie *util.Defense) (string, error) {
	user, err := w.Conn.NewCSRFUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return "", err
	}
	if user.Error != "" {
		return "", &errPkg.Errors{Alias: user.Error}
	}
	return user.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) GetIdByCookie(cookie *util.Defense) (int, error) {
	byCookie, err := w.Conn.GetIdByCookie(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return 0, err
	}
	if byCookie.Error != "" {
		return 0, &errPkg.Errors{Alias: byCookie.Error}
	}
	return int(byCookie.IdUser), nil
}
