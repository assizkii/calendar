package calendar

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Config struct {
	Host     string `yaml:"http_listen"`
	LogFile  string `yaml:"log_file"`
	LogLevel string `yaml:"log_level"`
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func RunServer(appConf Config) {
	http.HandleFunc("/", HomeRouterHandler)       // установим роутер
	err := http.ListenAndServe(appConf.Host, nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
