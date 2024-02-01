// Package narrative is defined for domain model.
package narrative

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Project.
type Repository interface {
	Save(narrative Narrative) error
	Get(ids []vo.ID) ([]Narrative, error)
	Delete(id vo.ID) error
}
