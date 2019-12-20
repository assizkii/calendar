package http

import (
	"calendar/internal/adapters/storages"
	"net/http"
	"strconv"
	"time"
)

type EventHandler struct{}

type HttpResponse struct {
	status int
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func (h *EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	event, err := prepareRequestData(r)

	if err != nil {
		result := &HttpResponse{http.StatusBadRequest, "", err.Error()}
		showResponse(result, w)
		return
	}

	eventStorage := storages.GetInstance()
	newId, err := eventStorage.Add(event)

	if err != nil {
		result := &HttpResponse{http.StatusBadRequest, "", err.Error()}
		showResponse(result, w)
		return
	}

	result := &HttpResponse{http.StatusOK, "new id " + strconv.Itoa(int(newId)), ""}

	showResponse(result, w)
}

func (h *EventHandler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		result := &HttpResponse{http.StatusBadRequest, "", err.Error()}
		showResponse(result, w)
		return
	}

	event, err := prepareRequestData(r)

	if err != nil {
		result := &HttpResponse{http.StatusOK, "", err.Error()}
		showResponse(result, w)
		return
	}

	eventStorage := storages.GetInstance()
	err = eventStorage.Update(int32(id), event)

	if err != nil {
		result := &HttpResponse{http.StatusOK, "", err.Error()}
		showResponse(result, w)
		return
	}

	result := &HttpResponse{http.StatusOK, "update success " + strconv.Itoa(id), ""}

	showResponse(result, w)
}

func (h *EventHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		result := &HttpResponse{http.StatusBadRequest, "", err.Error()}
		showResponse(result, w)
		return
	}

	eventStorage := storages.GetInstance()
	err = eventStorage.Delete(int32(id))

	if err != nil {
		result := &HttpResponse{http.StatusOK, "", err.Error()}
		showResponse(result, w)
		return
	}

	result := &HttpResponse{http.StatusOK, "delete success " + strconv.Itoa(id), ""}

	showResponse(result, w)
}

func (h *EventHandler) EventForDayHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	timeStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	timeEnd := timeStart.AddDate(1, 0, 0).Add(time.Nanosecond * -1)

	eventList := storages.GetInstance().FilterByDate(timeStart, timeEnd)
	result := &HttpResponse{http.StatusOK, eventList, ""}

	showResponse(result, w)
}

func (h *EventHandler) EventForWeekHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	weekStart := t.Day() - int(t.Weekday())
	timeStart := time.Date(t.Year(), t.Month(), weekStart, 0, 0, 0, 0, time.Local)
	timeEnd := timeStart.AddDate(0, 0, 7).Add(time.Nanosecond * -1)

	eventList := storages.GetInstance().FilterByDate(timeStart, timeEnd)
	result := &HttpResponse{http.StatusOK, eventList, ""}

	showResponse(result, w)
}

func (h *EventHandler) EventForMonthHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	timeStart := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, time.Local)
	timeEnd := timeStart.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	eventList := storages.GetInstance().FilterByDate(timeStart, timeEnd)
	result := &HttpResponse{http.StatusOK, eventList, ""}

	showResponse(result, w)
}
