// Package legend is defined for domain model.
package legend

import (
	"time"

	"github.com/w40141/UsdmApi/domain/request"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Legend is the largest unit of human management.
type Legend struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
}

var (
	_ request.ParentOfStory = (*Legend)(nil)
	_ request.Legender      = (*Legend)(nil)
)

// ParenOfScene implements request.LegendType.
func (*Legend) ParenOfScene() error {
	panic("unimplemented")
}

// ParentOfEpisode implements request.LegendType.
func (*Legend) ParentOfEpisode() error {
	panic("unimplemented")
}

// ParentOfStory implements story.ParentOfStory.
func (*Legend) ParentOfStory() error {
	panic("unimplemented")
}

// ParenOfNarrative implements narrative.ParentOfNarrative.
func (*Legend) ParenOfNarrative() error {
	panic("unimplemented")
}

// Description implements request.LegendType.
func (p Legend) Description() string {
	return p.description.String()
}

// ID implements request.LegendType.
func (p Legend) ID() string {
	return p.id.String()
}

// Title implements request.LegendType.
func (p *Legend) Title() string {
	return p.title.String()
}

// Update updates a Legend.
func (p *Legend) Update(
	title string,
	description string,
	reason string,
) (Legend, error) {
	if p == nil {
		return Create(title, description, reason)
	}
	return New(
		p.id.String(),
		title,
		description,
		reason,
	)
}

// Create creates a new Legend.
func Create(
	title string,
	description string,
	reason string,
) (Legend, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
	)
}

// Option is a functional option for Legend.
type Option func(*Legend)

// New creates a new Legend.
func New(
	id string,
	title string,
	description string,
	reason string,
	options ...Option,
) (Legend, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Legend{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Legend{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return Legend{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return Legend{}, e4
	}
	legend := Legend{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
	}
	for _, option := range options {
		option(&legend)
	}
	return legend, nil
}

// WithCreatedAt sets createdAt to Legend.
func WithCreatedAt(createdAt time.Time) Option {
	return func(p *Legend) {
		p.createdAt = &createdAt
	}
}

// WithUpdatedAt sets updatedAt to Legend.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(p *Legend) {
		p.updatedAt = &updatedAt
	}
}
