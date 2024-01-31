// Package story defines the entity object for user story.
package story

import "github.com/w40141/UsdmApi/domain/vo"

// Create creates a new Story.
func Create(
	title string,
	description string,
	reason string,
	legendID string,
	parentID string,
) (Story, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
		legendID,
		parentID,
	)
}

// Update updates a Story.
func (s Story) Update(
	title string,
	description string,
	reason string,
	parentID string,
) (Story, error) {
	return New(s.id.String(),
		title,
		description,
		reason,
		s.legendID.String(),
		parentID,
	)
}
