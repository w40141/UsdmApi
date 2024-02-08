// Package scene is defined for domain model.
package scene

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for user story or usecase.
type T struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	parentID    vo.ID
	id          vo.ID
	bookID      vo.ID
}

// C is a struct for creating a Scene.
type C struct {
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	bookID      vo.ID
	parentID    vo.ID
}

// M is an interface for the Scene.
type M interface {
	ID() vo.ID
	Title() vo.Title
	Description() vo.Sentence
	BookID() vo.ID
	ParentID() vo.ID
}

// D is a struct for deleting a Scene.
type D struct {
	id vo.ID
}

var (
	_ shelf.Scener = (*T)(nil)
	_ M            = (*T)(nil)
	_ M            = (*C)(nil)
)
