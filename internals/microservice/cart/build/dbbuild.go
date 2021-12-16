package build

import (
	confPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/myerror"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/orm"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDb(configDB confPkg.Database) (ormPkg.ConnectionInterface, error) {
	var err error
	addressPostgres := "postgres://" + configDB.UserName + ":" + configDB.Password +
		"@" + configDB.Host + ":" + configDB.Port + "/" + configDB.SchemaName

	conn, err := pgxpool.Connect(context.Background(), addressPostgres)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
		}
	}
	return conn, nil
}
