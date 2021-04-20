package controller

import (
	"encoding/json"
	"server/timeline/core/interactor"
	"server/timeline/entity"
)

type timelineController struct {
	timelineInteractor interactor.TimelineInteractor
}

type TimelineController interface {
	ValidateTimeline(t *entity.Timeline) error
	BuildTimeline(data []byte) (*entity.Timeline, error)
	Save(t *entity.Timeline) (*entity.Timeline, error)
	Delete(t *entity.Timeline) error
	Get(id uint) (*entity.Timeline, error)
}

func NewTimelineController(ti interactor.TimelineInteractor) TimelineController {
	return &timelineController{ti}
}

func (tc *timelineController) BuildTimeline(data []byte) (*entity.Timeline, error) {
	var timeline entity.Timeline
	err := json.Unmarshal(data, &timeline)
	if err != nil {
		return nil, err
	}
	return &timeline, nil
}

func (tc *timelineController) ValidateTimeline(t *entity.Timeline) error {
	return tc.timelineInteractor.Validate(t)
}

func (tc *timelineController) Save(t *entity.Timeline) (*entity.Timeline, error) {
	return tc.timelineInteractor.Save(t)
}

func (tc *timelineController) Delete(t *entity.Timeline) error {
	return tc.timelineInteractor.Delete(t)
}

func (tc *timelineController) Get(id uint) (*entity.Timeline, error) {
	return tc.timelineInteractor.Get(id)
}
