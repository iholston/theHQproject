package main

import (
	"fmt"
	"net/http"
)

var numchoices int

// You have to make a struct to hand to the template, you can't just hand a bunch of vars
type htmlpagevartransport struct {
	Var1 int
	Var2 int
	Var3 int
	Var4 int
	Place1 string
	Place2 string
	Place3 string
	Place4 string
	Place5 string

}

func main() {
	fmt.Println("Initiating web server\nListening (8080)...")
	http.HandleFunc("/", landingpage)
	http.HandleFunc("/hqthis", hqthis)
	err := http.ListenAndServe(":80", nil)
	if err != nil {panic(err)}
}
