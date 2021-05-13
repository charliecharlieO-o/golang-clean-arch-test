package repository

import (
	entity "server/timeline/entity"
)

type TimelineRepository interface {
	Save(t *entity.Timeline) (*entity.Timeline, error)
	Delete(t *entity.Timeline) error
	Get(t *entity.Timeline) (*entity.Timeline, error)
	SearchByName(name string) ([]*entity.Timeline, error)
	//SearchByTags(tags []string) ([]*entity.Timeline, error)
}
