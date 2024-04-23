package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const problemsFileName = "problem.csv"

func main() {
	file, err := os.Open(problemsFileName)
	if err != nil {
		fmt.Printf("failed to open the file: %v \n", err)
		return
	}

	defer file.Close()
	r := csv.NewReader(file)

	questions, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Something want wrong %v", err)
		return
	}

	var correctAnswers int
	for i, record := range questions {
		question, correctAnswer := record[0], record[1]

		//display one question at a time
		fmt.Printf("%d. %s ? \n", i+1, question)

		var asnwer string
		// get answer from user, thenm proceed to next question immediately
		if _, err := fmt.Scan(&asnwer); err != nil {
			fmt.Printf("failed to  scan: %v", err)
			return
		}
		if asnwer == correctAnswer {
			correctAnswers++
		}
	}
	// output the total number correct and how many
	fmt.Printf("Result: %d / %d\n", correctAnswers, len(questions))
}
