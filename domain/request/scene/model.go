// Package scene is defined for domain model.
package scene

import (
	"time"

	"github.com/w40141/UsdmApi/domain/request"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Scene is an entity object for user story or usecase.
type Scene struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	parentID    vo.ID
	id          vo.ID
	legendID    vo.ID
}

var _ request.Scener = (*Scene)(nil)

// Description implements request.SceneType.
func (s Scene) Description() string {
	return s.description.String()
}

// ID implements request.SceneType.
func (s Scene) ID() string {
	return s.id.String()
}

// Reason implements request.SceneType.
func (s Scene) Reason() string {
	return s.reason.String()
}

// Title implements request.SceneType.
func (s Scene) Title() string {
	return s.title.String()
}

// Create creates a new Scene.
func Create(
	title string,
	description string,
	reason string,
	legend request.Legender,
	story request.ParentOfScene,
) (Scene, error) {
	id := vo.NewID().String()
	return New(
		id,
		title,
		description,
		reason,
		legend.ID(),
		story.ID(),
	)
}

// Update updates a Scene.
func (s Scene) Update(
	title string,
	description string,
	reason string,
	story request.ParentOfScene,
) (Scene, error) {
	return New(
		s.id.String(),
		title,
		description,
		reason,
		s.legendID.String(),
		story.ID(),
	)
}

// Option is a functional option for Scene.
type Option func(*Scene)

// New creates a new Scene.
func New(
	id string,
	title string,
	description string,
	reason string,
	legendID string,
	storyID string,
	options ...Option,
) (Scene, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Scene{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Scene{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return Scene{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return Scene{}, e4
	}
	legendIDVo, e5 := vo.FromStringToID(legendID)
	if e5 != nil {
		return Scene{}, e5
	}
	storyIDVo, e6 := vo.FromStringToID(storyID)
	if e6 != nil {
		return Scene{}, e6
	}
	t := Scene{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		legendID:    legendIDVo,
		parentID:    storyIDVo,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// WithCreatedAt is a functional option for adding created at.
func WithCreatedAt(createdAt time.Time) Option {
	return func(t *Scene) {
		t.createdAt = &createdAt
	}
}

// WithUpdatedAt is a functional option for adding updated at.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(t *Scene) {
		t.updatedAt = &updatedAt
	}
}
