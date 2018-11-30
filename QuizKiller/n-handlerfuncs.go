package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func landingpage(w http.ResponseWriter, r *http.Request) {
	// Serve landing page
	pagevars := htmlpagevartransport{Var1: 0, Var2: 0, Var3: 0, Var4: 0}
	t, _ := template.ParseFiles("html/landingpage.html")
	err := t.Execute(w, pagevars)
	if err != nil {fmt.Println("Error executing template. Landing page.")}
}

func hqthis(w http.ResponseWriter, r *http.Request) {
	// Step 1: Set up
	var question string
	var choices [5]string
	var results [5][2]int // the first int is the "placement" the second is the "percentage"

	// Step 2: Get form values from landing page
	err := r.ParseForm()
	if err != nil {fmt.Println("Error parsing form")}
	question = r.FormValue("question")
	choices[0] = r.FormValue("answer1")
	choices[1] = r.FormValue("answer2")
	choices[2] = r.FormValue("answer3")
	choices[3] = r.FormValue("answer4")
	choices[4] = r.FormValue("answer5")

	switch {
	case choices[1] == "":
		numchoices = 1
	case choices[2] == "":
		numchoices = 2
	case choices[3] == "":
		numchoices = 3
	case choices[4] == "":
		numchoices = 4
	}

	// Step 3: Send question and choices to hqbot and get results
	results = hqbot(question, choices)

	// Step 4: Build new page and send results back to client
	pagevars := htmlpagevartransport{results[0][0], results[1][0], results[2][0], results[3][0],
	question, choices[0], choices[1], choices[2], choices[3]}
	t, _ := template.ParseFiles("html/resultspage.html")
	err = t.Execute(w, pagevars)
	if err != nil {fmt.Println("Error executing template. Hqthis.")}

	}

