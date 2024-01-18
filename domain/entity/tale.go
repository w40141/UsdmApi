// Package entity is ddefined for domain model.
package entity

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// Tale is an entity object for user story or usecase.
type Tale struct {
	id           vo.ID
	storyID      vo.ID
	plot         vo.Plot
	requirements []Requirement
}

// TaleOption is a functional option for Tale.
type TaleOption func(*Tale)

// NewTale creates a new Tale.
func NewTale(
	story Story,
	title string,
	description string,
	reason string,
	options ...TaleOption,
) (Tale, error) {
	newPlot, e1 := vo.NewPlot(title, description, reason)
	if e1 != nil {
		return Tale{}, e1
	}
	t := Tale{
		id:      vo.NewID(),
		plot:    newPlot,
		storyID: story.id,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// AddRequirements is a functional option for adding requirements.
func (t *Tale) AddRequirements(requirements []Requirement) func(*Tale) {
	return func(t *Tale) {
		t.requirements = requirements
	}
}
