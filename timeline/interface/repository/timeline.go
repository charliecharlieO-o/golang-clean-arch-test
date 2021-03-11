package repository

import (
	"server/timeline/entity"
)

type timelineRepository struct {
	timelineIdx   uint
	storynodeIdx  uint
	TimelineTable map[uint]*entity.Timeline
}

type TimelineRepository interface {
	Save(t *entity.Timeline) (*entity.Timeline, error)
}

func NewTimelineRepository() TimelineRepository {
	return &timelineRepository{
		timelineIdx:   0,
		storynodeIdx:  0,
		TimelineTable: make(map[uint]*entity.Timeline, 0),
	}
}

func (tr *timelineRepository) Save(t *entity.Timeline) (*entity.Timeline, error) {
	t.ID = tr.timelineIdx + 1
	tr.TimelineTable[t.ID] = t
	return t, nil
}
