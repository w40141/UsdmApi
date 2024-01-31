// Package narrative is defined for domain model.
package narrative

import (
	"errors"

	"github.com/w40141/UsdmApi/domain/request/legend"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Create is create a new Narrative.
func Create(
	title string,
	description string,
	reason string,
	legendEntity legend.Type,
	parent interface{},
) (Narrative, error) {
	parentID := ""
	switch p := parent.(type) {
	case legend.Type:
		parentID = p.ID()
	case Type:
		parentID = p.ID()
	default:
		return Narrative{}, errors.New("")
	}
	return New(
		vo.NewID().String(),
		title,
		description,
		reason,
		legendEntity.ID(),
		parentID,
	)
}

// Update updates a Narrative.
func (n Narrative) Update(
	title string,
	description string,
	reason string,
	parent interface{},
) (Narrative, error) {
	parentID := ""
	switch p := parent.(type) {
	case legend.Type:
		parentID = p.ID()
	case Type:
		parentID = p.ID()
	default:
		return Narrative{}, errors.New("")
	}
	return New(
		n.id.String(),
		title,
		description,
		reason,
		n.legendID.String(),
		parentID,
	)
}
