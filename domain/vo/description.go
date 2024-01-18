// Package vo is defined for value object.
package vo

import "errors"

// Description is a value object for Description.
// Description must be more than or equal 1 and less than or equal to 1000 characters.
type Description struct {
	value string
}

// NewDescription creates a new Description.
func NewDescription(value string) (Description, error) {
	if len(value) > 1000 {
		err := errors.New("title must be less than or equal to 1000 characters")
		return Description{}, err
	}
	return Description{value: value}, nil
}
