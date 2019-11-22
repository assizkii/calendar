package http_server

import (
	"html"
	"log"
	"net/http"
)

type Logger struct {
	handler http.Handler
}


func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.handler.ServeHTTP(w, r)

	if r.Method == "POST" {
		r.ParseForm()
		log.Printf("%q , body - %s", html.EscapeString(r.URL.Path), r.PostForm.Encode())
	} else {
		log.Printf("%q , queryParams - %s", html.EscapeString(r.URL.Path), r.URL.Query().Encode())
	}

}
func NewLogger(handlerToWrap http.Handler) *Logger {
	return  &Logger{handlerToWrap}
}