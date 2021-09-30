package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func exitHelper(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func parseToProblemStruct(theSlice [][]string) []Problem {
	var returnedProblem []Problem = make([]Problem, len(theSlice))
	for index, value := range theSlice {
		returnedProblem[index] = Problem{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}
	}
	return returnedProblem
}

func main() {
	var csvFlag *string = flag.String("csv", "problems.csv", "a csv formatted file filled with 'question, answer' format")
	var limitFlag *int = flag.Int("limit", 30, "insert your limit time for the quiz in seconds, the default value is 30s")

	flag.Parse()
	var readCsvFile, err = os.Open(*csvFlag)
	if err != nil {
		exitHelper(fmt.Sprintf("Failed to open csv file: %s", *csvFlag))
	}

	var streamFile *csv.Reader = csv.NewReader(readCsvFile)
	problemsInSlice, err := streamFile.ReadAll()

	if err != nil {
		exitHelper("Failed to parse the provided CSV")
	}

	var parseProblem []Problem = parseToProblemStruct(problemsInSlice)
	var timer = time.NewTimer(time.Duration(*limitFlag) * time.Second)

	var countScore int = 0
	for index, value := range parseProblem {
		fmt.Printf("Question #%d: %s = ", index+1, value.question)
		var answerChan chan string = make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou score %d out of %d question", countScore, len(parseProblem))
			return
		case answer := <-answerChan:
			if answer == value.answer {
				countScore++
			}
		}
	}
	fmt.Printf("\nYou score %d out of %d question", countScore, len(parseProblem))
}
