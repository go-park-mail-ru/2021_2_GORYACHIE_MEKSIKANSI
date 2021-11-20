package build

import (
	config "2021_2_GORYACHIE_MEKSIKANSI/configs"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Api"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Orm"
	Api2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart/Api"
	Application2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart/Application"
	Orm2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart/Orm"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyErrors"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interfaces"
	Api3 "2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Api"
	Application3 "2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Application"
	Orm3 "2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Orm"
	Api4 "2021_2_GORYACHIE_MEKSIKANSI/internal/Order/Api"
	Application4 "2021_2_GORYACHIE_MEKSIKANSI/internal/Order/Application"
	Orm4 "2021_2_GORYACHIE_MEKSIKANSI/internal/Order/Orm"
	Api5 "2021_2_GORYACHIE_MEKSIKANSI/internal/Profile/Api"
	Application5 "2021_2_GORYACHIE_MEKSIKANSI/internal/Profile/Application"
	Orm5 "2021_2_GORYACHIE_MEKSIKANSI/internal/Profile/Orm"
	Api6 "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant/Api"
	Application6 "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant/Application"
	Orm6 "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant/Orm"
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

func SetUp(connectionDB Interfaces.ConnectionInterface, logger errPkg.MultiLogger,
	uploader *s3manager.Uploader, nameBucket string) []interface{} {

	authWrapper := Orm.Wrapper{Conn: connectionDB}
	authApp := Application.Authorization{DB: &authWrapper}
	userInfo := Api.UserInfo{
		Application: &authApp,
		Logger:      logger,
	}
	var _ Interfaces.AuthorizationAPI = &userInfo

	profileWrapper := Orm5.Wrapper{
		Conn:       connectionDB,
		Uploader:   uploader,
		NameBucket: nameBucket,
	}
	profileApp := Application5.Profile{DB: &profileWrapper}
	profileInfo := Api5.InfoProfile{
		Application: &profileApp,
		Logger:      logger,
	}
	var _ Interfaces.ProfileAPI = &profileInfo

	midWrapper := Orm3.Wrapper{Conn: connectionDB}
	midApp := Application3.Middleware{DB: &midWrapper}
	infoMiddleware := Api3.InfoMiddleware{
		Application: &midApp,
		Logger:      logger,
	}
	var _ Interfaces.MiddlewareAPI = &infoMiddleware

	restWrapper := Orm6.Wrapper{Conn: connectionDB}
	restApp := Application6.Restaurant{DB: &restWrapper}
	restaurantInfo := Api6.InfoRestaurant{
		Application: &restApp,
		Logger:      logger,
	}
	var _ Interfaces.RestaurantAPI = &restaurantInfo

	cartWrapper := Orm2.Wrapper{Conn: connectionDB}
	cartApp := Application2.Cart{DB: &cartWrapper, DBRestaurant: &restWrapper}
	cartInfo := Api2.InfoCart{
		Application: &cartApp,
		Logger:      logger,
	}
	var _ Interfaces.CartApi = &cartInfo

	orderWrapper := Orm4.Wrapper{Conn: connectionDB}
	orderApp := Application4.Order{
		DB:           &orderWrapper,
		DBCart:       &cartWrapper,
		DBProfile:    &profileWrapper,
		DBRestaurant: &restWrapper,
	}
	orderInfo := Api4.InfoOrder{
		Application: &orderApp,
		Logger:      logger,
	}
	var _ Interfaces.OrderAPI = &orderInfo

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
		file, err := ioutil.ReadFile("build/PostgreSQL/DeleteTables.sql")
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

	file, err := ioutil.ReadFile("build/PostgreSQL/CreateTables.sql")
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
		file, err := ioutil.ReadFile("build/PostgreSQL/Fill.sql")
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
