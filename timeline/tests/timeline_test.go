package timeline

import (
	"server/timeline"
	"testing"
)

func TestValidateTimelineBuild(t *testing.T) {
	story := `{
		"name": "Test",
		"description": "an example of a story",
		"root": {
			"premise": "Root Premise",
			"outcomes": {
				"A": {"premise": "Consequence A"},
				"B": {"premise": "Consequence B"}
			}
		}
	}`
	service := timeline.NewTimelineService()
	controller := service.NewTimelineController()

	timeline, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Errorf(`Timeline Build Failed: %q`, err)
	}
	err = controller.ValidateTimeline(timeline)
	if err != nil {
		t.Errorf(`Validation Failed: %q`, err)
	}
}

func TestValidateTimelineNoTree(t *testing.T) {
	story := `{
		"name": "Test",
		"description": "an example of a story",
		"root": {
			"premise": "Root Premise"
		}
	}`
	service := timeline.NewTimelineService()
	controller := service.NewTimelineController()

	timeline, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Errorf(`Timeline Build Failed: %q`, err)
	}
	err = controller.ValidateTimeline(timeline)
	if err == nil {
		t.Error("Invalid timeline passed validation")
	}
	if err.Error() != "Test is not a valid timeline: Min depth for timeline is 1" {
		t.Errorf(`Wrong validation error: %q`, err)
	}
}
