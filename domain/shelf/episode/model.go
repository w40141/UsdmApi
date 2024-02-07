// Package episode defines the entity object for user episode.
package episode

import (
	"fmt"
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

// Description implements request.Episoder.
func (e T) Description() string {
	return e.description.String()
}

// ID implements request.Episoder.
func (e T) ID() vo.ID {
	return e.id
}

// ParentOfScene implements request.Episoder.
func (*T) ParentOfScene() error {
	panic("unimplemented")
}

// Reason implements request.Episoder.
func (e T) Reason() string {
	return e.reason.String()
}

// Title implements request.Episoder.
func (e T) Title() string {
	return e.title.String()
}

// Update updates a Episode.
func (e *T) Update(
	title string,
	description string,
	reason string,
	story shelf.Storyer,
) (T, error) {
	if e == nil {
		return T{}, fmt.Errorf("episode is nil")
	}
	return New(
		e.id.String(),
		title,
		description,
		reason,
		e.bookID.String(),
		story.ID().String(),
	)
}

// Delete deletes a Episode.
func (e *T) Delete(
	participant shelf.Participanter,
) (D, error) {
	if e == nil {
		return D{}, fmt.Errorf("episode is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: e.id,
	}, nil
}
