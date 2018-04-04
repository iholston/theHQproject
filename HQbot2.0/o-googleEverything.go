package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"regexp"
	"log"
	"sync"
)

func googleEverything(url string, answers [3]string, out chan<- [3]int, wg1 *sync.WaitGroup) {
	defer wg1.Done()
	var chans [10]chan [3]int
	var totalArray [3]int
	var wg sync.WaitGroup

	for i := range chans {
		chans[i] = make(chan [3]int)
	}

	// Step 1: Get body text from url
	resp, _ := http.Get("https://www.google.com/search?q=" + url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s := string(bytes)

	// Step 2: Make everything lowercase
	answers[0] = strings.ToLower(answers[0])
	answers[1] = strings.ToLower(answers[1])
	answers[2] = strings.ToLower(answers[2])

	urls := strings.Split(s, "<h3 class=\"r\"><a href=\"/url?q=")
	howMany := len(urls)

	if howMany < 11 {
		url1 := strings.Split(urls[1], "&amp;sa")
		url2 := strings.Split(urls[2], "&amp;sa")
		url3 := strings.Split(urls[3], "&amp;sa")
		url4 := strings.Split(urls[4], "&amp;sa")
		url5 := strings.Split(urls[5], "&amp;sa")

		wg.Add(5)
		go searchPage(url1[0], answers, chans[0], &wg)
		go searchPage(url2[0], answers, chans[1], &wg)
		go searchPage(url3[0], answers, chans[2], &wg)
		go searchPage(url4[0], answers, chans[3], &wg)
		go searchPage(url5[0], answers, chans[4], &wg)

		results1 := <-chans[0]
		results2 := <-chans[1]
		results3 := <-chans[2]
		results4 := <-chans[3]
		results5 := <-chans[4]
		wg.Wait()

		totalArray[0] = results1[0] + results2[0] + results3[0] + results4[0] + results5[0]
		totalArray[1] = results1[1] + results2[1] + results3[1] + results4[1] + results5[1]
		totalArray[2] = results1[2] + results2[2] + results3[2] + results4[2] + results5[2]
		out <- totalArray
	} else {
		url1 := strings.Split(urls[1], "&amp;sa")
		url2 := strings.Split(urls[2], "&amp;sa")
		url3 := strings.Split(urls[3], "&amp;sa")
		url4 := strings.Split(urls[4], "&amp;sa")
		url5 := strings.Split(urls[5], "&amp;sa")
		url6 := strings.Split(urls[6], "&amp;sa")
		url7 := strings.Split(urls[7], "&amp;sa")
		url8 := strings.Split(urls[8], "&amp;sa")
		url9 := strings.Split(urls[9], "&amp;sa")
		url10 := strings.Split(urls[10], "&amp;sa")

		wg.Add(10)
		go searchPage(url1[0], answers, chans[0], &wg)
		go searchPage(url2[0], answers, chans[1], &wg)
		go searchPage(url3[0], answers, chans[2], &wg)
		go searchPage(url4[0], answers, chans[3], &wg)
		go searchPage(url5[0], answers, chans[4], &wg)
		go searchPage(url6[0], answers, chans[5], &wg)
		go searchPage(url7[0], answers, chans[6], &wg)
		go searchPage(url8[0], answers, chans[7], &wg)
		go searchPage(url9[0], answers, chans[8], &wg)
		go searchPage(url10[0], answers, chans[9], &wg)

		results1 := <-chans[0]
		results2 := <-chans[1]
		results3 := <-chans[2]
		results4 := <-chans[3]
		results5 := <-chans[4]
		results6 := <-chans[5]
		results7 := <-chans[6]
		results8 := <-chans[7]
		results9 := <-chans[8]
		results10 := <-chans[9]
		wg.Wait()

		totalArray[0] = results1[0] + results2[0] + results3[0] + results4[0] + results5[0] +
			results6[0] + results7[0] + results8[0] + results9[0] + results10[0]
		totalArray[1] = results1[1] + results2[1] + results3[1] + results4[1] + results5[1] +
			results6[0] + results7[1] + results8[1] + results9[1] + results10[1]
		totalArray[2] = results1[2] + results2[2] + results3[2] + results4[2] + results5[2] +
			results6[0] + results7[2] + results8[2] + results9[2] + results10[2]
		out <- totalArray
	}
}

func searchPage(url string, answers [3]string, out chan<- [3]int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Step 1: Get body text from url
	resp, _ := http.Get(url)
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

		// Step 3.4: Add up findings
		totalNum := num1 + num2 + num3
		totalArray[i] = totalNum
	}
	// Step 4: Send out results
	out <- totalArray
}

