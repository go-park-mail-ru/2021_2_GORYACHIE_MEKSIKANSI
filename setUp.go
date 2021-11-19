package main

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	cart "2021_2_GORYACHIE_MEKSIKANSI/Cart"
	config "2021_2_GORYACHIE_MEKSIKANSI/Configs"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	interfaces "2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	order "2021_2_GORYACHIE_MEKSIKANSI/Order"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func setUp(connectionDB interfaces.ConnectionInterface, logger errPkg.MultiLogger,
	uploader *s3manager.Uploader, nameBucket string) []interface{} {

	authWrapper := auth.Wrapper{Conn: connectionDB}
	authApp := auth.Authorization{DB: &authWrapper}
	userInfo := auth.UserInfo{
		Application: &authApp,
		Logger:      logger,
	}
	var _ interfaces.AuthorizationAPI = &userInfo

	profileWrapper := profile.Wrapper{
		Conn:       connectionDB,
		Uploader:   uploader,
		NameBucket: nameBucket,
	}
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

	orderWrapper := order.Wrapper{Conn: connectionDB}
	orderApp := order.Order{
		DB:           &orderWrapper,
		DBCart:       &cartWrapper,
		DBProfile:    &profileWrapper,
		DBRestaurant: &restWrapper,
	}
	orderInfo := order.InfoOrder{
		Application: &orderApp,
		Logger:      logger,
	}
	var _ interfaces.OrderAPI = &orderInfo

	var result []interface{}
	result = append(result, userInfo)
	result = append(result, cartInfo)
	result = append(result, profileInfo)
	result = append(result, infoMiddleware)
	result = append(result, restaurantInfo)
	result = append(result, orderInfo)

	return result
}

func CreateDb() (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+config.DBLogin+":"+config.DBPassword+
			"@"+config.DBHost+":"+config.DBPort+"/"+config.DBName)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
		}
	}

	if config.Debug {
		file, err := ioutil.ReadFile("PostgreSQL/DeleteTables.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBDeleteFileNotFound,
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errPkg.Errors{
					Alias: errPkg.MCreateDBNotDeleteTables,
				}
			}
		}
	}

	file, err := ioutil.ReadFile("PostgreSQL/CreateTables.sql")
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBCreateFileNotFound,
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBNotCreateTables,
			}
		}
	}

	if config.Debug {
		file, err := ioutil.ReadFile("PostgreSQL/Fill.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBFillFileNotFound,
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errPkg.Errors{
					Alias: errPkg.MCreateDBNotFillTables,
				}
			}
		}
	}
	return conn, nil
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func ConnectAws() *session.Session {
	AccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey := GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion := GetEnvWithKey("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Endpoint: aws.String("fra1.digitaloceanspaces.com"),
			Region:   aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
