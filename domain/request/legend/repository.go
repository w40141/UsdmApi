// Package legend is defined for domain model.
package legend

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository interface for Legend.
type Repository interface {
	Save(legend Legend) error
	Get(ids []vo.ID) ([]Legend, error)
	Delete(id vo.ID) error
}
