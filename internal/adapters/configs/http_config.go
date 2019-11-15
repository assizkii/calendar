package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
	return appConf
}