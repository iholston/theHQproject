package main

import (
	"strings"
	"sync"
	"time"
)

func hqbot(question string, choices [5]string) [5][2]int {
	// Step 1: Set up
	var results [5][2]int // The format here is the first index is the placement of the
						  // choice and then the confidence percentage
	var wg sync.WaitGroup // This forces hqbot to wait for all routines to finish before moving on
	var resultchans [4]chan [5]int // Channel for each algorithm to put its responses into
	for i := range resultchans { // This has to be this way, have to initialize int arrays?
		resultchans[i] = make(chan [5]int)
	}

	// Step 2: Create url from question. Very simple.
	stringQuestion := string(question)
	arrayQuestion := strings.Fields(stringQuestion)
	url := strings.Join(arrayQuestion, "%20")
	url = "https://www.google.com/search?q=" + url

	// Check to make sure this is not stopping after 2 seconds
	timeout := make(chan bool, 10)
	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 10; i++ {
			timeout <- true
		}
	}()

	// Step 3: Hand url, choices, response channel, and wg's to the different algorithms, wait, and return
	wg.Add(2) // The wg.Done() is deferred in each func
	//go wikifp(url, choices, resultchans[0], &wg)
	go googlefp(url, choices, resultchans[1], &wg)
	//go googleEverthing(url, choices, resultchans[2], &wg, timeout)
	//go googlesr(url, choices, resultchans[3], &wg)
	go calculationtime(&results, resultchans, &wg) // This function does all the beautiful calculations
	wg.Wait()
	return results
}


func calculationtime (results *[5][2]int, channelArray [4]chan[5]int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Step 1: Listen to all channels and store results of algorithms
	googleFPResults := <- channelArray[1]
	//googleSRResults := <- channelArray[3]
	//resultsArr := [2][4]int{googleFPResults, googleSRResults}

	// Step 2: Gets messy. Sort each result
	// Step 2.1: This creates an array whose index represents the choice number and whose value
	// represents the correctness of the choice. In this case, 1 is most correct while 2 is second correct, etc.
	// And then stores it in bigOrder
	// NEEDS TO BE UPDATED IN THE EVEN THAT MULTIPLE CHOICES COME BACK WITH THE SAME # OF APPEARANCES ON THE PAGE
	/*var bigOrder [2][4]int
	for z, results := range resultsArr {
		maximum := 0
		maxindex := 0
		orderArray := [4]int{0, 0, 0, 0}
		for i := 1; i < 5; i++ {
			for j, result := range results {
				if result > maximum {
					maximum = result
					maxindex = j
				}
			}
			orderArray[maxindex] = i
			googleFPResults[maxindex] = -1 // Removes that max value so the next highest value becomes the next one
			maximum = 0
		}
		bigOrder[z] = orderArray
	}*/

	for i:= 0; i< 5; i++ {
		results[i][0] = googleFPResults[i]
	}

	// Step 2.2: Take bigOrder and massage the data and get percents based on the importance of each alg.

}