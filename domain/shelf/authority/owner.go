// Package authority defines the entity object for user authority.
package authority

// Owner is the authority of owner.
type Owner struct{}

var _ T = (*Owner)(nil)

// CanCreate returns whether the owner can create.
func (o *Owner) CanCreate() bool {
	return o != nil
}

// CanDelete returns whether the owner can delete.
func (o *Owner) CanDelete() bool {
	return o != nil
}

// NewOwner is create a new Owner.
func NewOwner() *Owner {
	return &Owner{}
}
