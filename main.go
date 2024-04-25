package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const problemsFileName = "problem.csv"

var (
	correctAnswers int
)

func main() {

	questions := readCSVFile(problemsFileName)

	go StartQuiz(questions)

	startTimer(20 * time.Second)

	fmt.Printf("Result: %d / %d\n", correctAnswers, len(questions))
}

func readCSVFile(problemsFileName string) [][]string {
	file, err := os.Open(problemsFileName)

	if err != nil {
		fmt.Printf("failed to open the file: %v \n", err)
		panic(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	questions, err := r.ReadAll()

	if err != nil {
		fmt.Printf("failed to read csv file:  %v", err)
		panic(err)
	}
	return questions
}

func StartQuiz(questions [][]string) {

	for index, record := range questions {
		question, correctAnswer := record[0], record[1]

		fmt.Printf("%d. %s ? \n", index+1, question)
		var asnwer string

		if _, err := fmt.Scan(&asnwer); err != nil {
			fmt.Printf("failed to  scan: %v", err)
			return
		}
		if asnwer == correctAnswer {
			correctAnswers++
		}
	}
	fmt.Printf("Result: %d / %d\n", correctAnswers, len(questions))
	os.Exit(0)
}

func startTimer(duration time.Duration) {
	time.Sleep(duration)
}
