package main

import (
	"github.com/assizkii/calendar/internal/adapters/servers/http_server"
	"github.com/assizkii/calendar/internal/domain/usecases"
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


