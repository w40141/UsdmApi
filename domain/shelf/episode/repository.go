// Package episode is defined for domain model.
package episode

import "github.com/w40141/UsdmApi/domain/vo"

// WriteRepository is a repository interface for Episode.
type WriteRepository interface {
	Create(episode C) (T, error)
	Update(episode T) (T, error)
}

// ReadRepository is a repository interface for Episode.
type ReadRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Episode.
type DeleteRepository interface {
	Delete(obj D) error
}
