package models

import "example.com/gin/db"

type Registration struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	EventID int64 `json:"event_id"`
}

func RegisterUserForEvent(userID, eventID int64) error {
	query := `
	INSERT INTO registrations (user_id, event_id) VALUES (?, ?)
	`
	_, err := db.DB.Exec(query, userID, eventID)
	return err
}

func CancelUserRegistration(userID, eventID int64) error {
	query := `
	DELETE FROM registrations WHERE user_id = ? AND event_id = ?
	`
	_, err := db.DB.Exec(query, userID, eventID)
	return err
}
