// Package participant defines the entity object for user participant.
package participant

import (
	"github.com/w40141/UsdmApi/domain/user"
	"github.com/w40141/UsdmApi/domain/vo"
)

// New is create a new Participant.
func New(
	memberID string,
	legendID string,
	authority user.Authoritier,
) (T, error) {
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
