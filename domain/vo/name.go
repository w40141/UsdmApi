// Package vo is defined for value object.
package vo

// Name is a value object for Name.
type Name struct {
	value string
}

// NewName creates a new Name.
func NewName(name string) (Name, error) {
	if name == "" {
		return Name{}, ErrNoName
	}
	return Name{value: name}, nil
}

func (n Name) String() string {
	return n.value
}
