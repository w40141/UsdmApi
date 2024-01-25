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
	id string,
	title string,
	description string,
	ownerID string,
	options ...Option,
) (Project, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Project{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Project{}, e2
	}
	descriptionVo, e3 := vo.NewDescription(description)
	if e3 != nil {
		return Project{}, e3
	}
	ownerIDVo, e4 := vo.FromStringToID(ownerID)
	if e4 != nil {
		return Project{}, e4
	}
	p := Project{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		ownerID:     ownerIDVo,
	}
	for _, option := range options {
		option(&p)
	}
	return p, nil
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
