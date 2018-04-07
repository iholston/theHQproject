package main

import (
	"github.com/go-vgo/robotgo"
	"os/exec"
	"io/ioutil"
	"fmt"
	"bytes"
	"os"
	"strconv"
)

func picToQuestion(index int) ([]byte, [3][]byte, bool) {
	// Step 1: Take ScreenShot, Crop it, Tesseract It, Move back to Desktop 2
	if useDefault == false {
		if testMode == true {
			Sleep(1)
		}
		arr := []string{"shift", "command"}
		robotgo.KeyTap("3", arr) // takes screen shot for tesseract
		SleepM(500)
		cmd1 := "convert ~/Desktop/QandA.png -crop 700x700+915+350 ~/Desktop/croppedPic.png"
		crop := exec.Command("bash", "-c", cmd1)
		crop.Run()
		SleepM(600)
		cmd2 := "tesseract ~/Desktop/croppedPic.png ~/Desktop/text"
		tesseract := exec.Command("bash", "-c", cmd2)
		tesseract.Run()
		SleepM(750)
		robotgo.KeyTap("2", "control")
	}

	// Step 2: Extract the Questions and Answers from the text file
	// byte No. 32 = Space, byte No. 10 = \n, byte No. 63 = ?
	var QandAs []byte
	var err error
	if useDefault == true {
		QandAs, err = ioutil.ReadFile("/Users/TheChosenOne/go/src/github.com/iholston/theHQProject/HQbot2.0/x_text.txt")
	} else {
		QandAs, err = ioutil.ReadFile("/Users/TheChosenOne/Desktop/text.txt")
	}
	if err != nil {
		fmt.Print(err)
	}

	if len(QandAs) < 10 {
		fmt.Println("ERROR: Input too small")
		var x []byte
		var y [3][]byte
		cmdX := "rm ~/Desktop/QandA.png ~/Desktop/croppedPic.png ~/Desktop/text.txt"
		movePNG := exec.Command("bash", "-c", cmdX)
		movePNG.Run()
		return x, y, true
	}

	// Step 2.2: Get Question into a byte slice
	var QandAindex int
	question := make([]byte, len(QandAs))
	for i := 0; i < len(QandAs); i++ {
		// This if statement changes the line breaks to spaces
		if QandAs[i] == 10 {
			question[i] = 32
			continue
		}
		question[i] = QandAs[i]
		// Once a Question Mark is reached end the loop
		if QandAs[i] == 63 {
			QandAindex = i + 1
			break
		}
	}

	// Steps 2.3-2.6 Get answers into their own byte slice
	answer1 := make([]byte, len(QandAs))
	for i := 0; i < len(QandAs); i++ {
		// This keeps new lines from being put in front of bytes slice and ends loop if at end of byte slice
		if QandAindex == len(QandAs) {
			var answers [3][]byte
			fmt.Println("ERROR: Next Question.")
			cmd3 := "rm ~/Desktop/QandA.png ~/Desktop/croppedPic.png ~/Desktop/text.txt"
			movePNG := exec.Command("bash", "-c", cmd3)
			movePNG.Run()
			return question, answers, true
		}
		if QandAs[QandAindex] == 10 {
			if answer1[0] == 0 {
				QandAindex++
				i--
				continue
			} else {
				break
			}
		}
		answer1[i] = QandAs[QandAindex]
		QandAindex++
	}
	answer2 := make([]byte, len(QandAs))
	for i := 0; i < len(QandAs); i++ {
		// This keeps new lines from being put in front of bytes slice and ends loop if at end of byte slice
		if QandAs[QandAindex] == 10 {
			if answer2[0] == 0 {
				QandAindex++
				i--
				continue
			} else {
				break
			}
		}
		answer2[i] = QandAs[QandAindex]
		QandAindex++
	}
	answer3 := make([]byte, len(QandAs))
	for i := 0; i < len(QandAs); i++ {
		// This keeps new lines from being put in front of bytes slice and ends loop if at end of byte slice
		if QandAs[QandAindex] == 10 {
			if answer3[1] == 0 {
				QandAindex++
				i--
				continue
			} else {
				break
			}
		}
		answer3[i] = QandAs[QandAindex]
		QandAindex++
	}

	// Step 2.7: Very Important! Cut out non-string contributing bytes
	question = bytes.Trim(question, "\x00")
	answer1 = bytes.Trim(answer1, "\x00")
	answer2 = bytes.Trim(answer2, "\x00")
	answer3 = bytes.Trim(answer3, "\x00")
	answers := [3][]byte{answer1, answer2, answer3}

	// Step 2.8: Log questions and answers raw from tesseract and parsed
	//			 in case something goes wrong
	var fileName string
	if noFile == false {
		if testMode == true {
			fileName = "/Users/TheChosenOne/Desktop/testlogs"
		} else {
			fileName = folderName + "logs"
		}
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		text := "Question No. " + strconv.Itoa(index) + " Tesseract Output: \nBytes: \n"
		f.WriteString(text)
		if _, err = f.Write(QandAs); err != nil {
			panic(err)
		}
		text = "\nString: \n"
		f.WriteString(text)
		if _, err = f.WriteString(string(QandAs)); err != nil {
			panic(err)
		}
		text = "\nParsed Question. Byte Length: " + strconv.Itoa(len(question)) + "\n"
		f.WriteString(text)
		f.WriteString(string(question))
		text = "\nParsed Answer 1. Byte Length: " + strconv.Itoa(len(answer1)) + "\n"
		f.WriteString(text)
		f.WriteString(string(answer1))
		text = "\nParsed Answer 2. Byte Length: " + strconv.Itoa(len(answer2)) + "\n"
		f.WriteString(text)
		f.WriteString(string(answer2))
		text = "\nParsed Answer 3. Byte Length: " + strconv.Itoa(len(answer3)) + "\n"
		f.WriteString(text)
		f.WriteString(string(answer3))
		text = "\nThat's all for the first question folks! \n\n"
		f.WriteString(text)
	}


	// Step 3: Cleanup. Move ScreenShots and text stuffs to log folder & return
	if testMode == true || noFile == true {
		cmd3 := "rm ~/Desktop/QandA.png ~/Desktop/croppedPic.png ~/Desktop/text.txt"
		movePNG := exec.Command("bash", "-c", cmd3)
		movePNG.Run()
	} else {
		cmd3 := "mv ~/Desktop/QandA.png /Users/TheChosenOne/go/src/github.com/iholston/theHQProject/gameTrials/"
		cmd3 += botStartTime + "/QandA" + strconv.Itoa(index) + ".png"
		movePNG := exec.Command("bash", "-c", cmd3)
		movePNG.Run()
		cmd4 := "mv ~/Desktop/text.txt /Users/TheChosenOne/go/src/github.com/iholston/theHQProject/gameTrials/"
		cmd4 += botStartTime + "/QandAtext" + strconv.Itoa(index) + ".txt"
		moveText := exec.Command("bash", "-c", cmd4)
		moveText.Run()
		cmd5 := "mv ~/Desktop/croppedPic.png /Users/TheChosenOne/go/src/github.com/iholston/theHQProject/gameTrials/"
		cmd5 += botStartTime + "/croppedPic" + strconv.Itoa(index) + ".png"
		moveCPNG := exec.Command("bash", "-c", cmd5)
		moveCPNG.Run()
	}
	return question, answers, false
}