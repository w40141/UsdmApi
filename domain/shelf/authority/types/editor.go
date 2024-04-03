// Package types defines the entity object for user authority.
package types

var _ (Authoritier) = (*Editor)(nil)

// Editor is the authority of editor.
type Editor struct{}

// Compare returns other
func (t *Editor) Compare(a Authoritier) int {
	if t.Value() > a.Value() {
		return 1
	}
	if t.Value() == a.Value() {
		return 0
	}
	return -1
}

// CanDelete returns whether the editor can delete.
func (t *Editor) CanDelete() bool {
	return false
}

// CanInvite returns whether the editor can invite.
func (t *Editor) CanInvite() bool {
	return false
}

// CanCreate returns whether the editor can create.
func (t *Editor) CanCreate() bool {
	return t != nil
}

// CanEdit returns whether the editor can edit.
func (t *Editor) CanEdit() bool {
	return t != nil
}

// Value returns the authority type.
func (t *Editor) Value() Type {
	return EditorType
}

// NewEditor is create a new editor.
func NewEditor() *Editor {
	return &Editor{}
}
