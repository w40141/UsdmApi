// Package narrative is defined for domain model.
package narrative

import (
	"time"

	"github.com/w40141/UsdmApi/domain/bookshelf"
	"github.com/w40141/UsdmApi/domain/bookshelf/participant"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Option is a functional option for Narrative.
type Option func(*N)

// New creates a new Narrative.
func New(
	id string,
	title string,
	description string,
	reason string,
	bookID string,
	options ...Option,
) (N, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return N{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return N{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return N{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return N{}, e3
	}
	bookIDVo, e5 := vo.FromStringToID(bookID)
	if e5 != nil {
		return N{}, e5
	}
	n := N{
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
	return func(e *N) {
		e.createdAt = &createdAt
	}
}

// WithUpdatedAt sets updatedAt to Narrative.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(e *N) {
		e.updatedAt = &updatedAt
	}
}

// Create is create a new Narrative.
func Create(
	title string,
	description string,
	reason string,
	book bookshelf.Booker,
	participant participant.P,
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
