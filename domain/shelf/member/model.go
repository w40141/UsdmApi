// Package member defines types.
package member

import (
	"github.com/w40141/UsdmApi/domain/shelf/authority/model"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// E is an entity object for member.
type E struct {
	id        vo.MemberID
	account   vo.AccountID
	authority model.Authoritier
	world     vo.WorldID
}

// C is a struct for creating a new member.
type C struct {
	account   vo.AccountID
	authority model.Authoritier
	world     vo.WorldID
}
