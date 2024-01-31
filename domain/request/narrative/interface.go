// Package narrative is defined for domain model.
package narrative

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Project.
type Repository interface {
	Save(narrative Narrative) error
	Get(ids []vo.ID) ([]Narrative, error)
	Delete(id vo.ID) error
}

// Type is a interface for Narrative.
type Type interface {
	Narrative() error
	Story() error
	Episode() error
	Scene() error
	ID() string
	Title() string
	Description() string
	Reason() string
}
