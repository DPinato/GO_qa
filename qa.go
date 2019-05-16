package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Question struct {
	QText       string   `json:"qText"`
	PossAnswers []string `json:"possAnswers"`
	Excerpt     string   `json:"excerpt"`
	Answers     []string `json:"answers"`
	Explanation string   `json:"explanation"`
}

func main() {
	totalQuestions := 0
	correctQuestions := 0
	wrongQuestions := 0

	log.Println("Starting ...")
	//	jsonQAFile := "C:\\Users\\Davide.Pinato\\Desktop\\Study things\\GO\\CCNP_300-101_book_proc.json"

	if len(os.Args) < 2 {
		log.Fatal("usage: go run qa.go <qa_file>")
	}

	jsonQAFile := os.Args[1]

	questions := readQuestionsFromFile(jsonQAFile)
	fmt.Printf("\nI have read %d questions", len(questions))
	fmt.Printf("\n")

	randomSeed := getRandomSeed()
	rand.Seed(int64(randomSeed))
	fmt.Printf("Seed %d\n", randomSeed)
	fmt.Printf("\n\n")

	for {
		// clear the screen
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()

		// choose next random question
		nextID := rand.Intn(len(questions))
		log.Printf("Showing question, %d\n\n", nextID)
		showQuestion(questions[nextID])

		// get answer from user and check if it is correct
		answer := getAnswerFromUser()
		answerState := checkAnswer(answer, questions[nextID])

		if answerState {
			fmt.Printf("\tCORRECT ")
			correctQuestions++
		} else {
			fmt.Printf("\tWRONG ")
			wrongQuestions++
		}
		fmt.Printf("(%s)\n\n\n", answer)

		// show the correct answer
		showCorrectAnswer(questions[nextID])
		totalQuestions++

		// show stats
		fmt.Printf("\nTotal %d\tCorrect %d\tIncorrect %d\n", totalQuestions, correctQuestions, wrongQuestions)

		fmt.Printf("\n\n(Press ENTER to continue)")
		enter := bufio.NewReader(os.Stdin)
		enter.ReadString('\n')

	}

}

func readQuestionsFromFile(file string) []Question {
	fmt.Println("readQuestionsFromFile()")
	fmt.Println(file)

	// read JSON from file
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// parse JSON
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var tmpQList []Question
	json.Unmarshal(byteValue, &tmpQList)
	return tmpQList

}

func getRandomSeed() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter random seed (default 0)")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n") // this will work on Windows
	tmpSeed, err := strconv.Atoi(text)
	if err != nil {
		text = strings.TrimSuffix(text, "\n") // this will work on Unix
		tmpSeed, err = strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
	}

	// fmt.Printf("\t%s\n", text)
	fmt.Printf("\n")

	if len(text) == 0 {
		return 0
	}

	return tmpSeed

}

func showQuestion(q Question) {
	fmt.Println(q.QText + "\n\n")
	for i := 0; i < len(q.PossAnswers); i++ {
		fmt.Println("- " + q.PossAnswers[i])
	}
	fmt.Printf("\n")
}

func getAnswerFromUser() string {
	fmt.Printf("\n\n")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	return text[:len(text)-1]
}

func checkAnswer(answer string, q Question) bool {
	tmpAns := strings.Join(q.Answers, "")
	// tmpAns = tmpAns[:len(tmpAns)-1]
	fmt.Printf("given:\t%s\t%d\n", tmpAns, len(tmpAns))
	fmt.Printf("answer:\t%s\t%d\n", answer, len(answer))

	if answer == tmpAns {
		return true
	}

	return false

}

func showCorrectAnswer(q Question) {
	tmpAns := strings.Join(q.Answers, "")
	fmt.Printf("Answer %s\t%s\n", tmpAns, q.Explanation)
}

func debug_ShowQuestions(qList []Question) {
	for i := 0; i < len(qList); i++ {
		fmt.Println("qTest: " + qList[i].QText)
	}
}
