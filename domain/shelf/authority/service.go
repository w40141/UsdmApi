// Package authority defines services for member authority.
package authority

import (
	"github.com/w40141/UsdmApi/domain/shelf/authority/types"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// New creates a new authority.
func New(t types.Type) (types.Authoritier, error) {
	switch t {
	case types.ViewerType:
		return types.NewViewer(), nil
	case types.EditorType:
		return types.NewEditor(), nil
	case types.OwnerType:
		return types.NewOwner(), nil
	case types.ManagerType:
		return types.NewManager(), nil
	default:
		return nil, vo.ErrInvalidAuthorityType
	}
}
