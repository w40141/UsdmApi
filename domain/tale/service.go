// Package tale is defined for domain model.
package tale

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Tale.
func Create(
	title string,
	description string,
	reason string,
	storyID string,
	ownerID string,
) (Tale, error) {
	id := vo.NewID().String()
	return New(
		id,
		title,
		description,
		reason,
		storyID,
		ownerID,
		WithCreatedAt(time.Now()),
		WithUpdatedAt(time.Now()),
	)
}

// Update updates a Tale.
func (t Tale) Update(
	title string,
	description string,
	reason string,
	storyID string,
	ownerID string,
) (Tale, error) {
	return New(
		t.id.String(),
		title,
		description,
		reason,
		storyID,
		ownerID,
		WithUpdatedAt(time.Now()),
	)
}
