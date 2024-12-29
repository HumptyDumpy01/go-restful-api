package models

import (
	"time"
)

type Event struct {
	ID          float64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      float64
}

var events = []Event{}

func (e Event) Save() {
	// TODO: Add it to the database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
