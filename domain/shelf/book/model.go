// Package book is defined for domain model.
package book

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/shelf/participant"
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

// C is a struct for creating a new Book.
type C struct {
	title       vo.Title
	description vo.Sentence
}

// D is a struct for deleting a Book.
type D struct {
	id vo.ID
}

// CanDelete implements shelf.Deletable.
func (*D) CanDelete() bool {
	return true
}

// ID implements shelf.Deletable.
func (*D) ID() vo.ID {
	panic("unimplemented")
}

// ID implements bookshelf.Book.
func (t T) ID() vo.ID {
	return t.id
}

// ParentOfEpisode implements bookshelf.Book.
func (*T) ParentOfEpisode() error {
	return nil
}

// ParentOfScene implements bookshelf.Book.
func (*T) ParentOfScene() error {
	return nil
}

// ParentOfStory implements bookshelf.Book.
func (*T) ParentOfStory() error {
	return nil
}

var (
	_ shelf.Booker    = (*T)(nil)
	_ shelf.Deletable = (*D)(nil)
)

// Update updates a Book.
func (t *T) Update(
	title string,
	description string,
	participant participant.T,
) (T, error) {
	if t == nil {
		return T{}, fmt.Errorf("book is nil")
	}
	if !participant.CanCreate() {
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
	participant participant.T,
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
