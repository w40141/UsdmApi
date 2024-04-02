// Package model defines the entity object for user authority.
package model

// Owner is the authority of owner.
type Owner struct{}

var _ Authoritier = (*Owner)(nil)

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

// Value returns the authority type.
func (t *Owner) Value() AurhorityType {
	return OwnerType
}

// NewOwner is create a new owner.
func NewOwner() *Owner {
	return &Owner{}
}
