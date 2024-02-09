// Package model defines the entity object for user authority.
package model

import "github.com/w40141/UsdmApi/domain/shelf/authority"

// Owner is the authority of owner.
type Owner struct{}

var _ authority.Authoritier = (*Owner)(nil)

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
