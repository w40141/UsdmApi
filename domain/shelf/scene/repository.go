// Package scene is defined for domain model.
package scene

import "github.com/w40141/UsdmApi/domain/vo"

// WriteRepository is a repository interface for Scene.
type WriteRepository interface {
	Save(scene M) (T, error)
}

// ReadRepository is a repository interface for Scene.
type ReadRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Scene.
type DeleteRepository interface {
	Delete(obj D) error
}
