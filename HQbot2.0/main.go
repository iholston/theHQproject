package main

import (
	"os/exec"
	"time"
	"fmt"
	"os"
	"bufio"
	"sync"
)

var botStartTime = time.Now().Format("MonJan22006@15:04:05")
var folderName = "/Users/TheChosenOne/go/src/github.com/iholston/theHQProject/gameTrials/" + botStartTime +"/"
var testMode = false
var useDefault = false
var testGame = false
var noFile = false
var terminalLog = "terminalLog"
var testFN = "/Users/TheChosenOne/Desktop/testgamenotes.txt"

func init() { // 1. Changes Default ScreenShot Names and 2. Creates the session folder and log file

	// Change Default ScreenShot Names. Makes it much easier to locate screenshots
	cmd1 := "defaults write com.apple.screencapture name \"QandA\""
	changeScreenShotNames := exec.Command("bash", "-c", cmd1)
	changeScreenShotNames.Run()
	cmd2 := "defaults write com.apple.screencapture \"include-date\" 0"
	removeTimeandDate := exec.Command("bash", "-c", cmd2)
	removeTimeandDate.Run()
	cmd3 := "killall SystemUIServer"
	restart := exec.Command("bash", "-c", cmd3)
	restart.Run()

	// Create New folder for Question.pngs and QuestionText.pngs and create log file
	fmt.Print("\nWould you like to make logs for this game? (Y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "Y\n" || input == "y\n" {
		cmd4 := "mkdir " + folderName
		createNewFolder := exec.Command("bash", "-c", cmd4)
		createNewFolder.Run()
		fileName := folderName + "logs"
		fileName2 := folderName + "terminalCapture.txt" // Text file of all terminal output
		_, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		_, err = os.Create(fileName2)
		if err != nil {
			panic(err)
		}
	} else {
		noFile = true
	}


	fmt.Println("Beginning main loop...")
	//Set up UI to say "Please let us know when the game is about to start//
}

// Program Logic
func main() {

	// Variables
	var Question []byte
	var Answers [3][]byte
	var err bool
	var wg sync.WaitGroup

	// Logic
	startUpDialog()
	gameMod()

	for i := 1; i < 100; i++ {
		// If Test Game
		if testGame == true {
			_ , err := os.Create(testFN)
			if err != nil {
				panic(err)
			}
		}

		// In loop variables
		googleEverythin := make(chan [3]int)
		var chans [4]chan [3]int
		for i := range chans {
			chans[i] = make(chan [3]int)
		}

		// Step 1: Get the Human to tell you when the question is on Screen
		if humanHandler(i) {  // returns true when user inputs "donzo" to end program
			break
		}

		// Step 2: Print out Question and Answers so users can see if q&a can be trusted
		Question, Answers, err = picToQuestion(i)
		if err {
			continue
		}
		fmt.Print("\nQuestion: " + string(Question) + "\nAnswers : ")
		fmt.Println(string(Answers[0]), string(Answers[1]), string(Answers[2]))
		AnswersString := [3]string{string(Answers[0]), string(Answers[1]), string(Answers[2])}


		// Step 3: Run original question and answers through different Algorithms
		wg.Add(5)
		go googleEverything(makeURL2(Question), AnswersString, googleEverythin, &wg)
		go googleFirstPageIt(makeURL2(Question), AnswersString, chans[0], &wg)
		go wikiFirstPageIt(makeURL2(Question), AnswersString, chans[1], &wg)
		go googleSR_Alg(makeURL2(Question), Answers, chans[2], &wg)
		go output(googleEverythin, chans[0], chans[1], chans[2], &wg)
		wg.Wait()


		// If Test Game Mode:
		if testGame == true {
			testGameQ(i)
		}
	}

	// Step 7: CleanUp after program
	returnScreenShotsToNormal()
}