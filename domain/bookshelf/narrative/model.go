// Package narrative is defined for domain model.
package narrative

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/bookshelf"
	"github.com/w40141/UsdmApi/domain/bookshelf/participant"
	"github.com/w40141/UsdmApi/domain/vo"
)

// N is an entity object for user N or usecase.
type N struct {
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
	_ bookshelf.ParentOfStory = (*N)(nil)
	_ bookshelf.Narrativer    = (*N)(nil)
)

// ParentOfScene implements request.Narrativer.
func (*N) ParentOfScene() error {
	panic("unimplemented")
}

// ParentOfStory implements story.ParentOfStory.
func (*N) ParentOfStory() error {
	panic("unimplemented")
}

// ID implements Narrativer.
func (n N) ID() vo.ID {
	return n.id
}

// Description implements Narrativer.
func (n N) Description() string {
	return n.description.String()
}

// Reason implements Narrativer.
func (n N) Reason() string {
	return n.reason.String()
}

// Title implements Narrativer.
func (n N) Title() string {
	return n.title.String()
}

// Update updates a Narrative.
func (n *N) Update(
	title string,
	description string,
	reason string,
) (N, error) {
	if n == nil {
		return N{}, fmt.Errorf("narrative is nil")
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
func (n *N) Delete(
	participant participant.P,
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
