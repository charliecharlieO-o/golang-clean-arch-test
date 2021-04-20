package repository

import (
	entity "server/timeline/entity"
)

type TimelineRepository interface {
	Save(t *entity.Timeline) (*entity.Timeline, error)
	Delete(t *entity.Timeline) error
	Get(id uint) (*entity.Timeline, error)
}
