package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	problemsFileName *string
	correctAnswers   int
)

func main() {
	problemsFileName = flag.String("csv", "problem.csv", "a csv file in the format of 'question,asnwer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	questions := readCSVFile(problemsFileName)

	go StartQuiz(questions)

	startTimer(time.Duration(*timeLimit) * time.Second)

	fmt.Printf("\nResult: %d / %d\n", correctAnswers, len(questions))
}

func readCSVFile(problemsFileName *string) [][]string {
	file, err := os.Open(*problemsFileName)

	if err != nil {
		exit(fmt.Sprintf("failed to open the file: %v \n", *problemsFileName))
	}

	defer file.Close()

	r := csv.NewReader(file)
	questions, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("failed to read csv file:  %v", *problemsFileName))
	}
	return questions
}

func StartQuiz(questions [][]string) {

	for index, record := range questions {
		question, correctAnswer := record[0], record[1]

		fmt.Printf("Problem #%d: %s = ", index+1, question)

		var asnwer string

		if _, err := fmt.Scan(&asnwer); err != nil {
			exit(fmt.Sprintf("failed to  scan: %v", err))
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
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
