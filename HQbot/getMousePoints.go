package main

import(
"fmt"
"time"
"github.com/go-vgo/robotgo")

func main() {
time.Sleep(2000 * time.Millisecond)
fmt.Println(robotgo.GetMousePos())	
}

