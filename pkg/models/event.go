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
	ID          int64     `reform:"id,pk" json:"-"`
	Title       string    `reform:"title" json:"title"`
	Description string    `reform:"description" json:"description"`
	Type        string    `reform:"type" json:"type"`
	Note        string    `reform:"note" json:"note"`
	CoachID     int64     `reform:"coach_id" json:"coach_id"`
	PlaceID     int64     `reform:"place_id" json:"place_id"`
	StartDate   time.Time `reform:"start_date" json:"start_date"`
	Duration    string    `reform:"duration" json:"duration"`
	CreatedAt   time.Time `reform:"created_at" json:"created_at"`
	UpdatedAt   time.Time `reform:"updated_at" json:"updated_at"`
}

func (e *Event) BeforeInsert() error {
	now := time.Now()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = now
	}
	if e.UpdatedAt.IsZero() {
		e.UpdatedAt = now
	}

	e.CreatedAt = e.CreatedAt.UTC().Truncate(time.Second)
	e.UpdatedAt = e.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}

// BeforeUpdate sets time.Time fields to UTC and marshals meta.
func (e *Event) BeforeUpdate() error {
	e.CreatedAt = e.CreatedAt.UTC().Truncate(time.Second)
	e.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	return nil
}

// AfterFind sets time.Time fields to UTC and unmarshals meta.
func (e *Event) AfterFind() error {
	e.CreatedAt = e.CreatedAt.UTC().Truncate(time.Second)
	e.UpdatedAt = e.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}
