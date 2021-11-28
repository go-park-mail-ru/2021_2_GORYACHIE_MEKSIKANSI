package config

type AppConfig struct {
	Port    string
	Cors    CorsConfig    `mapstructure:"cors"`
	Primary PrimaryConfig `mapstructure:"primary"`
}

type CorsConfig struct {
	Host string
	Port string
}

type PrimaryConfig struct {
	Debug      bool
	DeleteLogs bool
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

type AwsBucket struct {
	Region          string
	AccessKeyId     string
	SecretAccessKey string
	Name            string
	Endpoint        string
}

type AwsConfig struct {
	Aws AwsBucket `mapstructure:"aws"`
}

type MicroserviceConfig struct {
	Authorization AuthorizationMicroservice `mapstructure:"authorization"`
	Cart CartMicroservice `mapstructure:"cart"`
	Restaurant RestaurantMicroservice `mapstructure:"restaurant"`
}

type AuthorizationMicroservice struct {
	Host      string
	Port      string
	Network   string
}

type CartMicroservice struct {
	Host      string
	Port      string
	Network   string
}

type RestaurantMicroservice struct {
	Host      string
	Port      string
	Network   string
}


