package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.google.com/search?q=where+is+the+lincoln+memorial+located&oq=where+is+the+lincoln+memorial+located")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s := string(bytes)	
	fmt.Println(strings.Count("cheese", "E"))
	fmt.Println(strings.Count(s, "Washington, DC.")) // before & after each rune
	fmt.Println(strings.Count(s, "Washington, DC")) // before & after each rune
}

