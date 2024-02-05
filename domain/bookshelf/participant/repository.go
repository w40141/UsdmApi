// Package participant defines the entity object for user participant.
package participant

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// Repository is a repository for Participant.
type Repository interface {
	Save(P) error
	UpdateParticipant(P) error
	GetParticipants(vo.ID) ([]P, error)
	DeleteParticipant(vo.ID) error
}
