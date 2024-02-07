// Package authority defines the entity object for user authority.
package authority

import "github.com/w40141/UsdmApi/domain/shelf"

var _ shelf.Authoritier = (*Owner)(nil)

// Owner is the authority of owner.
type Owner struct{}

// CanDelete returns whether the owner can delete.
func (t *Owner) CanDelete() bool {
	return t != nil
}

// CanInvite returns whether the owner can invite.
func (t *Owner) CanInvite() bool {
	return t != nil
}

// CanCreate returns whether the owner can create.
func (t *Owner) CanCreate() bool {
	return t != nil
}

// CanEdit returns whether the owner can edit.
func (t *Owner) CanEdit() bool {
	return t != nil
}

// NewOwner is create a new Owner.
func NewOwner() *Owner {
	return &Owner{}
}
