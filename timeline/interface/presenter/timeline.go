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
	if t != nil {
		return errors.New(t.Name + " Timeline Error " + e.Error())
	}
	return errors.New("Timeline Error " + e.Error())
}

func (tp *timelinePresenter) TimelineResponse(t *entity.Timeline) *entity.Timeline {
	return t
}
