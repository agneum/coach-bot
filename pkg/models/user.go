package models

import (
	"time"
)

//go:generate reform

// User represents a row in user table.
//reform:user
type User struct {
	ID         int64     `reform:"id,pk" json:"-"`
	TelegramID string    `reform:"tg_id" json:"-"`
	FirstName  string    `reform:"first_name" json:"first_name"`
	LastName   string    `reform:"last_name" json:"last_name"`
	Username   string    `reform:"username" json:"-"`
	Bio        string    `reform:"bio" json:"bio"`
	CreatedAt  time.Time `reform:"created_at" json:"-"`
	UpdatedAt  time.Time `reform:"updated_at" json:"-"`
}

func (u *User) BeforeInsert() error {
	now := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = now
	}

	u.CreatedAt = u.CreatedAt.UTC().Truncate(time.Second)
	u.UpdatedAt = u.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}

// BeforeUpdate sets time.Time fields to UTC and marshals meta.
func (u *User) BeforeUpdate() error {
	u.CreatedAt = u.CreatedAt.UTC().Truncate(time.Second)
	u.UpdatedAt = time.Now().UTC().Truncate(time.Second)

	return nil
}

// AfterFind sets time.Time fields to UTC and unmarshals meta.
func (u *User) AfterFind() error {
	u.CreatedAt = u.CreatedAt.UTC().Truncate(time.Second)
	u.UpdatedAt = u.UpdatedAt.UTC().Truncate(time.Second)

	return nil
}
