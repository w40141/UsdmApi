// Package member defines services.
package member

import (
	"github.com/w40141/UsdmApi/domain/shelf/authority"
	"github.com/w40141/UsdmApi/domain/shelf/authority/model"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// New is a constructor for E.
func New(
	id vo.MemberID,
	account vo.AccountID,
	world vo.WorldID,
	authorityType model.AurhorityType,
) (E, error) {
	t, e := authority.New(authorityType)
	if e != nil {
		return E{}, e
	}
	return E{
		id:        id,
		account:   account,
		world:     world,
		authority: t,
	}, nil
}

// Create creates a new member.
func Create(
	account string,
	world string,
	authorityType model.AurhorityType,
	member E,
) (C, error) {
	if !member.CanInvite() {
		return C{}, vo.ErrCannotInvite
	}

	accountID, e1 := vo.FromStringToID(account)
	if e1 != nil {
		return C{}, e1
	}
	worldID, e2 := vo.FromStringToID(world)
	if e2 != nil {
		return C{}, e2
	}
	t, e3 := authority.New(authorityType)
	if e3 != nil {
		return C{}, e3
	}

	return C{
		account:   vo.AccountID(accountID),
		world:     vo.WorldID(worldID),
		authority: t,
	}, nil
}
