package entity

//Timeline ... Structure made out of story nodes
type Timeline struct {
	ID          uint
	Name        string
	Description string
	Root        *StoryNode
}
