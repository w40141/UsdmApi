// Package authority defines the entity object for user authority.
package authority

// T is an entity object for user Authority or usecase.
type T interface {
	CanCreate() bool
	CanDelete() bool
}
