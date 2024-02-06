// Package authority defines the entity object for user authority.
package authority

var _ T = (*Viewer)(nil)

// Viewer is an entity object for user Viewer or usecase.
type Viewer struct{}

// CanCreate returns whether the Viewer can create.
func (r *Viewer) CanCreate() bool {
	return false
}

// CanDelete returns whether the Viewer can delete.
func (r *Viewer) CanDelete() bool {
	return false
}

// NewViewer is create a new Viewer.
func NewViewer() *Viewer {
	return &Viewer{}
}
