// Package authority defines services for member authority.
package authority

import (
	"github.com/w40141/UsdmApi/domain/shelf/authority/model"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// New creates a new authority.
func New(t model.AurhorityType) (model.Authoritier, error) {
	switch t {
	case model.ViewerType:
		return model.NewViewer(), nil
	case model.EditorType:
		return model.NewEditor(), nil
	case model.OwnerType:
		return model.NewOwner(), nil
	case model.ManagerType:
		return model.NewManager(), nil
	default:
		return nil, vo.ErrInvalidAuthorityType
	}
}
