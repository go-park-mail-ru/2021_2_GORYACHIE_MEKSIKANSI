package build

import (
	Application "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/application"
	confPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/myerror"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/orm"
	servicePkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/service"
	"github.com/spf13/viper"
)

const (
	ConfNameMain   = "main"
	ConfNameDB = "database"
	ConfType   = "yml"
	ConfPath   = "./internal/microservice/cart/config/"
)

func SetUp (connectionDB ormPkg.ConnectionInterface) servicePkg.CartManager {
	authWrapper := ormPkg.Wrapper{Conn: connectionDB}
	authorizationManager := Application.Cart{DB: &authWrapper}
	authInfo := servicePkg.CartManager{Application: &authorizationManager}

	var _ servicePkg.CartManagerInterface = &authInfo

	return authInfo
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
