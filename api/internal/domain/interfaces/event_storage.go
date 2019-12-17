package interfaces

import (
	"github.com/assizkii/calendar/api/internal/domain/entities"
	"time"

)

type EventStorage interface {
	Validate(event entities.Event) error
	Get(id string) (entities.Event, error)
	Add(e entities.Event) (string, error)
	Update(id string, e entities.Event) error
	Delete(id string) error
	List() map[string]entities.Event
	FilterByDate(start time.Time) []entities.Event
}

