package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

// This function googles the question and counts instances of each choice on the page.
func googlefp(url string, choices [5]string, output chan<- [5]int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Set up

	// Step 1: Get body tet from url
	resp, err := http.Get(url)
	if err != nil {fmt.Println(err)}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {fmt.Println(string(bytes))}
	resp.Body.Close()
	s := strings.ToLower(string(bytes))

	// Step 2: Massage the Data
	choices[0] = strings.ToLower(choices[0])
	choices[1] = strings.ToLower(choices[1])
	choices[2] = strings.ToLower(choices[2])
	choices[3] = strings.ToLower(choices[3])
	fmt.Println(choices)

	// Step 3: Get matches.
	var totalArray [5]int
	for i := 0; i < 5; i++ {
		// Step 3.1: Count number of times answer appears in the body text directly
		num1 := strings.Count(s, choices[i])

		// Step 3.2: Remove special characters from answers and recount
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {log.Fatal(err)}
		processedAnswer := reg.ReplaceAllString(choices[i], " ")
		processedAnswer = strings.TrimSpace(processedAnswer)
		num2 := strings.Count(s, processedAnswer)

		// Step 3.3: Split answer up into individual parts and see how many times those appear
		ansW := strings.Split(processedAnswer, " ")
		num3 := strings.Count(s, ansW[0])

		totalNum := num1 + num2 + num3
		totalArray[i] = totalNum
	}

	// Step 4: Sort matches and send array to output chan
	output <- totalArray
}