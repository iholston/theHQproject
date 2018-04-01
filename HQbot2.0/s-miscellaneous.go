package main

import (
	"github.com/go-vgo/robotgo"
	"os/exec"
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
	"strconv"
)

func helperFunc(function string) {
	switch function {
	case "quickTimePlayerSetup" : {
	/*	robotgo.KeyTap("1", "control")
		quickTime := exec.Command("bash", "-c", "open /Applications/QuickTime\\ Player.app/")
		quickTime.Run()
		Sleep(2)
		mods := []string{"alt", "command"}
		robotgo.KeyTap("n", mods)
		Sleep(5)
		robotgo.MoveMouse(660, 584)
		Sleep(1)
		robotgo.MouseClick()
		Sleep(2)
		robotgo.MoveMouse(669, 639)
		robotgo.MouseClick()*/
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

func testGameQ(index int) {
	fmt.Println("\n---------------------------------")
	fmt.Print("\nIs there any input for this question? (Y/n, Default = No) ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "Y\n" || input == "y\n" {
		fmt.Print("\nWas the question correct? (Y/n) ")
		input, _ = reader.ReadString('\n')
		textFile, _ := os.Open("testgamenotes.txt")
		if input == "Y\n" || input == "y\n" {
			textFile.WriteString("Question: " + strconv.Itoa(index) + " was correct")
			textFile.WriteString("------------------------")
		} else {
			textFile.WriteString("Question: " + strconv.Itoa(index) + " was NOT correct")
			textFile.WriteString("------------------------")
		}
		for {
			fmt.Println("Type your notes below:")
			fmt.Println("---------------------------------")
			notes, _ := reader.ReadString('\n')
			fmt.Println("Are you sure you want to commit? (Y/n)")
			input, _ = reader.ReadString('\n')
			if input == "Y\n" || input == "y\n" {
				textFile.WriteString(notes)
				textFile.WriteString("------------------------\n\n")
				break
			} else {
				continue
			}
		}
	}
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

//func makeURL(question []byte) string {
/*newQuestion := make([]byte, len(question) + 100)
questionUp := 0
for i := 0; i < len(question); i++ {
	if question[i] == 32 {
		newQuestion[questionUp] = '%'
		newQuestion[questionUp + 1] = '2'
		newQuestion[questionUp + 2] = '0'
		questionUp = questionUp + 3
		continue
	}
	newQuestion[questionUp] = question[i]
	questionUp++
}
finalQuestion := string(newQuestion)
finalQuestion = strings.Trim(finalQuestion, "\x00")
return finalQuestion
}*/