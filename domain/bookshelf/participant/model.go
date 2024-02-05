// Package participant defines the entity object for user participant.
package participant

import (
	"fmt"

	"github.com/w40141/UsdmApi/domain/bookshelf/authority"
	"github.com/w40141/UsdmApi/domain/vo"
)

// P is an entity object for Participant.
type P struct {
	authority authority.T
	memberID  vo.ID
	legendID  vo.ID
}

// CanCreate returns whether the participant can create.
func (p *P) CanCreate() bool {
	if p.authority == nil {
		return false
	}
	return p.authority.CanCreate()
}

// CanDelete returns whether the participant can delete.
func (p *P) CanDelete() bool {
	if p.authority == nil {
		return false
	}
	return p.authority.CanDelete()
}

// Update is update a authority of Participant.
func (p *P) Update(
	authority authority.T,
) (P, error) {
	if p == nil {
		return P{}, fmt.Errorf("participant is nil")
	}
	return P{
		memberID:  p.memberID,
		legendID:  p.legendID,
		authority: authority,
	}, nil
}
