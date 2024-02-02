// Package narrative is defined for domain model.
package narrative

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/request"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Narrative is an entity object for user Narrative or usecase.
type Narrative struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
	legendID    vo.ID
}

var (
	_ request.ParentOfStory = (*Narrative)(nil)
	_ request.Narrativer    = (*Narrative)(nil)
)

// ParentOfScene implements request.Narrativer.
func (*Narrative) ParentOfScene() error {
	panic("unimplemented")
}

// ParentOfStory implements story.ParentOfStory.
func (*Narrative) ParentOfStory() error {
	panic("unimplemented")
}

// ID implements Narrativer.
func (n Narrative) ID() string {
	return n.id.String()
}

// Description implements Narrativer.
func (n Narrative) Description() string {
	return n.description.String()
}

// Reason implements Narrativer.
func (n Narrative) Reason() string {
	return n.reason.String()
}

// Title implements Narrativer.
func (n Narrative) Title() string {
	return n.title.String()
}

// Create is create a new Narrative.
func Create(
	title string,
	description string,
	reason string,
	legend request.Legender,
) (Narrative, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
		legend.ID(),
	)
}

// Update updates a Narrative.
func (n *Narrative) Update(
	title string,
	description string,
	reason string,
) (Narrative, error) {
	if n == nil {
		return Narrative{}, fmt.Errorf("narrative is nil")
	}
	return New(
		n.id.String(),
		title,
		description,
		reason,
		n.legendID.String(),
	)
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
	n := Narrative{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		legendID:    legendIDVo,
	}
	for _, option := range options {
		option(&n)
	}
	return n, nil
}

// WithCreatedAt sets createdAt to Narrative.
func WithCreatedAt(createdAt time.Time) Option {
	return func(e *Narrative) {
		e.createdAt = &createdAt
	}
}

// WithUpdatedAt sets updatedAt to Narrative.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(e *Narrative) {
		e.updatedAt = &updatedAt
	}
}
