// Package scene is defined for domain model.
package scene

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// Create creates a new Scene.
func Create(
	title string,
	description string,
	reason string,
	book shelf.Booker,
	parent shelf.ParentOfScene,
) (C, error) {
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return C{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return C{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return C{}, e4
	}
	return C{
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      book.ID(),
		parentID:    parent.ID(),
	}, nil
}

// NewT creates a new Scene.
func NewT(
	id string,
	title string,
	description string,
	reason string,
	bookID string,
	storyID string,
	createdAt time.Time,
	updatedAt time.Time,
) (T, error) {
	idVo, e1 := vo.FromStringToID(id)
	if e1 != nil {
		return T{}, nil
	}
	titleVo, e2 := vo.NewTitle(title)
	if e2 != nil {
		return T{}, e2
	}
	descriptionVo, e3 := vo.NewSentence(description)
	if e3 != nil {
		return T{}, e3
	}
	reasonVo, e4 := vo.NewSentence(reason)
	if e4 != nil {
		return T{}, e4
	}
	bookIDVo, e5 := vo.FromStringToID(bookID)
	if e5 != nil {
		return T{}, e5
	}
	storyIDVo, e6 := vo.FromStringToID(storyID)
	if e6 != nil {
		return T{}, e6
	}
	return T{
		id:          idVo,
		title:       titleVo,
		description: descriptionVo,
		reason:      reasonVo,
		bookID:      bookIDVo,
		parentID:    storyIDVo,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}
