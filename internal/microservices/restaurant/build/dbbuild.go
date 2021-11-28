package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant/orm"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDb(configDB config.Database) (ormPkg.ConnectionInterface, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+configDB.UserName+":"+configDB.Password+
			"@"+configDB.Host+":"+configDB.Port+"/"+configDB.SchemaName)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
		}
	}
	return conn, nil
}
