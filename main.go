package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Node struct {
	Start bool
	Edges map[string]int
}

type Solution struct {
	Data      map[string]Node
	StartTime time.Time
}

// Recursively sets up concurrent goroutines for each edge traveled.
// First waits the specified amount of time before printing out the current node,
// and then sets up a new WaitGroup for all available edges to traverse.
func (s *Solution) Run(curr string, waitTime int, wg *sync.WaitGroup) {
	// Lets the "parent" node know that all child goroutines have finished
	defer wg.Done()

	time.Sleep(time.Second * time.Duration(waitTime))

	fmt.Println(time.Since(s.StartTime).Seconds(), curr)

	// Sets up a WaitGroup that blocks this function from exiting before its children
	// return, so that the main thread does not prematurely exit.
	var newWg sync.WaitGroup
	for edge, wait := range s.Data[curr].Edges {
		newWg.Add(1)
		go s.Run(edge, wait, &newWg)
	}

	newWg.Wait()
}

func main() {
	// Keep track of the starting time for timestamp calculations
	start := time.Now()
	fmt.Println("Start:", start)

	// Read input from input.json and store it as a map[string]Node
	fileBytes, _ := os.ReadFile("input.json")
	var data map[string]Node
	err := json.Unmarshal(fileBytes, &data)
	if err != nil {
		panic(err)
	}

	// Determine starting node
	var startNode string
	for label, node := range data {
		if node.Start {
			startNode = label
			break
		}
	}

	// Setup solution struct and run the program with a preliminary WaitGroup
	// and a wait time of 0 seconds (first node prints immediately)
	sol := Solution{
		Data:      data,
		StartTime: start,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	sol.Run(startNode, 0, &wg)
	wg.Wait()

	// Display total time elapsed for the program
	fmt.Println("End:", time.Since(start))
}
