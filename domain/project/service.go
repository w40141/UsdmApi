// Package project is defined for domain model.
package project

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Project.
func Create(
	title string,
	description string,
	ownerID string,
) (Project, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		ownerID,
		WithCreatedAt(time.Now()),
		WithUpdatedAt(time.Now()),
	)
}

// Update updates a Project.
func (p Project) Update(
	title string,
	description string,
	ownerID string,
) (Project, error) {
	return New(
		p.id.String(),
		title,
		description,
		ownerID,
		WithUpdatedAt(time.Now()),
	)
}
