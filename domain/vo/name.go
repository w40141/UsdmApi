// Package vo is defined for value object.
package vo

import "errors"

// Name is a value object for Name.
type Name struct {
	value string
}

// NewName creates a new Name.
func NewName(name string) (Name, error) {
	if name == "" {
		err := errors.New("Name is required")
		return Name{}, err
	}
	return Name{value: name}, nil
}
