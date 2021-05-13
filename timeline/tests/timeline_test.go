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

	tl, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Errorf(`Timeline Build Failed: %q`, err)
	}
	err = controller.ValidateTimeline(tl)
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

	tl, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Errorf(`Timeline Build Failed: %q`, err)
	}
	err = controller.ValidateTimeline(tl)
	if err == nil {
		t.Error("Invalid timeline passed validation")
	}
	if err.Error() != "Test Timeline Error Min depth for timeline is 1" {
		t.Errorf(`Wrong validation error: %q`, err)
	}
}

func TestValidateTimelineSave(t *testing.T) {
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
	tl, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Error("Invalid timeline")
	}
	tl, err = controller.Save(tl)
	if tl.ID != 1 || err != nil {
		t.Error("Timeline was not saved correctly")
	}
}

func TestTimelineDelete(t *testing.T) {
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
	tl, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Error("Invalid timeline")
	}
	tl, err = controller.Save(tl)
	if tl.ID != 1 || err != nil {
		t.Error("Timeline was not saved correctly")
	}
	err = controller.Delete(tl)
	if err != nil {
		t.Errorf(`Timeline was not deleted: %q`, err)
	}
}

func TestGetTimeline(t *testing.T) {
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
	tl, err := controller.BuildTimeline([]byte(story))
	if err != nil {
		t.Error("Invalid timeline")
	}
	tl, err = controller.Save(tl)
	if tl.ID != 1 || err != nil {
		t.Error("Timeline was not saved correctly")
	}
	tl, err = controller.Get(tl)
	if tl == nil || err != nil {
		t.Errorf(`Timeline was not found: %q`, err)
	}
}
