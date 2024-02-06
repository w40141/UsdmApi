// Package narrative is defined for domain model.
package narrative

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/shelf/participant"
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

var (
	_ shelf.ParentOfStory = (*T)(nil)
	_ shelf.Narrativer    = (*T)(nil)
)

// ParentOfScene implements request.Narrativer.
func (*T) ParentOfScene() error {
	panic("unimplemented")
}

// ParentOfStory implements story.ParentOfStory.
func (*T) ParentOfStory() error {
	panic("unimplemented")
}

// ID implements Narrativer.
func (n T) ID() vo.ID {
	return n.id
}

// Description implements Narrativer.
func (n T) Description() string {
	return n.description.String()
}

// Reason implements Narrativer.
func (n T) Reason() string {
	return n.reason.String()
}

// Title implements Narrativer.
func (n T) Title() string {
	return n.title.String()
}

// Update updates a Narrative.
func (n *T) Update(
	title string,
	description string,
	reason string,
) (T, error) {
	if n == nil {
		return T{}, fmt.Errorf("narrative is nil")
	}
	return New(
		n.id.String(),
		title,
		description,
		reason,
		n.bookID.String(),
	)
}

// Delete deletes a Narrative.
func (n *T) Delete(
	participant participant.T,
) (D, error) {
	if n == nil {
		return D{}, fmt.Errorf("book is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: n.id,
	}, nil
}