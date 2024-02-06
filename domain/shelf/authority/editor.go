// Package authority defines the entity object for user authority.
package authority

var _ T = (*Editor)(nil)

// Editor is an entity object for user Editor or usecase.
type Editor struct{}

// CanCreate returns whether the editor can create.
func (e *Editor) CanCreate() bool {
	return e != nil
}

// CanDelete returns whether the editor can delete.
func (e *Editor) CanDelete() bool {
	return false
}

// NewEditor is create a new Editor.
func NewEditor() *Editor {
	return &Editor{}
}
