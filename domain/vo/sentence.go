// Package vo is defined for value object.
package vo

// Sentence is a value object for sentence.
// Sentence must be less than or equal to 1000 characters.
type Sentence struct {
	value string
}

// NewSentence creates a new Description.
func NewSentence(value string) (Sentence, error) {
	if len(value) > 1000 {
		return Sentence{}, ErrDescriptionOverLength
	}
	return Sentence{value: value}, nil
}

func (s Sentence) String() string {
	return s.value
}
