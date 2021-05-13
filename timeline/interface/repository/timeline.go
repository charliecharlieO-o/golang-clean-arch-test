package repository

import (
	"errors"
	"server/timeline/entity"
)

type timelineRepository struct {
	timelineIdx   uint
	TimelineTable map[uint]*entity.Timeline
	TagIdx        map[string][]*entity.Timeline
}

type TimelineRepository interface {
	Save(t *entity.Timeline) (*entity.Timeline, error)
	Delete(t *entity.Timeline) error
	Get(t *entity.Timeline) (*entity.Timeline, error)
	SearchByTags(tags []string) ([]*entity.Timeline, error)
}

func NewTimelineRepository() TimelineRepository {
	return &timelineRepository{
		timelineIdx:   0,
		TimelineTable: make(map[uint]*entity.Timeline, 0),
	}
}

func (tr *timelineRepository) addTags(t *entity.Timeline) {
	for i := 0; i < len(t.Tags); i++ {
		tag := t.Tags[i]
		_, exists := tr.TagIdx[tag]
		if exists {
			tr.TagIdx[tag] = append(tr.TagIdx[tag], t)
		} else {
			tr.TagIdx[tag] = []*entity.Timeline{t}
		}
	}
}

func (tr *timelineRepository) Save(t *entity.Timeline) (*entity.Timeline, error) {
	// Add to main table
	t.ID = tr.timelineIdx + 1
	tr.TimelineTable[t.ID] = t
	tr.timelineIdx++
	// Add timeline to tags
	tr.addTags(t)
	return t, nil
}

func (tr *timelineRepository) Delete(t *entity.Timeline) error {
	if t.ID > tr.timelineIdx || t.ID <= 0 {
		return errors.New("Timeline doesn't exist")
	}
	_, exists := tr.TimelineTable[t.ID]
	if !exists {
		return errors.New("Timeline doesn't exist")
	}
	tr.TimelineTable[t.ID] = nil
	return nil
}

func (tr *timelineRepository) Get(t *entity.Timeline) (*entity.Timeline, error) {
	if t.ID > tr.timelineIdx || t.ID <= 0 {
		return nil, errors.New("Timeline doesn't exist")
	}
	timeline, _ := tr.TimelineTable[t.ID]
	if timeline != nil {
		return timeline, nil
	}
	return nil, errors.New("Timeline doesn't exist")
}

func (tr *timelineRepository) SearchByTags(tags []string) ([]*entity.Timeline, error) {
	output := []*entity.Timeline{}
	for i := 0; i < len(tags); i++ {
		timelines, exists := tr.TagIdx[tags[i]]
		if exists {
			copy(output, timelines)
		}
	}
	return output, nil
}
