// Package story defines the entity object for user story.
package story

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf/scene"
	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// E is an entity object for story.
type E struct {
	id          vo.StoryID
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	book        vo.BookID
	sences      []scene.T
	// ancestor    *vo.StoryID
	// descendant  *vo.StoryID
	creator  vo.AccountID
	createAt *time.Time
	updateAt *time.Time
}

// C is a struct for creating a new story.
type C struct {
	title       vo.Title
	description vo.Sentence
	reason      vo.Sentence
	book        vo.BookID
	creator     vo.AccountID
}

// D is a struct for deleting a story.
type D struct {
	id vo.ID
}
