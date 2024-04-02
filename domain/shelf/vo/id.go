// Package vo defines value objects for ID types.
package vo

import "github.com/rs/xid"

type (
	// AccountID is a type for Account ID.
	AccountID ID
	// WorldID is a type for World ID.
	WorldID ID
	// AuthorityID is a type for Authority ID.
	AuthorityID ID
	// AbilityID is a type for Ability ID.
	AbilityID ID
	// SagaID is a type for Saga ID.
	SagaID ID
	// CharacterID is a type for Chapter ID.
	CharacterID ID
	// BookID is a type for Book ID.
	BookID ID
	// StoryID is a type for Story ID.
	StoryID ID
	// MemberID is a type for Member ID.
	MemberID ID
)

// ID is a value object to identify.
type ID struct {
	value xid.ID
}

// NewID creates a new ID.
func NewID() ID {
	return ID{value: xid.New()}
}

// FromStringToID reads an ID from its string representation
func FromStringToID(id string) (ID, error) {
	i, err := xid.FromString(id)
	if err != nil {
		return ID{}, err
	}
	return ID{value: i}, err
}

func (i *ID) String() string {
	if i == nil {
		return ""
	}
	return i.value.String()
}
