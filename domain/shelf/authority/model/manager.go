// Package model defines the entity object for user authority.
package model

// Manager is the authority of editor.
type Manager struct{}

var _ Authoritier = (*Manager)(nil)

// CanDelete returns whether the manager can delete.
func (t *Manager) CanDelete() bool {
	return false
}

// CanInvite returns whether the manager can invite.
func (t *Manager) CanInvite() bool {
	return t != nil
}

// CanCreate returns whether the manager can create.
func (t *Manager) CanCreate() bool {
	return t != nil
}

// CanEdit returns whether the manager can edit.
func (t *Manager) CanEdit() bool {
	return t != nil
}

// Value returns the authority type.
func (t *Manager) Value() AurhorityType {
	return ManagerType
}

// NewManager is create a new manager.
func NewManager() *Editor {
	return &Editor{}
}
