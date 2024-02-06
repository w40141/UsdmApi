// Package narrative is defined for domain model.
package narrative

import "github.com/w40141/UsdmApi/domain/vo"

// CreateRepository is a repository interface for Narrative.
type CreateRepository interface {
	Save(narrative C) (T, error)
}

// WriteRepository is a repository interface for Narrative.
type WriteRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Narrative.
type DeleteRepository interface {
	Delete(obj D) error
}
