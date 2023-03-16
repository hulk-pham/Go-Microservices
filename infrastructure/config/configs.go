package config

import (
	"github.com/spf13/viper"

	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Join(filepath.Dir(b), "../")
)

type Config struct {
	AppEnv              string `mapstructure:"APP_ENV"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	AppPort             string `mapstructure:"APP_PORT"`
	RpcPort             string `mapstructure:"RPC_PORT"`
	JwtKey              string `mapstructure:"JWT_KEY"`
	AwsKey              string `mapstructure:"AWS_KEY"`
	AwsSecret           string `mapstructure:"AWS_SECRET"`
	AwsRegion           string `mapstructure:"AWS_REGION"`
	S3BucketName        string `mapstructure:"S3_BUCKET_NAME"`
	AzureStorageAccount string `mapstructure:"AZURE_STORAGE_ACCOUNT"`
	AzureContainerName  string `mapstructure:"AZURE_CONTAINER_NAME"`
	RedisHost           string `mapstructure:"REDIS_HOST"`
	RedisPort           string `mapstructure:"REDIS_PORT"`
}

func LoadAppConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func AppConfig() (config Config) {
	config, err := LoadAppConfig(basepath)
	if err != nil {
		panic("failed to load app config")
	}
	return
}
