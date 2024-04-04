// Package saga defines services.
package saga

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// Create creates a new Story.
func Create(
	title string,
	description string,
	creator string,
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

// New creates a new Saga.
func New(
	id string,
	title string,
	description string,
	creator string,
	options ...Option,
) (E, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return E{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return E{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return E{}, e3
	}
	creatorID, e4 := vo.FromStringToID(creator)
	if e4 != nil {
		return E{}, e4
	}
	s := E{
		id:          vo.SagaID(idVo),
		title:       titleVo,
		description: descriptionVo,
		creator:     vo.MemberID(creatorID),
	}
	for _, option := range options {
		option(&s)
	}
	return s, nil
}

// Option is a functional option for Story.
type Option func(*E)

// WithCreateAt is a functional option for adding create time.
func WithCreateAt(createAt time.Time) func(*E) {
	return func(s *E) {
		s.createAt = &createAt
	}
}

// WithUpdateAt is a functional option for adding update time.
func WithUpdateAt(updateAt time.Time) func(*E) {
	return func(s *E) {
		s.updateAt = &updateAt
	}
}
