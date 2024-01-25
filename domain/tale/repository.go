// Package tale is defined for domain model.
package tale

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Tale.
type Repository interface {
	Save(tale Tale) error
	Get(ids []vo.ID) ([]Tale, error)
	Delete(id vo.ID) error
}
