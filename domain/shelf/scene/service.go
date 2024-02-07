// Package scene is defined for domain model.
package scene

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Scene.
func Create(
	title string,
	description string,
	reason string,
	book shelf.Booker,
	parent shelf.ParentOfScene,
) (C, error) {
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return C{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return C{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return C{}, e4
	}
	return C{
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      book.ID(),
		parentID:    parent.ID(),
	}, nil
}

// Option is a functional option for Scene.
type Option func(*T)

// New creates a new Scene.
func New(
	id string,
	title string,
	description string,
	reason string,
	bookID string,
	storyID string,
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
		return T{}, e4
	}
	bookIDVo, e5 := vo.FromStringToID(bookID)
	if e5 != nil {
		return T{}, e5
	}
	storyIDVo, e6 := vo.FromStringToID(storyID)
	if e6 != nil {
		return T{}, e6
	}
	t := T{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      bookIDVo,
		parentID:    storyIDVo,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// WithCreatedAt is a functional option for adding created at.
func WithCreatedAt(createdAt time.Time) Option {
	return func(t *T) {
		t.createdAt = &createdAt
	}
}

// WithUpdatedAt is a functional option for adding updated at.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(t *T) {
		t.updatedAt = &updatedAt
	}
}
