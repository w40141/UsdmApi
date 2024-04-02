// Package member defines methods.
package member

// CanDelete returns whether the member can delete.
func (m *E) CanDelete() bool {
	if m.authority == nil {
		return false
	}
	return m.authority.CanDelete()
}

// CanInvite returns whether the member can invite.
func (m *E) CanInvite() bool {
	if m.authority == nil {
		return false
	}
	return m.authority.CanInvite()
}

// CanCreate returns whether the member can create.
func (m *E) CanCreate() bool {
	if m.authority == nil {
		return false
	}
	return m.authority.CanCreate()
}

// CanEdit returns whether the member can edit.
func (m *E) CanEdit() bool {
	if m.authority == nil {
		return false
	}
	return m.authority.CanEdit()
}
