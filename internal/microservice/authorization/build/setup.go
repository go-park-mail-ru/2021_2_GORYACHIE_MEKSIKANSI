package build

import (
	Application "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/application"
	confPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/myerror"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/orm"
	servicePkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/service"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfType     = "yml"
	ConfPath     = "./internal/microservice/authorization/config/"
)

func SetUp(connectionDB ormPkg.ConnectionInterface) servicePkg.AuthorizationManager {
	authWrapper := ormPkg.Wrapper{Conn: connectionDB}
	authorizationManager := Application.AuthorizationApplication{DB: &authWrapper}
	authInfo := servicePkg.AuthorizationManager{Application: &authorizationManager}

	var _ servicePkg.AuthorizationManagerInterface = &authInfo

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
