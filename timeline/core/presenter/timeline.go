package presenter

import (
	entity "server/timeline/entity"
)

type TimelinePresenter interface {
	TimelineError(t *entity.Timeline, e error) error
	TimelineResponse(t *entity.Timeline) *entity.Timeline
	//TimelinesResponse(t []*entity.Timeline) []*entity.Timeline
}
