// Package shelf is defined for domain model.
package shelf

import "github.com/w40141/UsdmApi/domain/vo"

// Booker is a type of book.
type Booker interface {
	ParentOfStory
	ParentOfEpisode
	ParentOfScene
}

// Narrativer is a type of narrative.
type Narrativer interface {
	ParentOfStory
	ParentOfScene
}

// Storyer is a type of story.
type Storyer interface {
	ParentOfEpisode
	ParentOfScene
}

// Episoder is a type of episode.
type Episoder interface {
	ParentOfScene
}

// Scener is a type of scene.
type Scener interface {
	Parent
}

// ParentOfStory is a type that can be the parent of story.
type ParentOfStory interface {
	ParentOfStory() error
	Parent
}

// ParentOfEpisode is a type that can be the parent of episode.
type ParentOfEpisode interface {
	ParentOfEpisode() error
	Parent
}

// ParentOfScene is a type that can be the parent of scene.
type ParentOfScene interface {
	ParentOfScene() error
	Parent
}

// Parent is a type that can be the parent.
type Parent interface {
	ID() vo.ID
}

// Participanter is a type that can be a participant.
type Participanter interface {
	CanDelete() bool
	CanInvite() bool
	CanCreate() bool
	CanEdit() bool
}
