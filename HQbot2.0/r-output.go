package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"github.com/go-vgo/robotgo"
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