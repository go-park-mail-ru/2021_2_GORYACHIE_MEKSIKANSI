package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Api"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Application"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization/Orm"
	Api2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart/Api"
	Application2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart/Application"
	Orm2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart/Orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	resPoroto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/proto"
	Api3 "2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Api"
	Application3 "2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Application"
	Orm3 "2021_2_GORYACHIE_MEKSIKANSI/internal/Middleware/Orm"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
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
	defer grpcConnAuth.Close()

	authManager := authProto.NewAuthorizationServiceClient(grpcConnAuth)

	authCtx := context.Background()

	authWrapper := Orm.Wrapper{Conn: authManager, Ctx: authCtx}
	authApp := Application.Authorization{DB: &authWrapper}
	userInfo := Api.UserInfo{
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

	midWrapper := Orm3.Wrapper{Conn: connectionDB}
	midApp := Application3.Middleware{DB: &midWrapper}
	infoMiddleware := Api3.InfoMiddleware{
		Application: &midApp,
		Logger:      logger,
	}
	var _ Interface.MiddlewareAPI = &infoMiddleware

	grpcConnRes, errDial := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if errDial != nil {
		println("GG")
		return nil
	}
	defer grpcConnRes.Close()

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
	defer grpcConnCart.Close()

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
		DB:           &orderWrapper,
		//DBCart:       &cartWrapperOld,
		DBProfile:    &profileWrapper,
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
