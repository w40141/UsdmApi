// Package narrative is defined for domain model.
package narrative

import "github.com/w40141/UsdmApi/domain/vo"

// WriteRepository is a repository interface for Narrative.
type WriteRepository interface {
	Create(narrative C) (T, error)
	Update(narrative T) (T, error)
}

// ReadRepository is a repository interface for Narrative.
type ReadRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Narrative.
type DeleteRepository interface {
	Delete(obj D) error
}
