package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/api"
	authApiPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/api"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/application"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm"
	Api2 "2021_2_GORYACHIE_MEKSIKANSI/internals/cart/api"
	cartApiPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/cart/api"
	Application2 "2021_2_GORYACHIE_MEKSIKANSI/internals/cart/application"
	Orm2 "2021_2_GORYACHIE_MEKSIKANSI/internals/cart/orm"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
	Api3 "2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/api"
	midlApiPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/api"
	Application3 "2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/application"
	Orm3 "2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/orm"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	Api4 "2021_2_GORYACHIE_MEKSIKANSI/internals/order/api"
	orderApiPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/order/api"
	Application4 "2021_2_GORYACHIE_MEKSIKANSI/internals/order/application"
	Orm4 "2021_2_GORYACHIE_MEKSIKANSI/internals/order/orm"
	Api5 "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/api"
	profileApiPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/api"
	Application5 "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/application"
	Orm5 "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm"
	profileOrmPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm"
	Api6 "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/api"
	resApiPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/api"
	Application6 "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/application"
	Orm6 "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/orm"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	ConfNameMain         = "main"
	ConfNameDB           = "database"
	ConfNameBucket       = "bucket"
	ConfNameMicroservice = "microservice"
	ConfType             = "yml"
	ConfPath             = "./config/"
)

type installSetUp struct {
	User       api.UserInfo
	Profile    Api5.InfoProfile
	Midle      Api3.InfoMiddleware
	Restaraunt Api6.InfoRestaurant
	Cart       Api2.InfoCart
	Order      Api4.InfoOrder
}

func SetUp(connectionDB profileOrmPkg.ConnectionInterface, logger errPkg.MultiLogger,
	uploader *s3manager.Uploader, nameBucket string, microserviceConfig config.MicroserviceConfig, intCh chan authPkg.WebSocketOrder) *installSetUp {

	addressAuth := microserviceConfig.Authorization.Host + ":" + microserviceConfig.Authorization.Port
	grpcConnAuth, errDialAuth := grpc.Dial(
		addressAuth,
		grpc.WithInsecure(),
	)
	if errDialAuth != nil {
		logger.Errorf("Not connect %s , %s", addressAuth, errDialAuth.Error())
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
	var _ authApiPkg.AuthorizationApiInterface = &userInfo

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
	var _ profileApiPkg.ProfileApiInterface = &profileInfo

	countInternalServer := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "countInternalServer",
		Help: "Number internal processed",
	})
	hitsUrl := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hitsUrlApi",
		Help: "Number connect url",
	}, []string{"path"})
	timingUrl := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "timingUrlApi",
		Help: "request execution time",
	}, []string{"time", "path"})

	prometheus.MustRegister(countInternalServer, hitsUrl, timingUrl)

	midWrapper := Orm3.Wrapper{DBConn: connectionDB, Conn: authManager, Ctx: authCtx}
	midApp := Application3.Middleware{DB: &midWrapper}
	infoMiddleware := Api3.InfoMiddleware{
		Application:         &midApp,
		Logger:              logger,
		CountInternalMetric: countInternalServer,
		HitsMetric:          hitsUrl,
		TimingMetric:        timingUrl,
	}
	var _ midlApiPkg.MiddlewareApiInterface = &infoMiddleware

	addressRes := microserviceConfig.Restaurant.Host + ":" + microserviceConfig.Restaurant.Port
	grpcConnRes, errDialRes := grpc.Dial(
		addressRes,
		grpc.WithInsecure(),
	)
	if errDialRes != nil {
		logger.Errorf("Not connect %s , %s", addressRes, errDialRes.Error())
		return nil
	}
	resManager := resProto.NewRestaurantServiceClient(grpcConnRes)
	resCtx := context.Background()

	restWrapper := Orm6.Wrapper{Conn: resManager, Ctx: resCtx}
	restApp := Application6.Restaurant{DB: &restWrapper}
	restaurantInfo := Api6.InfoRestaurant{
		Application: &restApp,
		Logger:      logger,
	}
	var _ resApiPkg.RestaurantApiInterface = &restaurantInfo

	addressCart := microserviceConfig.Cart.Host + ":" + microserviceConfig.Cart.Port
	grpcConnCart, errDialCart := grpc.Dial(
		addressCart,
		grpc.WithInsecure(),
	)
	if errDialCart != nil {
		logger.Errorf("Not connect %s , %s", addressCart, errDialCart.Error())
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
	var _ cartApiPkg.CartApiInterface = &cartInfo

	orderWrapper := Orm4.Wrapper{Conn: connectionDB, ConnService: cartManager, Ctx: cartCtx}
	orderApp := Application4.Order{
		DB:        &orderWrapper,
		DBProfile: &profileWrapper,
		IntCh:     intCh,
	}
	orderInfo := Api4.InfoOrder{
		Application: &orderApp,
		Logger:      logger,
	}
	var _ orderApiPkg.OrderApiInterface = &orderInfo

	var result installSetUp
	result.User = userInfo
	result.Cart = cartInfo
	result.Profile = profileInfo
	result.Midle = infoMiddleware
	result.Restaraunt = restaurantInfo
	result.Order = orderInfo

	return &result
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
			Text: errNewSess.Error(),
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
			Text: errRead.Error(),
		}, nil
	}
	appConfig := config.AppConfig{}
	errUnmarshal := viper.Unmarshal(&appConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameMicroservice)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Text: errRead.Error(),
		}, nil
	}
	microserviceConfig := config.MicroserviceConfig{}
	errUnmarshal = viper.Unmarshal(&microserviceConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Text: errRead.Error(),
		}, nil
	}
	dbConfig := config.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameBucket)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Text: errRead.Error(),
		}, nil
	}
	awsConfig := config.AwsConfig{}
	errUnmarshal = viper.Unmarshal(&awsConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, appConfig)
	result = append(result, dbConfig)
	result = append(result, awsConfig)
	result = append(result, microserviceConfig)

	return nil, result
}
