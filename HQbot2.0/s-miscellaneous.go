package main

import (
	"github.com/go-vgo/robotgo"
	"os/exec"
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
)

func helperFunc(function string) {
	switch function {
	case "quickTimePlayerSetup" : {
		robotgo.KeyTap("1", "control")
		quickTime := exec.Command("bash", "-c", "open /Applications/QuickTime\\ Player.app/")
		quickTime.Run()
		Sleep(2)
		robotgo.KeyTap("space", "command")
		SleepM(500)
		robotgo.TypeStr("quickTime")
		SleepM(500)
		robotgo.KeyTap("enter")
		SleepM(500)
		mods := []string{"alt", "command"}
		robotgo.KeyTap("n", mods)
		Sleep(3)
		robotgo.MoveMouseSmooth(658, 581)
		Sleep(1)
		robotgo.MouseClick()
		Sleep(1)
		robotgo.MoveMouseSmooth(669, 639)
		robotgo.MouseClick()
		for { // Test the screenshots
			robotgo.KeyTap("1", "control")
			SleepM(500)
			arr := []string{"shift", "command"}
			robotgo.KeyTap("3", arr) // takes screen shot for tesseract
			Sleep(1)
			cmd1 := "convert ~/Desktop/QandA.png -crop 700x700+915+300 ~/Desktop/Test.png"
			crop := exec.Command("bash", "-c", cmd1)
			crop.Run()
			fmt.Print("Would you like to try again? (Y/n): ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			cmd2 := "rm ~/Desktop/QandA.png ~/Desktop/Test.png"
			clearScreen := exec.Command("bash", "-c", cmd2)
			clearScreen.Run()
			if input == "y\n" {
				continue
			} else {
				fmt.Println("-------------------------------------")
				fmt.Println("______________________________________________________________________________")
				return
			}
		}
	}

	}
}

func makeURL2(question []byte) string {
	stringQuestion := string(question)
	arrayQuestion := strings.Fields(stringQuestion)
	url := strings.Join(arrayQuestion, "%20")
	return url
}

func Sleep(tim time.Duration) {
	time.Sleep(tim * time.Second)
}

func SleepM(tim time.Duration){
	time.Sleep(tim * time.Millisecond)
}

func returnScreenShotsToNormal(){
	cmd1 := "defaults write com.apple.screencapture name \"Screen Shot\""
	changeScreenShotNames := exec.Command("bash", "-c", cmd1)
	changeScreenShotNames.Run()
	cmd2 := "killall SystemUIServer"
	restart := exec.Command("bash", "-c", cmd2)
	restart.Run()
	cmd3 := "defaults write com.apple.screencapture \"include-date\" 1"
	addDateandTimeBack := exec.Command("bash", "-c", cmd3)
	addDateandTimeBack.Run()
	restart.Run()
}

