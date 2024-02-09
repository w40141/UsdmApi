// Package narrative is defined for domain model.
package narrative

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for user T or usecase.
type T struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
	bookID      vo.ID
}

// C is a struct for creating a new Narrative.
type C struct {
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	bookID      vo.ID
}

// D is a struct for deleting a Narrative.
type D struct {
	id vo.ID
}

var _ shelf.Narrativer = (*T)(nil)
