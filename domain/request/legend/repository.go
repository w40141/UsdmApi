// Package legend is defined for domain model.
package legend

import "github.com/w40141/UsdmApi/domain/vo"

// CreateRepository is a repository interface for Legend.
type CreateRepository interface {
	Save(legend C) (T, error)
}

// WriteRepository is a repository interface for Legend.
type WriteRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Legend.
type DeleteRepository interface {
	Delete(obj D) error
}
