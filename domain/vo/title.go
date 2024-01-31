// Package vo is defined for value object.
package vo

// Title is a value object for title.
// Title must be more than 0 and less than or equal to 100 characters.
type Title struct {
	value string
}

// NewTitle creates a new Title.
func NewTitle(value string) (Title, error) {
	if value == "" {
		return Title{}, ErrNoTitle
	}
	if len(value) > 100 {
		return Title{}, ErrTitleOverLength
	}
	return Title{value: value}, nil
}

func (t Title) String() string {
	return t.value
}
