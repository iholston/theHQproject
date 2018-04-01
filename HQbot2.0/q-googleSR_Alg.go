package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"strconv"
	"fmt"
)

func googleSR_Alg(qPartofURL string, answers [3][]byte) int {
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
	if results[0] > results[1] {
		if results[0] > results[2] {
			fmt.Println("\nGoogle results:\n---------------\nAnswer 1: Correct")
			return 1 //results[0]
		} else {
			fmt.Println("\nGoogle results:\n---------------\nAnswer 3: Correct")
			return 3 //results[2]
		}
	} else {
		if results[1] > results[2] {
			fmt.Println("\nGoogle results:\n---------------\nAnswer 2: Correct")
			return 2 //results[1]
		} else {
			fmt.Println("\nGoogle results:\n---------------\nAnswer 3: Correct")
			return 3 //results[2]
		}
	}
}