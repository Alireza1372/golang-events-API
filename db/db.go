package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(5)

	createTable()

}

func createTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users Table")
	}

	createEventsTable := ` 
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NUll,
	description TEXT NOT NUll,
	location TEXT NOT NUll,
	dateTime DATETIME NOT NUll,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events Table")
	}

	createRegistrationTable := `
CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)

	)
`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("could not create registrations Table")
	}

}
