package models

import (
	"HumptyDumpy01/go-restful-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      float64
}

var events = []Event{}

func (e Event) Save() (int64, error) {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)  
	VALUES (?, ?, ? ,? ,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	e.ID = id
	return id, nil

	//events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
