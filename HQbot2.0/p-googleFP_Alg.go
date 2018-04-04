package main

import (
	"strings"
	"regexp"
	"log"
	"net/http"
	"io/ioutil"
	"sync"
)

func googleFirstPageIt(url string, answers [3]string, out chan<- [3]int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Step 1: Get body text from url
	resp, _ := http.Get("https://www.google.com/search?q=" + url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s := string(bytes)

	// Step 2: Make everything lowercase
	s = strings.ToLower(s)
	answers[0] = strings.ToLower(answers[0])
	answers[1] = strings.ToLower(answers[1])
	answers[2] = strings.ToLower(answers[2])


	// Step 3: Get matches. num1 = all direct matches, num2 = all matches from processed questions
	//						num3 = all matches from questions split up if they are more than one word
	var totalArray [3]int
	for i := 0; i < 3; i++ {
		// Step 3.1: Count number of times answer appears in the body text direcly
		num1 := strings.Count(s, answers[i])

		// Step 3.2: Remove special characters from answers and recount
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedAnswer := reg.ReplaceAllString(answers[i], " ")
		processedAnswer = strings.TrimSpace(processedAnswer)
		num2 := strings.Count(s, processedAnswer)

		// Step 3.3: Split answer up into individual parts and see how many times those appear
		ansW := strings.Split(processedAnswer, " ")
		j := 0
		var num3 int
		for range ansW {
			if len(ansW[j]) < 4 {
				j++
				continue
			}
			num3 += strings.Count(s, ansW[j])
			j++
		}

		// Step 3.4: Count all matches and put them into array and get a total value
		totalNum := num1 + num2 + num3
		totalArray[i] = totalNum

	}

	// Step 5: Output to output()
	out <- totalArray
}