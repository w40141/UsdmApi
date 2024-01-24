// Package entity is defined entity models to use this service.
package entity

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Task is a entity model for task.
type Task struct {
	id           vo.ID
	title        vo.Title
	description  vo.Description
	reason       vo.Description
	projectID    vo.ID
	ownerID      vo.ID
	workerID     vo.ID
	parentTaskID vo.ID
	createdAt    time.Time
	updatedAt    time.Time
}

type TaskParam struct {
	param Task
}

func TaskBuilder(
	id vo.ID,
	title vo.Title,
	description vo.Description,
	reason vo.Description,
	projectID vo.ID,
	ownerID vo.ID,
) *TaskParam {
	return &TaskParam{
		id:          id,
		title:       title,
		description: description,
		reason:      reason,
		projectID:   projectID,
		ownerID:     ownerID,
	}
}

// NewTask creates a new Task.
func NewTask(
	id vo.ID,
	title vo.Title,
	description vo.Description,
	reason vo.Description,
	projectID vo.ID,
	ownerID vo.ID,
	options ...TaskOption,
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

// TaskOption is a functional option for Task.
type TaskOption func(*Task)

// CreateTask creates a new Task.
func CreateTask(
	id vo.ID,
	title vo.Title,
	description vo.Description,
	reason vo.Description,
	ownerID vo.ID,
	options ...TaskOption,
) Task {
	t := Task{
		id:          id,
		title:       title,
		description: description,
		reason:      reason,
		ownerID:     ownerID,
	}
	for _, option := range options {
		option(&t)
	}
	return t
}

// WithParentTask is a functional option for adding parent task.
func (t *Task) WithParentTask(parentTaskID vo.ID) TaskOption {
	return func(t *Task) {
		t.parentTaskID = parentTaskID
	}
}

// WithWorker is a functional option for adding worker.
func (t *Task) WithWorker(workerID vo.ID) TaskOption {
	return func(t *Task) {
		t.workerID = workerID
	}
}

// WithCreatedAt is a functional option for adding created at.
func (t *Task) WithCreatedAt(createdAt time.Time) TaskOption {
	return func(t *Task) {
		t.createdAt = createdAt
	}
}

// WithUpdatedAt is a functional option for adding updated at.
func (t *Task) WithUpdatedAt(updatedAt time.Time) TaskOption {
	return func(t *Task) {
		t.updatedAt = updatedAt
	}
}

// TaskRepository is a repository for Task.
type TaskRepository interface {
	Create(task Task) error
	Update(task Task) error
	Get(ids []vo.ID) ([]Task, error)
	Delete(id vo.ID) error
}
