package inmemory_storage

import (
	"errors"
	"fmt"
	"github.com/assizkii/calendar/entities"
	"github.com/assizkii/calendar/api/internal/domain/interfaces"
	"github.com/google/uuid"
	"sort"
	"strings"
	"sync"
	"time"
)

type EventMemoryStorage struct {
	mx    sync.RWMutex
	store map[string]entities.Event
}

func New() interfaces.EventStorage {
	return &EventMemoryStorage{
		mx:    sync.RWMutex{},
		store: make(map[string]entities.Event),
	}
}

func (em *EventMemoryStorage) Get(id string) (entities.Event, error) {
	em.mx.RLock()
	defer em.mx.RUnlock()
	event, ok := em.store[id]

	if !ok {
		return event, fmt.Errorf("could not find id %d", id)
	}

	return event, nil
}

func (em *EventMemoryStorage) Add(e entities.Event) (string, error) {
	if err := em.Validate(e); err != nil {
		return "", err
	}

	em.mx.Lock()
	e.Id = uuid.New().String()
	em.store[e.Id] = e
	defer em.mx.Unlock()

	return e.Id, nil
}

func (em *EventMemoryStorage) Update(id string, e entities.Event) error {
	if _, err := em.Get(id); err != nil {
		return err
	}
	if err := em.Validate(e); err != nil {
		return fmt.Errorf("validate error: %s", err)
	}

	em.mx.Lock()
	defer em.mx.Unlock()

	em.store[e.Id] = e

	return nil
}

func (em *EventMemoryStorage) Delete(id string) error {
	if _, err := em.Get(id); err != nil {
		return err
	}

	em.mx.Lock()
	defer em.mx.Unlock()

	delete(em.store, id)

	return nil
}

func (em *EventMemoryStorage) List() (map[string]entities.Event, error ){
	em.mx.RLock()
	defer em.mx.RUnlock()

	return em.store, nil
}

func (em *EventMemoryStorage) Validate(e entities.Event) error {

	switch "" {
	case strings.TrimSpace(e.Description):
		return errors.New("description field cannot be empty string")
	case strings.TrimSpace(e.Title):
		return errors.New("title field cannot be empty string")
	}

	return nil
}

func (em *EventMemoryStorage) FilterByDate(from time.Time) ([]entities.Event, error) {

	var filteredEvents []entities.Event
	list, _ = em.List()
	for _, event := range list {
		if event.GetStart().GetSeconds() >= from.Unix() {
			filteredEvents = append(filteredEvents, event)
		}
	}

	sort.SliceStable(filteredEvents, func(i, j int) bool {
		return filteredEvents[i].GetStart().Seconds < filteredEvents[j].GetStart().Seconds
	})
	return filteredEvents, nil
}
