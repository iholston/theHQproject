package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"
)

func googleSR_Alg(qPartofURL string, answers [3][]byte, out chan<- [3]int, wg *sync.WaitGroup ) {
	defer wg.Done()
	var results [3]int
	var url [3]string
	url[0] = qPartofURL + "%20" + makeURL2(answers[0])
	url[1] = qPartofURL + "%20" + makeURL2(answers[1])
	url[2] = qPartofURL + "%20" + makeURL2(answers[2])

	for i := 0; i < 3; i++ {
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