// Package story defines the entity object for user story.
package story

import "github.com/w40141/UsdmApi/domain/vo"

// Repository is a repository for Story.
type Repository interface {
	Save(story Story) error
	Get(ids []vo.ID) ([]Story, error)
	Delete(id vo.ID) error
}
