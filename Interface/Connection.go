package Interface

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error) // TODO: replace on interface Tx
}
