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
	ownerID vo.ID,
	options ...Option,
) (Task, error) {
	id := vo.CreateID()
	titleVo, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Task{}, e1
	}
	descriptionVo, e2 := vo.NewDescription(description)
	if e2 != nil {
		return Task{}, e2
	}
	reasonVo, e3 := vo.NewDescription(reason)
	if e3 != nil {
		return Task{}, e3
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	t := Task{
		id:          id,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		ownerID:     ownerID,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// Update updates a Task.
func (t Task) Update(title string, description string, reason string) (Task, error) {
	titleVo, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Task{}, e1
	}
	descriptionVo, e2 := vo.NewDescription(description)
	if e2 != nil {
		return Task{}, e2
	}
	reasonVo, e3 := vo.NewDescription(reason)
	if e3 != nil {
		return Task{}, e3
	}
	return New(
		t.id,
		titleVo,
		descriptionVo,
		reasonVo,
		t.projectID,
		t.ownerID,
		WithCreatedAt(t.createdAt),
		WithUpdatedAt(time.Now()),
	), nil
}
