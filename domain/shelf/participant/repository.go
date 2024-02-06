// Package participant defines the entity object for user participant.
package participant

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// Repository is a repository for Participant.
type Repository interface {
	Save(T) error
	UpdateParticipant(T) error
	GetParticipants(vo.ID) ([]T, error)
	DeleteParticipant(vo.ID) error
}
