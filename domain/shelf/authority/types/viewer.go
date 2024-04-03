// Package types defines the entity object for user authority.
package types

var _ (Authoritier) = (*Viewer)(nil)

// Viewer is the authority of viewer.
type Viewer struct{}

// Compare returns other
func (t *Viewer) Compare(a Authoritier) int {
	if t.Value() > a.Value() {
		return 1
	}
	if t.Value() == a.Value() {
		return 0
	}
	return -1
}

// CanDelete returns whether the viewer can delete.
func (t *Viewer) CanDelete() bool {
	return false
}

// CanInvite returns whether the viewer can invite.
func (t *Viewer) CanInvite() bool {
	return false
}

// CanCreate returns whether the viewer can create.
func (t *Viewer) CanCreate() bool {
	return false
}

// CanEdit returns whether the viewer can edit.
func (t *Viewer) CanEdit() bool {
	return false
}

// Value returns the authority type.
func (t *Viewer) Value() Type {
	return ViewerType
}

// NewViewer is create a new viewer.
func NewViewer() *Viewer {
	return &Viewer{}
}
