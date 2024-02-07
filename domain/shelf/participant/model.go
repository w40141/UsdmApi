// Package participant defines the entity object for user participant.
package participant

import (
	"fmt"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/user"
	"github.com/w40141/UsdmApi/domain/vo"
)

// T is an entity object for Participant.
type T struct {
	authority user.Authoritier
	memberID  vo.ID
	legendID  vo.ID
}

var _ shelf.Participanter = (*T)(nil)

// CanDelete returns whether the participant can delete.
func (t *T) CanDelete() bool {
	return t.authority.CanDelete()
}

// CanInvite returns whether the participant can invite.
func (t *T) CanInvite() bool {
	return t.authority.CanInvite()
}

// CanCreate returns whether the participant can create.
func (t *T) CanCreate() bool {
	return t.authority.CanCreate()
}

// CanEdit returns whether the participant can edit.
func (t *T) CanEdit() bool {
	return t.authority.CanEdit()
}

// Update is update a authority of Participant.
func (t *T) Update(
	authority user.Authoritier,
) (T, error) {
	if !t.CanInvite() {
		return T{}, fmt.Errorf("participant is nil")
	}
	return T{
		memberID:  t.memberID,
		legendID:  t.legendID,
		authority: authority,
	}, nil
}

// Invite is create a new Participant.
func (t *T) Invite(
	memberID string,
	legendID string,
	authority user.Authoritier,
) (T, error) {
	if !t.CanInvite() {
		return T{}, fmt.Errorf("participant can not invite")
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
