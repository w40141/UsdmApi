// Package epic is defined for domain model.
package epic

import (
	"time"

	"github.com/w40141/UsdmApi/domain/vo"
)

// Create is create a new Epic.
func Create(
	projectID string,
	title string,
	description string,
	startAt string,
	endAt string,
) (Epic, error) {
	newProjectID, e1 := vo.FromStringToID(projectID)
	if e1 != nil {
		return Epic{}, e1
	}
	newTitle, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Epic{}, e2
	}
	newDescription, e3 := vo.NewDescription(description)
	if e3 != nil {
		return Epic{}, e3
	}
	newStartAt, e4 := vo.FromStringToTime(startAt)
	if e4 != nil {
		return Epic{}, e4
	}
	newEndAt, e5 := vo.FromStringToTime(endAt)
	if e5 != nil {
		return Epic{}, e5
	}
	if newStartAt.After(newEndAt) {
		return Epic{}, vo.ErrInvalidTime
	}
	return New(
		vo.CreateID(),
		newProjectID,
		newTitle,
		newDescription,
		newStartAt,
		newEndAt,
		WithCreatedAt(time.Now()),
		WithUpdatedAt(time.Now()),
	), nil
}

// Update updates a Epic.
func (e Epic) Update(
	projectID string,
	title string,
	description string,
	startAt string,
	endAt string,
) (Epic, error) {
	newProjectID, e1 := vo.FromStringToID(projectID)
	if e1 != nil {
		return Epic{}, e1
	}
	newTitle, e2 := vo.NewTitle(title)
	if e2 != nil {
		return Epic{}, e2
	}
	newDescription, e3 := vo.NewDescription(description)
	if e3 != nil {
		return Epic{}, e3
	}
	newStartAt, e4 := vo.FromStringToTime(startAt)
	if e4 != nil {
		return Epic{}, e4
	}
	newEndAt, e5 := vo.FromStringToTime(endAt)
	if e5 != nil {
		return Epic{}, e5
	}
	if newStartAt.After(newEndAt) {
		return Epic{}, vo.ErrInvalidTime
	}
	return New(
		e.id,
		newProjectID,
		newTitle,
		newDescription,
		newStartAt,
		newEndAt,
		WithCreatedAt(e.createdAt),
		WithUpdatedAt(time.Now()),
	), nil
}
