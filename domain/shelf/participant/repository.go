// Package participant defines the entity object for user participant.
package participant

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// WriteRepository is a repository interface for Episode.
type WriteRepository interface {
	Create(participant T) (T, error)
	UpdateParticipant(participan T) (T, error)
}

// ReadRepository is a repository interface for Episode.
type ReadRepository interface {
	GetParticipants(vo.ID) ([]T, error)
}

// DeleteRepository is a repository interface for Episode.
type DeleteRepository interface {
	DeleteParticipant(vo.ID) error
}
