// Package entity is ddefined for domain model.
package entity

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// Story is an entity object for user story or usecase.
type Story struct {
	id        vo.ID
	projectID vo.ID
	plot      vo.Plot
	tales     []Tale
}

// StoryOption is a functional option for Story.
type StoryOption func(*Story)

// NewStory creates a new Story.
func NewStory(
	project Project,
	title string,
	description string,
	reason string,
	options ...StoryOption,
) (Story, error) {
	newPlot, e1 := vo.NewPlot(title, description, reason)
	if e1 != nil {
		return Story{}, e1
	}
	s := Story{
		id:        vo.NewID(),
		plot:      newPlot,
		projectID: project.id,
	}
	for _, option := range options {
		option(&s)
	}
	return s, nil
}

// AddTales is a functional option for adding tales.
func (s *Story) AddTales(tale []Tale) func(*Story) {
	return func(s *Story) {
		s.tales = tale
	}
}
