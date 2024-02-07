// Package story defines the entity object for user story.
package story

import (
	"fmt"
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

// ParentOfEpisode implements request.ParentOfEpisode.
func (*T) ParentOfEpisode() error {
	panic("unimplemented")
}

// Description implements request.Storyer.
func (s T) Description() string {
	return s.description.String()
}

// ID implements request.Storyer.
func (s T) ID() vo.ID {
	return s.id
}

// Reason implements request.Storyer.
func (s T) Reason() string {
	return s.reason.String()
}

// Title implements request.Storyer.
func (s T) Title() string {
	return s.title.String()
}

// ParentID getter
func (s T) ParentID() string {
	return s.parentID.String()
}

// CreateAt getter
func (s T) CreateAt() (time.Time, bool) {
	return *s.createAt, s.createAt != nil
}

// UpdateAt getter
func (s T) UpdateAt() (time.Time, bool) {
	return *s.updateAt, s.updateAt != nil
}

// ParentOfScene implements request.Storyer.
func (*T) ParentOfScene() error {
	panic("unimplemented")
}

// Update updates a Story.
func (s *T) Update(
	title string,
	description string,
	reason string,
	parent shelf.ParentOfStory,
) (T, error) {
	if s == nil {
		return T{}, fmt.Errorf("story is nil")
	}
	return New(
		s.id.String(),
		title,
		description,
		reason,
		s.bookID.String(),
		parent.ID().String(),
	)
}

// Delete deletes a Story.
func (s *T) Delete(
	participant shelf.Participanter,
) (D, error) {
	if s == nil {
		return D{}, fmt.Errorf("story is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: s.id,
	}, nil
}
