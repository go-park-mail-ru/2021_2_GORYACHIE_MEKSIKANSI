//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/orm WrapperMiddlewareInterface,ConnectionMiddlewareInterface,ConnectionInterface,TransactionInterface
package orm

import (
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util/cast"
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

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
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
		return false, &errPkg.Errors{Text: user.Error}
	}
	return user.CheckResult, nil
}

func (w *Wrapper) NewCSRF(cookie *util.Defense) (string, error) {
	user, err := w.Conn.NewCSRFUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return "", err
	}
	if user.Error != "" {
		return "", &errPkg.Errors{Text: user.Error}
	}
	return user.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) GetIdByCookie(cookie *util.Defense) (int, error) {
	byCookie, err := w.Conn.GetIdByCookie(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return 0, err
	}
	if byCookie.Error != "" {
		return 0, &errPkg.Errors{Text: byCookie.Error}
	}
	return int(byCookie.IdUser), nil
}

func (w *Wrapper) CheckAccessWebsocket(websocket string) (bool, error) {
	contextTransaction := context.Background()
	tx, err := w.DBConn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.MCheckAccessWebsocketTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var exist *int32
	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM cookie WHERE websocket = $1", websocket).Scan(&exist)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = tx.Commit(contextTransaction)
			if err != nil {
				return false, &errPkg.Errors{
					Text: errPkg.MCheckAccessWebsocketNotCommit,
				}
			}
			return false, nil
		}
		return false, &errPkg.Errors{
			Text: errPkg.MCheckAccessWebsocketNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.MCheckAccessWebsocketNotCommit,
		}
	}

	return true, nil
}
