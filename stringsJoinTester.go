package main

import(
"fmt"
"strings"
"io/ioutil"
)

func main() {
	QandAs, err := ioutil.ReadFile("/Users/TheChosenOne/Desktop/text.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(makeURL2(QandAs))
}

func makeURL2(question []byte) string {
	stringQuestion := string(question)
	arrayQuestion := strings.Fields(stringQuestion)
	url := strings.Join(arrayQuestion, "%20")
	return url
}
