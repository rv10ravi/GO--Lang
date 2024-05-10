package main

import (
	"context"		//For managing cancellation and timeouts of go routines
	"encoding/json" // For marshalling and unmarshalling data to and from JSON format
	"fmt"			// For printing messages to the console
	"sync"			// For providing concurrency primitives
	"time"			// For handling time-related operation

)

// Vote struct to store voter information
type Vote struct {
	VoterID   int
	Timestamp time.Time
	Candidate string
}

// PollingStation struct to store vote data
type PollingStation struct {
	ID       int
	VoteData []Vote
}

// Election struct to store election data
type Election struct {
	Votes map[string]int 
	mu    sync.Mutex     
}

// TallyVotes method to tally votes for each candidate
func (e *Election) TallyVotes(ctx context.Context, p PollingStation) {
	e.mu.Lock() 
	defer e.mu.Unlock() 
	for _, v := range p.VoteData {
		select {
		case <-ctx.Done():
			return 
		default:
			e.Votes[v.Candidate]++
		}
	}
}

//Menu to display options
func mainMenu() {
	fmt.Println("-------------------------")
	fmt.Println("Election Management System")
	fmt.Println("1. Start Election")
	fmt.Println("2. View Election Results")
	fmt.Println("3. Marshal Election Data")
	fmt.Println("4. Unmarshal Election Data")
	fmt.Println("5. Run Unit Tests")
	fmt.Println("6. Exit")
	fmt.Println("-------------------------")
}

//Main function to run the program
func main() {
	election := Election{
		Votes: make(map[string]int), 
	}

	var choice int
	for {
		mainMenu()
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			startElection(&election)
		case 2:
			viewResults(&election)
		case 3:
			marshalData(&election)
		case 4:
			unmarshalData(&election)
		case 5:
			runUnitTests(&election)
		case 6:
			fmt.Println("Exiting...")
			fmt.Println("-------------------------")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

//Function to start the election
func startElection(election *Election) {
	pollingStations := []PollingStation{
		{
			ID: 1,
			VoteData: []Vote{
				{VoterID: 1, Timestamp: time.Now(), Candidate: "A"},
				{VoterID: 2, Timestamp: time.Now(), Candidate: "B"},
				{VoterID: 3, Timestamp: time.Now(), Candidate: "A"},
			},
		},
		{
			ID: 2,
			VoteData: []Vote{
				{VoterID: 4, Timestamp: time.Now(), Candidate: "B"},
				{VoterID: 5, Timestamp: time.Now(), Candidate: "C"},
				{VoterID: 6, Timestamp: time.Now(), Candidate: "A"},
			},
		},
	}

	// Create a context with cancellation capability
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, ps := range pollingStations {
		go election.TallyVotes(ctx, ps) //Concurrency Management with Context
	}

	fmt.Println("-------------------------")
	fmt.Println("Election started. Processing votes...")
	time.Sleep(1 * time.Second)
	fmt.Println("Election processing completed.")

}


//Function to view the election results
func viewResults(election *Election) {
	election.mu.Lock() 
	defer election.mu.Unlock() 

	fmt.Println("-------------------------")
	fmt.Println("Election Results:")
	for candidate, votes := range election.Votes {
		fmt.Printf("%s: %d\n", candidate, votes)
	}
	fmt.Println("-------------------------")
}

//Function to marshal the election data
func marshalData(election *Election) {
	jsonData, err := json.Marshal(election) // Marshalling of Data
	if err != nil { // Error Handling
		fmt.Println("Error:", err) 
		return
	}
	fmt.Println("-------------------------")
	fmt.Println("Marshalled Election Data:", string(jsonData))
}

//Function to unmarshal the election data
func unmarshalData(election *Election) {
	// Simulating unmarshalling from JSON data 
	jsonData := []byte(`{"Votes":{"A":3,"B":2,"C":1}}`)

	var newElection Election
	err := json.Unmarshal(jsonData, &newElection) // Unmarshalling of Data
	if err != nil { // Error Handling
		fmt.Println("Error:", err) 
		return
	}

	election.mu.Lock() 
	defer election.mu.Unlock() 
	election.Votes = newElection.Votes
	fmt.Println("-------------------------")
	fmt.Println("Unmarshalled Election Data:", election.Votes)
}


//Function to run unit tests
func runUnitTests(election *Election) {
	fmt.Println("Running unit tests...")
	// Simulate running unit tests
	if err := testTallyVotes(election); err != nil {
		fmt.Println("Unit test failed:", err)
		return
	}
	fmt.Println("-------------------------")
	fmt.Println("Unit tests passed.")
}

//Function to test the tallying of votes
func testTallyVotes(election *Election) error {
	// Unit test for tallying votes
	expected := map[string]int{"A": 3, "B": 2, "C": 1}
	if !compareMaps(election.Votes, expected) {
		return fmt.Errorf("Incorrect tallying results. Expected: %v, Got: %v", expected, election.Votes)
	}
	return nil
}

//Function to compare two maps
func compareMaps(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	return true
}


