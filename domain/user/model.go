// Package user is defined entity models to use this service.
package user

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// User is a entity model for user.
type User struct {
	createdAt time.Time
	updatedAt time.Time
	name      vo.Name
	id        vo.ID
}

// New creates a new User.
func New(
	id vo.ID,
	name vo.Name,
	options ...Option,
) User {
	u := User{
		id:   id,
		name: name,
	}
	for _, option := range options {
		option(&u)
	}
	return u
}

// Option is a functional option for User.
type Option func(*User)

// WithCreatedAt is a functional option for User.
func WithCreatedAt(createdAt time.Time) Option {
	return func(u *User) {
		u.createdAt = createdAt
	}
}

// WithUpdatedAt is a functional option for User.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(u *User) {
		u.updatedAt = updatedAt
	}
}
