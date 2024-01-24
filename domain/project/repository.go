// Package project is defined for domain model.
package project

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Project.
type Repository interface {
	Create(project Project) error
	Update(project Project) error
	Get(ids []vo.ID) ([]Project, error)
	Delete(id vo.ID) error
}
