// Package model defines the entity object for user authority.
package model

// Viewer is the authority of viewer.
type Viewer struct{}

var _ Authoritier = (*Viewer)(nil)

// CanDelete returns whether the viewer can delete.
func (r *Viewer) CanDelete() bool {
	return false
}

// CanInvite returns whether the viewer can invite.
func (r *Viewer) CanInvite() bool {
	return false
}

// CanCreate returns whether the viewer can create.
func (r *Viewer) CanCreate() bool {
	return false
}

// CanEdit returns whether the viewer can edit.
func (r *Viewer) CanEdit() bool {
	return false
}

// Value returns the authority type.
func (r *Viewer) Value() AurhorityType {
	return ViewerType
}

// NewViewer is create a new viewer.
func NewViewer() *Viewer {
	return &Viewer{}
}
