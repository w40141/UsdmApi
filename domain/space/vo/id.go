// Package vo defines value objects in space.
package vo

import "github.com/w40141/UsdmApi/domain/vo"

type (
	// AccountID is a value object to identify an account.
	AccountID vo.ID
	// WorldID is a value object to identify a world.
	WorldID vo.ID
	// AbilityID is a value object to identify an ability.
	AbilityID vo.ID
	// CharacterID is a value object to identify a character.
	CharacterID vo.ID
	// AuthorityID is a value object to identify an authority.
	AuthorityID vo.ID
	// OrganismID is a value object to identify an organism.
	OrganismID vo.ID
)

// AccountName is a value object for AccountName.
type AccountName vo.Name
