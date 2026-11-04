package models

import (
	"example/rest-api/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

var events = []Event{}

func GetAllEvents() ([]Event, error) {

	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Error querying events:", err)
		return nil, err
	}

	defer rows.Close()
	events = []Event{}

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		events = append(events, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	var event Event
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	return &event, err
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID")
		return err
	}
	e.ID = id

	return nil
}

func UpdateEvent(event *Event) (int64, error) {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, err
	}

	return rowsAffected, nil
}
