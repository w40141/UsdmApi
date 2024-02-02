// Package legend is defined for domain model.
package legend

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/request"
	"github.com/w40141/UsdmApi/domain/request/participant"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is the largest unit of human management.
type T struct {
	createdAt   *time.Time
	updatedAt   *time.Time
	title       vo.Title
	description vo.Sentence
	id          vo.ID
}

// C is a struct for creating a new Legend.
type C struct {
	title       vo.Title
	description vo.Sentence
}

// D is a struct for deleting a Legend.
type D struct {
	id vo.ID
}

// ID implements request.Legender.
func (p *T) ID() (string, error) {
	if p == nil {
		return "", fmt.Errorf("legend is nil")
	}
	return p.id.String(), nil
}

// ParentOfEpisode implements request.Legender.
func (*T) ParentOfEpisode() error {
	return nil
}

// ParentOfNarrative implements request.Legender.
func (*T) ParentOfNarrative() error {
	return nil
}

// ParentOfScene implements request.Legender.
func (*T) ParentOfScene() error {
	return nil
}

// ParentOfStory implements request.Legender.
func (*T) ParentOfStory() error {
	return nil
}

var _ request.Legender = (*T)(nil)

// Update updates a Legend.
func (p *T) Update(
	title string,
	description string,
	participant participant.T,
) (T, error) {
	if p == nil {
		return T{}, fmt.Errorf("legend is nil")
	}
	if !participant.CanCreate() {
		return T{}, fmt.Errorf("participant can not update")
	}
	return New(
		p.id.String(),
		title,
		description,
		WithCreatedAt(p.createdAt),
		WithUpdatedAt(p.updatedAt),
	)
}

// Delete deletes a Legend.
func (p *T) Delete(
	participant participant.T,
) (D, error) {
	if p == nil {
		return D{}, fmt.Errorf("legend is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: p.id,
	}, nil
}
