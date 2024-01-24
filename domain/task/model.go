// Package task is defined entity models to use this service.
package task

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Task is a entity model for task.
type Task struct {
	createdAt    time.Time
	updatedAt    time.Time
	title        vo.Title
	description  vo.Description
	reason       vo.Description
	id           vo.ID
	projectID    vo.ID
	ownerID      vo.ID
	workerID     vo.ID
	parentTaskID vo.ID
}

// New creates a new Task.
func New(
	id vo.ID,
	title vo.Title,
	description vo.Description,
	reason vo.Description,
	projectID vo.ID,
	ownerID vo.ID,
	options ...Option,
) Task {
	t := Task{
		id:          id,
		title:       title,
		description: description,
		reason:      reason,
		projectID:   projectID,
		ownerID:     ownerID,
	}
	for _, option := range options {
		option(&t)
	}
	return t
}

// Option is a functional option for Task.
type Option func(*Task)

// WithParentTask is a functional option for adding parent task.
func WithParentTask(parentTaskID vo.ID) Option {
	return func(t *Task) {
		t.parentTaskID = parentTaskID
	}
}

// WithWorker is a functional option for adding worker.
func WithWorker(workerID vo.ID) Option {
	return func(t *Task) {
		t.workerID = workerID
	}
}

// WithCreatedAt is a functional option for adding created at.
func WithCreatedAt(createdAt time.Time) Option {
	return func(t *Task) {
		t.createdAt = createdAt
	}
}

// WithUpdatedAt is a functional option for adding updated at.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(t *Task) {
		t.updatedAt = updatedAt
	}
}
