package presenter

import (
	"errors"
	entity "server/timeline/entity"
)

type timelinePresenter struct {
}

type TimelinePresenter interface {
	TimelineError(t *entity.Timeline, e error) error
	TimelineResponse(t *entity.Timeline) *entity.Timeline
}

func NewTimelinePresenter() TimelinePresenter {
	return &timelinePresenter{}
}

func (tp *timelinePresenter) TimelineError(t *entity.Timeline, e error) error {
	if e == nil {
		return nil
	}
	return errors.New(t.Name + " is not a valid timeline: " + e.Error())
}

func (tp *timelinePresenter) TimelineResponse(t *entity.Timeline) *entity.Timeline {
	return t
}
