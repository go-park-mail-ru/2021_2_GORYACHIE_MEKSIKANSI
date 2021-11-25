package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"github.com/spf13/viper"
)

const (
	ConfNameDB = "database"
	ConfType   = "yml"
	ConfPath   = "./config/"
)

func InitConfig() (error, []interface{}) {
	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameDB)
	errRead := viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	dbConfig := config.DBConfig{}
	errUnmarshal := viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, dbConfig)

	return nil, result
}
