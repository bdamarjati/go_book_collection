package util

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBNet      string `mapstructure:"DB_NET"`
	DBAddr     string `mapstructure:"DB_ADDR"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	_, ci := os.LookupEnv("APP_CI")
	if !ci {
		viper.SetConfigFile(path + "/.env")

		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
	} else {
		config.DBDriver = os.Getenv("DB_DRIVER")
		config.DBUser = os.Getenv("DB_USER")
		config.DBPassword = os.Getenv("DB_PASSWORD")
		config.DBAddr = "127.0.0.1:3306"
		config.DBNet = "tcp"
		config.DBName = os.Getenv("DB_NAME")
	}
	return
}
