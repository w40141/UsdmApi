// Package story defines the entity object for user story.
package story

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Story.
func Create(
	title string,
	description string,
	reason string,
	book shelf.Booker,
	parent shelf.ParentOfStory,
	participant shelf.Participanter,
) (C, error) {
	if !participant.CanCreate() {
		return C{}, fmt.Errorf("participant can not create")
	}
	titleVo, e1 := vo.NewTitle(title)
	if e1 != nil {
		return C{}, e1
	}
	descriptionVo, e2 := vo.NewSentence(description)
	if e2 != nil {
		return C{}, e2
	}
	reasonVo, e3 := vo.NewSentence(reason)
	if e3 != nil {
		return C{}, e3
	}
	return C{
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      book.ID(),
		parentID:    parent.ID(),
	}, nil
}

// Option is a functional option for Story.
type Option func(*T)

// New creates a new Story.
func New(
	id string,
	title string,
	description string,
	reason string,
	bookID string,
	parentID string,
	options ...Option,
) (T, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return T{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return T{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return T{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return T{}, e3
	}
	bookIDVo, e5 := vo.FromStringToID(bookID)
	if e5 != nil {
		return T{}, e5
	}
	parentIDVo, e6 := vo.FromStringToID(parentID)
	if e6 != nil {
		return T{}, e6
	}
	s := T{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      bookIDVo,
		parentID:    parentIDVo,
	}
	for _, option := range options {
		option(&s)
	}
	return s, nil
}

// WithCreateAt is a functional option for adding create time.
func WithCreateAt(createAt time.Time) func(*T) {
	return func(s *T) {
		s.createAt = &createAt
	}
}

// WithUpdateAt is a functional option for adding update time.
func WithUpdateAt(updateAt time.Time) func(*T) {
	return func(s *T) {
		s.updateAt = &updateAt
	}
}
