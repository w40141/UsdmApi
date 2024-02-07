// Package episode is defined for domain model.
package episode

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Episode.
type Repository interface {
	Save(episode T) error
	Get(ids []vo.ID) ([]T, error)
	Delete(id vo.ID) error
}
