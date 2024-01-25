// Package user is defined entity models to use this service.
package user

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for User.
type Repository interface {
	Save(user User) error
	Get(ids []vo.ID) ([]User, error)
	Delete(id vo.ID) error
}
