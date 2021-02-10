package main

import (
	"encoding/json"
	"fmt"
	tmln "server/timeline"
)

func main() {
	// Timeline saved in a JSON
	storyJSON := `{
		"name": "Sample",
		"description": "an example of a story",
		"root": {
			"premise": "alguien enveneno el abrevadero!",
			"outcomes": {
				"limpiarlo": {"premise": "$$$"},
				"dejarlo": {"premise": "muerte"}
			}
		}
	}`

	// Loading the timeline into a Timeline struct and validating it
	var story tmln.Timeline
	err := json.Unmarshal([]byte(storyJSON), &story)
	err = story.ValidateTimeline()
	if err == nil {
		fmt.Println("Valid Timeline :D")
	}

	// MiniGame - Do this until we run out of possibilities
	fmt.Println("Want to play?...")

	// Get first premise
	currentNode := story.Root
	for currentNode != nil && len(currentNode.Outcomes) > 0 {
		// Display Premise
		fmt.Println(currentNode.Premise)
		for choice, _ := range currentNode.Outcomes {
			fmt.Println(" - ", choice)
		}
		// Receive input
		var decision string
		fmt.Println("What will it be peasant?")
		fmt.Print("> ")
		fmt.Scanln(&decision)
		// Get next premise
		currentNode, _ = currentNode.Outcomes[decision]
	}
	if currentNode != nil {
		fmt.Println("Final outcome: " + currentNode.Premise)
	}
	fmt.Println("You reached the end of the game :(")
}
