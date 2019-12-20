package http

import (
	"encoding/json"
	"calendar/internal/domain/entities"
	"github.com/golang/protobuf/ptypes"
	"net/http"
	"time"
)

func showResponse(result *HttpResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(status)
	w.Write(response)
}

func prepareRequestData(r *http.Request) (entities.Event, error) {
	var event entities.Event
	event.Title = r.FormValue("title")
	event.Description = r.FormValue("description")

	timeStart, err := time.Parse("2006-01-02", r.FormValue("start"))
	if err != nil {
		return event, err
	}

	timeEnd, err := time.Parse("2006-01-02", r.FormValue("end"))
	if err != nil {
		return event, err
	}

	event.From, err =  ptypes.TimestampProto(timeStart)
	if err != nil {
		return event, err
	}

	event.To, err =  ptypes.TimestampProto(timeEnd)
	if err != nil {
		return event, err
	}

	return event, nil
}


