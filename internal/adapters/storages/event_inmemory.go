package storages

import (
	"calendar/internal/domain/entities"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)


type EventMemoryStorage struct {
	counter int32
	mx    *sync.RWMutex
	store map[int32]entities.Event
}

var instance *EventMemoryStorage
var once sync.Once

func GetInstance() *EventMemoryStorage {
	once.Do(func() {
		instance =  &EventMemoryStorage{
			counter : 1,
			mx   : &sync.RWMutex{},
			store: make(map[int32]entities.Event),
		}
	})
	return instance
}

func (em *EventMemoryStorage) Get(id int32) (entities.Event, error) {
	em.mx.RLock()
	defer em.mx.RUnlock()
	event, ok := em.store[id]

	if !ok {
		return event, fmt.Errorf("could not find id %d", id)
	}

	return event, nil
}

func (em *EventMemoryStorage) Add(e entities.Event) (int32, error) {
	if err := em.Validate(e); err != nil {
		return 0, err
	}

	em.mx.Lock()
	e.Id = em.GetNewId()
	em.store[e.Id] = e
	defer em.mx.Unlock()

	return e.Id, nil
}

func (em *EventMemoryStorage) Update(id int32, e entities.Event) error {
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

func (em *EventMemoryStorage) Delete(id int32) error {
	if _, err := em.Get(id); err != nil {
		return err
	}

	em.mx.Lock()
	defer em.mx.Unlock()

	delete(em.store, id)

	return nil
}

func (em *EventMemoryStorage) List() map[int32]entities.Event {
	em.mx.RLock()
	defer em.mx.RUnlock()

	return em.store
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

func (em *EventMemoryStorage) FilterByDate(from time.Time, to time.Time) []entities.Event {

	var filteredEvents []entities.Event

	for _, event := range em.List() {
		if  event.GetFrom().GetSeconds() >= from.Unix() && event.GetFrom().GetSeconds() <= to.Unix() ||
			event.GetTo().GetSeconds() >= from.Unix() && event.GetTo().GetSeconds() <= to.Unix() {
			filteredEvents = append(filteredEvents, event)
		}
	}

	sort.SliceStable(filteredEvents, func(i, j int) bool {
		return filteredEvents[i].GetFrom().Seconds < filteredEvents[j].GetFrom().Seconds
	})
	return filteredEvents
}


func (em *EventMemoryStorage) GetNewId() int32 {
	id := em.counter
	em.counter++
	return id
}

