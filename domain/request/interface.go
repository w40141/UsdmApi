// Package request is defined for domain model.
package request

// ParentOfStory is a type that can be the parent of story.
type ParentOfStory interface {
	ParentOfStory() error
	ID() string
}

// Legender is a type of legend.
type Legender interface {
	ID() string
	Description() string
	Title() string
	ParentOfStory() error
	ParenOfNarrative() error
	ParentOfEpisode() error
	ParenOfScene() error
}

// Narrativer is a type of narrative.
type Narrativer interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParentOfStory() error
	ParenOfScene() error
}

// Storyer is a type of story.
type Storyer interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParenOfScene() error
	ParenOfEpisode() error
}

// Episoder is a type of episode.
type Episoder interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParenOfScene() error
}

// Scener is a type of scene.
type Scener interface {
	ID() string
	Description() string
	Title() string
	Reason() string
}

// ParentOfScene is a type that can be the parent of scene.
type ParentOfScene interface {
	ParenOfScene() error
	ID() string
}

// ParentOfEpisode is a type that can be the parent of episode.
type ParentOfEpisode interface {
	ParenOfEpisode() error
	ID() string
}
