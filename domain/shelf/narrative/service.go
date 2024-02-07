// Package narrative is defined for domain model.
package narrative

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Option is a functional option for Narrative.
type Option func(*T)

// New creates a new Narrative.
func New(
	id string,
	title string,
	description string,
	reason string,
	bookID string,
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
	n := T{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      bookIDVo,
	}
	for _, option := range options {
		option(&n)
	}
	return n, nil
}

// WithCreatedAt sets createdAt to Narrative.
func WithCreatedAt(createdAt time.Time) Option {
	return func(e *T) {
		e.createdAt = &createdAt
	}
}

// WithUpdatedAt sets updatedAt to Narrative.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(e *T) {
		e.updatedAt = &updatedAt
	}
}

// Create is create a new Narrative.
func Create(
	title string,
	description string,
	reason string,
	book shelf.Booker,
	participant shelf.Participanter,
) (C, error) {
	if !participant.CanCreate() {
		return C{}, nil
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
	}, nil
}
