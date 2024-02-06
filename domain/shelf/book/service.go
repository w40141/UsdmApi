// Package book is defined for domain model.
package book

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf/participant"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Option is a functional option for Book.
type Option func(*T)

// Create creates a new Book.
func Create(
	title string,
	description string,
	participant participant.T,
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
	return C{
		title:       titleVo,
		description: descriptionVo,
	}, nil
}

// New creates a new Book.
func New(
	id string,
	title string,
	description string,
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
	Book := T{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
	}
	for _, option := range options {
		option(&Book)
	}
	return Book, nil
}

// WithCreatedAt sets createdAt to Book.
func WithCreatedAt(createdAt *time.Time) Option {
	return func(p *T) {
		p.createdAt = createdAt
	}
}

// WithUpdatedAt sets updatedAt to Book.
func WithUpdatedAt(updatedAt *time.Time) Option {
	return func(p *T) {
		p.updatedAt = updatedAt
	}
}
