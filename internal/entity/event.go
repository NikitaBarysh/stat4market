package entity

import "time"

type Event struct {
	EventID   int64     `json:"event_id"`
	EventType string    `json:"eventType"`
	UserID    int64     `json:"userID"`
	EventTime time.Time `json:"eventTime"`
	Payload   string    `json:"payload"`
}
