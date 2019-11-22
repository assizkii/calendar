package interfaces

import (
	"calendar/internal/domain/entities"
)

type EventStorage interface {
	Validate(event entities.Event) error
	Get(id int32) (entities.Event, error)
	Add(e entities.Event) error
	Update(id int32, e entities.Event) error
	Delete(id int32) error
	List() map[int32]entities.Event
}
