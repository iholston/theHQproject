package main

import (
	"os/exec"
	"time"
	"fmt"
	"os"
	"bufio"
)

var botStartTime = time.Now().Format("MonJan22006@15:04:05")
var folderName = "/Users/TheChosenOne/go/theHQProject/gameTrials/" + botStartTime +"/"
var testMode = false
var useDefault = false
var testGame = false
var terminalLog = "terminalLog"
var testFN = "/Users/TheChosenOne/Desktop/testgamenotes.txt"

func startUp() {
	//Set up UI to say "Please Make sure phone is connected and unlocked"//
	fmt.Print("Initializing....\n\nWhich type of start?\n" +
		"---------------------------------\nf: full start\ns: skip QTP\n" +
		"t: SQ test mode\nz: FG test mode\n---------------------------------\nEnter: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// If its a TEST
	if input == "t\n" {
		fmt.Println("\nEntering Test Mode...")
		fmt.Print("Would you like to use default picture? (Y/n) ")
		input, _ = reader.ReadString('\n')
		if input == "Y\n" || input == "y\n" {
			useDefault = true
		}
		fmt.Print("In this mode, log files appear on desktop and QTP setup is skipped.\n" +
			"-------------------------------------------------------------------\n")
		testMode = true
		return
	} else if input == "z\n" {
		fmt.Println("\nTest Game Activated. Each question will have notes activated. \n" +
			"Otherwise the game should run as normal. ~QTP Skipped")
		testGame = true
	} else if input == "s\n" {
		fmt.Println("Skipping QuickTimePlayer setup function only. Everything else is normal.")
	}
	fmt.Print("----------------------------------------------------------------\n")
	fmt.Print("Please make sure phone is connected and unlocked.\n")
	fmt.Println("\nInitiating bot starting procedures...")
	// Start QuickTimePlayer and get it to be using iphone as camera
	// put this in helperFunc cuz a lot of debugging requires me to comment all that out
	if input != "s\n" && input != "z\n" {
		helperFunc("quickTimePlayerSetup")
	}
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
	cmd4 := "mkdir " + folderName
	createNewFolder := exec.Command("bash", "-c", cmd4)
	createNewFolder.Run()
	fileName := folderName + "logs"
	fileName2 := folderName + "terminalCapture.txt" // Text file of all terminal output
	_ , err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_ , err = os.Create(fileName2)
	if err != nil {
		panic(err)
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


	// Logic
	startUp() // Func: startUP() - 1. Starts QTP, 2. Changes Default ScreenShot Names, and 3. Creates the session folder and log file
	for i := 1; i < 100; i++ {

		// Test Game
		if testGame == true {
			_ , err := os.Create(testFN)
			if err != nil {
				panic(err)
			}
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

			// Step 3.1: Google question and see how many times each answer is on the first page
		googleFirstPageIt(makeURL2(Question), AnswersString) // Returns [3][5]int and an int

			// Step 3.2: Google question + "wikipedia", grab the wikipedia article that pops up
			//			 see how many times each answer appears on that page
		wikiFirstPageIt(makeURL2(Question), AnswersString)

			// Step 3.3: Google question with answer and see how many results it returns
		googleSR_Alg(makeURL2(Question), Answers)


		// Step 4: Create "Better" question and run it through the previous algorithms
		//createBetterQuestion(Question)

			// Step 4.1: Same as 3.1 but with "BETTER QUESTION"
		//googleFirstPageIt(makeURL(Question), AnswersString) // Returns [3][5]int and an int

			// Step 4.2: Same as 3.2 but with "BETTER QUESTION"
		//wikiFirstPageIt(makeURL(Question), AnswersString)

			// Step 4.3: Same as 3.3 but with "BETTER QUESTION"
		//googleResultsIt()


		// Step 5: Take all data and determine the final answer
		//FinalGuess()


		// Step 6: Display Answer
		//displayAnswers()

		// Test Game Mode:
		if testGame == true {
			testGameQ(i)
		}
	}

	// Step 7: CleanUp after program
	//returnScreenShotsToNormal()
}