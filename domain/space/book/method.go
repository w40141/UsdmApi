// Package book is defined for domain model.
package book

import (
	"fmt"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// ID implements shelf.Deletable.
func (d *D) ID() vo.ID {
	return d.id
}

var _ shelf.Booker = (*T)(nil)

// ID implements shelf.Booker.
func (t T) ID() vo.ID {
	return t.id
}

// Book implements shelf.Booker.
func (*T) Book() error {
	return nil
}

// ParentOfEpisode implements shelf.Booker.
func (*T) ParentOfEpisode() error {
	return nil
}

// ParentOfScene implements shelf.Booker.
func (*T) ParentOfScene() error {
	return nil
}

// ParentOfStory implements shelf.Booker.
func (*T) ParentOfStory() error {
	return nil
}

// Update updates a Book.
func (t *T) Update(
	title string,
	description string,
	participant shelf.Participanter,
) (T, error) {
	if t == nil {
		return T{}, fmt.Errorf("book is nil")
	}
	if !participant.CanEdit() {
		return T{}, fmt.Errorf("participant can not update")
	}
	return New(
		t.id.String(),
		title,
		description,
		WithCreatedAt(t.createdAt),
		WithUpdatedAt(t.updatedAt),
	)
}

// Delete deletes a Book.
func (t *T) Delete(
	participant shelf.Participanter,
) (D, error) {
	if t == nil {
		return D{}, fmt.Errorf("book is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: t.id,
	}, nil
}
