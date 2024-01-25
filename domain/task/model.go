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
	id string,
	title string,
	description string,
	reason string,
	projectID string,
	ownerID string,
	workerID string,
	options ...Option,
) (Task, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Task{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Task{}, e2
	}
	descriptionVo, e3 := vo.NewDescription(description)
	if e3 != nil {
		return Task{}, e3
	}
	reasonVo, e4 := vo.NewDescription(reason)
	if e4 != nil {
		return Task{}, e4
	}
	projectIDVo, e5 := vo.FromStringToID(projectID)
	if e5 != nil {
		return Task{}, e5
	}
	ownerIDVo, e6 := vo.FromStringToID(ownerID)
	if e6 != nil {
		return Task{}, e6
	}
	workerIDVo, e7 := vo.FromStringToID(workerID)
	if e7 != nil {
		return Task{}, e7
	}
	t := Task{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		projectID:   projectIDVo,
		ownerID:     ownerIDVo,
		workerID:    workerIDVo,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// Option is a functional option for Task.
type Option func(*Task)

// WithParentTask is a functional option for adding parent task.
func WithParentTask(parentTaskID vo.ID) Option {
	return func(t *Task) {
		t.parentTaskID = parentTaskID
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
