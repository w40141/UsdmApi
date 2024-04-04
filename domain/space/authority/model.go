// Package authority defines the authorities of a user in a space.
package authority

// Type is an int representing authorities.
// It is a 4-bit integer, with each bit representing a different authority.
// The rightmost bit is the viewer authority, the next bit is the editor authority, and so on.
type Type int

const (
	// OwnerType is the authority of owner.
	OwnerType Type = 0b1111
	// AdminType is the authority of admin.
	AdminType Type = 0b0111
	// ManagerType is the authority of manager.
	ManagerType Type = 0b0011
	// EditorType is the authority of editor.
	EditorType Type = 0b0001
	// ViewerType is the authority of viewer.
	ViewerType Type = 0b0000

	canDelete Type = 0b1000
	canInvite Type = 0b0100
	canCreate Type = 0b0010
	canEdit   Type = 0b0001
)

// E is a struct with a Type field.
type E struct {
	power Type
}
