// Package story defines the entity object for user story.
package story

import "github.com/w40141/UsdmApi/domain/vo"

// WriteRepository is a repository interface for Episode.
type WriteRepository interface {
	Create(story C) (T, error)
	Update(story T) (T, error)
}

// ReadRepository is a repository interface for Episode.
type ReadRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Episode.
type DeleteRepository interface {
	Delete(obj D) error
}
