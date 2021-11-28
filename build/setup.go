package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/api"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization/orm"
	Api2 "2021_2_GORYACHIE_MEKSIKANSI/internal/cart/api"
	Application2 "2021_2_GORYACHIE_MEKSIKANSI/internal/cart/application"
	Orm2 "2021_2_GORYACHIE_MEKSIKANSI/internal/cart/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/proto"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/cart/proto"
	resPoroto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/restaurant/proto"
	Api3 "2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/api"
	Application3 "2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/application"
	Orm3 "2021_2_GORYACHIE_MEKSIKANSI/internal/middleware/orm"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	Api4 "2021_2_GORYACHIE_MEKSIKANSI/internal/order/api"
	Application4 "2021_2_GORYACHIE_MEKSIKANSI/internal/order/application"
	Orm4 "2021_2_GORYACHIE_MEKSIKANSI/internal/order/orm"
	Api5 "2021_2_GORYACHIE_MEKSIKANSI/internal/profile/api"
	Application5 "2021_2_GORYACHIE_MEKSIKANSI/internal/profile/application"
	Orm5 "2021_2_GORYACHIE_MEKSIKANSI/internal/profile/orm"
	Api6 "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/api"
	Application6 "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/application"
	Orm6 "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/orm"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	ConfNameMain   = "main"
	ConfNameDB     = "database"
	ConfNameBucket = "bucket"
	ConfType       = "yml"
	ConfPath       = "./config/"
)

func SetUp(connectionDB Interface.ConnectionInterface, logger errPkg.MultiLogger,
	uploader *s3manager.Uploader, nameBucket string) []interface{} {

	grpcConnAuth, errDial := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if errDial != nil {
		println("GG")
		return nil
	}

	authManager := authProto.NewAuthorizationServiceClient(grpcConnAuth)

	authCtx := context.Background()

	authWrapper := orm.Wrapper{Conn: authManager, Ctx: authCtx}
	authApp := application.Authorization{DB: &authWrapper}
	userInfo := api.UserInfo{
		Application: &authApp,
		Logger:      logger,
	}
	var _ Interface.AuthorizationAPI = &userInfo

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
	var _ Interface.ProfileAPI = &profileInfo

	midWrapper := Orm3.Wrapper{Conn: authManager, Ctx: authCtx}
	midApp := Application3.Middleware{DB: &midWrapper}
	infoMiddleware := Api3.InfoMiddleware{
		Application: &midApp,
		Logger:      logger,
	}
	var _ Interface.MiddlewareAPI = &infoMiddleware

	grpcConnRes, errDial := grpc.Dial(
		"127.0.0.1:8084",
		grpc.WithInsecure(),
	)
	if errDial != nil {
		println("GG")
		return nil
	}

	resManager := resPoroto.NewRestaurantServiceClient(grpcConnRes)

	resCtx := context.Background()

	restWrapper := Orm6.Wrapper{Conn: resManager, Ctx: resCtx}
	restApp := Application6.Restaurant{DB: &restWrapper}
	restaurantInfo := Api6.InfoRestaurant{
		Application: &restApp,
		Logger:      logger,
	}
	var _ Interface.RestaurantAPI = &restaurantInfo

	grpcConnCart, errDial := grpc.Dial(
		"127.0.0.1:8082",
		grpc.WithInsecure(),
	)
	if errDial != nil {
		println("GG")
		return nil
	}

	cartManager := cartProto.NewCartServiceClient(grpcConnCart)

	cartCtx := context.Background()

	cartWrapper := Orm2.Wrapper{Conn: cartManager, Ctx: cartCtx}
	cartApp := Application2.Cart{DB: &cartWrapper}
	cartInfo := Api2.InfoCart{
		Application: &cartApp,
		Logger:      logger,
	}
	var _ Interface.CartApi = &cartInfo

	//cartWrapperOld := Orm2.Wrapper{}

	orderWrapper := Orm4.Wrapper{Conn: connectionDB, ConnService: cartManager, Ctx: cartCtx}
	orderApp := Application4.Order{
		DB: &orderWrapper,
		//DBCart:       &cartWrapperOld,
		DBProfile: &profileWrapper,
		//DBRestaurant: &restWrapper,
	}
	orderInfo := Api4.InfoOrder{
		Application: &orderApp,
		Logger:      logger,
	}
	var _ Interface.OrderAPI = &orderInfo

	var result []interface{}
	result = append(result, userInfo)
	result = append(result, cartInfo)
	result = append(result, profileInfo)
	result = append(result, infoMiddleware)
	result = append(result, restaurantInfo)
	result = append(result, orderInfo)

	return result
}

func ConnectAws(config config.AwsBucket) (error, *session.Session) {
	sess, errNewSess := session.NewSession(
		&aws.Config{
			Endpoint: aws.String(config.Endpoint),
			Region:   aws.String(config.Region),
			Credentials: credentials.NewStaticCredentials(
				config.AccessKeyId,
				config.SecretAccessKey,
				"",
			),
		})
	if errNewSess != nil {
		return &errPkg.Errors{
			Alias: errNewSess.Error(),
		}, nil
	}
	return nil, sess
}

func InitConfig() (error, []interface{}) {
	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameMain)
	errRead := viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	appConfig := config.AppConfig{}
	errUnmarshal := viper.Unmarshal(&appConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	dbConfig := config.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameBucket)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	awsConfig := config.AwsConfig{}
	errUnmarshal = viper.Unmarshal(&awsConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, appConfig)
	result = append(result, dbConfig)
	result = append(result, awsConfig)

	return nil, result
}
