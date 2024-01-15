// Package entity is ddefined for domain model.
package entity

import (
	"github.com/w40141/ProManeApp/domain/vo"
)

// Project is the largest unit of human management.
type Project struct {
	id          vo.ID
	title       vo.Title
	description vo.Description
	// addreviator vo.Addreviator
}

// NewProject creates a new Project.
func NewProject(
	title string,
	description string,
	// addreviator string,
) (Project, error) {
	newTitle, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Project{}, e1
	}
	newDescription, e2 := vo.NewDescription(description)
	if e2 != nil {
		return Project{}, e2
	}

	return Project{
		id:          vo.NewID(),
		title:       newTitle,
		description: newDescription,
		// addreviator: addreviator,
	}, nil
}
