package main

import (
	"flag"
	"github.com/assizkii/calendar/pkg/calendar"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

//go:generate protoc --go_out=paths=source_relative:. pkg/mngtservice/Event.proto

var (
	appConf  calendar.Config
	confPath string
)

func init() {
	flag.StringVar(&confPath, "config", "", "configs/conf.yaml")
	flag.Parse()
	readConfig()
}

func main() {
	file, err := os.OpenFile(appConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	calendar.RunServer(appConf)
	calendar.GenerateEvents()
}

func readConfig() {

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &appConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
