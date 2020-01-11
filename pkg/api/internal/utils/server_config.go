package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type ServerConfig struct {
	Host       string `yaml:"host"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`
	DbHost     string `yaml:"db_host"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

func GetAppConfig() ServerConfig {

	var appConfig ServerConfig

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	setLogFile(appConfig)
	return appConfig
}

func setLogFile(appConf ServerConfig) {
	file, err := os.OpenFile(appConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
}
