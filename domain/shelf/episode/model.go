// Package episode defines the entity object for user episode.
package episode

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for Episode.
type T struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	storyID     vo.ID
	id          vo.ID
	bookID      vo.ID
}

// C is a struct for creating a new Episode.
type C struct {
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	storyID     vo.ID
	bookID      vo.ID
}

// D is a struct for deleting a Episode.
type D struct {
	id vo.ID
}

var _ shelf.Episoder = (*T)(nil)
