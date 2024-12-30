package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// DB by using this var you can interact with the db in other files
var DB *sql.DB

func InitDB() {
	var err error
	// api.db is a local path. It would be created automatically when
	// this file does not exist.
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	// how many pools(open connections) we can have to this db
	DB.SetMaxOpenConns(10)
	// how many connections to keep open when no one uses the app.
	DB.SetMaxIdleConns(5)

	// use the snippet "dbCreateTable".
	createTables()
}

func createTables() {
	createUsersTables := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,                  
	    password TEXT NOT NULL
)`
	_, err := DB.Exec(createUsersTables)
	if err != nil {
		panic("Could not create user table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
	    FOREIGN KEY(user_id) REFERENCES users(id)
)`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
}
