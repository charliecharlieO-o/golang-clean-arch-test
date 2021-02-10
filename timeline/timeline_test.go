package timeline

import "testing"

func TestValidateTimelineBasic(t *testing.T) {
	tl := Timeline{Name: "Test", Description: "test"}
	tl.Root = &StoryNode{
		Premise:  "Testing timeline",
		Outcomes: make(map[string]*StoryNode),
	}
	tl.Root.Outcomes["A"] = &StoryNode{
		Premise: "Consequence A",
	}
	tl.Root.Outcomes["B"] = &StoryNode{
		Premise: "Consequence B",
	}
	err := tl.ValidateTimeline()
	if err != nil {
		t.Errorf(`Validation failed: %q`, err)
	}
}

func TestValidateTimelineNoTree(t *testing.T) {
	tl := Timeline{Name: "Test", Description: "test"}
	tl.Root = &StoryNode{
		Premise: "Testing timeline",
	}
	err := tl.ValidateTimeline()
	if err == nil {
		t.Error("Invalid timeline passed validation")
	}
	if err.Error() != "Min depth for timeline is 1" {
		t.Errorf(`Wrong validation error: %q`, err)
	}
}

// TODO: Test max story depth exceeded
// TODO: Test max story depth limit
// TODO: Test empty premise in a node
// TODO: Test max outcomes exceeded
// TODO: Test outcomes required
