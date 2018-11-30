package main

// Goes through each page of google and counts answer


/*
func googleEverything(url string, answers [4]string, out chan<- [4]int, wg *sync.WaitGroup, timeout <-chan bool) {
	defer wg.Done()
	var chans [10]chan [3]int // Ten channels for the ten go routines to talk
	var totalArray [3]int	  // Final array to return
	var resulties [10][3]int  // Array that holds the return values from all the go routines
	for i := range chans {
		chans[i] = make(chan [3]int, 2)
	}
	for i := 0; i < 10; i++ {
		resulties[i] = [3]int{0,0,0}
	}

	// Step 1: Get body text from url
	resp, _ := http.Get("https://www.google.com/search?q=" + url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	s := string(bytes)

	// Step 2: Make everything lowercase
	answers[0] = strings.ToLower(answers[0])
	answers[1] = strings.ToLower(answers[1])
	answers[2] = strings.ToLower(answers[2])

	urls := strings.Split(s, "<h3 class=\"r\"><a href=\"/url?q=")
	howMany := len(urls)

	if howMany < 11 { // if google returns less than ten urls then only do five, rather than trying to find out how many
		url1 := strings.Split(urls[1], "&amp;sa")
		url2 := strings.Split(urls[2], "&amp;sa")
		url3 := strings.Split(urls[3], "&amp;sa")
		url4 := strings.Split(urls[4], "&amp;sa")
		url5 := strings.Split(urls[5], "&amp;sa")

		go searchPage(url1[0], answers, chans[0], 1)
		go searchPage(url2[0], answers, chans[1], 2)
		go searchPage(url3[0], answers, chans[2], 3)
		go searchPage(url4[0], answers, chans[3], 4)
		go searchPage(url5[0], answers, chans[4], 5)

		timedout := false
		for i := 0; i < 10; i++ {
			if timedout == true {
				break
			}
			select {
			case resulties[0] = <-chans[0]:
			case resulties[1] = <-chans[1]:
			case resulties[2] = <-chans[2]:
			case resulties[3] = <-chans[3]:
			case resulties[4] = <-chans[4]:
			case resulties[5] = <-chans[5]:
			case timedout = <-timeout:
				break
			}
		}
		for i := 0; i < 5; i++ {
			totalArray[0] += resulties[i][0]
			totalArray[1] += resulties[i][1]
			totalArray[2] += resulties[i][2]
		}
		out <- totalArray
	} else {
		url1 := strings.Split(urls[1], "&amp;sa")
		url2 := strings.Split(urls[2], "&amp;sa")
		url3 := strings.Split(urls[3], "&amp;sa")
		url4 := strings.Split(urls[4], "&amp;sa")
		url5 := strings.Split(urls[5], "&amp;sa")
		url6 := strings.Split(urls[6], "&amp;sa")
		url7 := strings.Split(urls[7], "&amp;sa")
		url8 := strings.Split(urls[8], "&amp;sa")
		url9 := strings.Split(urls[9], "&amp;sa")
		url10 := strings.Split(urls[10], "&amp;sa")

		//go getValues(chans, timeout, resulties)
		go searchPage(url1[0], answers, chans[0], 1)
		go searchPage(url2[0], answers, chans[1], 2)
		go searchPage(url3[0], answers, chans[2], 3)
		go searchPage(url4[0], answers, chans[3], 4)
		go searchPage(url5[0], answers, chans[4], 5)
		go searchPage(url6[0], answers, chans[5], 6)
		go searchPage(url7[0], answers, chans[6], 7)
		go searchPage(url8[0], answers, chans[7], 8)
		go searchPage(url9[0], answers, chans[8], 9)
		go searchPage(url10[0], answers, chans[9], 10)

		timedout := false
		for i := 0; i < 10; i++ {
			if timedout == true {
				break
			}
			select {
			case resulties[0] = <-chans[0]:
			case resulties[1] = <-chans[1]:
			case resulties[2] = <-chans[2]:
			case resulties[3] = <-chans[3]:
			case resulties[4] = <-chans[4]:
			case resulties[5] = <-chans[5]:
			case resulties[6] = <-chans[6]:
			case resulties[7] = <-chans[7]:
			case resulties[8] = <-chans[8]:
			case resulties[9] = <-chans[9]:
			case timedout = <-timeout:
				break
			}
		}
		for i := 0; i < 10; i++ {
			totalArray[0] += resulties[i][0]
			totalArray[1] += resulties[i][1]
			totalArray[2] += resulties[i][2]
		}
		out <- totalArray
	}
}

func searchPage(url string, answers [4]string, out chan<- [4]int, searchN int) {
	// Step 1: Get body text from url
	resp, err := http.Get(url)
	if err != nil {
		totalA := [3]int{0,0,0}
		fmt.Println("ERROR1 go searchPage #" + strconv.Itoa(searchN))
		out <-totalA
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		totalA := [3]int{0,0,0}
		fmt.Println("ERROR2 go searchPage #" + strconv.Itoa(searchN))
		out <-totalA
		return
	}
	resp.Body.Close()
	s := string(bytes)

	// Step 2: Make everything lowercase
	s = strings.ToLower(s)
	answers[0] = strings.ToLower(answers[0])
	answers[1] = strings.ToLower(answers[1])
	answers[2] = strings.ToLower(answers[2])
	answers[3] = strings.ToLower(answers[3])


	// Step 3: Get matches. num1 = all direct matches, num2 = all matches from processed questions
	//						num3 = all matches from questions split up if they are more than one word
	var totalArray [4]int
	for i := 0; i < 4; i++ {
		// Step 3.1: Count number of times answer appears in the body text direcly
		num1 := strings.Count(s, answers[i])

		// Step 3.2: Remove special characters from answers and recount
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedAnswer := reg.ReplaceAllString(answers[i], " ")
		processedAnswer = strings.TrimSpace(processedAnswer)
		num2 := strings.Count(s, processedAnswer)

		// Step 3.3: Split answer up into individual parts and see how many times those appear
		//ansW := strings.Split(processedAnswer, " ")
		//num3 := strings.Count(s, ansW[0])

		// Step 3.4: Add up findings
		totalNum := num1 + num2 //+ num3
		totalArray[i] = totalNum
	}
	// Step 4: Send out results
	out <- totalArray
}
*/