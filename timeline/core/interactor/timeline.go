package interactor

import (
	"errors"
	"server/timeline/core/presenter"
	"server/timeline/core/repository"
	"server/timeline/entity"
)

const (
	// TimelineMaxDepth marks how deep a story can go
	TimelineMaxDepth = 3
	// OutcomeMax indicates the maximum outcomes per node
	OutcomeMax = 3
	// OutcomeMin indicates the minimum outcomes per node
	OutcomeMin = 2
)

type timelineInteractor struct {
	TimelineRepository repository.TimelineRepository
	TimelinePresenter  presenter.TimelinePresenter
}

type TimelineInteractor interface {
	Validate(t *entity.Timeline) error
	Save(t *entity.Timeline) (*entity.Timeline, error)
	Delete(t *entity.Timeline) error
	Get(id uint) (*entity.Timeline, error)
}

func NewTimelineInteractor(r repository.TimelineRepository, p presenter.TimelinePresenter) TimelineInteractor {
	return &timelineInteractor{r, p}
}

//validateStoryline ... Validates the nodes under the selected root node
func validateStoryline(node *entity.StoryNode, depth int) error {
	outcomeCount := len(node.Outcomes)
	if outcomeCount == 0 && depth == 0 {
		return errors.New("Min depth for timeline is 1")
	}
	if depth > TimelineMaxDepth {
		return errors.New("Max story depth exceeded")
	}
	if len(node.Premise) == 0 {
		return errors.New("Missing premise")
	}
	if outcomeCount != 0 && (outcomeCount < OutcomeMin || outcomeCount > OutcomeMax) {
		return errors.New("Outcomes out of bounds: " + node.Premise)
	}
	for _, node := range node.Outcomes {
		err := validateStoryline(node, depth+1)
		if err != nil {
			return err
		}
	}
	return nil
}

//ValidateTimeline ... Validates all nodes of a timeline tree
func (ti *timelineInteractor) Validate(t *entity.Timeline) error {
	err := validateStoryline(t.Root, 0)
	return ti.TimelinePresenter.TimelineError(t, err)
}

//Save ... Saves a timeline into a timeline repository
func (ti *timelineInteractor) Save(t *entity.Timeline) (*entity.Timeline, error) {
	saved, err := ti.TimelineRepository.Save(t)
	if err != nil {
		return nil, ti.TimelinePresenter.TimelineError(t, err)
	}
	return ti.TimelinePresenter.TimelineResponse(saved), nil
}

func (ti *timelineInteractor) Delete(t *entity.Timeline) error {
	err := ti.TimelineRepository.Delete(t)
	if err != nil {
		return ti.TimelinePresenter.TimelineError(t, err)
	}
	return nil
}

func (ti *timelineInteractor) Get(id uint) (*entity.Timeline, error) {
	timeline, err := ti.TimelineRepository.Get(id)
	if err != nil {
		return nil, ti.TimelinePresenter.TimelineError(nil, err)
	}
	return ti.TimelinePresenter.TimelineResponse(timeline), nil
}
