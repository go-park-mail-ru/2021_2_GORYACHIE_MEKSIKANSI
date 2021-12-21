package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"strings"
)

func CreateDb(configDB config.Database) (*pgxpool.Pool, error) {
	var err error
	addressPostgres := "postgres://" + configDB.UserName + ":" + configDB.Password +
		"@" + configDB.Host + ":" + configDB.Port + "/" + configDB.SchemaName

	conn, err := pgxpool.Connect(context.Background(), addressPostgres)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.MCreateDBNotConnect,
		}
	}
	return conn, nil
}

func FillDb(conn *pgxpool.Pool) error {
	contextTransaction := context.Background()
	tx, err := conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.MCreateDBTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	file, errRead := ioutil.ReadFile("./build/postgresql/deletetables.sql")
	if errRead != nil {
		return &errPkg.Errors{
			Text: errPkg.MCreateDBDeleteFileNotFound,
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = tx.Exec(context.Background(), request)
		if err != nil {
			return &errPkg.Errors{
				Text: errPkg.MCreateDBNotDeleteTables,
			}
		}
	}

	file, err = ioutil.ReadFile("./build/postgresql/fill.sql")
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.MCreateDBFillFileNotFound,
		}
	}

	requests = strings.Split(string(file), ";")
	for _, request := range requests {

		_, err = tx.Exec(context.Background(), request)
		if err != nil {
			return &errPkg.Errors{
				Text: errPkg.MCreateDBNotFillTables,
			}
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.MCreateDBNotCommit,
		}
	}
	return nil
}
