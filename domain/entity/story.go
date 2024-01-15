// Package entity is ddefined for domain model.
package entity

import "github.com/w40141/ProManeApp/domain/vo"

// Story is an entity object for user story or usecase.
type Story struct {
	title  vo.Title
	reason vo.Description
}

// NewStory creates a new Story.
func NewStory(title string, reason string) (Story, error) {
	newTitle, e1 := vo.NewTitle(title)
	if e1 != nil {
		return Story{}, e1
	}
	newReason, e2 := vo.NewDescription(reason)
	if e2 != nil {
		return Story{}, e2
	}
	return Story{
		title:  newTitle,
		reason: newReason,
	}, nil
}
