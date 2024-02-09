// Package authority defines the entity object for user.
package authority

// Authoritier is an interface for user authority.
type Authoritier interface {
	CanDelete() bool
	CanInvite() bool
	CanCreate() bool
	CanEdit() bool
}
