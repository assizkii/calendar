package servers

import (
	"calendar/internal/adapters/configs"
	"fmt"
	"html"
	"log"
	"net/http"
)


func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func RunHttpServer(appConf configs.HttpConfig) {
	http.HandleFunc("/", HomeRouterHandler)       // установим роутер
	err := http.ListenAndServe(appConf.Host, nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
