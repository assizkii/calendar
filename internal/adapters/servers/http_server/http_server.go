package http_server

import (
	"log"
	"net/http"
)


func RunHttpServer(appConf HttpConfig) {

	handler := &EventHandler{}

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", handler.CreateEventHandler)
	mux.HandleFunc("/update_event", handler.UpdateEventHandler)
	mux.HandleFunc("/delete_event", handler.DeleteEventHandler)
	mux.HandleFunc("/events_for_day", handler.EventForDayHandler)
	mux.HandleFunc("/events_for_week", handler.EventForWeekHandler)
	mux.HandleFunc("/events_for_month", handler.EventForMonthHandler)

	wrappedMux := NewLogger(mux)

	log.Printf("server is listening at %s", appConf.Host)

	err := http.ListenAndServe(appConf.Host, wrappedMux)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
