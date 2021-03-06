package build

import (
	Application "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/application"
	confPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/myerror"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/orm"
	servicePkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/service"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfType     = "yml"
	ConfPath     = "./internals/microservice/restaurant/config/"
)

func SetUp(connectionDB ormPkg.ConnectionInterface) servicePkg.RestaurantManager {
	authWrapper := ormPkg.Wrapper{Conn: connectionDB}
	authorizationManager := Application.Restaurant{DB: &authWrapper}
	authInfo := servicePkg.RestaurantManager{Application: &authorizationManager}

	var _ servicePkg.RestaurantManagerInterface = &authInfo

	return authInfo
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
	appConfig := confPkg.AppConfig{}
	errUnmarshal := viper.Unmarshal(&appConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Text: errRead.Error(),
		}, nil
	}
	dbConfig := confPkg.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, dbConfig)
	result = append(result, appConfig)

	return nil, result
}
