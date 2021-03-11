package main

import (
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

	tservice := tmln.NewTimelineService()
	controller := tservice.NewTimelineController()

	// Loading the timeline into a Timeline struct and validating it
	timeline, _ := controller.BuildTimeline([]byte(storyJSON))
	verr := controller.ValidateTimeline(timeline)

	if verr != nil {
		fmt.Println("Invalid timeline", timeline)
	}

	// Save Timeline
	controller.Save(timeline)

	// MiniGame - Do this until we run out of possibilities
	fmt.Println("Want to play?...")

	// Get first premise
	currentNode := timeline.Root
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
