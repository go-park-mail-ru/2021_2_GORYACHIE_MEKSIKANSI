package config

type AppConfig struct {
	Primary PrimaryConfig `mapstructure:"primary"`
}

type PrimaryConfig struct {
	Host    string
	Port    string
	Network string
}

type Database struct {
	Host       string
	Port       string
	UserName   string
	Password   string
	SchemaName string
}

type DBConfig struct {
	Db Database `mapstructure:"db"`
}
