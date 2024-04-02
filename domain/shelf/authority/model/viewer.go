// Package model defines the entity object for user authority.
package model

import "github.com/w40141/UsdmApi/domain/shelf/authority"

// Viewer is the authority of viewer.
type Viewer struct{}

var _ authority.Authoritier = (*Viewer)(nil)

// CanDelete returns whether the Viewer can delete.
func (r *Viewer) CanDelete() bool {
	return false
}

// CanInvite implements T.
func (*Viewer) CanInvite() bool {
	return false
}

// CanCreate returns whether the Viewer can create.
func (r *Viewer) CanCreate() bool {
	return false
}

// CanEdit implements T.
func (*Viewer) CanEdit() bool {
	return false
}

// NewViewer is create a new Viewer.
func NewViewer() *Viewer {
	return &Viewer{}
}