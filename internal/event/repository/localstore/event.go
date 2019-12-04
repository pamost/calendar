package localstore

import (
	"sync"

	"github.com/pamost/calendar/internal/entities"
	"github.com/pamost/calendar/internal/event"
)

type EventLocalStore struct {
	events map[string]*entities.Event
	mutex  *sync.Mutex
}

func NewEventLocalStore() *EventLocalStore {
	return &EventLocalStore{
		events: make(map[string]*entities.Event),
		mutex:  new(sync.Mutex),
	}
}

// Create event in local storage
func (els *EventLocalStore) Create(e *entities.Event) error {
	els.mutex.Lock()
	els.events[e.ID] = e
	els.mutex.Unlock()

	return nil
}

// Read event from local storage
func (els *EventLocalStore) Read() ([]*entities.Event, error) {
	store := make([]*entities.Event, 0)

	els.mutex.Lock()
	for _, e := range els.events {
		store = append(store, e)
	}
	els.mutex.Unlock()

	return store, nil
}

// Update event in local storage
func (els *EventLocalStore) Update(e *entities.Event) error {
	els.mutex.Lock()
	defer els.mutex.Unlock()

	if _, ok := els.events[e.ID]; ok {
		els.events[e.ID].Author = e.Author
		els.events[e.ID].Title = e.Title
		els.events[e.ID].Description = e.Description
		els.events[e.ID].Start = e.Start
		els.events[e.ID].End = e.End

		return nil
	}

	return event.ErrEventNotFound
}

// Delete event from local storage
func (els *EventLocalStore) Delete(id string) error {
	els.mutex.Lock()
	defer els.mutex.Unlock()

	if _, ok := els.events[id]; ok {
		delete(els.events, id)
		return nil
	}

	return event.ErrEventNotFound
}
