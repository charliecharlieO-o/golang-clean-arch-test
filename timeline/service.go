package timeline

import (
	"server/timeline/core/interactor"
	tp "server/timeline/core/presenter"
	tr "server/timeline/core/repository"
	"server/timeline/interface/controller"
	ip "server/timeline/interface/presenter"
	ir "server/timeline/interface/repository"
)

type timelineService struct{}

type TimelineService interface {
	NewTimelineController() controller.TimelineController
}

func NewTimelineService() TimelineService {
	return &timelineService{}
}

func (ts *timelineService) NewTimelineController() controller.TimelineController {
	return controller.NewTimelineController(ts.NewTimelineInteractor())
}

func (ts *timelineService) NewTimelineInteractor() interactor.TimelineInteractor {
	return interactor.NewTimelineInteractor(ts.NewTimelineRepository(), ts.NewTimelinePresenter())
}

func (ts *timelineService) NewTimelineRepository() tr.TimelineRepository {
	return ir.NewTimelineRepository()
}

func (ts *timelineService) NewTimelinePresenter() tp.TimelinePresenter {
	return ip.NewTimelinePresenter()
}
