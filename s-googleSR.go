package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// The least useful algorithm. This takes the question, and googles it with each answer then ranks the choices
// based on the number of search results returned.

func googlesr(url1 string, answers [4]string, out chan<- [4]int, wg *sync.WaitGroup) {
	defer wg.Done()
	var results [4]int
	var url [3]string
	url[0] = url1 + "%20" + makeURL(answers[0])
	url[1] = url1 + "%20" + makeURL(answers[1])
	url[2] = url1 + "%20" + makeURL(answers[2])

	for i := 0; i < 4; i++ {
		resp, _ := http.Get("https://www.google.com/search?q=" + url[i])
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		s := string(bytes)
		r, _ := regexp.Compile("id=\"resultStats.*results")
		s = r.FindString(s)
		var k string
		for i := 0; i < len(s); i++{
			if '0' <= s[i] && s[i] <= '9' {
				k += string(s[i])
			}
		}
		results[i], _ = strconv.Atoi(k)
	}
	out <- results
}

func makeURL(question string) string {
	arrayQuestion := strings.Fields(question)
	url := strings.Join(arrayQuestion, "%20")
	return url
}