package models

import (
	"time"
)

//reform:template
type Template struct {
	ID          int64
	Title       string
	Description string
	Type        string
	Note        string
	CoachID     int64
	PlaceID     int64
	DayOfWeek   int
	StartTime   time.Time
	Duration    time.Duration
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
