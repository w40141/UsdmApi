// Package saga defines types.
package saga

import (
	"time"

	"github.com/w40141/UsdmApi/domain/shelf/vo"
)

// E is an entity object for Saga.
type E struct {
	id          vo.SagaID
	title       vo.Title
	description vo.Sentence
	creator     vo.MemberID
	createAt    *time.Time
	updateAt    *time.Time
}

// C is a struct for creating a new Saga.
type C struct {
	title       vo.Title
	description vo.Sentence
	creator     vo.MemberID
}

// D is a struct for deleting a Story.
type D struct {
	id vo.ID
}
