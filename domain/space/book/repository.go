// Package book is defined for domain model.
package book

import "github.com/w40141/UsdmApi/domain/vo"

// WriteRepository is a repository interface for Book.
type WriteRepository interface {
	Create(book C) (T, error)
	Update(book T) (T, error)
}

// ReadRepository is a repository interface for Book.
type ReadRepository interface {
	Get(ids []vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Book.
type DeleteRepository interface {
	Delete(obj D) error
}
