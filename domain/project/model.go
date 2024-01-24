// Package project is defined for domain model.
package project

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Project is the largest unit of human management.
type Project struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Description
	id          vo.ID
	ownerID     vo.ID
}

// Option is a functional option for Project.
type Option func(*Project)

// New creates a new Project.
func New(
	id vo.ID,
	title vo.Title,
	description vo.Description,
	ownerID vo.ID,
	options ...Option,
) Project {
	p := Project{
		id:          id,
		title:       title,
		description: description,
		ownerID:     ownerID,
	}
	for _, option := range options {
		option(&p)
	}
	return p
}

// WithCreatedAt sets createdAt to Project.
func WithCreatedAt(createdAt time.Time) Option {
	return func(p *Project) {
		p.createdAt = createdAt
	}
}

// WithUpdatedAt sets updatedAt to Project.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(p *Project) {
		p.updatedAt = updatedAt
	}
}
