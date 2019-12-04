package usecase

import (
	"time"

	"github.com/pamost/calendar/internal/entities"
	"github.com/pamost/calendar/internal/event"
)

type EventUseCase struct {
	eventRepository event.Repository
}

func NewEventUseCase(eventRepository event.Repository) *EventUseCase {
	return &EventUseCase{
		eventRepository: eventRepository,
	}
}

func (e EventUseCase) Add(author, title, description string, start, end time.Time) error {
	eventData := &entities.Event{
		Author:      author,
		Title:       title,
		Description: description,
		Start:       start,
		End:         end,
	}

	return e.eventRepository.Create(eventData)
}

func (e EventUseCase) Get() ([]*entities.Event, error) {
	return e.eventRepository.Read()
}

func (e EventUseCase) Edit(event *entities.Event) error {
	return e.eventRepository.Update(event)
}

func (e EventUseCase) Delete(id string) error {
	return e.eventRepository.Delete(id)
}
