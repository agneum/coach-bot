package models

import (
	"time"
)

//go:generate reform

//reform:place
type Place struct {
	ID          int64     `reform:"id,pk" json:"-"`
	Name        string    `reform:"name" json:"name"`
	Address     string    `reform:"address" json:"address"`
	Description string    `reform:"description" json:"description"`
	CreatedAt   time.Time `reform:"created_at" json:"created_at"`
	UpdatedAt   time.Time `reform:"updated_at" json:"updated_at"`
}

func (p *Place) BeforeInsert() error {
	now := time.Now()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = now
	}

	p.CreatedAt = p.CreatedAt.UTC().Truncate(time.Second)
	p.UpdatedAt = p.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}

// BeforeUpdate sets time.Time fields to UTC and marshals meta.
func (p *Place) BeforeUpdate() error {
	p.CreatedAt = p.CreatedAt.UTC().Truncate(time.Second)
	p.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	return nil
}

// AfterFind sets time.Time fields to UTC and unmarshals meta.
func (p *Place) AfterFind() error {
	p.CreatedAt = p.CreatedAt.UTC().Truncate(time.Second)
	p.UpdatedAt = p.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}
