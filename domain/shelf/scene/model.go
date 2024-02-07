// Package scene is defined for domain model.
package scene

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for user story or usecase.
type T struct {
	createdAt   *time.Time
	updatedAt   *time.Time
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

// D is a struct for deleting a Scene.
type D struct {
	id vo.ID
}

var _ shelf.Scener = (*T)(nil)

// Description implements request.SceneType.
func (s T) Description() string {
	return s.description.String()
}

// ID implements request.SceneType.
func (s T) ID() vo.ID {
	return s.id
}

// Reason implements request.SceneType.
func (s T) Reason() string {
	return s.reason.String()
}

// Title implements request.SceneType.
func (s T) Title() string {
	return s.title.String()
}

// Update updates a Scene.
func (s *T) Update(
	title string,
	description string,
	reason string,
	parent shelf.ParentOfScene,
) (T, error) {
	if s == nil {
		return T{}, fmt.Errorf("scene is nil")
	}
	return New(
		s.ID().String(),
		title,
		description,
		reason,
		s.bookID.String(),
		parent.ID().String(),
	)
}

// Delete deletes a Scene.
func (s *T) Delete(
	participant shelf.Participanter,
) (D, error) {
	if s == nil {
		return D{}, fmt.Errorf("scene is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: s.id,
	}, nil
}
