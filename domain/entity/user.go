// Package entity is defined entity models to use this service.
package entity

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// User is a entity model for user.
type User struct {
	id        vo.ID
	name      vo.Name
	createdAt time.Time
	updatedAt time.Time
}

// NewUser creates a new User.
func NewUser(
	id vo.ID,
	name vo.Name,
	options ...UserOption,
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

// UserOption is a functional option for User.
type UserOption func(*User)

// WithCreatedAt is a functional option for User.
func WithCreatedAt(createdAt time.Time) UserOption {
	return func(u *User) {
		u.createdAt = createdAt
	}
}

// WithUpdatedAt is a functional option for User.
func WithUpdatedAt(updatedAt time.Time) UserOption {
	return func(u *User) {
		u.updatedAt = updatedAt
	}
}

// CreateUser creates a new User.
func CreateUser(name vo.Name) User {
	id := vo.NewID()
	createdAt := time.Now()
	updatedAt := time.Now()
	return NewUser(id, name, WithCreatedAt(createdAt), WithUpdatedAt(updatedAt))
}

// Update updates a User.
func (u User) Update(name vo.Name) User {
	return NewUser(u.id, name, WithCreatedAt(u.createdAt), WithUpdatedAt(time.Now()))
}

// UserRepository is a repository interface for User.
type UserRepository interface {
	Create(user User) error
	Update(user User) error
	Get(ids []vo.ID) ([]User, error)
	Delete(id vo.ID) error
}
