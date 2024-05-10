package main

import (
    "context"
    "testing"
    "time"
)

func TestTallyVotes(t *testing.T) {
    election := Election{
        Votes: make(map[string]int), // Error Handling (map initialization)
    }

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
        go election.TallyVotes(ctx, ps) // Concurrency Management with Context
    }

    time.Sleep(1 * time.Second)

    // Perform assertions to check the tallying results
    expected := map[string]int{"A": 3, "B": 2, "C": 1}
    for candidate, expectedVotes := range expected {
        if election.Votes[candidate] != expectedVotes {
            t.Errorf("Incorrect tallying result for candidate %s. Expected: %d, Got: %d", candidate, expectedVotes, election.Votes[candidate])
        }
    }
}
