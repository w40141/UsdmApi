// Package entity is ddefined for domain model.
package entity

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Project is the largest unit of human management.
type Project struct {
	id          vo.ID
	title       vo.Title
	description vo.Description
	ownerID     vo.ID
	createdAt   time.Time
	updatedAt   time.Time
}

// NewProject creates a new Project.
func NewProject(
	title vo.Title,
	description vo.Description,
	ownerID vo.ID,
) Project {
	return Project{
		id:          vo.NewID(),
		title:       title,
		description: description,
		ownerID:     ownerID,
	}
}

// ProjectRepository is a repository interface for Project.
type ProjectRepository interface {
	create(project Project) error
	update(project Project) error
	get(ids []vo.ID) ([]Project, error)
	delete(id vo.ID) error
}
