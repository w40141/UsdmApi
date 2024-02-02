// Package story defines the entity object for user story.
package story

import (
	"time"

	"github.com/w40141/UsdmApi/domain/request"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Story is an entity object for user story or usecase.
type Story struct {
	createAt    *time.Time
	updateAt    *time.Time
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	id          vo.ID
	legendID    vo.ID
	parentID    vo.ID
}

var (
	_ request.Storyer         = (*Story)(nil)
	_ request.ParentOfScene   = (*Story)(nil)
	_ request.ParentOfEpisode = (*Story)(nil)
)

// ParenOfEpisode implements request.ParentOfEpisode.
func (*Story) ParenOfEpisode() error {
	panic("unimplemented")
}

// Description implements request.StoryType.
func (s Story) Description() string {
	return s.description.String()
}

// ID implements request.StoryType.
func (s Story) ID() string {
	return s.id.String()
}

// Reason implements request.StoryType.
func (s Story) Reason() string {
	return s.reason.String()
}

// Title implements request.StoryType.
func (s Story) Title() string {
	return s.title.String()
}

// LegendID getter
func (s Story) LegendID() string {
	return s.legendID.String()
}

// ParentID getter
func (s Story) ParentID() string {
	return s.parentID.String()
}

// CreateAt getter
func (s Story) CreateAt() (time.Time, bool) {
	return *s.createAt, s.createAt != nil
}

// UpdateAt getter
func (s Story) UpdateAt() (time.Time, bool) {
	return *s.updateAt, s.updateAt != nil
}

// ParenOfScene implements request.StoryType.
func (*Story) ParenOfScene() error {
	panic("unimplemented")
}

// Create creates a new Story.
func Create(
	title string,
	description string,
	reason string,
	legend request.Legender,
	parent request.ParentOfStory,
) (Story, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
		legend.ID(),
		parent.ID(),
	)
}

// Update updates a Story.
func (s Story) Update(
	title string,
	description string,
	reason string,
	parent request.ParentOfStory,
) (Story, error) {
	return New(s.id.String(),
		title,
		description,
		reason,
		s.legendID.String(),
		parent.ID(),
	)
}

// Option is a functional option for Story.
type Option func(*Story)

// New creates a new Story.
func New(
	id string,
	title string,
	description string,
	reason string,
	legendID string,
	parentID string,
	options ...Option,
) (Story, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Story{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Story{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return Story{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return Story{}, e3
	}
	legendIDVo, e5 := vo.FromStringToID(legendID)
	if e5 != nil {
		return Story{}, e5
	}
	parentIDVo, e6 := vo.FromStringToID(parentID)
	if e6 != nil {
		return Story{}, e6
	}
	s := Story{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		legendID:    legendIDVo,
		parentID:    parentIDVo,
	}
	for _, option := range options {
		option(&s)
	}
	return s, nil
}

// WithCreateAt is a functional option for adding create time.
func WithCreateAt(createAt time.Time) func(*Story) {
	return func(s *Story) {
		s.createAt = &createAt
	}
}

// WithUpdateAt is a functional option for adding update time.
func WithUpdateAt(updateAt time.Time) func(*Story) {
	return func(s *Story) {
		s.updateAt = &updateAt
	}
}
