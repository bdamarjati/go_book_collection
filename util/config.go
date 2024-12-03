package util

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBNet         string `mapstructure:"DB_NET"`
	DBAddr        string `mapstructure:"DB_ADDR"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	MySqlSource   string
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
	cfg := mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  config.DBNet,
		Addr:                 config.DBAddr,
		DBName:               config.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	config.MySqlSource = cfg.FormatDSN()
	return
}
