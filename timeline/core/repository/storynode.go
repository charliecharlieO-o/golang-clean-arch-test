package repository

import (
	entity "server/timeline/entity"
)

type StoryNodeRepository interface {
	Save(n *entity.StoryNode) (*entity.StoryNode, error)
	Delete(n *entity.StoryNode) error
	Get(n *entity.StoryNode) (*entity.StoryNode, error)
}
