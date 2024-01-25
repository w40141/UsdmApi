// Package epic is defined for domain model.
package epic

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create is create a new Epic.
func Create(
	title string,
	description string,
	projectID string,
) (Epic, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		projectID,
		WithCreatedAt(time.Now()),
		WithUpdatedAt(time.Now()),
	)
}

// Update updates a Epic.
func (e Epic) Update(
	title string,
	description string,
	projectID string,
) (Epic, error) {
	return New(
		e.id.String(),
		title,
		description,
		projectID,
		WithUpdatedAt(time.Now()),
	)
}
