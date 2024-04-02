// Package model defines interface.
package model

// Authoritier is an interface for user authority.
type Authoritier interface {
	CanDelete() bool
	CanInvite() bool
	CanCreate() bool
	CanEdit() bool
	Value() AurhorityType
}

// AurhorityType is an int representing authorities.
type AurhorityType int

const (
	// EditorType is the authority of editor.
	EditorType AurhorityType = 1
	// ManagerType is the authority of manager.
	ManagerType AurhorityType = 2
	// OwnerType is the authority of owner.
	OwnerType AurhorityType = 3
	// ViewerType is the authority of viewer.
	ViewerType AurhorityType = 4
)
