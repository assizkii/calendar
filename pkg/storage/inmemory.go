package storage

import (
	"calendar/pkg/mngtservice"
	"errors"
	"fmt"
	"strings"
	"sync"
)

type EventMemoryStorage struct {
	mx    sync.RWMutex
	store map[int32]*mngtservice.Event
}

func New() EventStorage {
	return &EventMemoryStorage{store: map[int32]*mngtservice.Event{}}
}

func (em *EventMemoryStorage) Get(id int32) (*mngtservice.Event, error) {
	em.mx.RLock()
	defer em.mx.RUnlock()
	event, ok := em.store[id]

	if !ok {
		return nil, fmt.Errorf("could not find id %d", id)
	}

	return event, nil
}

func (em *EventMemoryStorage) Add(e *mngtservice.Event) error {
	if err := em.Validate(e); err != nil {
		return err
	}

	em.mx.Lock()
	defer em.mx.Unlock()

	em.store[e.Id] = e

	return nil
}

func (em *EventMemoryStorage) Update(id int32, e mngtservice.Event) error {
	if _, err := em.Get(id); err != nil {
		return err
	}
	if err := em.Validate(&e); err != nil {
		return fmt.Errorf("validate error: %s", err)
	}

	em.mx.Lock()
	defer em.mx.Unlock()

	em.store[e.Id] = &e

	return nil
}

func (em *EventMemoryStorage) Delete(id int32) error {
	if _, err := em.Get(id); err != nil {
		return err
	}

	em.mx.Lock()
	defer em.mx.Unlock()

	delete(em.store, id)

	return nil
}

func (em *EventMemoryStorage) List() (map[int32]*mngtservice.Event, error) {
	em.mx.RLock()
	defer em.mx.RUnlock()

	return em.store, nil
}

func (em *EventMemoryStorage) Validate(e *mngtservice.Event) error {
	if e.Id == 0 {
		return errors.New("ID field cannot be 0")
	}

	switch "" {
	case strings.TrimSpace(e.Description):
		return errors.New("description field cannot be empty string")
	case strings.TrimSpace(e.Title):
		return errors.New("title field cannot be empty string")
	}

	return nil
}

