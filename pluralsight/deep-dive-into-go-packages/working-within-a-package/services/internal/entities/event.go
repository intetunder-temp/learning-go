package entities

import "time"

type Event struct {
	ID          int
	Start       *time.Time
	Duration    *time.Duration
	Name        string
	Description string
}
