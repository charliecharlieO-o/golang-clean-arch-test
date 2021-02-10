package timeline

import "errors"

const (
	// StoryMaxDepth marks how deep a story can go
	StoryMaxDepth = 3
	// OutcomeMax indicates the maximum outcomes per node
	OutcomeMax = 3
	// OutcomeMin indicates the minimum outcomes per node
	OutcomeMin = 2
)

//Timeline ... Structure made out of story nodes
type Timeline struct {
	Name        string
	Description string
	Root        *StoryNode
}

//StoryNode ... The game's basic unit
type StoryNode struct {
	Premise     string
	Outcomes    map[string]*StoryNode
	SideEffects []interface{}
}

//ValidateStoryline ... Validates the nodes under the selected root node
func ValidateStoryline(node *StoryNode, depth int) error {
	outcomeCount := len(node.Outcomes)
	if outcomeCount == 0 && depth == 0 {
		return errors.New("Min depth for timeline is 1")
	}
	if depth > StoryMaxDepth {
		return errors.New("Max story depth exceeded")
	}
	if len(node.Premise) == 0 {
		return errors.New("Missing premise")
	}
	if outcomeCount != 0 && (outcomeCount < OutcomeMin || outcomeCount > OutcomeMax) {
		return errors.New("Outcomes out of bounds for premise: " + node.Premise)
	}
	for _, node := range node.Outcomes {
		err := ValidateStoryline(node, depth+1)
		if err != nil {
			return err
		}
	}
	return nil
}

//ValidateTimeline ... Validates all nodes of a timeline tree
func (t *Timeline) ValidateTimeline() error {
	return ValidateStoryline(t.Root, 0)
}

/* TODO
func (t *Timeline) loadNodeIndex() {
}

func (t *Timeline) AddNode(parent *StoryNode, child *StoryNode) error {
}

func (t *Timeline) RemoveNode(node *StoryNode) error {
}
*/

//ApplySideEffects ... Passes side effects to an applier function to be applied
func (n *StoryNode) ApplySideEffects(applier func([]interface{}) (bool, error)) (bool, error) {
	return applier(n.SideEffects)
}
