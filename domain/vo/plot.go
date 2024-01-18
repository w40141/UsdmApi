// Package vo is defined for value object.
package vo

// Plot is a value object for Plot.
type Plot struct {
	title       Title
	description Description
	reason      Description
}

// NewPlot creates a new Plot.
func NewPlot(title string, description string, reason string) (Plot, error) {
	newTitle, e1 := NewTitle(title)
	if e1 != nil {
		return Plot{}, e1
	}
	newDescription, e2 := NewDescription(description)
	if e2 != nil {
		return Plot{}, e2
	}
	newReason, e3 := NewDescription(reason)
	if e3 != nil {
		return Plot{}, e3
	}
	return Plot{
		title:       newTitle,
		description: newDescription,
		reason:      newReason,
	}, nil
}
