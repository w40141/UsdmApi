// Package task is defined entity models to use this service.
package task

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository for Task.
type Repository interface {
	Save(task Task) error
	Get(ids []vo.ID) ([]Task, error)
	Delete(id vo.ID) error
}
