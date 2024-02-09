// Package participant defines the entity object for user participant.
package participant

import (
	"fmt"

	"github.com/w40141/UsdmApi/domain/shelf/authority"
)

// CanDelete returns whether the participant can delete.
func (t *T) CanDelete() bool {
	if t == nil {
		return false
	}
	return t.authority.CanDelete()
}

// CanInvite returns whether the participant can invite.
func (t *T) CanInvite() bool {
	if t == nil {
		return false
	}
	return t.authority.CanInvite()
}

// CanCreate returns whether the participant can create.
func (t *T) CanCreate() bool {
	if t == nil {
		return false
	}
	return t.authority.CanCreate()
}

// CanEdit returns whether the participant can edit.
func (t *T) CanEdit() bool {
	if t == nil {
		return false
	}
	return t.authority.CanEdit()
}

// Update is update a authority of Participant.
func (t *T) Update(
	authority authority.Authoritier,
) (T, error) {
	if !t.CanInvite() {
		return T{}, fmt.Errorf("participant can not update)")
	}
	return T{
		memberID:  t.memberID,
		bookID:    t.bookID,
		authority: authority,
	}, nil
}

// Invite is create a new Participant.
func (t *T) Invite(participant T) (T, error) {
	if !t.CanInvite() {
		return T{}, fmt.Errorf("participant can not invite")
	}
	return T{
		memberID:  participant.memberID,
		bookID:    participant.bookID,
		authority: participant.authority,
	}, nil
}

// Delete is delete a Participant.
func (t *T) Delete() (D, error) {
	if !t.CanInvite() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		memberID: t.memberID,
		bookID:   t.bookID,
	}, nil
}
