// Package participant defines the entity object for user participant.
package participant

import (
	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/shelf/authority"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for Participant.
type T struct {
	authority authority.Authoritier
	memberID  vo.ID
	bookID    vo.ID
}

// D is a data object for Participant.
type D struct {
	memberID vo.ID
	bookID   vo.ID
}

var _ shelf.Participanter = (*T)(nil)
