// Package member defines methods.
package member

import (
	"github.com/w40141/UsdmApi/domain/shelf/authority"
	"github.com/w40141/UsdmApi/domain/shelf/authority/types"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// isSameAccount returns whether the account of the member is different.
func (e *E) isSameAccount(id vo.AccountID) bool {
	if e == nil {
		return false
	}
	return e.account == id
}

// isSameWorld returns whether the world exist member is different.
func (e *E) isSameWorld(id vo.WorldID) bool {
	if e == nil {
		return false
	}
	return e.world == id
}

// isSameMember returns whether the member is different.
func (e *E) isSameMember(id vo.MemberID) bool {
	if e == nil {
		return false
	}
	return e.id == id
}

// canDelete returns whether the member can delete.
func (e *E) canDelete() bool {
	return e.authority.CanDelete()
}

// canInvite returns whether the member can invite.
func (e *E) canInvite() bool {
	return e.authority.CanInvite()
}

// canCreate returns whether the member can create.
// func (e *E) canCreate() bool {
// 	return e.authority.CanCreate()
// }

// canEdit returns whether the member can edit.
func (e *E) canEdit() bool {
	return e.authority.CanEdit()
}

// ChangeAuthority changes the authority of the member.
func (e *E) ChangeAuthority(
	m E,
	authorityType types.Type,
) (E, error) {
	authority, e1 := authority.New(authorityType)
	if e1 != nil {
		return E{}, e1
	}
	if e.isSameMember(m.id) {
		return E{}, vo.ErrSameMember
	}
	if !e.canEdit() {
		return E{}, vo.ErrCannotEdit
	}
	if !e.isSameWorld(m.world) {
		return E{}, vo.ErrDiffWorld
	}
	if e.authority.Compare(authority) > 0 {
		return E{}, vo.ErrHigherAuthority
	}

	return E{
		id:        m.id,
		account:   m.account,
		world:     m.world,
		authority: authority,
	}, nil
}

// Create creates a new member.
func (e *E) Create(
	account string,
	authorityType types.Type,
) (C, error) {
	accountIDVo, e1 := vo.FromStringToID(account)
	if e1 != nil {
		return C{}, e1
	}
	accountID := vo.AccountID(accountIDVo)
	if e.isSameAccount(accountID) {
		return C{}, vo.ErrAccountSame
	}

	t, e3 := authority.New(authorityType)
	if e3 != nil {
		return C{}, e3
	}

	if !e.canInvite() {
		return C{}, vo.ErrCannotInvite
	}

	return C{
		account:   accountID,
		world:     e.world,
		authority: t,
	}, nil
}

// Delete deletes a member.
func (e *E) Delete(m E) (D, error) {
	if e.isSameMember(m.id) {
		return D{}, vo.ErrSameMember
	}
	if !e.isSameWorld(m.world) {
		return D{}, vo.ErrDiffWorld
	}
	if !e.canDelete() {
		return D{}, vo.ErrCannotDelete
	}
	return D{
		id: m.id,
	}, nil
}
