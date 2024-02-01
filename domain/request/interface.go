// Package request is defined for domain model.
package request

// ParentOfStory is a type that can be the parent of story.
type ParentOfStory interface {
	ParentOfStory() error
	ID() string
}

// LegendType is a type of legend.
type LegendType interface {
	ID() string
	Description() string
	Title() string
	ParentOfStory() error
	ParenOfNarrative() error
	ParentOfEpisode() error
	ParenOfScene() error
}

// NarrativeType is a type of narrative.
type NarrativeType interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParentOfStory() error
	ParenOfScene() error
}

// StoryType is a type of story.
type StoryType interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParenOfScene() error
	ParenOfEpisode() error
}

// EpisodeType is a type of episode.
type EpisodeType interface {
	ID() string
	Description() string
	Title() string
	Reason() string
	ParenOfScene() error
}

// SceneType is a type of scene.
type SceneType interface {
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
