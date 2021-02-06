package models

import (
	"time"
)

//go:generate reform

//reform:coach
type Coach struct {
	ID        int64     `reform:"id,pk" json:"-"`
	Name      string    `reform:"name" json:"name"`
	CreatedAt time.Time `reform:"created_at" json:"created_at"`
	UpdatedAt time.Time `reform:"updated_at" json:"updated_at"`
}

func (c *Coach) BeforeInsert() error {
	now := time.Now()
	if c.CreatedAt.IsZero() {
		c.CreatedAt = now
	}
	if c.UpdatedAt.IsZero() {
		c.UpdatedAt = now
	}

	c.CreatedAt = c.CreatedAt.UTC().Truncate(time.Second)
	c.UpdatedAt = c.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}

// BeforeUpdate sets time.Time fields to UTC and marshals meta.
func (c *Coach) BeforeUpdate() error {
	c.CreatedAt = c.CreatedAt.UTC().Truncate(time.Second)
	c.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	return nil
}

// AfterFind sets time.Time fields to UTC and unmarshals meta.
func (c *Coach) AfterFind() error {
	c.CreatedAt = c.CreatedAt.UTC().Truncate(time.Second)
	c.UpdatedAt = c.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}
