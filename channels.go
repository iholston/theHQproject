package main

import (
"fmt"
//"time"
)
func message(out chan<-string, fi string) { 
twoplustwo := fi + fi
var chickn [2]string
chickn[0] = twoplustwo
chickn[1] = "fructius"
out <- chickn
out <- "fructius"
}

func receiver(in <-chan string) {
	msg := <-in
	fmt.Println(msg + " receiver")
}

func message2(out chan<-string, fi string) {
twoplus4 := fi + fi + fi
out <- twoplus4
}

func main() {
	messages := make(chan string)
	chickn := "checkn is heckn good"
//	nuggies := "nuggiesDelicious "
	go message(messages, chickn)
	//go message2(messages, nuggies)
	receiver(messages)
	//time.Sleep(time.Second)	
}
