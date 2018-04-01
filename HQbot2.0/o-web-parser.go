package main

import (
	"net/http"
	"io/ioutil"
	"strings"
)

func web_Parser(url string) string {
	var bodyText []string

	// Step 1: Get URL
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s := string(bytes)

	// Step 2: Split based on <p> tags. splitPage[0] is the only element without a </p> in it
	splitPage := strings.Split(s, "<p>")

	// Step 3: parse each element of splitPage to the </p> using split N and but text into "bodyText"
	for i := range splitPage {
		if i == 0 {
			continue // the first element of split page has no <p>..bodyText..<p/>
		}
		subsetText := strings.SplitN(splitPage[i], "</p>", 2) // splits page into everything before </p> and everything after
		bodyText = append(bodyText, subsetText[0])	// grabs the text before </p> which would have been after <p>
	}
	return strings.Join(bodyText, " ")
}