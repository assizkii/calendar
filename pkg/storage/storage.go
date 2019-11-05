package storage

import (
	"calendar/pkg/mngtservice"
)

type EventStorage interface {
	Validate(event *mngtservice.Event) error
	Get(id int32) (*mngtservice.Event, error)
	Add(e *mngtservice.Event) error
	Update(id int32, e mngtservice.Event) error
	Delete(id int32) error
	List() (map[int32]*mngtservice.Event, error)
}
