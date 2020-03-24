package data

import (
	"github.com/intetunder-temp/learning-go/pluralsight/deep-dive-into-go-packages/working-within-a-package/services/internal/entities"
)

// GetAll returns all of the events that the system is tracking.
func GetAll() ([]entities.Event, error) {
	return []entities.Event{}, nil
}

// GetByID returns the requested event.
// If an event with the requested ID does not exist, then an error is returned.
func GetByID(id int) (entities.Event, error) {
	return entities.Event{}, nil
}
