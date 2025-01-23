package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// var events []Event = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	return nil
}

// GET All Events
func GetAllEvents() ([]Event, error) {
	var events []Event
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) //or exec but exec for changing
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil

}

func GetEventByID(id int64) (*Event, error) {
	var event Event
	query := "SELECT * FROM events WHERE id= ?"

	row := db.DB.QueryRow(query, id) //get one row its safe
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	fmt.Println(err)

	return err

}

func (event Event) DELETE() error {
	query := "DELETE FROM events WHERE id = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.ID)
	return err

}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id,user_id) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err

}
