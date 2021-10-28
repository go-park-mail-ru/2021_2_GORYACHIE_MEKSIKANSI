package Utils

import (
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"strings"

	"time"
)

func CreateDb() (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+config.DBLogin+":"+config.DBPassword+
			"@"+config.DBHost+":"+config.DBPort+"/"+config.DBName)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrNotConnect,
			Time: time.Now(),
		}
	}

	if config.DEBUG {
		file, err := ioutil.ReadFile("PostgreSQL/DeleteTables.sql")
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrDeleteFileNotFound,
				Time: time.Now(),
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errorsConst.Errors{
					Text: errorsConst.ErrNotDeleteTables,
					Time: time.Now(),
				}
			}
		}
	}

	file, err := ioutil.ReadFile("PostgreSQL/CreateTables.sql")
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrFileNotFound,
			Time: time.Now(),
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrNotCreateTables,
				Time: time.Now(),
			}
		}
	}

	if config.DEBUG {
		file, err := ioutil.ReadFile("PostgreSQL/Fill.sql")
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.ErrFillFileNotFound,
				Time: time.Now(),
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errorsConst.Errors{
					Text: errorsConst.ErrNotFillTables,
					Time: time.Now(),
				}
			}
		}
	}
	return conn, nil
}
