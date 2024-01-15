// Package vo is defined for value object.
package vo

import (
	"github.com/rs/xid"
)

// ID is a value object for id.
type ID struct {
	value string
}

// NewID creates a new ID.
func NewID() ID {
	return ID{value: xid.New().String()}
}
