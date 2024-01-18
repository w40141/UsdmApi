// Package entity is ddefined for domain model.
package entity

import (
	"github.com/w40141/UsdmApi/domain/vo"
)

// Requirement is an entity object for user story or usecase.
type Requirement struct {
	id   vo.ID
	plot vo.Plot
}

// NewRequirement creates a new Requirement.
func NewRequirement(title string, description string, reason string) (Requirement, error) {
	newPlot, e1 := vo.NewPlot(title, description, reason)
	if e1 != nil {
		return Requirement{}, e1
	}
	return Requirement{
		id:   vo.NewID(),
		plot: newPlot,
	}, nil
}
