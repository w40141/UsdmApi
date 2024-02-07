// Package scene is defined for domain model.
package scene

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for scene.
type Repository interface {
	Save(scene T) error
	Get(ids []vo.ID) ([]T, error)
	Delete(id vo.ID) error
}
