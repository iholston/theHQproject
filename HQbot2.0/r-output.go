package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"github.com/go-vgo/robotgo"
	"sync"
)

func startUpDialog() {

	//Set up UI to say "Please Make sure phone is connected and unlocked"//
	fmt.Print("\nInitialize Module\n---------------------------------\n" +
		"Which type of start?\nf: full start\ns: skip QTP\n" +
		"t: SQ test mode\nz: FG test mode\nEnter: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// If its a TEST
	if input == "t\n" {
		fmt.Println(">> Entering Test Mode")
		fmt.Print("---------------------------------\n")
		fmt.Println("\nTest Mode Module")
		fmt.Print("---------------------------------\n")
		fmt.Print("Would you like to use default picture? (Y/n) ")
		input, _ = reader.ReadString('\n')
		if input == "Y\n" || input == "y\n" {
			useDefault = true
			testMode = true
			fmt.Print("In this mode, log files appear on desktop and QTP setup is skipped.\n" +
				"-------------------------------------------------------------------\n")
			return
		}
		fmt.Print("In this mode, log files appear on desktop and QTP setup is skipped.\n" +
			"-------------------------------------------------------------------\n")
		testMode = true
		return
	} else if input == "z\n" {
		fmt.Print(">> Test Game Activated\n---------------------------------\n")
		fmt.Print("\nTest Game Module\n-------------------------------------------------------------")
		fmt.Println("\nTest Game Activated. Each question will have notes activated. \n" +
			"Otherwise the game should run as normal. ~QTP Skipped")
		fmt.Print("-----------------------------------------------------\n")
		testGame = true
	} else if input == "s\n" {
		fmt.Println(">> Skipping QTP function only.")
		fmt.Println("---------------------------------")
		return
	}
	if input == "\n" || input == "f\n"{
		fmt.Println(">> Full Game Mode")
		fmt.Print("---------------------------------\n")
	}
	fmt.Println("\nTest Screen Module")
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Please make sure phone is connected and unlocked.\n")
	fmt.Println("Taking test screen captures...")
	helperFunc("quickTimePlayerSetup")
}

func gameMod() {
	fmt.Print("\n" +
	"     __            /^\\" + "\n" +
	"   .'  \\          / :.\\" + "\n" +
	"  /     \\         | :: \\" +"\n" +
	" /   /.  \\       / ::: |" +"\n" +
	"|    |::. \\     / :::'/" +"\n" +
	"|   / \\::. |   / :::'/" +"\n" +
	" `--`   \\'  `~~~ ':'/`" +"\n" +
	"         /         (" +"\n" +
	"        /   0 _ 0   \\" +"\n" +
	"      \\/     \\_/     \\/" +"\n" +
	"    -== '.'   |   '.' ==-" +"\n" +
	"      /\\    '-^-'    /\\" +"\n" +
	"        \\   _   _   /" +"\n" +
	"       .-`-((\\o/))-`-." +"\n" +
	"  _   /     //^\\\\     \\   _" +"\n" +
	".\"o\".(    , .:::. ,    ).\"o\"." +"\n" +
	"|o  o\\     \\:::::/    //o  o| " +"\n" +
	" \\    \\\\   |:::::|   //    /   " +"\n" +
	"  \\    \\\\__/:::::\\__//    /   " +"\n" +
	"   \\ .:.\\  `':::'`  /.:. /" +"\n" +
	"    \\':: |_       _| ::'/" +"\n" +
	"     `---` `\"\"\"\"\"` `---`")
	fmt.Println("\n         Game Module")
	fmt.Print("---------------------------------\n")
}

func humanHandler(index int) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nReady for question #" + strconv.Itoa(index) + ": ")
	input, _ :=reader.ReadString('\n')
	fmt.Print("---------------------------------\n")
	exitWord := "donzo\n"
	if input == exitWord {
		return true
	}
	fmt.Println("Capturing question from screen...")
	if testMode == false {
		robotgo.KeyTap("1", "control") // switches to Desktop 1
	}
	Sleep(1)
	return false
}

func testGameQ(index int) {
	fmt.Println("\n---------------------------------")
	fmt.Print("\nIs there any input for this question? (Y/n, Default = No) ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "Y\n" || input == "y\n" {
		fmt.Print("\nWas the question correct? (Y/n) ")
		input, _ = reader.ReadString('\n')
		textFile, _ := os.Open("/Users/TheChosenOne/Desktop/testgamenotes.txt")
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

func output(googErethen <-chan [3]int, googFP <-chan [3]int, wikiFP <-chan [3]int, googSR <-chan [3]int, wg *sync.WaitGroup) {
	defer wg.Done()
	googEreAnswers := <-googErethen
	googFPAnswers := <-googFP
	wikiFPAnswers := <-wikiFP
	googSRAnswers := <-googSR

	//googleEverything() output
	fmt.Println("\nGoogle Deep Search")
	fmt.Println("-------------------")
	fmt.Println("Answer 1: " + strconv.Itoa(googEreAnswers[0]))
	fmt.Println("Answer 2: " + strconv.Itoa(googEreAnswers[1]))
	fmt.Println("Answer 3: " + strconv.Itoa(googEreAnswers[2]))

	//googleFP() output
	fmt.Println("\nGoogle F_P Search")
	fmt.Println("-------------------")
	fmt.Println("Answer 1: " + strconv.Itoa(googFPAnswers[0]))
	fmt.Println("Answer 2: " + strconv.Itoa(googFPAnswers[1]))
	fmt.Println("Answer 3: " + strconv.Itoa(googFPAnswers[2]))

	//wikiFP() output
	fmt.Println("\nWiki First_P Search")
	fmt.Println("-------------------")
	fmt.Println("Answer 1: " + strconv.Itoa(wikiFPAnswers[0]))
	fmt.Println("Answer 2: " + strconv.Itoa(wikiFPAnswers[1]))
	fmt.Println("Answer 3: " + strconv.Itoa(wikiFPAnswers[2]))

	//googSR() output
	if googSRAnswers[0] > googSRAnswers[1] {
		if googSRAnswers[0] > googSRAnswers[2] {
			fmt.Println("\nGoogle About Search\n-------------------\nAnswer 1: Correct")
		} else {
			fmt.Println("\nGoogle About Search\n-------------------\nAnswer 3: Correct")
		}
	} else {
		if googSRAnswers[1] > googSRAnswers[2] {
			fmt.Println("\nGoogle About Search\n-------------------\nAnswer 2: Correct")
		} else {
			fmt.Println("\nGoogle About Search\n-------------------\nAnswer 3: Correct")
		}
	}

}
