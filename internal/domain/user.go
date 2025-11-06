package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserEvent struct {
	EventType string
	UserID    string
	Name      string
	Email     string
	Time      time.Time
}
