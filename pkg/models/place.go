package models

import (
	"time"
)

//reform:place
type Place struct {
	ID          int64
	Name        string
	Address     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
