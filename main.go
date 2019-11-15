package main

import (
	"calendar/internal/adapters/configs"
	"calendar/internal/adapters/servers"
	"calendar/internal/domain/usecases"
	"flag"
	"log"
	"os"
)

//go:generate protoc --go_out=paths=source_relative:. internal/domain/entities/Event.proto

var confPath string

func init() {
	flag.StringVar(&confPath, "config", "", "configs/conf.yaml")
	flag.Parse()

}

func main() {

	appConf := configs.ReadHttpConfig(confPath)

	file, err := os.OpenFile(appConf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	servers.RunHttpServer(appConf)
	usecases.GenerateEvents()
}


