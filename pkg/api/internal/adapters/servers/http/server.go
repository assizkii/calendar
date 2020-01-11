package http

import (
	"log"
	"net/http"
)


func Run(appConf configs.ServerConfig) {

	handler := &EventHandler{}

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)
	mux.HandleFunc("/events_for_day", EventForDayHandler)
	mux.HandleFunc("/events_for_week", EventForWeekHandler)
	mux.HandleFunc("/events_for_month", EventForMonthHandler)

	wrappedMux := NewLogger(mux)

	log.Printf("server is listening at %s", appConf.Host)

	err := http.ListenAndServe(appConf.Host, wrappedMux)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
