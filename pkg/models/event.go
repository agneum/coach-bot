package models

import (
	"time"
)

type Event struct {
	ID           int64
	Title        string
	Description  string
	TypeActivity string
	StartDate    time.Time
	Duration     time.Duration
	CreatedAt    time.Time
}
