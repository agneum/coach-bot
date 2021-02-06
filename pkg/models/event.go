package models

import (
	"time"
)

//go:generate reform

const (
	DefaultDuration = 2 * time.Hour

	ClassicType = "classic"
	BeachType   = "beach"
	PartyType   = "party"

	NoteSelf     = "self"
	NoteWomanNet = "woman_net"
	NoteGame     = "game"
)

//reform:event
type Event struct {
	ID          int64
	Title       string
	Description string
	Type        string
	Note        string
	CoachID     int64
	PlaceID     int64
	StartDate   time.Time
	Duration    time.Duration
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
