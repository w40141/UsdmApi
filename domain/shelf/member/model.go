// Package member defines types.
package member

import (
	"github.com/w40141/UsdmApi/domain/shelf/authority"
	"github.com/w40141/UsdmApi/domain/shelf/authority/types"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// E is an entity object for member.
type E struct {
	id        vo.MemberID
	account   vo.AccountID
	authority types.Authoritier
	world     vo.WorldID
}

// C is a struct for creating a new member.
type C struct {
	account   vo.AccountID
	authority types.Authoritier
	world     vo.WorldID
}

// D is a struct for deleting a member.
type D struct {
	id vo.MemberID
}

// New is a constructor for E.
func New(
	id vo.MemberID,
	account vo.AccountID,
	world vo.WorldID,
	authorityType types.Type,
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
