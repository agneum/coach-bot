package models

import (
	"time"
)

//go:generate reform

//reform:template
type Template struct {
	ID          int64     `reform:"id,pk" json:"-"`
	Title       string    `reform:"title" json:"title"`
	Description string    `reform:"description" json:"description"`
	Type        string    `reform:"type" json:"type"`
	Note        string    `reform:"note" json:"note"`
	CoachID     int64     `reform:"coach_id" json:"coach_id"`
	PlaceID     int64     `reform:"place_id" json:"place_id"`
	Weekday     int       `reform:"weekday" json:"weekday"`
	StartTime   string    `reform:"start_time" json:"start_time"`
	Duration    string    `reform:"duration" json:"duration"`
	CreatedAt   time.Time `reform:"created_at" json:"created_at"`
	UpdatedAt   time.Time `reform:"updated_at" json:"updated_at"`
}

func (t *Template) BeforeInsert() error {
	now := time.Now()
	if t.CreatedAt.IsZero() {
		t.CreatedAt = now
	}
	if t.UpdatedAt.IsZero() {
		t.UpdatedAt = now
	}

	t.CreatedAt = t.CreatedAt.UTC().Truncate(time.Second)
	t.UpdatedAt = t.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}

// BeforeUpdate sets time.Time fields to UTC and marshals meta.
func (t *Template) BeforeUpdate() error {
	t.CreatedAt = t.CreatedAt.UTC().Truncate(time.Second)
	t.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	return nil
}

// AfterFind sets time.Time fields to UTC and unmarshals meta.
func (t *Template) AfterFind() error {
	t.CreatedAt = t.CreatedAt.UTC().Truncate(time.Second)
	t.UpdatedAt = t.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}
