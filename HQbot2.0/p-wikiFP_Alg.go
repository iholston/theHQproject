package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"regexp"
	"log"
	"sync"
)

func wikiFirstPageIt(url string, answers [3]string, out chan<- [3]int, wg *sync.WaitGroup) {
	defer wg.Done()
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
		return
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
		return
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
	fmt.Println("WikiLink: " + finalURL)
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
		// Update: Changed to only look for the first word multiple times cuz the wrong answers
		//			can be very wrong with small four letter words and throw off the algorithm
		ansW := strings.Split(processedAnswer, " ")
		num3 := strings.Count(s, ansW[0])
		totalNum := num1 + num2 + num3
		totalArray[i] = totalNum
	}

	// Step 5: Determine without super analytics which one had most results
	out <- totalArray
}

func web_Parser(url string) string {
	var bodyText []string

	// Step 1: Get URL
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s := string(bytes)

	// Step 2: Split based on <p> tags. splitPage[0] is the only element without a </p> in it
	splitPage := strings.Split(s, "<p>")

	// Step 3: parse each element of splitPage to the </p> using split N and put text into "bodyText"
	for i := range splitPage {
		if i == 0 {
			continue // the first element of split page has no <p>..bodyText..<p/>
		}
		subsetText := strings.SplitN(splitPage[i], "</p>", 2) // splits page into everything before </p> and everything after
		bodyText = append(bodyText, subsetText[0])	// grabs the text before </p> which would have been after <p>
	}

	// Step 4: Find text in wikitable
	wikitable := strings.Split(s, "wikitable")

	// Step 5: parse each element of wikitable to the </table> using split Na dn put text into bodyText
	for i := range wikitable {
		if i == 0 { // the first element of split page has no <p>..bodyText..<p/>
			continue
		}
		subsetText := strings.SplitN(wikitable[i], "</table>", 2)
		bodyText = append(bodyText, subsetText[0])
	}

	// Step 6: Return text
	return strings.Join(bodyText, " ")
}