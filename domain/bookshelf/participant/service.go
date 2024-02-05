// Package participant defines the entity object for user participant.
package participant

import (
	"github.com/w40141/UsdmApi/domain/bookshelf/authority"
	"github.com/w40141/UsdmApi/domain/vo"
)

// New is create a new Participant.
func New(
	memberID string,
	legendID string,
	authority authority.T,
) (P, error) {
	memberIDVo, e1 := vo.FromStringToID(memberID)
	if e1 != nil {
		return P{}, e1
	}
	lengendIDVo, e2 := vo.FromStringToID(legendID)
	if e2 != nil {
		return P{}, e2
	}
	return P{
		memberID:  memberIDVo,
		legendID:  lengendIDVo,
		authority: authority,
	}, nil
}
