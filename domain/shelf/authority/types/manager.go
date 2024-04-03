// Package types defines the entity object for user authority.
package types

var _ (Authoritier) = (*Manager)(nil)

// Manager is the authority of editor.
type Manager struct{}

// Compare returns other
func (t *Manager) Compare(a Authoritier) int {
	if t.Value() > a.Value() {
		return 1
	}
	if t.Value() == a.Value() {
		return 0
	}
	return -1
}

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
func (t *Manager) Value() Type {
	return ManagerType
}

// NewManager is create a new manager.
func NewManager() *Editor {
	return &Editor{}
}
