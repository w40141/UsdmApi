// Package vo is defined for value object.
package vo

import "errors"

var (
	// ErrInvalidTime is an error for invalid time.
	ErrInvalidTime = errors.New("start date must be before end date")
	// ErrNoName is an error for no name.
	ErrNoName = errors.New("Name is required")
	// ErrNoTitle is an error for no title.
	ErrNoTitle = errors.New("title is required")
	// ErrTitleOverLength is an error for title over length.
	ErrTitleOverLength = errors.New("title must be less than or equal to 100 characters")
	// ErrDescriptionOverLength is an error for description over length.
	ErrDescriptionOverLength = errors.New(
		"description must be less than or equal to 1000 characters",
	)
	// ErrParticipantCantDelete is an error for participant can't delete.
	ErrParticipantCantDelete = errors.New("participant can't delete")
	// ErrNilScene is an error for nil scene.
	ErrNilScene = errors.New("scene is nil")
	// ErrInvalidAuthorityType is an error for invalid authority type.
	ErrInvalidAuthorityType = errors.New("invalid authority type")
	// ErrCannotInvite is an error for cannot invite.
	ErrCannotInvite = errors.New("cannot invite")
)
