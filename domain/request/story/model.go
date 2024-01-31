// Package story defines the entity object for user story.
package story

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Story is an entity object for user story or usecase.
type Story struct {
	createAt    time.Time
	updateAt    time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
	legendID    vo.ID
	parentID    vo.ID
}

// Option is a functional option for Story.
type Option func(*Story)

// New creates a new Story.
func New(
	id string,
	title string,
	description string,
	reason string,
	legendID string,
	parentID string,
	options ...Option,
) (Story, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Story{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Story{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return Story{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return Story{}, e3
	}
	legendIDVo, e5 := vo.FromStringToID(legendID)
	if e5 != nil {
		return Story{}, e5
	}
	parentIDVo, e6 := vo.FromStringToID(parentID)
	if e6 != nil {
		return Story{}, e6
	}
	s := Story{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		legendID:    legendIDVo,
		parentID:    parentIDVo,
	}
	for _, option := range options {
		option(&s)
	}
	return s, nil
}

// WithCreateAt is a functional option for adding create time.
func WithCreateAt(createAt time.Time) func(*Story) {
	return func(s *Story) {
		s.createAt = createAt
	}
}

// WithUpdateAt is a functional option for adding update time.
func WithUpdateAt(updateAt time.Time) func(*Story) {
	return func(s *Story) {
		s.updateAt = updateAt
	}
}
