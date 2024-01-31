// Package legend is defined for domain model.
package legend

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Legend.
func Create(
	title string,
	description string,
	reason string,
) (Legend, error) {
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
	)
}

// Update updates a Legend.
func (p Legend) Update(
	title string,
	description string,
	reason string,
) (Legend, error) {
	return New(
		p.id.String(),
		title,
		description,
		reason,
	)
}
