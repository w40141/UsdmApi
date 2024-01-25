// Package task is defined entity models to use this service.
package task

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Task.
func Create(
	title string,
	description string,
	reason string,
	projectID string,
	ownerID string,
	workerID string,
	options ...Option,
) (Task, error) {
	options = append(options, WithCreatedAt(time.Now()))
	options = append(options, WithUpdatedAt(time.Now()))
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
		projectID,
		ownerID,
		workerID,
		options...,
	)
}

// Update updates a Task.
func (t Task) Update(
	title string,
	description string,
	reason string,
	projectID string,
	ownerID string,
	workerID string,
	options ...Option,
) (Task, error) {
	options = append(options, WithUpdatedAt(time.Now()))
	return New(
		t.id.String(),
		title,
		description,
		reason,
		projectID,
		ownerID,
		workerID,
		options...,
	)
}
