// Package entity is ddefined for domain model.
package entity

import "github.com/w40141/UsdmApi/domain/vo"

// Backlog is an entity object for user story or usecase.
type Backlog struct {
	projectID vo.ID
	stories   []Story
}

// BacklogOption is a functional option for Backlog.
type BacklogOption func(*Backlog)

// NewBacklog creates a new Backlog.
func NewBacklog(project Project, options ...BacklogOption) (Backlog, error) {
	b := Backlog{
		projectID: project.id,
	}
	for _, option := range options {
		option(&b)
	}
	return b, nil
}

// AddStories is a functional option for adding stories.
func (b *Backlog) AddStories(stories []Story) func(*Backlog) {
	return func(b *Backlog) {
		b.stories = stories
	}
}
