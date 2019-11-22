package http_server

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type HttpConfig struct {
	Host     string `yaml:"http_listen"`
	LogFile  string `yaml:"log_file"`
	LogLevel string `yaml:"log_level"`
}

func ReadHttpConfig(confPath string) HttpConfig {

	var appConf  = HttpConfig{}
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &appConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	setLogFile(appConf)

	return appConf
}

func setLogFile(appConf HttpConfig)  {
	file, err := os.OpenFile(appConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
}