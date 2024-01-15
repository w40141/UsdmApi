// Package vo is defined for value object.
package vo

import "errors"

// Title is a value object for title.
// Title must be more than 0 and less than or equal to 100 characters.
type Title struct {
	value string
}

// NewTitle creates a new Title.
func NewTitle(value string) (Title, error) {
	if value == "" {
		err := errors.New("title is required")
		return Title{}, err
	}
	if len(value) > 100 {
		err := errors.New("title must be less than or equal to 100 characters")
		return Title{}, err
	}
	return Title{value: value}, nil
}
