// Package entity is ddefined for domain model.
package entity

import "github.com/w40141/UsdmApi/domain/vo"

// Epic is an entity object for user story or usecase.
type Epic struct {
	id          vo.ID
	projectID   vo.ID
	title       vo.Title
	description vo.Description
}

// NewEpic creates a new Epic.
func NewEpic(project Project, title string, description string) (Epic, error) {
	newTitle, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Epic{}, e1
	}
	newDescription, e2 := vo.NewDescription(description)
	if e2 != nil {
		return Epic{}, e2
	}
	return Epic{
		id:          vo.NewID(),
		projectID:   project.id,
		title:       newTitle,
		description: newDescription,
	}, nil
}
