// Package authority defines the entity object for user authority.
package authority

import "github.com/w40141/UsdmApi/domain/user"

// Editor is the authority of editor.
type Editor struct{}

var _ user.Authoritier = (*Editor)(nil)

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

// NewEditor is create a new Editor.
func NewEditor() *Editor {
	return &Editor{}
}
