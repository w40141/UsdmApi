// Package vo is defined for value object.
package vo

import (
	"github.com/rs/xid"
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

func (i ID) String() string {
	return i.value.String()
}
