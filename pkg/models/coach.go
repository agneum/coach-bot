package models

import (
	"time"
)

//reform:coach
type Coach struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
