// Package authority defines the entity object for user authority.
package authority

import "github.com/w40141/UsdmApi/domain/shelf"

var _ shelf.Authoritier = (*Editor)(nil)

// Editor is the authority of editor.
type Editor struct{}

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
