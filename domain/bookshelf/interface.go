// Package bookshelf is defined for domain model.
package bookshelf

import "github.com/w40141/UsdmApi/domain/vo"

// Booker is a type of legend.
type Booker interface {
	ID() vo.ID
	ParentOfStory() error
	ParentOfEpisode() error
	ParentOfScene() error
}

// Narrativer is a type of narrative.
type Narrativer interface {
	ID() vo.ID
	ParentOfStory() error
	ParentOfScene() error
}

// Storyer is a type of story.
type Storyer interface {
	ID() vo.ID
	ParentOfScene() error
	ParentOfEpisode() error
}

// Episoder is a type of episode.
type Episoder interface {
	ID() vo.ID
	ParentOfScene() error
}

// Scener is a type of scene.
type Scener interface {
	ID() vo.ID
}

// ParentOfStory is a type that can be the parent of story.
type ParentOfStory interface {
	ParentOfStory() error
	ID() vo.ID
}

// ParentOfScene is a type that can be the parent of scene.
type ParentOfScene interface {
	ParentOfScene() error
	ID() vo.ID
}

// ParentOfEpisode is a type that can be the parent of episode.
type ParentOfEpisode interface {
	ParentOfEpisode() error
	ID() vo.ID
}

type Deletable interface {
	ID() vo.ID
	CanDelete() error
}
