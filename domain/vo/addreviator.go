// Package vo is defined for value object.
package vo

import "errors"

// Addreviator is a value object for Addreviator.
// Addreviator must be more than or equal 1 and less than or equal to 1000 characters.
type Addreviator struct {
	value string
}

// NewAddreviator creates a new Addreviator.
func NewAddreviator(value string) (Addreviator, error) {
	if value == "" {
		err := errors.New("Addreviator is required")
		return Addreviator{}, err
	}
	if len(value) > 1000 {
		err := errors.New("title must be less than or equal to 1000 characters")
		return Addreviator{}, err
	}
	return Addreviator{value: value}, nil
}
