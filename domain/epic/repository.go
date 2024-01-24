// Package epic is defined for domain model.
package epic

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Project.
type Repository interface {
	Create(epic Epic) error
	Update(epic Epic) error
	Get(ids []vo.ID) ([]Epic, error)
	Delete(id vo.ID) error
}
