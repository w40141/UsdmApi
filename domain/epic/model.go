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
	startAt     time.Time
	endAt       time.Time
	title       vo.Title
	description vo.Description
	id          vo.ID
	projectID   vo.ID
}

// Option is a functional option for Epic.
type Option func(*Epic)

// New creates a new Epic.
func New(
	id vo.ID,
	projectID vo.ID,
	title vo.Title,
	description vo.Description,
	startAt time.Time,
	endAt time.Time,
	options ...Option,
) Epic {
	e := Epic{
		id:          id,
		projectID:   projectID,
		title:       title,
		description: description,
		startAt:     startAt,
		endAt:       endAt,
	}
	for _, option := range options {
		option(&e)
	}
	return e
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
