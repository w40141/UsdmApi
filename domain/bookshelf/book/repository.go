// Package book is defined for domain model.
package book

import "github.com/w40141/UsdmApi/domain/vo"

// CreateRepository is a repository interface for Book.
type CreateRepository interface {
	Save(book C) (T, error)
}

// WriteRepository is a repository interface for Book.
type WriteRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Book.
type DeleteRepository interface {
	Delete(obj D) error
}
