package configs

import "github.com/spf13/viper"

type Config struct {
	DBSource     string `mapstructure:"DB_SOURCE"`
	AppPort      string `mapstructure:"APP_PORT"`
	JwtKey       string `mapstructure:"JWT_KEY"`
	AwsKey       string `mapstructure:"AWS_KEY"`
	AwsSecret    string `mapstructure:"AWS_SECRET"`
	AwsRegion    string `mapstructure:"AWS_REGION"`
	S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
}

func LoadAppConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
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
	config, err := LoadAppConfig(".")
	if err != nil {
		panic("failed to local config")
	}
	return
}
