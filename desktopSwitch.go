package main

import(
"github.com/go-vgo/robotgo"
"fmt"
)

func main() {
fmt.Println("about to hit ^C")
robotgo.KeyTap("1", "control")
fmt.Println("ddnt Work")
}
