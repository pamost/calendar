package event

import (
	"time"

	"github.com/pamost/calendar/internal/entities"
)

type UseCase interface {
	Add(author, title, description string, start, end time.Time) error
	Get() ([]*entities.Event, error)
	Edit(event *entities.Event) error
	Delete(id string) error
}
