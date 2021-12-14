package build

import (
	Application "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/application"
	confPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/myerror"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/orm"
	servicePkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/service"
	promoProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/promocode/proto"
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfType     = "yml"
	ConfPath     = "./internal/microservice/cart/config/"
)

func SetUp(connectionDB ormPkg.ConnectionInterface, logger errPkg.MultiLogger) *servicePkg.CartManager {

	addressCart := "127.0.0.1:8085"
	grpcConnCart, errDialCart := grpc.Dial(
		addressCart,
		grpc.WithInsecure(),
	)
	if errDialCart != nil {
		logger.Errorf("Not connect %s , %s", addressCart, errDialCart.Error())
		return nil
	}
	promoManager := promoProto.NewPromocodeServiceClient(grpcConnCart)
	promoCtx := context.Background()

	cartWrapper := ormPkg.Wrapper{Conn: connectionDB, ConnPromoService: promoManager, Ctx: promoCtx}
	cartManager := Application.Cart{DB: &cartWrapper}
	cartInfo := servicePkg.CartManager{Application: &cartManager}

	var _ servicePkg.CartManagerInterface = &cartInfo

	return &cartInfo
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
	appConfig := confPkg.AppConfig{}
	errUnmarshal := viper.Unmarshal(&appConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	dbConfig := confPkg.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, dbConfig)
	result = append(result, appConfig)

	return nil, result
}
