package models

import (
	"time"
)

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	// TODO: Add it to the database
	events = append(events, e)
}
