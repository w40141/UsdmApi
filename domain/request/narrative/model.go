// Package narrative is defined for domain model.
package narrative

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Narrative is an entity object for user Narrative or usecase.
type Narrative struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
	legendID    vo.ID
	parentID    vo.ID
}

var _ Type = (*Narrative)(nil)

// Description implements Type.
func (*Narrative) Description() string {
	panic("unimplemented")
}

// Episode implements Type.
func (*Narrative) Episode() error {
	panic("unimplemented")
}

// ID implements Type.
func (*Narrative) ID() string {
	panic("unimplemented")
}

// Narrative implements Type.
func (*Narrative) Narrative() error {
	panic("unimplemented")
}

// Reason implements Type.
func (*Narrative) Reason() string {
	panic("unimplemented")
}

// Scene implements Type.
func (*Narrative) Scene() error {
	panic("unimplemented")
}

// Story implements Type.
func (*Narrative) Story() error {
	panic("unimplemented")
}

// Title implements Type.
func (*Narrative) Title() string {
	panic("unimplemented")
}

// Option is a functional option for Narrative.
type Option func(*Narrative)

// New creates a new Narrative.
func New(
	id string,
	title string,
	description string,
	reason string,
	legendID string,
	parentID string,
	options ...Option,
) (Narrative, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Narrative{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Narrative{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return Narrative{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return Narrative{}, e3
	}
	legendIDVo, e5 := vo.FromStringToID(legendID)
	if e5 != nil {
		return Narrative{}, e5
	}
	parentIDVo, e6 := vo.FromStringToID(parentID)
	if e6 != nil {
		return Narrative{}, e6
	}
	n := Narrative{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		legendID:    legendIDVo,
		parentID:    parentIDVo,
	}
	for _, option := range options {
		option(&n)
	}
	return n, nil
}

// WithCreatedAt sets createdAt to Narrative.
func WithCreatedAt(createdAt time.Time) Option {
	return func(e *Narrative) {
		e.createdAt = createdAt
	}
}

// WithUpdatedAt sets updatedAt to Narrative.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(e *Narrative) {
		e.updatedAt = updatedAt
	}
}
