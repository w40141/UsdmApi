// Package epic is defined for domain model.
package epic

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Epic is an entity object for user story or usecase.
type Epic struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Description
	id          vo.ID
	projectID   vo.ID
}

// Option is a functional option for Epic.
type Option func(*Epic)

// New creates a new Epic.
func New(
	id string,
	title string,
	description string,
	projectID string,
	options ...Option,
) (Epic, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Epic{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Epic{}, e2
	}
	descriptionVo, e3 := vo.NewDescription(description)
	if e3 != nil {
		return Epic{}, e3
	}
	projectIDVo, e4 := vo.FromStringToID(projectID)
	if e4 != nil {
		return Epic{}, e4
	}
	e := Epic{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		projectID:   projectIDVo,
	}
	for _, option := range options {
		option(&e)
	}
	return e, nil
}

// WithCreatedAt sets createdAt to Epic.
func WithCreatedAt(createdAt time.Time) Option {
	return func(e *Epic) {
		e.createdAt = createdAt
	}
}

// WithUpdatedAt sets updatedAt to Epic.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(e *Epic) {
		e.updatedAt = updatedAt
	}
}
