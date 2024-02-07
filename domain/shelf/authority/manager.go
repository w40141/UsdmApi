// Package authority defines the entity object for user authority.
package authority

import "github.com/w40141/UsdmApi/domain/shelf"

var _ shelf.Authoritier = (*Manager)(nil)

// Manager is the authority of editor.
type Manager struct{}

// CanDelete returns whether the editor can delete.
func (t *Manager) CanDelete() bool {
	return false
}

// CanInvite returns whether the editor can invite.
func (t *Manager) CanInvite() bool {
	return t != nil
}

// CanCreate returns whether the editor can create.
func (t *Manager) CanCreate() bool {
	return t != nil
}

// CanEdit returns whether the editor can edit.
func (t *Manager) CanEdit() bool {
	return t != nil
}

// NewManager is create a new Manager.
func NewManager() *Editor {
	return &Editor{}
}
