package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	cart "2021_2_GORYACHIE_MEKSIKANSI/Cart"
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	errors "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"strings"
	"time"
)

func setUp(connectionDB interfaces.ConnectionInterface, logger errors.MultiLogger) []interface{} {
	authWrapper := auth.Wrapper{Conn: connectionDB}
	authApp := auth.Authorization{DB: &authWrapper}
	userInfo := auth.UserInfo{
		Application: &authApp,
		Logger:      logger,
	}
	var _ interfaces.AuthorizationAPI = &userInfo

	profileWrapper := profile.Wrapper{Conn: connectionDB}
	profileApp := profile.Profile{DB: &profileWrapper}
	profileInfo := profile.InfoProfile{
		Application: &profileApp,
		Logger:      logger,
	}
	var _ interfaces.ProfileAPI = &profileInfo

	midWrapper := mid.Wrapper{Conn: connectionDB}
	midApp := mid.Middleware{DB: &midWrapper}
	infoMiddleware := mid.InfoMiddleware{
		Application: &midApp,
		Logger:      logger,
	}
	var _ interfaces.MiddlewareAPI = &infoMiddleware

	restWrapper := restaurant.Wrapper{Conn: connectionDB}
	restApp := restaurant.Restaurant{DB: &restWrapper}
	restaurantInfo := restaurant.InfoRestaurant{
		Application: &restApp,
		Logger:      logger,
	}
	var _ interfaces.RestaurantAPI = &restaurantInfo

	cartWrapper := cart.Wrapper{Conn: connectionDB}
	cartApp := cart.Cart{DB: &cartWrapper, DBRestaurant: &restWrapper}
	cartInfo := cart.InfoCart{
		Application: &cartApp,
		Logger:      logger,
	}
	var _ interfaces.CartApi = &cartInfo

	var result []interface{}
	result = append(result, userInfo)
	result = append(result, cartInfo)
	result = append(result, profileInfo)
	result = append(result, infoMiddleware)
	result = append(result, restaurantInfo)

	return result
}

func CreateDb() (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+config.DBLogin+":"+config.DBPassword+
			"@"+config.DBHost+":"+config.DBPort+"/"+config.DBName)
	if err != nil {
		return nil, &errors.Errors{
			Text: errors.UCreateDBNotConnect,
			Time: time.Now(),
		}
	}

	if config.DEBUG {
		file, err := ioutil.ReadFile("PostgreSQL/DeleteTables.sql")
		if err != nil {
			return nil, &errors.Errors{
				Text: errors.UCreateDBDeleteFileNotFound,
				Time: time.Now(),
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errors.Errors{
					Text: errors.UCreateDBNotDeleteTables,
					Time: time.Now(),
				}
			}
		}
	}

	file, err := ioutil.ReadFile("PostgreSQL/CreateTables.sql")
	if err != nil {
		return nil, &errors.Errors{
			Text: errors.UCreateDBCreateFileNotFound,
			Time: time.Now(),
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errors.Errors{
				Text: errors.UCreateDBNotCreateTables,
				Time: time.Now(),
			}
		}
	}

	if config.DEBUG {
		file, err := ioutil.ReadFile("PostgreSQL/Fill.sql")
		if err != nil {
			return nil, &errors.Errors{
				Text: errors.UCreateDBFillFileNotFound,
				Time: time.Now(),
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errors.Errors{
					Text: errors.UCreateDBNotFillTables,
					Time: time.Now(),
				}
			}
		}
	}
	return conn, nil
}
