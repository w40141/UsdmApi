// Package user is defined entity models to use this service.
package user

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new User.
func Create(name string) (User, error) {
	id := vo.NewID()
	nameVo, err := vo.NewName(name)
	if err != nil {
		return User{}, err
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	return New(id, nameVo, WithCreatedAt(createdAt), WithUpdatedAt(updatedAt)), nil
}

// Update updates a User.
func (u User) Update(name string) (User, error) {
	nameVo, err := vo.NewName(name)
	if err != nil {
		return User{}, err
	}
	return New(u.id, nameVo, WithCreatedAt(u.createdAt), WithUpdatedAt(time.Now())), nil
}
