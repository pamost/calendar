package entities

import "time"

// Event
type Event struct {
	ID          string
	Author      string
	Title       string
	Description string
	Start       time.Time
	End         time.Time
}
