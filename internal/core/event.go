package core

import "time"

type Event struct {
	EventType string
	Time      time.Time
	UserIP    string
}
