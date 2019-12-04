package event

import (
	"github.com/pamost/calendar/internal/entities"
)

type Repository interface {
	Create(event *entities.Event) error
	Read() ([]*entities.Event, error)
	Update(event *entities.Event) error
	Delete(id string) error
}
