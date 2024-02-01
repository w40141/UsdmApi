// Package episode defines the entity object for user episode.
package episode

import (
	"time"

	"github.com/w40141/UsdmApi/domain/request"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Episode is an entity object for user story or usecase.
type Episode struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	storyID     vo.ID
	id          vo.ID
	legendID    vo.ID
}

var _ request.EpisodeType = (*Episode)(nil)

// Description implements request.EpisodeType.
func (e Episode) Description() string {
	return e.description.String()
}

// ID implements request.EpisodeType.
func (e Episode) ID() string {
	return e.id.String()
}

// ParenOfScene implements request.EpisodeType.
func (*Episode) ParenOfScene() error {
	panic("unimplemented")
}

// Reason implements request.EpisodeType.
func (e Episode) Reason() string {
	return e.reason.String()
}

// Title implements request.EpisodeType.
func (e Episode) Title() string {
	return e.title.String()
}

// Create creates a new Episode.
func Create(
	title string,
	description string,
	reason string,
	legend request.LegendType,
	story request.ParentOfEpisode,
) (Episode, error) {
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

// Update updates a Episode.
func (e Episode) Update(
	title string,
	description string,
	reason string,
	story request.ParentOfEpisode,
) (Episode, error) {
	return New(
		e.id.String(),
		title,
		description,
		reason,
		e.legendID.String(),
		story.ID(),
	)
}

// Option is a functional option for Episode.
type Option func(*Episode)

// New creates a new Episode.
func New(
	id string,
	title string,
	description string,
	reason string,
	legendID string,
	storyID string,
	options ...Option,
) (Episode, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Episode{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Episode{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return Episode{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return Episode{}, e4
	}
	legendIDVo, e5 := vo.FromStringToID(legendID)
	if e5 != nil {
		return Episode{}, e5
	}
	storyIDVo, e6 := vo.FromStringToID(storyID)
	if e6 != nil {
		return Episode{}, e6
	}
	t := Episode{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		legendID:    legendIDVo,
		storyID:     storyIDVo,
	}
	for _, option := range options {
		option(&t)
	}
	return t, nil
}

// WithCreatedAt is a functional option for adding created at.
func WithCreatedAt(createdAt time.Time) Option {
	return func(t *Episode) {
		t.createdAt = createdAt
	}
}

// WithUpdatedAt is a functional option for adding updated at.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(t *Episode) {
		t.updatedAt = updatedAt
	}
}
