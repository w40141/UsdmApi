// Package project is defined for domain model.
package project

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Project.
func Create(title string, description string, ownerID string) (Project, error) {
	titleVo, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Project{}, e1
	}
	descriptionVo, e2 := vo.NewDescription(description)
	if e2 != nil {
		return Project{}, e2
	}
	ownerIDVo, e3 := vo.FromStringToID(ownerID)
	if e3 != nil {
		return Project{}, e3
	}
	return New(
		vo.CreateID(),
		titleVo,
		descriptionVo,
		ownerIDVo,
		WithCreatedAt(time.Now()),
		WithUpdatedAt(time.Now()),
	), nil
}

// Update updates a Project.
func (p Project) Update(title string, description string, ownerID string) (Project, error) {
	titleVo, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Project{}, e1
	}
	descriptionVo, e2 := vo.NewDescription(description)
	if e2 != nil {
		return Project{}, e2
	}
	ownerIDVo, e3 := vo.FromStringToID(ownerID)
	if e3 != nil {
		return Project{}, e3
	}
	return New(
		p.id,
		titleVo,
		descriptionVo,
		ownerIDVo,
		WithCreatedAt(p.createdAt),
		WithUpdatedAt(time.Now()),
	), nil
}
