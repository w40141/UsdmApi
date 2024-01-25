// Package user is defined entity models to use this service.
package user

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new User.
func Create(name string) (User, error) {
	id := vo.NewID().String()
	return New(
		id,
		name,
		WithCreatedAt(time.Now()),
		WithUpdatedAt(time.Now()),
	)
}

// Update updates a User.
func (u User) Update(name string) (User, error) {
	return New(
		u.id.String(),
		name,
		WithUpdatedAt(time.Now()),
	)
}
