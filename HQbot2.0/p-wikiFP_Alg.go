package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"regexp"
	"log"
)

func wikiFirstPageIt(url string, answers [3]string) ([3][3]int, int) {
	fmt.Println("\nWikipedia says:\n------------------")
	var finalURL string

	// Step 1: Get url
	wikiURL := "https://www.google.com/search?q=" + url + "%20wikipedia"
	resp2, _ := http.Get(wikiURL)
	bytes2, _ := ioutil.ReadAll(resp2.Body)
	resp2.Body.Close()
	s2 := string(bytes2)

	//fmt.Println("\n\n\n\n\n")
	//fmt.Println(s)

	//Step 2: Get the wikipedia link that showed up first
	lookFor := "https://en.wikipedia.org/wiki/"
	processingArray := strings.Split(s2, lookFor)

	endOfUrl1 := strings.SplitN(processingArray[1], "&", 2)
	endOfUrl2 := strings.SplitN(processingArray[1], "%", 2)
	endOfUrl3 := strings.SplitN(processingArray[1], "+", 2)
	endOfUrl4 := strings.SplitN(processingArray[1], "\"", 2)
	endOfUrl5 := strings.SplitN(processingArray[2], "&", 2)
	endOfUrl6 := strings.SplitN(processingArray[2], "%", 2)
	endOfUrl7 := strings.SplitN(processingArray[2], "+", 2)
	endOfUrl8 := strings.SplitN(processingArray[2], "\"", 2)

	// Step 2.2: Get the actual ending from first parsed group
	smallestArray := [4]int{len(endOfUrl1[0]), len(endOfUrl2[0]), len(endOfUrl3[0]), len(endOfUrl4[0])}
	min := 500
	var smollestIndex int
	var urlEnd string
	for i := range smallestArray {
		if smallestArray[i] < min {
			min = smallestArray[i]
			smollestIndex = i
		}
	}
	if min == 500 {
		fmt.Println("Error: not getting any url endings")
		var chicken [3][3]int
		var nuggets int
		return chicken, nuggets
	} else {
		switch {
		case smollestIndex == 0:
			{
				urlEnd = endOfUrl1[0]
			}
		case smollestIndex == 1:
			{
				urlEnd = endOfUrl2[0]
			}
		case smollestIndex == 2:
			{
				urlEnd = endOfUrl3[0]
			}
		case smollestIndex == 3:
			{
				urlEnd = endOfUrl4[0]
			}
		}
	}

	// Step 2.3: Get the actual ending from second parsed group
	smallestArray2 := [4]int{len(endOfUrl5[0]), len(endOfUrl6[0]), len(endOfUrl7[0]), len(endOfUrl8[0])}
	min2 := 500
	var smollestIndex2 int
	var urlEnd2 string
	for i := range smallestArray2 {
		if smallestArray2[i] < min2 {
			min2 = smallestArray2[i]
			smollestIndex2 = i
		}
	}
	if min2 == 500 {
		fmt.Println("Error: not getting any url endings")
		var chicken [3][3]int
		var nuggets int
		return chicken, nuggets
	} else {
		switch {
		case smollestIndex2 == 0:
			{
				urlEnd2 = endOfUrl5[0]
			}
		case smollestIndex2 == 1:
			{
				urlEnd2 = endOfUrl6[0]
			}
		case smollestIndex2 == 2:
			{
				urlEnd2 = endOfUrl7[0]
			}
		case smollestIndex2 == 3:
			{
				urlEnd2 = endOfUrl8[0]
			}
		}
	}

	// Step 3: Check if they are equal and if not tell boss
	if urlEnd == urlEnd2 {
		//fmt.Println("Eurika!")
		finalURL = lookFor + urlEnd
	} else {
		fmt.Println("Suck my peen")
		finalURL = lookFor + "suck_my_peen"
	}
	fmt.Println(finalURL)
	// Step 4: Do exactly what gfirst_page_alg.go does but with new url
	// Step 1: Get body text from url
	s := web_Parser(finalURL)

	// Step 2: Make everything lowercase
	s = strings.ToLower(s)
	answers[0] = strings.ToLower(answers[0])
	answers[1] = strings.ToLower(answers[1])
	answers[2] = strings.ToLower(answers[2])


	// Step 3: Get matches. num1 = all direct matches, num2 = all matches from processed questions
	//						num3 = all matches from questions split up if they are more than one word
	var answerArray [3][3]int // holds three responses ([3]int) for all three questions (the first [3])
	var totalArray [3]int
	for i := 0; i < 3; i++ {
		num1 := strings.Count(s, answers[i])
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedAnswer := reg.ReplaceAllString(answers[i], " ")
		processedAnswer = strings.TrimSpace(processedAnswer)
		num2 := strings.Count(s, processedAnswer)
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
		totalNum := num1 + num2 + num3
		totalArray[i] = totalNum
		answerArray[i][0] = num1
		answerArray[i][1] = num2
		answerArray[i][2] = num3

		// Step 4: Output all of the answers
		fmt.Println("Answer " + strconv.Itoa(i + 1) + ": " + strconv.Itoa(totalNum))
		//output("termLog","Answer " + strconv.Itoa(i + 1) + ": " + strconv.Itoa(totalNum))
	}

	// Step 5: Determine without super analytics which one had most results
	if totalArray[0] > totalArray[1] {
		if totalArray[0] > totalArray[2] {
			return answerArray, 1
		} else {
			return answerArray, 3
		}
	} else if totalArray[1] > totalArray[2] {
		return answerArray, 2
	} else {
		return answerArray, 3
	}
}