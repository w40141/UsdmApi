// Package authority defines the entity object for user authority.
package authority

// T is an entity object for user Authority or usecase.
type T interface {
	CanCreate() bool
	CanDelete() bool
}

// Owner is an entity object for user Owner or usecase.
type Owner struct{}

// CanCreate returns whether the owner can create.
func (o *Owner) CanCreate() bool {
	return o != nil
}

// CanDelete returns whether the owner can delete.
func (o *Owner) CanDelete() bool {
	return o != nil
}

// NewOwner is create a new Owner.
func NewOwner() *Owner {
	return &Owner{}
}

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

// Reader is an entity object for user Reader or usecase.
type Reader struct{}

// CanCreate returns whether the reader can create.
func (r *Reader) CanCreate() bool {
	return false
}

// CanDelete returns whether the reader can delete.
func (r *Reader) CanDelete() bool {
	return false
}

// NewReader is create a new Reader.
func NewReader() *Reader {
	return &Reader{}
}
