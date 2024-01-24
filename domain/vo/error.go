// Package vo is defined for value object.
package vo

import "errors"

// ErrInvalidTime is an error for invalid time.
var ErrInvalidTime = errors.New("start date must be before end date")

// ErrNoName is an error for no name.
var ErrNoName = errors.New("Name is required")

// ErrNoTitle is an error for no title.
var ErrNoTitle = errors.New("title is required")

// ErrTitleOverLength is an error for title over length.
var ErrTitleOverLength = errors.New("title must be less than or equal to 100 characters")

// ErrDescriptionOverLength is an error for description over length.
var ErrDescriptionOverLength = errors.New(
	"description must be less than or equal to 1000 characters",
)
