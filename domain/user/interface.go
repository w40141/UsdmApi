// Package user defines the entity object for user.
package user

// Authoritier is an interface for user authority.
type Authoritier interface {
	CanDelete() bool
	CanInvite() bool
	CanCreate() bool
	CanEdit() bool
}
