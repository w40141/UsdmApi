// Package episode is defined for domain model.
package episode

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Legend.
type Repository interface {
	Save(legend Episode) error
	Get(ids []vo.ID) ([]Episode, error)
	Delete(id vo.ID) error
}
