// Package story defines the entity object for user story.
package story

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for Story.
type T struct {
	createAt    *time.Time
	updateAt    *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
	bookID      vo.ID
	parentID    vo.ID
}

// C is a struct for creating a new Story.
type C struct {
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	bookID      vo.ID
	parentID    vo.ID
}

// D is a struct for deleting a Story.
type D struct {
	id vo.ID
}

var _ shelf.Storyer = (*T)(nil)
