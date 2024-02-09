// Package book is defined for domain model.
package book

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// T is the largest unit of human management.
type T struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	id          vo.ID
}

// C is a struct for creating a new Book.
type C struct {
	title       vo.Title
	description vo.Sentence
}

// D is a struct for deleting a Book.
type D struct {
	id vo.ID
}
