// Package legend is defined for domain model.
package legend

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Legend is the largest unit of human management.
type Legend struct {
	createdAt   time.Time
	updatedAt   time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
}

var _ Type = (*Legend)(nil)

// Legend implements Type.
func (*Legend) Legend() error {
	panic("unimplemented")
}

// Episode implements Type.
func (*Legend) Episode() error {
	panic("unimplemented")
}

// Narrative implements Type.
func (*Legend) Narrative() error {
	panic("unimplemented")
}

// Scene implements Type.
func (*Legend) Scene() error {
	panic("unimplemented")
}

// Story implements Type.
func (*Legend) Story() error {
	panic("unimplemented")
}

// ID implements Type.
func (p Legend) ID() string {
	return p.id.String()
}

// Description implements Type.
func (p Legend) Description() string {
	return p.description.String()
}

// Title implements Type.
func (p Legend) Title() string {
	return p.title.String()
}

// Reason implements Type.
func (p Legend) Reason() string {
	return p.reason.String()
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
		p.createdAt = createdAt
	}
}

// WithUpdatedAt sets updatedAt to Legend.
func WithUpdatedAt(updatedAt time.Time) Option {
	return func(p *Legend) {
		p.updatedAt = updatedAt
	}
}
