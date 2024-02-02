// Package member defines the entity object for user member.
package member

import (
	"fmt"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Member is an entity object for user member.
type Member struct {
	name vo.Name
	id   vo.ID
}

// ID getter
func (m Member) ID() string {
	return m.id.String()
}

// Name getter
func (m Member) Name() string {
	return m.name.String()
}

// Create creates a new Member.
func Create(name string) (Member, error) {
	id := vo.NewID().String()
	return New(id, name)
}

// Update updates a Member.
func (m *Member) Update(name string) (Member, error) {
	if m == nil {
		return Member{}, fmt.Errorf("member is nil")
	}
	return New(m.ID(), name)
}

// New creates a new Member.
func New(id string, name string) (Member, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return Member{}, e1
	}
	nameVo, e2 := vo.NewName(name)
	if e2 != nil {
		return Member{}, e2
	}
	m := Member{
		id:   idVo,
		name: nameVo,
	}
	return m, nil
}
