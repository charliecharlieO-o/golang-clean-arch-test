package entity

//StoryNode ... The game's basic unit
type StoryNode struct {
	ID          uint
	Premise     string
	Outcomes    map[string]*StoryNode
	SideEffects []interface{}
}
