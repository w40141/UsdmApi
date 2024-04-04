// Package authority defines services for member authority.
package authority

// CanDelete returns whether the editor can delete.
func (e *E) CanDelete() bool {
	if e == nil {
		return false
	}
	return e.power&canDelete != 0
}

// CanInvite returns whether the editor can invite.
func (e *E) CanInvite() bool {
	if e == nil {
		return false
	}
	return e.power&canInvite != 0
}

// CanCreate returns whether the editor can create.
func (e *E) CanCreate() bool {
	if e == nil {
		return false
	}
	return e.power&canCreate != 0
}

// CanEdit returns whether the editor can edit.
func (e *E) CanEdit() bool {
	if e == nil {
		return false
	}
	return e.power&canEdit != 0
}

// Value returns the authority type.
func (e *E) Value() Type {
	return e.power
}

// Compare compares the authority.
func (e *E) Compare(a E) int {
	if e == nil {
		return -1
	}
	if e.power > a.power {
		return 1
	}
	if e.power == a.power {
		return 0
	}
	return -1
}

// New is create a new authority.
func New(power Type) E {
	return E{power}
}
