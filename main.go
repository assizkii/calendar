package main

import (
	"calendar/internal/adapters/servers/http_server"
	"calendar/internal/domain/usecases"
	"flag"
)

//go:generate protoc --go_out=paths=source_relative:. internal/domain/entities/Event.proto

var confPath string

func init() {
	flag.StringVar(&confPath, "config", "", "configs/conf.yaml")
	flag.Parse()

}

func main() {

	appConf := http_server.ReadHttpConfig("configs/conf.yaml")

	http_server.RunHttpServer(appConf)

	usecases.GenerateEvents()
}


