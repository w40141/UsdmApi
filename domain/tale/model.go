// Package tale is defined for domain model.
package tale

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Tale is an entity object for user story or usecase.
type Tale struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Description
	reason      vo.Description
	storyID     vo.ID
	id          vo.ID
	ownerID     vo.ID
}

// Option is a functional option for Tale.
type Option func(*Tale)

// New creates a new Tale.
func New(
	id string,
	title string,
	description string,
	reason string,
	storyID string,
	ownerID string,
	options ...Option,
) (Tale, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Tale{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Tale{}, e2
	}
	descriptionVo, e3 := vo.NewDescription(description)
	if e3 != nil {
		return Tale{}, e3
	}
	reasonVo, e4 := vo.NewDescription(reason)
	if e4 != nil {
		return Tale{}, e4
	}
	storyIDVo, e5 := vo.FromStringToID(storyID)
	if e5 != nil {
		return Tale{}, e5
	}
	ownerIDVo, e6 := vo.FromStringToID(ownerID)
	if e6 != nil {
		return Tale{}, e6
	}
	t := Tale{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		storyID:     storyIDVo,
		ownerID:     ownerIDVo,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// WithCreatedAt is a functional option for adding created at.
func WithCreatedAt(createdAt time.Time) Option {
	return func(t *Tale) {
		t.createdAt = createdAt
	}
}

// WithUpdatedAt is a functional option for adding updated at.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(t *Tale) {
		t.updatedAt = updatedAt
	}
}
