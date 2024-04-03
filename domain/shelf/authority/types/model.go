// Package types defines interface.
package types

// Authoritier is an interface for user authority.
type Authoritier interface {
	CanDelete() bool
	CanInvite() bool
	CanCreate() bool
	CanEdit() bool
	Value() Type
	Compare(a Authoritier) int
}

// Type is an int representing authorities.
type Type int

const (
	// EditorType is the authority of editor.
	EditorType Type = 1
	// ManagerType is the authority of manager.
	ManagerType Type = 2
	// OwnerType is the authority of owner.
	OwnerType Type = 3
	// ViewerType is the authority of viewer.
	ViewerType Type = 4
)
