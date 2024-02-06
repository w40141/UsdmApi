// Package participant defines the entity object for user participant.
package participant

import (
	"fmt"

	"github.com/w40141/UsdmApi/domain/shelf/authority"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for Participant.
type T struct {
	authority authority.T
	memberID  vo.ID
	legendID  vo.ID
}

// CanCreate returns whether the participant can create.
func (p *T) CanCreate() bool {
	if p.authority == nil {
		return false
	}
	return p.authority.CanCreate()
}

// CanDelete returns whether the participant can delete.
func (p *T) CanDelete() bool {
	if p.authority == nil {
		return false
	}
	return p.authority.CanDelete()
}

// Update is update a authority of Participant.
func (p *T) Update(
	authority authority.T,
) (T, error) {
	if p == nil {
		return T{}, fmt.Errorf("participant is nil")
	}
	return T{
		memberID:  p.memberID,
		legendID:  p.legendID,
		authority: authority,
	}, nil
}

// Create is create
func (p *T) Create(
	memberID string,
	legendID string,
	authority authority.T,
) (T, error) {
	if !p.CanCreate() {
		return T{}, fmt.Errorf("participant can not create")
	}
	memberIDVo, e1 := vo.FromStringToID(memberID)
	if e1 != nil {
		return T{}, e1
	}
	lengendIDVo, e2 := vo.FromStringToID(legendID)
	if e2 != nil {
		return T{}, e2
	}
	return T{
		memberID:  memberIDVo,
		legendID:  lengendIDVo,
		authority: authority,
	}, nil
}
