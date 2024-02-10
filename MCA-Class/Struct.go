/* A team of archaeologists exploring an ancient temple stumbles upon a series of sealed chambers, each containing valuable artifacts and clues to unlock the next chamber. These chambers are arranged in a linear fashion, with each chamber containing a unique artifact and a riddle or puzzle to solve. The team must use their wits and knowledge of ancient civilizations to navigate through the chambers and retrieve the ultimate treasure hidden within the final chamber.
The team possesses an array named chambers where each element represents a chamber. Each chamber is represented by an object with the following properties:artifact: A description of the valuable artifact found within the chamber.puzzle: A riddle or puzzle presented to the team to unlock the next chamber.solved: A boolean value indicating whether the puzzle of the chamber has been solved or not.
The team's task is to create a function exploreTemple(chambers) that takes in the array of chambers and iterates through them, solving each puzzle sequentially to progress to the next chamber. Once the team successfully solves the puzzle of the final chamber, the function should return the ultimate treasure hidden within.
Create the function exploreTemple(chambers) and provide an example of how it would be used with a sample array of chambers. */

package main

import "fmt"

type chamber struct {
	artifact string
	puzzle   string
	solved   bool
	puzzleAnswer string
}

func exploreTemple(chambers []chamber) string {
	for i := 0; i < len(chambers); i++ {
		fmt.Println("Artifact: ", chambers[i].artifact)
		fmt.Println("Puzzle: ", chambers[i].puzzle)
		var answer string
		fmt.Println("Enter the answer of the puzzle: ")
		fmt.Scanln(&answer)
		if chambers[i].puzzleAnswer == answer {
			chambers[i].solved = true
			fmt.Println("Correct answer")
			fmt.Println("Moving to next chamber")
			fmt.Println("=====================================")
		} else {
			fmt.Println("Wrong answer")
			i--
		}
	}

	return " Tutankhamun's Death Mask :  The ultimate treasure unlocked!"
}

func main() {
	chambers := []chamber{
		{artifact: "Phaistos Disc", puzzle: "Investigate the purpose behind the massive geoglyphs in the desert", solved: false , puzzleAnswer: "Ceremonial"},
		{artifact: " Nazca Lines", puzzle: "Decrypt the unknown language and understand the illustrations.", solved: false , puzzleAnswer: "Enigmatic"},
		{artifact: "Voynich Manuscript", puzzle: " Investigate the craftsmanship and symbolism of the golden mask.", solved: false , puzzleAnswer: "Tutankhamun"},
	}
	fmt.Println(exploreTemple(chambers))
}



