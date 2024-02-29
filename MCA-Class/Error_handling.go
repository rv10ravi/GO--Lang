//You're tasked with creating a program for a highly sensitive government agency that tracks top-secret missions. The program must manage a collection of mission data, including mission names, objectives, and statuses. However, there's a catch: the program must be designed to self-destruct (trigger a panic) if anyone tries to access a mission labeled with a certain codeword. This requirement is in place to prevent unauthorized access to missions of extreme importance.//
//Your task is to design and implement this program in Go, ensuring that:Mission data is stored securely and can only be accessed through proper channels.The program includes error handling mechanisms to gracefully handle any unexpected issues.A panic is triggered if the user tries to access a mission labeled with the specified codeword, but the program should also include a recovery mechanism to prevent the entire program from crashing.Defer statements are used appropriately to clean up resources and ensure the program runs smoothly under normal conditions.

package main

import (
	"fmt"
	"log"
)

type mission struct {
	name      string
	objective string
	status    string
}

func main() {
	missions := map[string]mission{
		"001": {"Mission 001", "Objective 001", "In Progress"},
		"002": {"Mission 002", "Objective 002", "Completed"},
		"003": {"Mission 003", "Objective 003", "In Progress"},
	}

	// Accessing mission data
	accessMission(missions, "001")
	accessMission(missions, "002")
	accessMission(missions, "003")
	accessMission(missions, "004")
}

func accessMission(missions map[string]mission, missionID string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic occurred:", r)
		}
	}()

	if missionID == "004" {
		panic("Access denied: Unauthorized mission")
	}

	mission, ok := missions[missionID]
	if !ok {
		log.Println("Mission not found")
		return
	}

	fmt.Println("Mission Name:", mission.name)
	fmt.Println("Objective:", mission.objective)
	fmt.Println("Status:", mission.status)
	fmt.Println()
}


