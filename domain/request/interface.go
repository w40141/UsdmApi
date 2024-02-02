// Package request is defined for domain model.
package request

// ParentOfStory is a type that can be the parent of story.
type ParentOfStory interface {
	ParentOfStory() error
	ID() string
}

// Legender is a type of legend.
type Legender interface {
	ID() (string, error)
	ParentOfStory() error
	ParentOfNarrative() error
	ParentOfEpisode() error
	ParentOfScene() error
}

// Narrativer is a type of narrative.
type Narrativer interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParentOfStory() error
	ParentOfScene() error
}

// Storyer is a type of story.
type Storyer interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParentOfScene() error
	ParentOfEpisode() error
}

// Episoder is a type of episode.
type Episoder interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParentOfScene() error
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
	ParentOfScene() error
	ID() string
}

// ParentOfEpisode is a type that can be the parent of episode.
type ParentOfEpisode interface {
	ParentOfEpisode() error
	ID() string
}
