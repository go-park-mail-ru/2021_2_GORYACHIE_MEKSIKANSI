//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/orm WrapperMiddlewareInterface,ConnectionMiddlewareInterface,ConnectionInterface
package orm

import (
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"

	"context"
)

type WrapperMiddlewareInterface interface {
	CheckAccess(cookie *util.Defense) (bool, error)
	NewCSRF(cookie *util.Defense) (string, error)
	GetIdByCookie(cookie *util.Defense) (int, error)
	CheckAccessWebsocket(cookie string) (bool, error)
}

type ConnectionMiddlewareInterface interface {
	CheckAccessUser(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.CheckAccess, error)
	NewCSRFUser(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
	GetIdByCookie(ctx context.Context, in *authProto.Defense, opts ...grpc.CallOption) (*authProto.IdClientResponse, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Wrapper struct {
	Conn   ConnectionMiddlewareInterface
	DBConn ConnectionInterface
	Ctx    context.Context
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

func (w *Wrapper) CheckAccessWebsocket(websocket string) (bool, error) {
	contextTransaction := context.Background()
	tx, err := w.DBConn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.OGetOrderTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var exist *int32
	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM cookie WHERE websocket = $1", websocket).Scan(&exist)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.OGetOrderNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Alias: errPkg.OGetOrderNotCommit,
		}
	}

	if exist != nil {
		return true, nil
	}

	return false, nil
}
