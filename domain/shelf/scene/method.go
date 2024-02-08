// Package scene is defined for domain model.
package scene

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// ID implements shelf.Scener.
func (s T) ID() vo.ID {
	return s.id
}

// BookID implements M.
func (s *T) BookID() vo.ID {
	return s.bookID
}

// Description implements M.
func (s *T) Description() vo.Sentence {
	return s.description
}

// ParentID implements M.
func (s *T) ParentID() vo.ID {
	return s.parentID
}

// Title implements M.
func (s *T) Title() vo.Title {
	return s.title
}

// BookID implements M.
func (s *C) BookID() vo.ID {
	return s.bookID
}

// Description implements M.
func (s *C) Description() vo.Sentence {
	return s.description
}

// ID implements M.
func (*C) ID() vo.ID {
	return vo.NewID()
}

// ParentID implements M.
func (s *C) ParentID() vo.ID {
	return s.parentID
}

// Title implements M.
func (s *C) Title() vo.Title {
	return s.title
}

// Update updates a Scene.
func (s *T) Update(
	title string,
	description string,
	reason string,
	parent shelf.ParentOfScene,
) (T, error) {
	if s == nil {
		return T{}, vo.ErrNilScene
	}
	return NewT(
		s.ID().String(),
		title,
		description,
		reason,
		s.bookID.String(),
		parent.ID().String(),
		s.createdAt,
		time.Now(),
	)
}

// Delete deletes a Scene.
func (s *T) Delete(
	participant shelf.Participanter,
) (D, error) {
	if s == nil {
		return D{}, vo.ErrNilScene
	}
	if !participant.CanDelete() {
		return D{}, vo.ErrParticipantCantDelete
	}
	return D{
		id: s.id,
	}, nil
}
